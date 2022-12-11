package main

import (
	"flashStudyAPI/controllers"
	"flashStudyAPI/middlewares"
	"flashStudyAPI/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.POST("/passwordreset", controllers.PasswordReset)

	protected := r.Group("api/user")

	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/", controllers.CurrentUser)
	protected.POST("/update", controllers.UpdateUser)

	protectedGroup := r.Group("api/group")
	protectedGroup.Use(middlewares.JwtAuthMiddleware())
	protectedGroup.GET("/", controllers.GetGroups)
	protectedGroup.POST("/create", controllers.CreateGroup)
	protectedGroup.POST("/update", controllers.UpdateGroup)
	protectedGroup.POST("/delete", controllers.DeleteGroup)

	r.Run(":8080")

}
