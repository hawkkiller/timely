package bot

import (
	"github.com/mymmrac/telego"
	"github.com/pkg/errors"
)

func Start(token string) (*telego.Bot, <-chan telego.Update, error) {
	bot, err := telego.NewBot(token)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create bot")
	}

	updates, err := bot.UpdatesViaLongPulling(nil)

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get updates")
	}

	return bot, updates, nil
}
