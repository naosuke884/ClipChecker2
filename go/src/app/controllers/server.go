package controllers

import (
	"ClipChecker2/config"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	http.HandleFunc("/favicon.ico", FaviconHandler)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
