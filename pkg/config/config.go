package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	AuthServerUrl     string
	TelegramBotUrl    string `mapstructure:"bot_url"`
	DBpath            string `mapstructure:"db_file"`
	Messages          Messages
}

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidUrl   string `mapstructure:"invalid_url"`
	Unauthorized string `mapstructure:"unauthorized"`
	UnableToSave string `mapstructure:"unable_to_save"`
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	SaveSuccessfully  string `mapstructure:"save_success"`
	UnknownCommand    string `mapstructure:"unknown_command"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	if err := os.Setenv("TOKEN", "5187131287:AAH7x1R1GzEIpOK_RCgz9xieOqjzIRVmhug"); err != nil {
		log.Fatal(err)
	}
	os.Setenv("CONSUMER_KEY", "100583-074d230a1c7f4fa6165fe80")
	os.Setenv("AUTH_SERVER_URL", "http://127.0.0.1:8000/")
	os.Setenv("BOT_REDIRECT", "https://t.me/goEduPocketBot/")
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}
	if err := viper.BindEnv("auth_server_url"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("token")
	cfg.PocketConsumerKey = viper.GetString("consumer_key")
	cfg.AuthServerUrl = viper.GetString("auth_server_url")
	cfg.TelegramBotUrl = viper.GetString("bot_redirect")
	return nil
}
