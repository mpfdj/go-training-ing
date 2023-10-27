package tests

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

//CREATE TABLE books (
//	id SERIAL PRIMARY KEY,
//	title VARCHAR(255) NOT NULL,
//	isbn VARCHAR(255) NOT NULL,
//	language VARCHAR(255) NOT NULL,
//	publisher VARCHAR(255) NOT NULL,
//	num_pages INTEGER NOT NULL
//);

type Books struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Isbn      string
	Language  string
	Publisher string
	Num_pages int
}

func TestWithPostgres(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Amsterdam"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Books{})

	var books Books
	db.First(&books, "id = ?", -1)

	if books.Id == 0 &&
		books.Title == "" &&
		books.Isbn == "" &&
		books.Language == "" &&
		books.Publisher == "" &&
		books.Num_pages == 0 {
		fmt.Println("Returning a 404")

	}

	fmt.Printf("%#v\n", books)

}
