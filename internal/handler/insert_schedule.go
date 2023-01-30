package handler

import (
	"fmt"

	"github.com/hawkkiller/timely/pkg/logger"
	"github.com/mymmrac/telego"
)

func (h *handler) handleInsertSchedule(upd telego.Update) {
	l := logger.GetLogger()
	doc := upd.Message.Document

	if doc == nil || doc.MimeType != "application/json" {
		return
	}

	file, err := h.bot.GetFile(&telego.GetFileParams{
		FileID: doc.FileID,
	})

	if err != nil {
		l.Warnf("can't get file with id %s", doc.FileID)
		return
	}

	token := h.bot.Token()

	link := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, file.FilePath)

	if err := h.service.InsertSchedule(link); err != nil {
		l.Warnf("can't insert schedule: %s", err)
		_, err := h.bot.SendMessage(&telego.SendMessageParams{
			ChatID: telego.ChatID{ID: upd.Message.Chat.ID},
			Text:   "Can't insert schedule",
		})
		if err != nil {
			l.Fatalf("can't send message: %s", err)
		}
		return
	}

	_, err = h.bot.SendMessage(&telego.SendMessageParams{
		ChatID: telego.ChatID{ID: upd.Message.Chat.ID},
		Text:   "Schedule inserted",
	})

	if err != nil {
		l.Fatalf("can't send message: %s", err)
	}
}
