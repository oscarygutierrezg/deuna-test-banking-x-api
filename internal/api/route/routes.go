package api_route

import (
	"banking-api/internal/service"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.RouterGroup, s *service.Services) {
	paymentApi(r.Group("/payments"), s)
	refundApi(r.Group("/refunds"), s)
}
