package router

import (
	"net/http"
)

func EchoHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Echo"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
