package controllers

import (
	"flashStudyAPI/models"
	"flashStudyAPI/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupInput struct {
	Name        string `json:"name" binding:"required,min=3,max=50"`
	Description string `json:"description" binding:"min=0,max=500"`
	IsPublic    *bool  `json:"isPublic" binding:"required"`
	Color       string `json:"color"`
}

type GroupUpdateInput struct {
	Id          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required,min=3,max=50"`
	Description string `json:"description" binding:"min=0,max=500"`
	IsPublic    *bool  `json:"isPublic" binding:"required"`
	Color       string `json:"color"`
}

// @Summary Get search results
// @Param value query string true "word"
// @Router /search [get]
func Search(c *gin.Context) {
	query, err := c.GetQuery("value")
	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Param 'value' missing!"})
		return
	}

	data, errs := models.Search(query)

	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

// @Summary Get popular groups
// @Param count query string true "count"
// @Router /popularGroups  [get]
func PopularGroups(c *gin.Context) {
	query, err := c.GetQuery("count")
	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Param 'count' missing!"})
		return
	}

	count, _ := strconv.Atoi(query)
	data, errs := models.PopularGroups(count)

	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

// @Summary Delete a group
// @Param id query string true "Group Id"
// @Router /group/delete   [get]
func DeleteGroup(c *gin.Context) {
	id, err := c.GetQuery("id")
	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Param 'id' missing!"})
		return
	}

	gid, _ := strconv.Atoi(id)

	response, errs := models.DeleteGroup(uint(gid))

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": response})

}

// @Summary Update data of a group
// @Description Insert all the values even if they are not new
// @Param body json {object} true "Group Id"
// @Router /group/delete   [get]
func UpdateGroup(c *gin.Context) {

	var input GroupUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g := models.Group{}

	g.Name = input.Name
	g.Description = input.Description
	g.IsPublic = *input.IsPublic
	g.Color = input.Color

	response, err := models.UpdateGroup(input.Id, &g)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": response})
}

func CreateGroup(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input GroupInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g := models.Group{}

	g.Name = input.Name
	g.Description = input.Description
	g.IsPublic = *input.IsPublic
	g.Color = input.Color
	g.UserId = int(user_id)
	g.IsLiked = false

	_, errs := models.SaveGroup(&g)

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "group created"})
}

func GetGroups(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g, err := models.GetGroups(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := 0; i < len(g); i++ {
		err = models.IsLiked(user_id, g[i].ID)
		if err == nil {
			g[i].IsLiked = true
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": g})
}
