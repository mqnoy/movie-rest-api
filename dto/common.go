package dto

import (
	"github.com/mqnoy/movie-rest-api/model"
	"github.com/mqnoy/movie-rest-api/util"
)

type Timestamp struct {
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func ParseTimestampResponse(m model.TimestampColumn) Timestamp {
	return Timestamp{
		CreatedAt: util.DateToEpoch(m.CreatedAt),
		UpdatedAt: util.DateToEpoch(m.UpdatedAt),
	}
}

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
	TotalItems int64 `json:"total_items"`
	Offset     int   `json:",omitempty"`
}

type ListParam[T any] struct {
	Filters    T
	Orders     string
	Pagination Pagination
}

type ListResponse[T any] struct {
	Rows     []*T       `json:"rows"`
	MetaData Pagination `json:"meta_data"`
}

type SelectAndCount[M any] struct {
	Rows  []*M
	Count int64
}
