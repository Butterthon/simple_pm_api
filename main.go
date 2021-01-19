package main

import (
	"os"
	"simple_pm_api/pkg/core"
	v1ApiRouter "simple_pm_api/routers/v1"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title SIMPLE PM
// @version 1.0.0
// @description API仕様書

// @contact.name d.tabata

// @host localhost:8000
// @BasePath /api/
func main() {
	router := gin.New()

	// 環境変数ファイル読み込み
	core.LoadEnv(".env")

	// CORSの設定
	attachCORS(router)

	// v1APIルーター
	v1ApiRouter.Router(router)

	if debugMode, _ := strconv.ParseBool(os.Getenv("DEBUG")); debugMode {
		router.RunTLS(":8000", "ssl/server.crt", "ssl/server.key")
	} else {
		router.Run(":8000")
	}
}

// attachCORS CORSの設定
func attachCORS(router *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	corsConfig.AllowHeaders = strings.Split(os.Getenv("ALLOW_HEADERS"), ",")
	router.Use(cors.New(corsConfig))
}
