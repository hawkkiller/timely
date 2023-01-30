package main

import (
	"os"

	"github.com/hawkkiller/timely/internal/db"
	"github.com/hawkkiller/timely/pkg/conn"

	"github.com/hawkkiller/timely/internal/bot"
	"github.com/hawkkiller/timely/internal/handler"
	"github.com/hawkkiller/timely/internal/service"
	"github.com/hawkkiller/timely/pkg/logger"
)

// main is the entry point for the application.
func main() {
	start()
}

func start() {
	logger.Init()
	l := logger.GetLogger()
	l.Info("Get token from environment...")
	token := os.Getenv("TELEGRAM_TOKEN")
	l.Info("Connecting to database...")
	connection, err := conn.Connect()
	if err != nil {
		l.Fatal(err)
	}
	l.Info("Start bot configuration...")
	bot, updates, err := bot.Start(token)
	if err != nil {
		l.Fatal(err)
	}
	l.Info("Create DB client...")
	scheduleDB := db.NewScheduleDB(connection)
	l.Info("Create service...")
	svc := service.NewScheduleService(scheduleDB)
	l.Info("Register handler...")
	h := handler.NewHandler(svc, bot)
	l.Info("Start bot...")
	h.StartListening(updates)
}
