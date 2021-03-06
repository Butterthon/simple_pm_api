package project

import (
	"net/http"
	"simple_pm_api/services"
	"simple_pm_api/services/project_service"

	"github.com/gin-gonic/gin"
)

// Gets 全件取得
// @Summary プロジェクト全件取得
// @Producs json
// @Success 200 {object} []models.Project
// @Failure 400 {object} exceptions.APIException(detail={})
// @Failure 422 {object} exceptions.APIException(detail={})
// @Failure 500 {object} exceptions.APIException(detail={})
// @Router /project/ [get]
func Gets(context *gin.Context) {
	projects, err := project_service.Gets(services.GetOrm(), context)
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, projects)
}

// Create プロジェクト登録
// @Summary プロジェクト登録
// @Producs json
// @Param project body models.Project true "..."
// @Success 200 {object} responses.JSONResponse(data={})
// @Failure 400 {object} exceptions.APIException(detail={})
// @Failure 422 {object} exceptions.APIException(detail={})
// @Failure 500 {object} exceptions.APIException(detail={})
// @Router /project/ [post]
func Create(context *gin.Context) {
	transaction := services.GetOrm().Begin()

	project, err := project_service.Create(transaction, context)
	if err != nil {
		transaction.Rollback()
		return
	}
	context.JSON(http.StatusOK, project)

	transaction.Commit()
}
