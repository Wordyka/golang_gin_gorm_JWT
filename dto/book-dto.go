package dto


// Berisi informasi mengenai data apa saja yang akan dikirim antar client server

//struct BookUpdateDTO untuk digunakan pada pengubahan data buku
type BookUpdateDTO struct {
	ID          uint64 `json:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id, omitempty"`
}

// structc BookCreateDTO untuk digunakan pada pembuatan data buku
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}


