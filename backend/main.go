package main

import (
	"backend/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

		
		router.GET("/problems", func(c *gin.Context) {
			problems, err := data.LoadProblems()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read problems"})
				return
			}
			c.JSON(http.StatusOK, problems)
		})
	

	router.Run(":8500")
}
