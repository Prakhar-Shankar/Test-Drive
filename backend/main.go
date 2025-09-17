package main

import (
	"backend/data"
	"context"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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
		type req struct {
			Tests string `json:"tests"`
		}
		var body req
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"output": "invalid request"})
			return
		}

		// Write tests into problems/user_test.go
		problemsDir, _ := filepath.Abs(filepath.Join("..", "problems"))
		if err := os.WriteFile(filepath.Join(problemsDir, "user_test.go"), []byte(body.Tests), 0644); err != nil {
			c.JSON(http.StatusOK, gin.H{"output": "write failed: " + err.Error()})
			return
		}

		// Run exactly what you do in the terminal: cd problems && go test
		ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, "go", "test") // exactly: cd problems && go test
		cmd.Dir = problemsDir
		// Optional: force module mode and clean env
		cmd.Env = append(os.Environ(), "GO111MODULE=on")

		out, err := cmd.CombinedOutput()
		// Do NOT append err.Error() to the output; it would add "exit status N"
		// Use err only to compute exit code.
		exit := 0
		if err != nil {
			if ee, ok := err.(*exec.ExitError); ok {
				exit = ee.ExitCode()
			} else if ctx.Err() == context.DeadlineExceeded {
				exit = 124
			} else {
				exit = -1
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"ok":       err == nil,
			"exitCode": exit,
			"dir":      cmd.Dir,     // debug
			"args":     cmd.Args,    // debug
			"output":   string(out), // exact terminal output of `go test` in problems
		})
	})

	router.Run(":8500")
}
