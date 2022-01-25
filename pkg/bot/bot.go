package telegram

import (
	"github.com/cortez1488/pocket-telegram-bot/pkg/config"
	"github.com/cortez1488/pocket-telegram-bot/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	pocketClient    *pocket.Client
	tokenRepository repository.TokenRepository
	redirectUrl     string
	config          *config.Config
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, redirectUrl string, tr repository.TokenRepository, cfg *config.Config) *Bot {
	return &Bot{bot: bot, pocketClient: pocketClient, redirectUrl: redirectUrl, tokenRepository: tr, config: cfg}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}
	b.handleFunc(updates)
	return nil
}

func (b *Bot) handleFunc(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					b.handleError(update.Message.Chat.ID, err)
				}
				continue
			}
			if err := b.handleMessage(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}
		}
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}
