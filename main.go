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

	r.Static("/image", "./public/images-slide")

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

	protectedSlide := r.Group("api/slide")
	protectedSlide.Use(middlewares.JwtAuthMiddleware())
	protectedSlide.GET("/", controllers.GetSlide)
	protectedSlide.POST("/create", controllers.CreateSlide)
	protectedSlide.POST("/update", controllers.UpdateSlide)
	protectedSlide.POST("/delete", controllers.DeleteSlide)

	protectedNote := r.Group("api/note")
	protectedNote.Use(middlewares.JwtAuthMiddleware())
	protectedNote.GET("/", controllers.GetNote)
	protectedNote.POST("/create", controllers.CreateNote)
	protectedNote.POST("/update", controllers.UpdateNote)
	protectedNote.POST("/delete", controllers.DeleteNote)

	protectedExam := r.Group("api/exam")
	protectedExam.Use(middlewares.JwtAuthMiddleware())
	protectedExam.GET("/", controllers.GetExam)
	protectedExam.POST("/create", controllers.CreateExam)
	protectedExam.POST("/update", controllers.UpdateExam)
	protectedExam.POST("/delete", controllers.DeleteExam)

	protectedStats := r.Group("api/stats")
	protectedStats.Use(middlewares.JwtAuthMiddleware())
	protectedStats.GET("/", controllers.GetStats)
	protectedStats.GET("/create", controllers.CreateStats)

	r.Run(":8080")

}
