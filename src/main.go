package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/naosuke884/ClipChecker2/later/config"
	"github.com/naosuke884/ClipChecker2/later/infrastructure/provider"
)

func main() {
	config.Load()
	// db.InitializeSchema(true)
	// presentation.MainServe()
	client := resty.New()
	twitch := provider.NewTwitchProvider(client)
	fmt.Println(twitch.GetToken())
}
