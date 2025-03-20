package exceptions

import (
	"errors"
	"fmt"
	"net/http"

	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"github.com/go-playground/validator/v10"
)

var (
	ErrorEmailAlreadyExists = errors.New("email already registered")
)

var errorStatusMap = map[error]int{
	ErrorEmailAlreadyExists: http.StatusConflict,
}

func ErrorHandler(err error) communication.ResponseDTO {
	if status, exists := errorStatusMap[err]; exists {
		return communication.ResponseDTO{
			StatusCode: status,
			Message:    err.Error(),
		}
	}

	return communication.ResponseDTO{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
	}
}

func CreateValidationErrorMessages(err error) []string {
	var verr validator.ValidationErrors

	if errors.As(err, &verr) {
		fmt.Println(err)
		out := make([]string, 0, len(verr))
		for _, fe := range verr {
			var msg string
			switch fe.Tag() {
			case "required":
				msg = fmt.Sprintf("Field '%s' is required.", fe.Field())
			case "email":
				msg = fmt.Sprintf("Field '%s' should be a valid e-mail.", fe.Field())
			case "min":
				msg = fmt.Sprintf("Field '%s' should have at least %s characters.", fe.Field(), fe.Param())
			default:
				msg = fmt.Sprintf("Field '%s' is invalid.", fe.Field())
			}
			out = append(out, msg)
		}
		return out
	}
	return make([]string, 0)
}
