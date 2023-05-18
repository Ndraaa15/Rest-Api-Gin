package book

type BooksRequest struct {
	ID    int    `json:"ID" binding:"required"`
	Title string `json:"title" binding:"required"`
	//Fungsi json.Number adalah membuat semua data yang di ambil dalam price apabila mengandung angka maka akan diambil
	Price int `json:"price" binding:"required,number"`
	// Subtitle string `json:"sub_title"` //Subtitle  berfungsi untuk mengambil data json yang bernama sub_title
	Description string `json:"description" binding:"required,number"`
	Rating      int    `json:"rating" binding:"required,number"`
	Discount    int    `json:"discount" binding:"required,number"`
}
