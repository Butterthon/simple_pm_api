package core

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv 環境変数ファイル読み込み
// TODO: ロギングの設定
func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("環境変数ファイルの読み込みに失敗しました")
	}
}
