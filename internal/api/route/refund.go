package api_route

import (
	"banking-api/internal/api/controller"
	"banking-api/internal/api/middleware"
	"banking-api/internal/service"
	"github.com/gin-gonic/gin"
)

func refundApi(r *gin.RouterGroup, s *service.Services) {

	r.POST("",
		api_middleware.Refund.CreateValidation,
		api_controller.Refund.Create(s),
	)
}
