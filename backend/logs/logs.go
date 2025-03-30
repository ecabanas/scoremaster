package logs

import (
	"log"
	"os"
)

type CustomLogger struct {
	*log.Logger
}

func (l *CustomLogger) FatalIf(err error, message string) {
	if err != nil {
		l.Printf("%s: %v", message, err)
		os.Exit(1)
	}
}

var (
	Info    *CustomLogger
	Warning *CustomLogger
	Error   *CustomLogger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Info = &CustomLogger{log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)}
	Warning = &CustomLogger{log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)}
	Error = &CustomLogger{log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)}
}
