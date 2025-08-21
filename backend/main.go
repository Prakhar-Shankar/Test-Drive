package main

import (
	"backend/data"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
		
		router.GET("/problems", func(c *gin.Context) {
			problems, err := data.LoadProblems()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read problems"})
				return
			}
			c.JSON(http.StatusOK, problems)
		})
		router.POST("/run", func(c *gin.Context) {
			var body struct {
				Tests string `json:"tests"`
			}
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
				return
			}
		
			// Save test file (always overwrite for now)
			testFile := "../problems/user_test.go"
			if err := os.WriteFile(testFile, []byte(body.Tests), 0644); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "could not write test file"})
				return
			}
		
			// Run `go test`
			cmd := exec.Command("go", "test", "../problems/...")
			out, err := cmd.CombinedOutput()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error":  err.Error(),
					"output": string(out),
				})
				return
			}
		
			c.JSON(http.StatusOK, gin.H{"output": string(out)})
		})
	

	router.Run(":8500")
}
