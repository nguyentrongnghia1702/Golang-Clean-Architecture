package middleware

import (
	"bytes"
	"io"
	"mcs-nghiadeptrai/mcs-common/logger"

	"github.com/gin-gonic/gin"
)

func LogRequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Body != nil {
			body, err := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			if err == nil {
				logger.LogInfo(string(body), c)
			} else {
				c.IndentedJSON(500, nil)
			}
		}

		c.Next()
	}

}
