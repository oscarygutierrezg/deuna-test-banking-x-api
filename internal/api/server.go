package api

import (
	"banking-api/internal/api/route"
	"banking-api/internal/service"
	"banking-api/pkg/umdw"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewServer(s *service.Services) *gin.Engine {
	r := gin.Default()
	version := r.Group("/v1")

	version.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Type", " Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowWildcard:    true,
		AllowWebSockets:  true,
	}))

	version.Use(umdw.BodyContext)

	version.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"name":   "Banking API",
			"status": "running",
		})
	})

	api_route.SetRoutes(version, s)

	return r
}
