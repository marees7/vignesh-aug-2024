package helpers

import (
	"fmt"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/common/constants"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"golang.org/x/crypto/bcrypt"
)

func StringConvertion(jobIDStr string) (int, error) {
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		loggers.ErrorData.Println("Error occured while String Convertion,Please check properly")
		return 0, fmt.Errorf("error occured while String Convertion,Please check properly")
	}
	return jobID, nil
}

// Hashing the password here
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// by using limit and offset using pagination
func Pagination(offsetStr, limitStr string) (limit, offset int) {
	limit, _ = strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = constants.DefaultLimit
	}

	offset, _ = strconv.Atoi(offsetStr)
	switch {
	case offset <= 0:
		offset = constants.DefaultOffset
	}
	return limit, offset
}
