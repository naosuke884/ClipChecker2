package config

import (
	"io/ioutil"
	"log"

	"github.com/naosuke884/ClipChecker2/later/utils"
	"gopkg.in/yaml.v2"
)

type ConfigList struct {
	Web struct {
		Port    string `yaml:"port"`
		LogFile string `yaml:"log-file"`
	}

	Db struct {
		Driver    string `yaml:"driver"`
		Sqlite3Db string `yaml:"sqlite3-db"`
	}
}

var Config ConfigList

func Load() {
	loadConfig()
	utils.LoadEnv()
	utils.LoggingSettings(Config.Web.LogFile)
}

func loadConfig() {
	content, _ := ioutil.ReadFile("config.yml")
	err := yaml.Unmarshal(content, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
