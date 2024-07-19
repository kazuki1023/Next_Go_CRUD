package model

import "time"

type (
	CreateRequest struct {
		Title      string    `json:"title"`
		Generation string    `json:"generation"`
		StartDate  time.Time `json:"start_date"`
		FinishDate time.Time `json:"finish_date"`
		Detail     string    `json:"detail"`
	}

	UpdateRequest struct {
		Title      string    `json:"title"`
		Generation string    `json:"generation"`
		StartDate  time.Time `json:"start_date"`
		FinishDate time.Time `json:"finish_date"`
		Detail     string    `json:"detail"`
	}

	EventResponse struct {
		ID         int       `json:"id"`
		Title      string    `json:"title"`
		Generation string    `json:"generation"`
		StartDate  time.Time `json:"start_date"`
		FinishDate time.Time `json:"finish_date"`
		Detail     string    `json:"detail"`
	}
)
