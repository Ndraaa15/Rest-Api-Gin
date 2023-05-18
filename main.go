package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//Connect to database
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection error")
	}
	//Untuk metode lainnya bisa dilihat di dokumentasinya.
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	bookRequest := book.BooksRequest{
		Title: "$100 Start Up",
		Price: 95000,
	}

	bookService.Create(bookRequest)

	router := gin.Default()
	//API versioning digunakan apabila kita menambahkan fitur baru pada apps atau website sehingga agar tidak mengubah semua url maka cukup dengan menggunakan
	//API versioning.
	v1 := router.Group("/v1")
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	//Untuk mengambil data query dengan path dibawah maka url harus berbentuk ?=".."
	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBooksHandler)
	//Ingin mengganti localhost
	router.Run(":8888")

}
