package entity

type Note struct {
	ID          int    `json:"id" example:"1"`
	Title       string `json:"title" example:"Example text"`
	Description string `json:"description" example:"Example description"`
}
