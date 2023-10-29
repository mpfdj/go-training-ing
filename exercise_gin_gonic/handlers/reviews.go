package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-training/exercise_gin_gonic/globals"
	"net/http"
	"strconv"
)

type Reviews struct {
	Id      int `gorm:"primaryKey"`
	BookId  int
	Rating  int
	Comment string
}

func GetReviews(c *gin.Context) {
	bookId := c.Param("bookId")
	id, _ := strconv.Atoi(bookId)

	db := globals.Database
	db.AutoMigrate(&Reviews{})

	var reviews []Reviews
	result := db.Where("book_id = ?", id).Find(&reviews)

	fmt.Printf("%#v\n", result)
	fmt.Printf("%#v\n", reviews)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
