package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/hello",func(contex *gin.Context) {
		contex.String(200, "hello ankit")
	})
	if err:=router.Run(":8080");err!=nil{
		fmt.Println("Failed to start server ",err)
	}
}