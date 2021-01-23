package project_service

import (
	"net/http"
	"simple_pm_api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Project プロジェクトの構造体
type Project models.Project

// Gets 全件取得
func Gets(transaction *gorm.DB, context *gin.Context) ([]Project, error) {
	var projects []Project
	if err := transaction.Find(&projects).Error; err != nil {
		return projects, err
	}
	return projects, nil
}

// Create プロジェクト登録
func Create(transaction *gorm.DB, context *gin.Context) (Project, error) {
	var project Project

	// // パラメータが不正な場合
	// if err := context.ShouldBindBodyWith(&project, binding.JSON); err != nil {
	// 	println(err.Error())
	// 	context.JSON(http.StatusUnprocessableEntity, gin.H{"code": "00001", "message": "パラメータが不正です！"})
	// 	return project, err
	// }
	// _ = context.BindJSON(&project)

	// パラメータが不正な場合
	if err := context.ShouldBind(&project); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"code": "00001", "message": "パラメータが不正です"})
		return project, err
	}

	err := transaction.Transaction(func(transction *gorm.DB) error {
		if err := transction.Create(&project).Error; err != nil {
			return err
		}
		return nil
	})

	// 登録に失敗した場合
	if err != nil {
		println(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"code": "AAA", "message": "msg"})
		return project, err
	}

	return project, nil
}
