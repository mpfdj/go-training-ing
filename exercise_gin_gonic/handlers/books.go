package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-training/exercise_gin_gonic/globals"
	"net/http"
	"strconv"
)

type Books struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Isbn      string
	Language  string
	Publisher string
	Num_pages int
}

func GetBook(c *gin.Context) {
	bookId := c.Param("bookId")
	id, _ := strconv.Atoi(bookId)

	db := globals.Database
	db.AutoMigrate(&Books{})

	var books Books
	result := db.First(&books, "id = ?", id)

	fmt.Printf("%#v\n", result)
	fmt.Printf("%#v\n", books)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, books)
}
