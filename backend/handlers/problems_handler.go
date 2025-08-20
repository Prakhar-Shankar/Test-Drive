// backend/handlers/problems.go
package handlers

import (
	"net/http"

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
