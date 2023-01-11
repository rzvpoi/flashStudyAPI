package main

import (
	"flashStudyAPI/controllers"
	"flashStudyAPI/middlewares"
	"flashStudyAPI/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "flashStudyAPI/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title FlashStudy API
// @version         1.0
// @description     FlashStudy API in Go using Gin framework.

// @contact.name   Tudor Poienariu
// @contact.url    https://linkedin.com/in/tudor-poienariu-635a48232
// @contact.email  razvanpoienariu@gmail.com

// @host      localhost:8081
// @BasePath  /api
func main() {

	models.ConnectDataBase()

	r := gin.Default()

	r.Static("/image", "./public/images-slide")
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "PUT", "PATH", "DELETE", "UPDATE", "GET"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.POST("/passwordreset", controllers.PasswordReset)
	public.GET("/popularGroups", controllers.PopularGroups)
	public.GET("/search", controllers.Search)

	protected := r.Group("api/user")

	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/", controllers.CurrentUser)
	protected.PUT("/update", controllers.UpdateUser)

	protectedGroup := r.Group("api/group")
	protectedGroup.Use(middlewares.JwtAuthMiddleware())
	protectedGroup.GET("/", controllers.GetGroups)
	protectedGroup.POST("/create", controllers.CreateGroup)
	protectedGroup.PUT("/update", controllers.UpdateGroup)
	protectedGroup.DELETE("/delete", controllers.DeleteGroup)
	protectedGroup.POST("/like", controllers.LikeGroup)

	protectedSlide := r.Group("api/slide")

	protectedSlide.Use(middlewares.JwtAuthMiddleware())
	protectedSlide.GET("/", controllers.GetSlide)
	protectedSlide.POST("/create", controllers.CreateSlide)
	protectedSlide.PUT("/update", controllers.UpdateSlide)
	protectedSlide.DELETE("/delete", controllers.DeleteSlide)

	protectedNote := r.Group("api/note")

	protectedNote.Use(middlewares.JwtAuthMiddleware())
	protectedNote.GET("/", controllers.GetNote)
	protectedNote.POST("/create", controllers.CreateNote)
	protectedNote.PUT("/update", controllers.UpdateNote)
	protectedNote.DELETE("/delete", controllers.DeleteNote)

	protectedExam := r.Group("api/exam")

	protectedExam.Use(middlewares.JwtAuthMiddleware())
	protectedExam.GET("/", controllers.GetExam)
	protectedExam.POST("/create", controllers.CreateExam)
	protectedExam.PUT("/update", controllers.UpdateExam)
	protectedExam.DELETE("/delete", controllers.DeleteExam)

	protectedStats := r.Group("api/stats")

	protectedStats.Use(middlewares.JwtAuthMiddleware())
	protectedStats.GET("/", controllers.GetStats)
	protectedStats.GET("/create", controllers.CreateStats)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8081")

}
