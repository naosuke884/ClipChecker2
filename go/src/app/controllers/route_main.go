package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("URL: "+r.URL.Path, "Method:"+r.Method)
}

func top(w http.ResponseWriter, r *http.Request) {
	baseHandler(w, r)
	fmt.Println(color.BlueString("top2"))
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	baseHandler(w, r)
	// http.ServeFile(w, r, "static/favicon.ico")
}
