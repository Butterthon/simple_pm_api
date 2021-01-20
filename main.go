package main

import (
	"os"
	"simple_pm_api/pkg/core"
	v1ApiRouter "simple_pm_api/routers/v1"
	"strconv"

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

	// ミドルウェアの設定
	router.Use(core.CORSMiddleware())
	router.Use(core.ProcessRequestMiddleware())

	// v1APIルーター
	v1ApiRouter.Router(router)

	if debugMode, _ := strconv.ParseBool(os.Getenv("DEBUG")); debugMode {
		router.RunTLS(":8000", "ssl/server.crt", "ssl/server.key")
	} else {
		router.Run(":8000")
	}
}
