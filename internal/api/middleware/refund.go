package api_middleware

import (
	"banking-api/pkg/uhttp"
	"banking-api/pkg/umdw"
	"banking-api/pkg/util"
	"github.com/gin-gonic/gin"
)

var Refund httpRefundMdw

type httpRefundMdw struct{}

func (httpRefundMdw) CreateValidation(c *gin.Context) {
	require := []string{
		"transactionId",
		"amount",
		"currency",
	}

	verify := umdw.VerificationFunctions{}

	err := umdw.BodyVerifyFields(c, require, verify)
	if err != nil {
		uhttp.Error(c, &api_util.RequiredFieldError{Message: err.Error()})
		return
	}

	c.Next()
}
