package tgbot

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
)

// Start bot
func Start(token string) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	options := []bot.Option{
		bot.WithMessageTextHandler("/select", bot.MatchTypeExact, CommandHandler),
		bot.WithDefaultHandler(MyDefaultHandler),
		//bot.WithCallbackQueryDataHandler("btn_", bot.MatchTypePrefix, utils.CallbackHandler),

	}

	b, err := bot.New(token, options...)
	if err != nil {
		fmt.Println("Error while starting bot:", err)
		return
	}

	// Handlers:
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, MyStartHandler)

	b.Start(ctx)
}
