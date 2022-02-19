package structs

import "time"

// Order represent the model for an order
type Order struct {
	OrderID      int       `json:"orderId" example:"1"`
	CostumerName string    `json:"costumerName" example:"Arif Kurniawa"`
	OrderAt      time.Time `json:"orderAt" example:"2022-02-17T21:16:38.147532+07:00"`
	Items        []Item    `json:"items"`
}

// Item represent the model for an item in the order
type Item struct {
	Item_id     int    `json:"itemId" example:"1"`
	Item_code   string `json:"itemCode" example:"AAA321"`
	Description string `json:"description" example:"A random description"`
	Quantity    int    `json:"quantity" example:"10"`
	Order_id    int    `json:"orderId" example:"1"`
}
