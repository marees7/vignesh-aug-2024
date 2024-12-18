package helpers

import (
	"fmt"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
)

func StringConvertion(jobIDStr string) (int, error) {
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		loggers.ErrorData.Println("error occured while String Convertion,Please check properly")
		return 0, fmt.Errorf("error occured while String Convertion,Please check properly")
	}
	return jobID, nil
}
