package routes

import (
	"database/sql"
	"go-structure-project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Use anonymous functions to adapt controller functions to gin.HandlerFunc
	router.GET("/books", func(c *gin.Context) {
		controllers.GetBooks(c, db)
	})
	router.GET("/books/:id", func(c *gin.Context) {
		controllers.GetBookByID(c, db)
	})
	router.PUT("/books/:id", func(c *gin.Context) {
		controllers.UpdateBooks(c, db)
	})
	router.POST("/books", func(c *gin.Context) {
		controllers.CreateBook(c, db)
	})
	router.DELETE("/books/:id", func(c *gin.Context) {
		controllers.DeleteBook(c, db)
	})
}
