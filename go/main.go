package main

import (
	"ClipChecker2/config"
	"ClipChecker2/driver"
)

func main() {
	config.Load()
	driver.ServerStart()
}
