package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/naosuke884/ClipChecker2/later/config"
	"github.com/naosuke884/ClipChecker2/later/infrastructure/provider"
)

func main() {
	config.Load()
	err := godotenv.Load("env/dev.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// db.InitializeSchema(true)
	// presentation.MainServe()

	client := resty.New()
	twitch := provider.NewTwitchProvider(client)
	fmt.Println(twitch.GetToken())
}
