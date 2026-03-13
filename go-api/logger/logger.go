package logger

import (
	"log"
	"os"
)

var Log = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

func init() {
	Log.Println("Logger initialized")
}
