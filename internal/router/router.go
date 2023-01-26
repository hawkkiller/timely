package router

import "github.com/julienschmidt/httprouter"

func ConfigureRouter() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc("GET", "/echo", EchoHandler)
	return router
}
