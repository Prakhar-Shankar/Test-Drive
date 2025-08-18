package handlers

import (
	"net/http"
	"backend/services"
	"github.com/gin-gonic/gin"
)

func GetProblems(c *gin.Context) {
	problems, err := services.LoadProblems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, problems)
}
