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

type ReviewJson struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
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

func PostReview(c *gin.Context) {
	bookId := c.Param("bookId")
	id, _ := strconv.Atoi(bookId)

	db := globals.Database
	db.AutoMigrate(&Books{})

	var books Books
	result := db.First(&books, "id = ?", id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	}

	var reviewJson ReviewJson
	if err := c.BindJSON(&reviewJson); err != nil {
		return
	}
	fmt.Printf("%#v\n", reviewJson)

	review := Reviews{
		BookId:  id,
		Rating:  reviewJson.Rating,
		Comment: reviewJson.Comment,
	}

	db.AutoMigrate(&Reviews{})
	result = db.Create(&review)
	fmt.Printf("%#v\n", result)

	var lastReview []Reviews
	result = db.Last(&lastReview)

	fmt.Printf("%#v\n", result)
	fmt.Printf("%#v\n", lastReview)

	// Return lastReview[0]

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
