package services

import (
	"database/sql"
	"go-structure-project/models"
)

// GetBooksFromDB retrieves a list of books from the database.
func GetBooksFromDB(db *sql.DB) ([]models.Book, error) {
	var books []models.Book
	query := "SELECT id, title, author, quantity FROM books"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

// GetBookByID retrieves a book by its ID from the database.
func GetBookByIDFromDB(db *sql.DB, id string) (*models.Book, error) {
	var book models.Book
	query := "SELECT id, title, author, quantity FROM books WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func UpdateBooksInDB(db *sql.DB, taskID string, editedBook models.Book) (models.Book, error) {
	query := "UPDATE books SET title = ?, author = ?, quantity = ? WHERE id = ?"
	_, err := db.Exec(query, editedBook.Title, editedBook.Author, editedBook.Quantity, taskID)
	if err != nil {
		return models.Book{}, err
	}

	return editedBook, nil
}

// CreateBookInDB creates a new book in the database.
func CreateBookInDB(db *sql.DB, newBook models.Book) error {
	query := "INSERT INTO books (id, title, author, quantity) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, newBook.ID, newBook.Title, newBook.Author, newBook.Quantity)
	return err
}

func DeleteBookByID(db *sql.DB, taskID string) error {
	_, err := db.Exec("DELETE FROM books WHERE id = ?", taskID)
	if err != nil {
		return err
	}

	return nil
}
