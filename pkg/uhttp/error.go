package uhttp

import (
	"banking-api/internal/service"
	"banking-api/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CustomError(c *gin.Context, code int, data interface{}) {
	var response = Response{
		Data: data,
	}
	response.reply(c, code)
}

func Error(c *gin.Context, err error) {
	var status int
	var customErr *api_util.RequiredFieldError
	switch {
	case errors.Is(err, service.InvalidCardDetails), errors.As(err, &customErr):
		status = http.StatusBadRequest
	case errors.Is(err, service.InsufficientFunds):
		status = http.StatusPaymentRequired
	case errors.Is(err, service.TransactionAlreadyRefunded):
		status = http.StatusConflict
	case errors.Is(err, service.TransactionNotFound):
		status = http.StatusNotFound
	default:
		status = http.StatusInternalServerError
	}
	CustomError(c, status, err.Error())
}
