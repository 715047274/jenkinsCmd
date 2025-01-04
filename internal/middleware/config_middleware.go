package middleware

import (
	"github.com/715047274/jenkinsCmd/internal/config"
	"github.com/gin-gonic/gin"
)

// ConfigMiddleware binds the configuration to the Gin context
func ConfigMiddleware(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", cfg)
		// log.Panicln("------MiddleWare Called-------")
		c.Next()
	}
}
