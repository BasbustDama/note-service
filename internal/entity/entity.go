package entity

type Customer struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"BasbustDama"`
	Password string `json:"password" example:"qwerty123"`
}

type Note struct {
	ID          int    `json:"id" example:"1"`
	CustomerID  int    `json:"customer_id" example:"1"`
	Title       string `json:"title" example:"Example text"`
	Description string `json:"description" example:"Example description"`
}
