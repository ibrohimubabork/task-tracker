package models

type Response[T any] struct {
	Status string          `json:"status"`         // Contoh: "success"
	Data   T               `json:"data,omitempty"` // Data utama (bisa object atau slice)
	Meta   *Meta           `json:"meta,omitempty"` // Opsional: Untuk pagination
	Errors []ErrorResponse `json:"errors,omitempty"`
}

// Meta digunakan jika data yang dikembalikan berupa list/pagination
type Meta struct {
	CurrentPage int   `json:"current_page" example:"1"`
	TotalPage   int   `json:"total_page" example:"10"`
	PerPage     int   `json:"per_page" example:"20"`
	TotalData   int64 `json:"total_data" example:"195"`
}

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
