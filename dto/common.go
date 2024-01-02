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
