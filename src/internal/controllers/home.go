package controllers

import (
	"golang-web-server/src/web/templates/pages/home"
	"net/http"
)

type homeHanders struct {} 

func(h *homeHanders) homeHandler(w http.ResponseWriter, r *http.Request) error {
	return home.Home(home.PageData{Title: "Home page", Message: "hello world!"}).Render(r.Context(), w)
}