package main

import (
	"github.com/gin-first-app/db"
	"github.com/gin-first-app/router"
)

func main() {
	//DBのopen処理とcloseを遅延
	db.Initialize()
	defer db.Close()

	//ルーティング設定
	router.Router()
}
