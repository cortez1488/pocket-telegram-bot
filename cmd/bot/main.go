package main

import (
	"github.com/boltdb/bolt"
	telegram "github.com/cortez1488/pocket-telegram-bot/pkg/bot"
	"github.com/cortez1488/pocket-telegram-bot/pkg/config"
	"github.com/cortez1488/pocket-telegram-bot/pkg/server"
	"github.com/cortez1488/pocket-telegram-bot/repository"
	bolt2 "github.com/cortez1488/pocket-telegram-bot/repository/boltdb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := initDB(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	tokenRepo := bolt2.NewBoltRepository(db)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal("NewBotApi: ", err)
	}

	pocketClient, err := pocket.NewClient(cfg.PocketConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	telegramBot := telegram.NewBot(bot, pocketClient, cfg.AuthServerUrl, tokenRepo, cfg)

	authServer := server.NewAuthorizationServer(pocketClient, tokenRepo, cfg.TelegramBotUrl)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authServer.Start(); err != nil {
		log.Fatal(err.Error())
	}

}

func initDB(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.DBpath, 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessToken))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
