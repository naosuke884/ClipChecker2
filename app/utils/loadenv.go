package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(fmt.Sprintf("env/%v.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}
}
