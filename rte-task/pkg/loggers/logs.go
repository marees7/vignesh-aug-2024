package loggers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	ErrorData *log.Logger
	WarnData  *log.Logger
	InfoData  *log.Logger
)

func LoggerFiles() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filepath.Join(filepath.Dir(wd), os.Getenv("LOG_FILE_PATH")), os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Error Occured", err)
		return
	}
	InfoData = log.New(file, "INFO: ", log.Lshortfile|log.LstdFlags)
	ErrorData = log.New(file, "ERROR: ", log.Lshortfile|log.LstdFlags)
	WarnData = log.New(file, "WARN:", log.Lshortfile|log.LstdFlags)
}
