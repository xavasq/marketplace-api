package dto

type ProductResponseDTO struct {
	ID       int64  `json:"idd"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}
