package middlewares

import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	gin.LoggerWithFormatter()
}
