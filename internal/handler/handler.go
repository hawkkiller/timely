package handler

import (
	"strings"

	"github.com/hawkkiller/timely/internal/service"
	"github.com/mymmrac/telego"
)

type handler struct {
	service service.ScheduleService
	bot     *telego.Bot
}

type Handler interface {
	StartListening(updates <-chan telego.Update)
}

func NewHandler(service service.ScheduleService, bot *telego.Bot) Handler {
	return &handler{
		service: service,
		bot:     bot,
	}
}

func (h *handler) StartListening(updates <-chan telego.Update) {
	for upd := range updates {
		// start goroutine for each update
		go func(upd telego.Update) {
			text := upd.Message.Text
			if strings.HasPrefix(text, "/schedule") {
				h.handleGetSchedule(text)
			}
			if strings.HasPrefix(text, "/insert_schedule") {
				h.handleInsertSchedule(upd)
			}
		}(upd)
	}
}
