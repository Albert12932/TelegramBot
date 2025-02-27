package tgbot

import (
	"context"
	"github.com/Albert12932/TelegramBot/pkg/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

// MyStartHandler handles "/start" command with "Hello, name"
func MyStartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Hello, " + bot.EscapeMarkdown(update.Message.From.FirstName),
		ParseMode: models.ParseModeMarkdown,
	})

	if err != nil {
		log.Printf("Error %s in %v", err, utils.GetFunctionName())
		return
	}

	log.Printf("send message to Chat %v: %v", update.Message.Chat.ID, message.Text)
}

func MyDefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Default Handler",
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Printf("Error %s in %v", err, utils.GetFunctionName())
		return
	}

	log.Printf("send message to Chat %v: %v", update.Message.Chat.ID, message.Text)

}

// Multiple select handler
//func CommandHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
//	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
//		ChatID:      update.Message.Chat.ID,
//		Text:        "Select multiple options",
//		ReplyMarkup: utils.BuildKeyboard(),
//	})
//	if err != nil {
//		log.Printf("Error %s in %v", err, utils.GetFunctionName())
//		return
//	}
//
//	log.Printf("send message to Chat %v: %v", update.Message.Chat.ID, message.Text)
//}

func FootballHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Enter",
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Printf("Error %s in %v", err, utils.GetFunctionName())
		return
	}

	log.Printf("send message to Chat %v: %v", update.Message.Chat.ID, message.Text)
}
