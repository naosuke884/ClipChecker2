package config

import (
	"ClipChecker2/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/michaeljs1990/sqlitestore"
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config ConfigList

var Store *sqlitestore.SqliteStore

func Load() {
	loadConfig()
	loadEnv()
	setStore()
	utils.LoggingSettings(Config.LogFile)
}

func loadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}

func loadEnv() {
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setStore() {
	var err error
	Store, err = sqlitestore.NewSqliteStore("ClipChecker.db", "sessions", "/", 3600, []byte(os.Getenv("SESSION_KEY")))
	if err != nil {
		log.Fatal(err)
	}
}
