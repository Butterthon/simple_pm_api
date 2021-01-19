package v1

import (
	"net/http"
	"os"
	"strconv"

	_ "simple_pm_api/docs" // swag initで生成したdocsフォルダ

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Router ルーター
func Router(router *gin.Engine) {
	// 疎通確認用
	router.GET("/", test)

	// Swagger-UIの設定
	debugMode, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if debugMode {
		url := ginSwagger.URL(os.Getenv("ASSET_DOMAIN") + "swagger/doc.json")
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}

// @Summary テスト
// @Producs json
// @Success 200 {object} responses.JSONResponse(data={})
// @Failure 400 {object} exceptions.APIException(detail={})
// @Failure 500 {object} exceptions.APIException(detail={})
// @Router /test [get]
func test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
