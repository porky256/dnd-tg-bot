package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/porky256/dnd-tg-bot/internal/database"
	"github.com/porky256/dnd-tg-bot/internal/database/postgres"
	"github.com/porky256/dnd-tg-bot/internal/handlers"
	"go.uber.org/zap"
	"os"
	"time"
)

var config database.DBConfig
var token string

func main() {
	pg, err := postgres.ConnectPGSQL(config)
	log, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("logger is not created: " + err.Error())
	}
	if err != nil {
		log.Fatal("error have occurred while connecting to database: " + err.Error())
	}
	log.Info("data Provider connected successfully")
	handler := handlers.NewHandlers(pg, 3*time.Second, log)
	b, err := bot.New(token, bot.WithDefaultHandler(handler.DefaultHandler))
	if err != nil {
		log.Fatal("error have occurred while connecting to bot: " + err.Error())
	}

	log.Info("Bot connected successfully")
	b = handler.Register(b)
	log.Info("Bot started")
	b.Start(context.Background())
}

func init() {
	dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode, telegramToken :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_CONTAINER_NAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("TELEGRAM_BOT_TOKEN")

	config = database.DBConfig{
		DriverName:    "postgres",
		Host:          dbHost,
		Port:          dbPort,
		Name:          dbName,
		User:          dbUser,
		Password:      dbPassword,
		SSLMode:       dbSSLMode,
		MaxIdleDBConn: 5,
		MaxOpenDBConn: 10,
	}

	token = telegramToken
}
