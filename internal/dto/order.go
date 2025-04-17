package dto

type OrderRequestDTO struct {
	ProductID int64 `json:"product_id" validate:"required,gt=0"`
	Quantity  int   `json:"quantity" validate:"required,gt=0"`
}

type OrderResponseDTO struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
	Price     int64 `json:"price"`
}
