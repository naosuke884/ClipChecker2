package main

import (
	"fmt"

	"github.com/naosuke884/ClipChecker2/later/config"
	"github.com/naosuke884/ClipChecker2/later/db"
	"github.com/naosuke884/ClipChecker2/later/presentation"
)

func main() {
	config.Load()
	fmt.Printf("%v", config.Config)
	db.InitializeSchema(true)
	presentation.MainServe()
}
