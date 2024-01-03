package dto

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
	Timestamp
}

type MovieCreatePayload struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float32 `json:"rating" validate:"required,numeric,min=0,max=1"`
	Image       string  `json:"image" validate:"required"`
}

type MovieCreateParam struct {
	Payload MovieCreatePayload
}

type MovieDetailParam struct {
	ID int `uri:"id" binding:"required"`
}

type FilterMovieParams struct {
	Title string
}

type MovieUpdatePayload struct {
	ID          int
	Title       string   `json:"title" validate:"omitempty"`
	Description string   `json:"description" validate:"omitempty"`
	Rating      *float32 `json:"rating" validate:"omitempty,numeric,min=0,max=1"`
	Image       string   `json:"image" validate:"omitempty"`
}

type MovieUpdateParam struct {
	Payload MovieUpdatePayload
}
