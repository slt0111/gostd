package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "root"
	dbname   = "go_test"
)

type User struct {
	ID         int64     `gorm:"type:bigint(20);column:id;primary_key"`
	UserName   string    `gorm:"type:varchar(255);column:username"`
	NickName   string    `gorm:"type:varchar(255);column:nickname"`
	PassWord   string    `gorm:"type:varchar(255);column:password"`
	Phone      string    `gorm:"type:varchar(255);column:phone"`
	CreateTime time.Time `gorm:"column:createtime"`
	UpdateTime time.Time `gorm:"column:updatetime"`
}

var db *gorm.DB
var connErr error

func initDB() (conn *gorm.DB) {
	osqlInfo := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, password, host, port, dbname)
	fmt.Println(osqlInfo)
	db, connErr = gorm.Open("mysql", osqlInfo+"?charset=utf8&parseTime=true&loc=Local")

	if connErr != nil {
		panic(connErr)
	}
	//defer db.Close()'
	db.LogMode(true)
	db.SingularTable(true)
	db.AutoMigrate(&User{})

	return db
}

//新增
func add() {
	test := &User{
		ID: 3,
	}
	fmt.Println("add")
	db.Create(test)
}

//查询全部
func selectA() {
	fmt.Println("select")
	var res []User
	//db.Where("username = ?","a").First(&res)
	db.Find(&res)
	for _, v := range res {
		fmt.Println(v)
	}
	//fmt.Println("res:",res)
}

//自定义查询
func selectC() {
	//查询第一条记录
	var user User
	db.First(&user, "username = ?", "a")
	fmt.Println("第一条记录：", user)

	//通过map查询
	var selectMap map[string]interface{}
	selectMap = make(map[string]interface{})
	selectMap["ID"] = 4
	selectMap["UserName"] = "a"

	var users []User
	users = selectByMap(selectMap)
	fmt.Println("通过map查询：")
	for _, v := range users {
		fmt.Println(v)
	}

	var usera []User
	usera = selectSql("1")
	fmt.Println("通过sql语句查询：")
	for _, v := range usera {
		fmt.Println(v)
	}
	//fmt.Println(selectByMap(selectMap))
}

func selectByMap(condition map[string]interface{}) []User {
	var users []User
	db.Where(condition).Find(&users)
	return users
}

func selectSql(typr string) []User {
	var users []User
	db.Exec("select * from user").Find(&users)
	return users
}

//删除
func delete() {
	fmt.Println("delete1")

}

func update() {
	fmt.Println("update")

}
func main() {
	initDB()
	selectC()
}
