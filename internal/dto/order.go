package dto

type OrderRequestDTO struct {
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

type OrderResponseDTO struct {
	ID        int64 `json:"ID"`
	UserID    int64 `json:"UserID"`
	ProductID int64 `json:"ProductID"`
	Quantity  int   `json:"quantity"`
	Price     int64 `json:"Price"`
}
