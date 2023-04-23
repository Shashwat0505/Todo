package main

import(
	"servers/models"
	"servers/initializers"
	"fmt"
)
func init(){
	initializers.DBconnection()
}
func main(){
	initializers.DB.AutoMigrate(&models.Activity{})
	fmt.Println("migration done successful!!")
}


