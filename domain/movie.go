package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/mqnoy/movie-rest-api/dto"
	"github.com/mqnoy/movie-rest-api/model"
	"github.com/mqnoy/movie-rest-api/pkg/cerror"
)

type MovieUseCase interface {
	CreateMovie(ctx *gin.Context, param dto.MovieCreateParam) (*dto.Movie, *cerror.CustomError)
	DetailMovie(ctx *gin.Context, param dto.MovieDetailParam) (*dto.Movie, *cerror.CustomError)
	ListMovies(ctx *gin.Context, param dto.ListParam[dto.FilterMovieParams]) (*dto.ListResponse[dto.Movie], *cerror.CustomError)
	RemoveMovie(ctx *gin.Context, param dto.MovieDetailParam) *cerror.CustomError
	UpdateMovie(ctx *gin.Context, param dto.MovieUpdateParam) (*dto.Movie, *cerror.CustomError)
}

type MovieRepository interface {
	InsertMovie(row model.Movie) (*model.Movie, error)
	SelectMovieById(id int) (*model.Movie, error)
	SelectAndCountUnit(param dto.ListParam[dto.FilterMovieParams]) (*dto.SelectAndCount[model.Movie], error)
	DeleteMovieById(id int) error
	UpdateMovieById(id int, values interface{}) error
}
