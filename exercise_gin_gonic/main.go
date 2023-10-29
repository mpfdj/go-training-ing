package main

import (
	"github.com/gin-gonic/gin"
	"go-training/exercise_gin_gonic/globals"
	"go-training/exercise_gin_gonic/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Amsterdam"
	globals.Database, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func main() {
	router := gin.Default()
	router.GET("/api/books/:bookId", handlers.GetBook)
	router.GET("/api/books/:bookId/reviews", handlers.GetReviews)
	router.POST("/api/books/:bookId/reviews", handlers.PostReview)
	router.Run()
}
