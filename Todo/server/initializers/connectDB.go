package initializers

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

func DBconnection() {

	dsn := "root:root@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err!=nil{
		log.Fatal(err)
	}else{
		fmt.Println("Database connected successful!!")
	}
	DB=db
}