package shutdown

import (
	"io"
	"os"
	"os/signal"

	"github.com/hawkkiller/timely/pkg/logger"
)

func Graceful(signals []os.Signal, closeItems ...io.Closer) {
	l := logger.GetLogger()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, signals...)
	sig := <-sigc
	l.Infof("Caught signal %s. Shutting down...", sig)

	// Here we can do graceful shutdown (close connections and etc)
	for _, closer := range closeItems {
		if err := closer.Close(); err != nil {
			l.Errorf("failed to close %v: %v", closer, err)
		}
	}
}
