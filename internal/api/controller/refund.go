package api_controller

import (
	"banking-api/internal/entity"
	"banking-api/internal/service"
	"banking-api/pkg/uhttp"
	"banking-api/pkg/umdw"
	"github.com/gin-gonic/gin"
)

var Refund httpRefund

type httpRefund struct{}

func (httpRefund) Create(s *service.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req entity.RefundRequest
		_ = umdw.BodyParse(&req, c)

		res, err := s.Refund.Create(&req)
		if err != nil {
			uhttp.Error(c, err)
			return
		}

		uhttp.Success(c, "Refund created successfully.", res)
	}
}
