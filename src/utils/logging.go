package utils

import (
	"io"
	"log"
	"os"

	"github.com/natefinch/lumberjack"
)

func LoggingSettings(logFile string) {
	writer := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    100, // megabytes
		MaxBackups: 10,
		MaxAge:     28, //days
		Compress:   true,
	}
	multiLogFile := io.MultiWriter(os.Stdout, writer)    //標準出力とファイル出力を同時に行う
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //ログのフォーマットの指定
	log.SetOutput(multiLogFile)                          //出力先指定
}
