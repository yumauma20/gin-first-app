package main

import "gin-first-app/db"

func main() {
	//DBのopen処理とcloseを遅延
	db.Initialize()
	defer db.Close()

	//ルーティング設定
	// router.Router()
}
