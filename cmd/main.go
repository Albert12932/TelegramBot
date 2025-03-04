package main

import (
	"github.com/Albert12932/TelegramBot/internal/tgbot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env ", err)
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	webHookURL := os.Getenv("WEBHOOK_URL")
	if err := tgbot.Start(token, webHookURL); err != nil {
		log.Fatal("Ошибка при запуске бота", err)
	}

}
