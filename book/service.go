package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(bookRequest BooksRequest) (Book, error)
	Update(ID int, bookRequest BooksRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BooksRequest) (Book, error) {
	// price, err := bookRequest.Price.Int64()
	// discount, err := bookRequest.Discount.Int64()
	// rating, err := bookRequest.Rating.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		Discount:    bookRequest.Discount,
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BooksRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	book.Title = bookRequest.Title
	book.Price = bookRequest.Price
	book.Description = bookRequest.Description
	book.Rating = bookRequest.Rating
	book.Discount = bookRequest.Discount

	newBook, err := s.repository.Update(book)
	return newBook, err
}
