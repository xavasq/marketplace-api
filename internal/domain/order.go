package domain

type Order struct {
	ID        int64
	UserID    int64
	ProductID int64
	Quantity  int
	Price     int64
}
