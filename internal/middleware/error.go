package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/elSyarif/posnote-api.git/internal/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	var errMsgs []string
	for _, ginErr := range ctx.Errors {
		var verr validator.ValidationErrors
		if errors.As(ginErr.Err, &verr) {
			for _, er := range verr {
				errMsg := fmt.Sprintf("Error on field %s, condition: %s", er.Field(), er.ActualTag())
				errMsgs = append(errMsgs, errMsg)
			}
			helper.HTTPResponseError(ctx, http.StatusUnprocessableEntity, "fail", "validation error", errMsgs)
			return
		} else {
			helper.HTTPResponseError(ctx, http.StatusInternalServerError, "fail", ginErr.Err.Error(), nil)
			continue
		}

	}

}
