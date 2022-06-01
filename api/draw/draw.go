package draw

import (
	customer "go-api-auction/api/customer"
	product "go-api-auction/api/product"
	"time"
)

type DrawData struct {
	DrawID            string                  `json:"draw_id"`
	Product           product.ProductData     `json:"product"`
	Participants      []customer.CustomerData `json:"participants"`
	ParticipantsCount int                     `json:"participant_count"`
	CreateDate        time.Time               `json:"create_date"`
	EndDate           time.Time               `json:"end_date"`
	Status            string                  `json:"status"`
}
