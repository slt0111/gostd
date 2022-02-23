package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

const (
	host   = "123"
	port   = 1
	user   = "123"
	pwd    = "123"
	dbname = "123"
)

type User struct {
	ID int64 ``
}

func initDB(db *gorm.DB) {
	osqInfo := fmt.Sprintln("")
	db, connErro := gorm.Open("mysql", osqInfo+"?charset=utf8&parseTime=True&loc=Local")
	if connErro != nil {
		panic(connErro)
	}
	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate()

	return
}
func main() {
	fmt.Println("aaa")
}
