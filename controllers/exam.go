package controllers

import (
	"flashStudyAPI/models"
	"flashStudyAPI/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExamInput struct {
	Name        string `json:"name" binding:"required,min=3,max=200"`
	Description string `json:"description" binding:"required"`
	ExamDate    string `json:"examDate" binding:"required"`
}

type ExamUpdateInput struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required,min=3,max=200"`
	Description string `json:"description" binding:"required"`
	ExamDate    string `json:"examDate" binding:"required"`
}

func DeleteExam(c *gin.Context) {
	id, err := c.GetQuery("id")

	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'id' param missing!"})
		return
	}
	nid, _ := strconv.Atoi(id)

	errn := models.DeleteExam(uint(nid))

	if errn != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errn.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exam Deleted!"})
}

func UpdateExam(c *gin.Context) {
	var input ExamUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n := models.Exam{}

	n.Name = input.Name
	n.Description = input.Description
	n.ExamDate = input.ExamDate

	err := models.UpdateExam(&n, uint(input.Id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exam Updated!"})
}

func CreateExam(c *gin.Context) {
	uid, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input ExamInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n := models.Exam{}

	n.Name = input.Name
	n.Description = input.Description
	n.ExamDate = input.ExamDate
	n.UserId = int(uid)

	err = models.CreateExam(&n)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Exam Created!"})
}

func GetExam(c *gin.Context) {
	// get user id
	uid, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := models.GetExam(uid)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": note})
}
