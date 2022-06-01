package product

import "time"

type ProductData struct {
	ProductID   string    `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
}
