// backend/handlers/problems.go
package handlers

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"backend/data"
)

func GetProblems(c *gin.Context) {
	problems, err := data.LoadProblems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read problems"})
		return
	}
	c.JSON(http.StatusOK, problems)
}

func RunCode(c *gin.Context) {
	// Optional: If you want to accept user code via request body
	// code := c.PostForm("code")
	// Save code into problems/user_test.go if needed

	// Run `go test` in the problems folder
	cmd := exec.Command("go", "test", "./problems/...")
	out, err := cmd.CombinedOutput()

	if err != nil {
		// Return both error + output (Go puts failures here)
		c.String(http.StatusOK, string(out))
		return
	}

	// Return successful test output
	c.String(http.StatusOK, string(out))
}
