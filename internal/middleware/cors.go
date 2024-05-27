package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() func(c *gin.Context) {
	// Default() allows all origins
	return cors.Default()
}
