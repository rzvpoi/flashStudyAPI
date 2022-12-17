package controllers

import (
	"flashStudyAPI/models"
	"flashStudyAPI/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// swager:model
type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required"`
}

// swager:model
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// swager:model
type PasswordResetInput struct {
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required"`
	PasswordRepeat string `json:"passwordrepeat" binding:"required"`
}

// swager:model
type UserInput struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
}

// @Summary Update data of an user
// @Description  !!! Insert all the values even if they are not new
// @Tags User
// @Param user body UserInput true "User Update JSON"
// @Router /api/user/update [put]
func UpdateUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)

	var input UserInput

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := models.UpdateUser(user_id, input.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": response})

}

// @Summary Reset password
// @Tags User
// @Param password body PasswordResetInput true "Password Reset JSON"
// @Router /api/passwordreset [post]
func PasswordReset(c *gin.Context) {

	var input PasswordResetInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.PasswordRepeat != input.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords don't match"})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

	response, err := models.ResetPassword(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": response})

}

// @Summary Get user data
// @Tags User
// @Router /api/user [get]
func CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

// @Summary User Login
// @Tags User
// @Param login body LoginInput true "Login User JSON"
// @Router /api/login [post]
func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

	token, err := models.LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

// @Summary User Register
// @Tags User
// @Param register body RegisterInput true "Register User JSON"
// @Router /api/register [post]
func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
