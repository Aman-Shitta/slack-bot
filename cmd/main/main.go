package main

import (
	"fmt"
	"os"

	"github.com/Aman-Shitta/slack-bot/go-bot"
	uploader "github.com/Aman-Shitta/slack-bot/go-file-upload"
)

func main() {
	// Ensure at least one argument is provided
	if len(os.Args) < 2 {
		fmt.Println("❌ Usage: main [bot | upload]")
		os.Exit(1)
	}

	// Get the command
	command := os.Args[1]

	// Handle different commands
	switch command {
	case "bot":
		fmt.Println("🤖 Starting Slack Bot...")
		bot.Run()
	case "upload":
		fmt.Println("📤 Starting File Uploader...")
		uploader.Run()
	default:
		fmt.Printf("⚠️  Error: Unknown command '%s'\n", command)
		fmt.Println("✅ Available commands: bot | upload")
		os.Exit(1)
	}
}
