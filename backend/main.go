package main

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	router.GET("/problems", handlers.GetProblems)

	router.Run(":8500")
}