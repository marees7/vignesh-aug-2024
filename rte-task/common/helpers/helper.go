package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/common/constants"
	"github.com/Vigneshwartt/golang-rte-task/common/dto"
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
func Pagination(offsetStr, limitStr string) (int, int, *dto.ErrorResponse) {
	offset, err := strconv.Atoi(offsetStr)
	if offsetStr == "" {
		offset = constants.DefaultOffset
	} else if err != nil {
		return 0, 0, &dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
	}

	limit, err := strconv.Atoi(limitStr)
	if limitStr == "" {
		limit = constants.DefaultLimit
	} else if err != nil {
		return 0, 0, &dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		}
	}
	offset = (offset - 1) * limit

	return limit, offset, nil
}
