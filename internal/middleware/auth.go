package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wupyweb/realworld-kit/ent"
	"github.com/wupyweb/realworld-kit/internal/utils"
)

func AuthMiddleware(client *ent.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 从请求头中获取Authorization
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // 终止请求
			return
		}
		// 验证token
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		payload, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort() // 终止请求
			return
		}
		// 验证token成功后，获取当前用户,将用户id存储到上下文中
		c.Set("user_id", payload.UserID)
		c.Set("username", payload.UserName)
	

		// 请求前

		c.Next()

		// 请求后
	}
}
