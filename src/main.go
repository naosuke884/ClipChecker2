package main

import (
	"github.com/naosuke884/ClipChecker2/later/config"
	"github.com/naosuke884/ClipChecker2/later/driver"
)

func main() {
	config.Load()
	driver.ServerStart()
}
