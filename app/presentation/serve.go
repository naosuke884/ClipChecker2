package presentation

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/naosuke884/ClipChecker2/later/infrastructure"
	"github.com/naosuke884/ClipChecker2/later/service"
)

var db *sql.DB
var e *echo.Echo

func MainServe() {
	// echoのクライアントを作成
	e = echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}\n",
	}))

	//dbのクライアントを作成
	var err error
	db, err = sql.Open("sqlite3", "./db/clipchecker.db")
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// ユーザーのプレゼンターを作成
	presenter := getUserPresenter()
	presenter.ServeUser()

	// サーバーを起動
	e.Logger.Fatal(e.Start(":1323"))
}

func getUserPresenter() IUserPresenter {
	repository := infrastructure.NewUserSqlite3Repository(db)
	service := service.NewUserApplicationService(repository)
	echo := e
	presenter := NewUserPresenter(service, repository, echo)
	return presenter
}
