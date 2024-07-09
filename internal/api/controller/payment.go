package api_controller

import (
	"banking-api/internal/entity"
	"banking-api/internal/service"
	"banking-api/pkg/uhttp"
	"banking-api/pkg/umdw"
	"github.com/gin-gonic/gin"
)

var Payment httpPayment

type httpPayment struct{}

func (httpPayment) Create(s *service.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req entity.PaymentRequest
		_ = umdw.BodyParse(&req, c)

		res, err := s.Payment.Create(&req)
		if err != nil {
			uhttp.Error(c, err)
			return
		}

		uhttp.Success(c, "Payment created successfully.", res)
	}
}
