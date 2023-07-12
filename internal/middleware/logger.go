package middleware

import (
	"fmt"
	"strings"
	"time"
	"v1/pkg"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		// raw := c.Request.URL.RawQuery
		method := c.Request.Method
		c.Next()
		end := time.Now()
		pkg.Log.Info("",
			zap.String("method", method),
			zap.Int("status", c.Writer.Status()),
			zap.String("path", path),
			zap.String("take time", strings.Trim(fmt.Sprintf("%13v", end.Sub(start)), " ")),
		)
	}
}
