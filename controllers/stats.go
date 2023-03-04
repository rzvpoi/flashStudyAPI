package controllers

import (
	"flashStudyAPI/models"
	"flashStudyAPI/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:model
type StatsInput struct {
	CorrectAnswer int `json:"correctAnswer" binding:"required"`
	WrongAnswer   int `json:"wrongAnswer" binding:"required"`
	GroupId       int `json:"groupId" binding:"required"`
}

// @Summary Get stats for user
// @Tags Stats
// @Router /stats [get]
func GetStats(c *gin.Context) {
	// get group id
	id, errg := c.GetQuery("id")
	if !errg {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'id' param missing!"})
		return
	}
	gid, _ := strconv.Atoi(id)

	// get user id
	uid, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stats, err := models.GetStats(uid, uint(gid))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": stats})
}

// @Summary Create a stats
// @Tags Stats
// @Param stats body StatsInput true "Stats Create JSON"
// @Router /stats/create [post]
func CreateStats(c *gin.Context) {
	uid, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input StatsInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n := models.Stats{}

	n.CorrectAnswer = input.CorrectAnswer
	n.WrongAnswer = input.WrongAnswer
	n.GroupId = input.GroupId
	n.Grade = float64(n.CorrectAnswer) / (float64(n.CorrectAnswer) + float64(n.WrongAnswer)) * 100
	n.UserId = int(uid)

	err = models.CreateStats(&n)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stat created!"})
}
