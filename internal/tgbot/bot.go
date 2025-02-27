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
		bot.WithDefaultHandler(MyDefaultHandler),
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, MyStartHandler),
		//bot.WithMessageTextHandler("/select", bot.MatchTypeExact, utils.CommandHandler),
		//bot.WithCallbackQueryDataHandler("btn_", bot.MatchTypePrefix, utils.CallbackHandler),
		bot.WithMessageTextHandler("/foot", bot.MatchTypeExact, FootballHandler),
	}

	b, err := bot.New(token, options...)
	if err != nil {
		fmt.Println("Error while starting bot:", err)
		return
	}

	b.Start(ctx)
}
