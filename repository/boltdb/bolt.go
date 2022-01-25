package boltdb

import (
	"errors"
	"github.com/boltdb/bolt"
	"github.com/cortez1488/pocket-telegram-bot/repository"
	"strconv"
)

type BoltRepository struct {
	db *bolt.DB
}

func NewBoltRepository(db *bolt.DB) *BoltRepository {
	return &BoltRepository{db: db}
}

func (repo *BoltRepository) Save(chatID int64, token string, bucket repository.Bucket) error {
	err := repo.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(strconv.AppendInt(nil, chatID, 10), []byte(token))
	})
	return err
}

func (repo *BoltRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var token string

	err := repo.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(strconv.AppendInt(nil, chatID, 10))
		token = string(data)
		return nil
	})
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", errors.New("not found token for chatID")
	}

	return token, nil
}
