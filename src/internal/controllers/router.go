package controllers

import (
	"fmt"
	"net/http"
)


type MuxRouter struct {
	mux *http.ServeMux
}

type handleFunc func(w http.ResponseWriter, r *http.Request) error

func NewMuxRouter() IRouter {
	return &MuxRouter{
		mux: http.NewServeMux(),
	}
}

func(mr *MuxRouter) GetHandlers() http.Handler {

	home := homeHanders{}

	mr.mux.HandleFunc("/", handleHandler(home.homeHandler))

	return mr.mux
}

func handleHandler(handler handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w,r); err != nil {
			handleError(w,r,err)
		}
	}
}


func handleError(_ http.ResponseWriter, _ *http.Request, err error) {
	fmt.Errorf("error while handling request: %s", err.Error())
}