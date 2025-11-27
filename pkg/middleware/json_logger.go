package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONBindLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Đọc body request
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("[BindJSON] ERROR: cannot read body:", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
			return
		}

		// Phục hồi body cho Gin (vì ReadAll sẽ "consume" body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Gọi next handler
		c.Next()

		// Kiểm tra error gắn vào context (nếu có)
		for _, err := range c.Errors {
			if err.Type == gin.ErrorTypeBind {
				log.Println("=========================================")
				log.Println("[BindJSON] Error:", err.Err)
				log.Println("[BindJSON] Request Body:", string(bodyBytes))
				log.Println("[BindJSON] Method:", c.Request.Method, "URL:", c.Request.URL.Path)
				log.Println("=========================================")
			}
		}
	}
}
