package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Indra Brata",
		"status":  "college student",
		"college": "Brawijaya University",
	})
}

func (h *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"bio":         "Learning API",
		"exploration": "Perjalanan untuk belajar restAPI untuk menjadi bagian dari BCC",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	//Mengambil id params
	//Bisa mengambil lebih dari 1 params
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	//Mengambil title query
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

// Membuat endpoint
func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	booResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": booResponse,
	})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BooksRequest

	//Saat ngebind dengan json jangan lupa untuk menambahkan pointer
	err := c.ShouldBindJSON(&bookRequest)

	//Cara membuat error handling
	if err != nil {
		errorMessages := []string{}
		//Digunakan untuk error handling serta menampilkan lebih dari satu error
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
		// "sub_title": bookInput.Subtitle,
	})
}

func (h *bookHandler) UpdateBooksHandler(c *gin.Context) {
	var bookRequest book.BooksRequest

	//Saat ngebind dengan json jangan lupa untuk menambahkan pointer
	err := c.ShouldBindJSON(&bookRequest)

	//Cara membuat error handling
	if err != nil {
		errorMessages := []string{}
		//Digunakan untuk error handling serta menampilkan lebih dari satu error
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
		// "sub_title": bookInput.Subtitle,
	})
}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		Title:       b.Title,
		ID:          b.ID,
		Price:       b.Price,
		Description: b.Description,
		Discount:    b.Discount,
		Rating:      b.Rating,
	}
}
