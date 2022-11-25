package presenter

import (
	"html/template"
	"log"
	"net/http"
)

type Presenter func(w http.ResponseWriter, r *http.Request)

func TestPresenter() Presenter {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("static/test.html")
		if err != nil {
			log.Fatal(err)
		}
		if err := t.Execute(w, nil); err != nil {
			log.Fatal(err)
		}
	}
}
