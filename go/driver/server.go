package driver

import (
	"ClipChecker2/config"
	"ClipChecker2/presenter"
	"net/http"
)

func ServerStart() {
	http.HandleFunc("/", presenter.TestPresenter())
	http.ListenAndServe(":"+config.Config.Port, nil)
}
