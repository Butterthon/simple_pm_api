package core

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware CORSミドルウェア
func CORSMiddleware() gin.HandlerFunc {
	log.Println("CORSMiddleware")
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	corsConfig.AllowHeaders = strings.Split(os.Getenv("ALLOW_HEADERS"), ",")
	return cors.New(corsConfig)
}

// ProcessRequestMiddleware リクエストミドルウェア
func ProcessRequestMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
