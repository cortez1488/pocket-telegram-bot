package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"net/url"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return errInvalidURL
	}
	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return errUnauthorized
	}

	input := pocket.AddInput{URL: message.Text, AccessToken: accessToken}
	if err := b.pocketClient.Add(context.Background(), input); err != nil {
		return errUnableToSave
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	msg.Text = b.config.Messages.Responses.SaveSuccessfully
	b.bot.Send(msg)
	return nil
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}

}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthorizationProcess(message)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Messages.Responses.AlreadyAuthorized)
		b.bot.Send(msg)
	}
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.config.Messages.Responses.UnknownCommand)
	_, err := b.bot.Send(msg)
	return err
}
