package presenter

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/naosuke884/ClipChecker2/later/domain"
	"github.com/naosuke884/ClipChecker2/later/usecase"
)

type UserPresenter struct {
	usecase usecase.UserUseCase
}

func NewUserPresenter(usecase usecase.UserUseCase) *UserPresenter {
	return &UserPresenter{usecase: usecase}
}

type Presenter func(w http.ResponseWriter, r *http.Request)

func (presenter *UserPresenter) TestPresenter() Presenter {
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

func (presenter *UserPresenter) RegisterPresenter() Presenter {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			user := &domain.User{
				TwitchID: r.FormValue("twitch_id"),
				Name:     r.FormValue("name"),
				Email:    r.FormValue("email"),
				ImageURL: r.FormValue("image_url"),
				Token:    r.FormValue("token"),
			}
			if err := presenter.usecase.Register(user); err != nil {
				log.Fatal(err)
			}

			res_map := map[string]string{
				"status": "success",
			}
			res, err := json.Marshal(res_map)
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		}
	}
}

func (presenter *UserPresenter) LoginPresenter() Presenter {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := presenter.usecase.Login()

		if err != nil {
			log.Fatal(err)
		}

		res, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func (presenter *UserPresenter) LogoutPresenter() Presenter {
	return func(w http.ResponseWriter, r *http.Request) {
		err := presenter.usecase.Logout()

		if err != nil {
			log.Fatal(err)
		}
	}
}
