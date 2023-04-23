package router

import (
	// "log"
	// "net"
	"net/http"
	"servers/controllers"

	// "fmt"

	"github.com/gorilla/mux"
)
var r =mux.NewRouter()
func init() {
	
	r.HandleFunc("/home",controllers.Handler).Methods("GET")
	r.HandleFunc("/demo",controllers.Submit).Methods("POST")
	
}
func Run(){
	http.ListenAndServe(":5000",r)
	
}