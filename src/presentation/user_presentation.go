package presentation

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/naosuke884/ClipChecker2/later/domain"
)

type IUserPresenter interface {
	Register() echo.HandlerFunc
	Get() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	ServeUser()
}

type UserPresenter struct {
	service    domain.IUserApplicationService
	repository domain.IUserRepository
	echo       *echo.Echo
}

func NewUserPresenter(service domain.IUserApplicationService, repository domain.IUserRepository, echo *echo.Echo) IUserPresenter {
	return &UserPresenter{
		service:    service,
		repository: repository,
		echo:       echo,
	}
}

func (p *UserPresenter) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := domain.User{
			TwitchID: "twitch_id",
			Name:     "name",
			Email:    "email",
			ImageURL: "image_url",
		}
		err := p.service.Register(&user)
		if err != nil {
			e.Logger.Fatal(err)
		}

		return c.String(http.StatusOK, "register")
	}
}

func (p *UserPresenter) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			e.Logger.Fatal(err)
		}
		user, err := p.service.Get(domain.ID(id))
		if err != nil {
			e.Logger.Fatal(err)
		}
		return c.JSON(http.StatusOK, user)
	}
}

func (p *UserPresenter) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			e.Logger.Fatal(err)
		}
		user := domain.User{
			ID:       domain.ID(id),
			TwitchID: "twitch_id2",
			Name:     "name2",
			Email:    "email2",
			ImageURL: "image_url2",
		}
		err = p.service.Update(&user)
		if err != nil {
			e.Logger.Fatal(err)
		}
		return c.String(http.StatusOK, "update")
	}
}

func (p *UserPresenter) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			e.Logger.Fatal(err)
		}
		err = p.service.Delete(domain.ID(id))
		if err != nil {
			e.Logger.Fatal(err)
		}
		return c.String(http.StatusOK, "delete")
	}
}

func (p *UserPresenter) ServeUser() {
	p.echo.POST("/user/register", p.Register())
	p.echo.GET("/user/:id", p.Get())
	p.echo.POST("/user/update/:id", p.Update())
	p.echo.DELETE("/user/:id", p.Delete())
}
