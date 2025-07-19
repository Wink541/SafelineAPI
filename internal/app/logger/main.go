package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

var (
	Success *log.Logger
	Error   *log.Logger
	Warning *log.Logger
)

func LogInit() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix(fmt.Sprintf("%s[INFO]%s    ", Cyan, Reset))

	Success = log.New(os.Stdout, fmt.Sprintf("%s[SUCCESS]%s ", Green, Reset), log.Ldate|log.Ltime)
	Error = log.New(os.Stdout, fmt.Sprintf("%s[ERROR]%s   ", Red, Reset), log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, fmt.Sprintf("%s[WARNING]%s ", Yellow, Reset), log.Ldate|log.Ltime)
}
