package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
}

type repository struct {
	//Agar bisa melakukan akses ke dalam database.
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Debug().Find(&books).Error

	return books, err
}

func (r *repository) FindByID(ID int) (Book, error) {
	var book Book
	// err := r.db.Find(&book, ID).Error
	err := r.db.Debug().Where("id = ?", ID).First(&book).Error
	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	r.db.Save(book)
	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(book).Error
	return book, err
}
