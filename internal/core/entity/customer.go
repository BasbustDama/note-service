package entity

type Customer struct {
	ID       int    `json:"id" example:"123"`
	Name     string `json:"name" example:"BasbustDama"`
	Email    string `json:"email" example:"example@gmail.com"`
	Password string `json:"password" example:"qwerty123"`
}
