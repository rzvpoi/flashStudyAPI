package controllers

import (
	"flashStudyAPI/models"
	"flashStudyAPI/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LikeGroup(c *gin.Context) {
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

	err = models.LikeGroup(uid, uint(gid))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
