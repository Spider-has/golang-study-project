package controllers

import "net/http"


type IRouter interface {
	GetHandlers() http.Handler
} 