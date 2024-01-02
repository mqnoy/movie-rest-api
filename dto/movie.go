package dto

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
	Timestamp
}
