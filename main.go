package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/root-man/gobuy/config"
)

type Bot struct {
	Config      *config.BotConfig
	Bot         *telego.Bot
	UpdatesChan <-chan telego.Update
}

func main() {

	botConfig, err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bot, err := initBot(botConfig)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
		))
	}, th.CommandEqual("start"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Unknown command, use /start",
		))
	}, th.AnyCommand())

	bh.Start()
}

func initBot(config *config.BotConfig) (*telego.Bot, error) {
	token := config.GetToken()

	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger(), telego.WithHealthCheck())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return bot, nil
}
