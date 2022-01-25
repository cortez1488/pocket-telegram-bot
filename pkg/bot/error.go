package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidURL   = errors.New("url is not valid")
	errUnauthorized = errors.New("user in not authorized")
	errUnableToSave = errors.New("unable to save")
)

//
//
//
func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "")
	switch err {
	case errInvalidURL:
		msg.Text = b.config.Messages.Errors.InvalidUrl
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = b.config.Messages.Errors.Unauthorized
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = b.config.Messages.Errors.UnableToSave
		b.bot.Send(msg)
	default:
		msg.Text = b.config.Messages.Errors.Default
		b.bot.Send(msg)
	}

}
