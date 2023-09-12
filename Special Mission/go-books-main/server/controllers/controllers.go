package controllers

import (
	"go-structure-project/models"
	"go-structure-project/services"
	"net/http"
	"strconv"

	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context, db *sql.DB) {
	books, err := services.GetBooksFromDB(db)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id") // Get the status from the URL parameter

	tasks, err := services.GetBookByIDFromDB(db, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func UpdateBooks(c *gin.Context, db *sql.DB) {
	bookIDStr := c.Param("id")

	var editedBook models.Book
	if err := c.ShouldBindJSON(&editedBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	editedBook, err := services.UpdateBooksInDB(db, bookIDStr, editedBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	editedBook.ID = bookID

	c.IndentedJSON(http.StatusOK, editedBook)
}

func CreateBook(c *gin.Context, db *sql.DB) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateBookInDB(db, newBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newBook)
}

func DeleteBook(c *gin.Context, db *sql.DB) {
	taskID := c.Param("id") // Get the task ID from the URL parameter
	// Call the service to delete the task by its ID
	if err := services.DeleteBookByID(db, taskID); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil) // Respond with HTTP 204 (No Content)
}
