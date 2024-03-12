package main

import (
  	"github.com/gin-gonic/gin"
	"libraryBackend/controllers"
	"libraryBackend/models"
	"libraryBackend/middlewares"
)

func main() {

	models.ConnectDataBase()
	
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login",controllers.Login)

	// Reader := r.Group("api/reader")
	// Reader.GET("/searchbook",controllers.Searchbook)
	// Reader.POST("/issuerequest",controllers.Issuerequest)

	Admin := r.Group("/api/admin")
	Admin.Use(middlewares.JwtAuthMiddleware())
	Admin.GET("/user",controllers.CurrentUser)
	Admin.POST("/bookpost",controllers.Bookpost)
	// Admin.PUT("/bookupdate",controllers.Bookupdate)
	// Admin.DELETE("/bookdelete",controllers.Bookdelete)
	// Admin.GET("/booklist",controllers.Booklist)
	// Admin.POST("issuebook",controllers.Issuebook)
	// Admin.PUT("/updateissuebook",controllers.Updateissuebook)

	// Owner := r.Group("/api/Owner")
	// Owner.POST("/createlibrary",controllers.Library)

	r.Run(":8088")

}