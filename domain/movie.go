package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/mqnoy/movie-rest-api/dto"
	"github.com/mqnoy/movie-rest-api/model"
	"github.com/mqnoy/movie-rest-api/pkg/cerror"
)

type MovieUseCase interface {
	CreateMovie(ctx *gin.Context, param dto.MovieCreateParam) (*dto.Movie, *cerror.CustomError)
}

type MovieRepository interface {
	InsertMovie(row model.Movie) (*model.Movie, error)
}
