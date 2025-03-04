package tgbot

import (
	"errors"
	"fmt"
	"github.com/go-telegram/bot"
	"log"
	"net/http"
)

var BotInstance *bot.Bot

// Start bot
func Start(token, webHookURL string) error {
	var err error

	BotInstance, err = bot.New(token)
	if err != nil {
		fmt.Println("Error while starting bot:", err)
		return errors.New("Error while starting bot!")
	}

	err = setWebhook(token, webHookURL)
	if err != nil {
		fmt.Println("Error while setting webhook:", err)
	}

	http.HandleFunc("/webhook", webhookHandler)
	log.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

func setWebhook(token, url string) error {
	webHookURL := fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", token, url)
	resp, err := http.Get(webHookURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Println("Successfully set webhook to", url)
	return nil
}
