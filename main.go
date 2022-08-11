package main

import (
	"httpserver/book"
	"httpserver/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=Golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}
	db.AutoMigrate(&book.Book{})
	// bookRepository := book.NewRepository(db)

	// books, err := bookRepository.FindAll()
	// book, err := bookRepository.FindByID(1)

	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }
	// fmt.Println("Title :", book.Title)

	book := book.Book{
		Title:       "test",
		Description: "tes aj",
		Price:       10000,
		rating:      4,
		Discount:    0,
	}

	//Create data

	// book := book.Book{}
	// book.Title = "Man cap"
	// book.Price = 200000
	// book.Discount = 10
	// book.Rating = 5
	// book.Description = "ngetes aja"
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error creating book")
	// }

	// var books []book.Book

	// GET DATA
	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error creating book")
	// }
	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	// GET DATA WITH ID

	// var book book.Book
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println((err))
	// }

	// UPDATE WITH ID
	// book.Title = "Man ula (revised Edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// DELETE WITH ID
	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println(err)
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)
	router.Run()
}
