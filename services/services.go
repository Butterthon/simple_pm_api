package services

import (
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // インポートしておかないと「unknown driver "postgresql"」でエラーになる
	"github.com/jinzhu/gorm"
)

var orm *gorm.DB

// DBConnectionSetup DBコネクション接続
func DBConnectionSetup() {
	var err error
	orm, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("DB接続に失敗しました", err)
	}

	// orm.SetMaxIdleConns(10) // アイドル状態のコネクションプール最大数
	// orm.DB().SetMaxOpenConns(100) // オープンコネクションプール最大数
	// orm.DB().SetConnMaxLifetime(time.Hour) // コネクションを再利用できる最大時間
}

// GetOrm ...
func GetOrm() *gorm.DB {
	return orm
}

// CloseDBConnection DBコネクションを閉じる
func CloseDBConnection() {
	defer orm.Close()
}

// Begin トランザクション開始
func Begin() *gorm.DB {
	return orm.Begin()
}

// Commit コミット
func Commit() {
	orm.Commit()
}

// Rollback ロールバック
func Rollback() {
	orm.Rollback()
}
