package core

import (
	"strings"
)

// TelegramMarkdownV2Escape escapes special characters for Telegram's MarkdownV2 format.
// Characters that need escaping: _ * [ ] ( ) ~ ` > # + - = | { } . !
func TelegramMarkdownV2Escape(text string) string {
	// Order matters - escape backslash first to avoid double-escaping
	specialChars := []string{"\\", "_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
	
	result := text
	for _, ch := range specialChars {
		result = strings.ReplaceAll(result, ch, "\\"+ch)
	}
	return result
}

// TelegramMarkdownV2EscapeCode escapes text for use inside code blocks (where only ` needs escaping)
func TelegramMarkdownV2EscapeCode(text string) string {
	return strings.ReplaceAll(text, "`", "\\`")
}
