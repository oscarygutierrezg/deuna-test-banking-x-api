package api_route

import (
	"banking-api/internal/api/controller"
	"banking-api/internal/api/middleware"
	"banking-api/internal/service"
	"github.com/gin-gonic/gin"
)

func paymentApi(r *gin.RouterGroup, s *service.Services) {

	r.POST("",
		api_middleware.Payment.CreateValidation,
		api_controller.Payment.Create(s),
	)
}
