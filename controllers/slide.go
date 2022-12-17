package controllers

import (
	"flashStudyAPI/models"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:parameter
type SlideInput struct {
	Question string               `form:"question" binding:"required,min=3,max=500"`
	Answer   string               `form:"answer" binding:"required,max=500"`
	Tags     string               `form:"tags"`
	GroupId  string               `form:"groupId" binding:"required"`
	Image    multipart.FileHeader `form:"image" swaggerignore:"true"`
}

// swagger:parameter
type SlideUpdateInput struct {
	Id       string               `form:"id" binding:"required"`
	Question string               `form:"question" binding:"required,min=3,max=500"`
	Answer   string               `form:"answer" binding:"required,max=500"`
	Tags     string               `form:"tags"`
	Image    multipart.FileHeader `form:"image" swaggerignore:"true"`
}

// @Summary Update data of a slide
// @Description  !!! Insert all the values even if they are not new
// @Tags Slide
// @Param slide formData SlideUpdateInput true "Slide Create JSON"
// @Param image formData file true "the picture file"
// @Router /slide/update [put]
func UpdateSlide(c *gin.Context) {

	var input SlideUpdateInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Image.Size > 5000000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File Size over 5MB!"})
		return
	}

	s := models.Slide{}

	sid, _ := strconv.ParseUint(input.Id, 10, 0)

	if len(input.Image.Filename) < 1 {
		s.Image = ""
	} else {
		s.Image = input.Image.Filename
	}

	s.Question = input.Question
	s.Answer = input.Answer
	s.Tags = input.Tags

	filename, err := models.UpdateSlide(uint(sid), &s)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if len(s.Image) > 0 {
		if err := c.SaveUploadedFile(&input.Image, "./public/images-slide/"+filename); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Slide updated!"})
}

// @Summary Delete a slide
// @Param id query string true "Slide Delete Query"
// @Tags Slide
// @Router /slide/delete   [delete]
func DeleteSlide(c *gin.Context) {
	id, err := c.GetQuery("id")

	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'id' param missing!"})
		return
	}
	sid, _ := strconv.Atoi(id)

	s := models.Slide{}
	s.ID = uint(sid)

	response, errs := models.DeleteSlide(s)

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": response})

}

// @Summary Get slide for user
// @Tags Slide
// @Router /slide [get]
func GetSlide(c *gin.Context) {
	id, err := c.GetQuery("gid")

	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'gid' param missing!"})
		return
	}
	gid, _ := strconv.Atoi(id)

	s, errs := models.GetSlide(gid)

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": s})

}

// @Summary Create a slide
// @Tags Slide
// @Param slide formData SlideInput true "Slide Create JSON"
// @Param image formData file true "the picture file"
// @Router /slide/create [post]
func CreateSlide(c *gin.Context) {
	var input SlideInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Image.Size > 5000000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File Size over 5MB!"})
		return
	}

	s := models.Slide{}

	if len(input.Image.Filename) < 1 {
		s.Image = ""
	} else {
		s.Image = input.Image.Filename
	}

	s.Image = input.Image.Filename
	s.Question = input.Question
	s.Answer = input.Answer
	s.Tags = input.Tags
	s.GroupId, _ = strconv.Atoi(input.GroupId)

	filename, err := models.CreateSlide(&s)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(s.Image) > 0 {
		if err := c.SaveUploadedFile(&input.Image, "./public/images-slide/"+filename); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Unable to save the file"})
			models.DeleteSlide(s)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Slide Created!"})

}
