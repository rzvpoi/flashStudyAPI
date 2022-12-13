package controllers

import (
	"flashStudyAPI/models"
	"flashStudyAPI/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NoteInput struct {
	Title   string `json:"title" binding:"required,min=3,max=30"`
	Text    string `json:"text" binding:"required"`
	GroupId int    `json:"groupId" binding:"required"`
}

type NoteUpdateInput struct {
	Id    int    `json:"id" binding:"required"`
	Title string `json:"title" binding:"required,min=3,max=30"`
	Text  string `json:"text" binding:"required"`
}

func DeleteNote(c *gin.Context) {
	id, err := c.GetQuery("id")

	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'id' param missing!"})
		return
	}
	nid, _ := strconv.Atoi(id)

	errn := models.DeleteNote(uint(nid))

	if errn != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errn.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note Deleted!"})
}

func UpdateNote(c *gin.Context) {
	var input NoteUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n := models.Note{}

	n.Text = input.Text
	n.Title = input.Title

	err := models.UpdateNote(&n, uint(input.Id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note Updated!"})
}

func CreateNote(c *gin.Context) {
	uid, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input NoteInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n := models.Note{}

	n.Title = input.Title
	n.Text = input.Text
	n.GroupId = input.GroupId
	n.UserId = int(uid)

	err = models.CreateNote(&n)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note Created!"})

}

func GetNote(c *gin.Context) {
	// get group id
	id, errg := c.GetQuery("id")
	if !errg {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'id' param missing!"})
		return
	}
	gid, _ := strconv.Atoi(id)

	uid, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := models.GetNote(uid, uint(gid))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": note})

}
