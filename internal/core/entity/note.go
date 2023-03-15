package entity

type Note struct {
	ID          int    `json:"id" example:"123"`
	CustomerID  int    `json:"customer_id" example:"123"`
	Title       string `json:"title" example:"example"`
	Description string `json:"description" example:"example"`
}
