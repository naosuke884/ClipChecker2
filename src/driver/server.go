package driver

import (
	"net/http"

	"github.com/naosuke884/ClipChecker2/later/config"
	"github.com/naosuke884/ClipChecker2/later/presenter"
	"github.com/naosuke884/ClipChecker2/later/service"
)

func ServerStart() {
	user := service.NewUserAppService()
	http.HandleFunc("/", presenter.TestPresenter())
	http.HandleFunc("/login", user.Login())
	http.ListenAndServe(":"+config.Config.Port, nil)
}
