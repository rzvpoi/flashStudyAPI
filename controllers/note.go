package controllers

import (
	"flashStudyAPI/models"
	"flashStudyAPI/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:model
type NoteInput struct {
	Title   string `json:"title" binding:"required,min=3,max=30"`
	Text    string `json:"text" binding:"required"`
	GroupId int    `json:"groupId" binding:"required"`
}

// swagger:model
type NoteUpdateInput struct {
	Id    int    `json:"id" binding:"required"`
	Title string `json:"title" binding:"required,min=3,max=30"`
	Text  string `json:"text" binding:"required"`
}

// @Summary Delete a note
// @Param id query string true "Note Delete Query"
// @Tags Note
// @Router /note/delete   [delete]
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

// @Summary Update data of a note
// @Description  !!! Insert all the values even if they are not new
// @Tags Note
// @Param note body NoteUpdateInput true "Note Update JSON"
// @Router /note/update [put]
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

// @Summary Create a note
// @Description Create New Note
// @Tags Note
// @Param note body NoteInput true "Note Create JSON"
// @Router /note/create [post]
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

// @Summary Get note for user
// @Tags Note
// @Router /note [get]
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
