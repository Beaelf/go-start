package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(context *gin.Context) {
		s := time.Now()
		// path, response code, log latency
		logger.Info("incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))
	}, func(c *gin.Context) {
		c.Set("requestId", rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}

		if rid, exists := c.Get("requestId"); exists {
			h["requestId"] = rid
		}

		c.JSON(200, h)
	})
	
	r = gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run()
}
