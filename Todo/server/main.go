package main

import (
	"fmt"
	"servers/initializers"
	"servers/router"
)

func init(){
	initializers.DBconnection()
}
func main(){
	fmt.Println("inside main")
router.Run()

}
