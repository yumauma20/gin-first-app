package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.db

func Initialize() {
	// 宣言済みのグローバル変数dbを（:=）で初期化しようとするとローカル変数dbを初期化することになるので注意

	//dbのコネクション接続する
	db, _ = gorm.Open("sqlite3", "task.db")

	//ログの有効化
	db.LogMode(true)

	//task構造体（models）をもとにマイグレーションを実行
	db.AutoMigrate(&models.Task{})
}

func Get() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
