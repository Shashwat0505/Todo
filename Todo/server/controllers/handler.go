package controllers

import (
	// "fmt"
	"fmt"
	"html/template"
	"net/http"
	"servers/initializers"
	"servers/models"
	
)

func Handler(w http.ResponseWriter,r *http.Request){
	 var tmplt= template.Must(template.ParseFiles("../template/main.html"))


	tmplt.Execute(w,nil)
}

func Submit(w http.ResponseWriter,r *http.Request){
	
	// fmt.Println(r.ParseForm())
	r.ParseForm()
	var task=r.FormValue("task")
	fmt.Println(task)

	var acti =models.Activity{
		ActivityName: task,
	}

	initializers.DB.Create(&acti)
	

	

}
