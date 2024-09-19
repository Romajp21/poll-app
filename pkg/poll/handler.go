package poll

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPollsHandler(c *gin.Context) {
	polls, err := GetPolls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener encuestas"})
		return
	}
	c.JSON(http.StatusOK, polls)
}

func VoteHandler(c *gin.Context) {
	optionIdStr := c.Param("id")
	optionId, err := strconv.Atoi(optionIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = Vote(optionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al votar"})
		return
	}

	c.Status(http.StatusOK)
}
