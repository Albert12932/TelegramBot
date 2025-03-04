package tgbot

import (
	"context"
	"encoding/json"
	"github.com/go-telegram/bot/models"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var update models.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	// ✅ Теперь мы вручную вызываем хэндлер для команды "/start"
	if update.Message != nil && update.Message.Text == "/start" {
		MyStartHandler(context.Background(), BotInstance, &update)
	}
	if update.Message != nil && update.Message.Text == "/foot" {
		FootballHandler(context.Background(), BotInstance, &update)
	}
	if update.Message != nil {
		MyDefaultHandler(context.Background(), BotInstance, &update)
	}

	// HTTP 200 OK
	w.WriteHeader(http.StatusOK)
}
