package loggers

import (
	"fmt"
	"log"
	"os"
)

var (
	ErrorData *log.Logger
	WarnData  *log.Logger
	InfoData  *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Error Occured", err)
		return
	}
	InfoData = log.New(file, "INFO: ", log.Lshortfile|log.LstdFlags)
	ErrorData = log.New(file, "ERROR: ", log.Lshortfile|log.LstdFlags)
	WarnData = log.New(file, "WARN:", log.Lshortfile|log.LstdFlags)
}
