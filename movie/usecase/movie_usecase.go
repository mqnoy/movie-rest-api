package usecase

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mqnoy/movie-rest-api/domain"
	"github.com/mqnoy/movie-rest-api/dto"
	"github.com/mqnoy/movie-rest-api/model"
	"github.com/mqnoy/movie-rest-api/pkg/cerror"
	"github.com/mqnoy/movie-rest-api/pkg/logger"
	"gorm.io/gorm"
)

type movieUseCase struct {
	movieRepo domain.MovieRepository
}

func New(movieRepo domain.MovieRepository) domain.MovieUseCase {
	return &movieUseCase{
		movieRepo: movieRepo,
	}
}

func (u *movieUseCase) CreateMovie(ctx *gin.Context, param dto.MovieCreateParam) (*dto.Movie, *cerror.CustomError) {
	payload := param.Payload

	// validate title
	title := strings.TrimSpace(payload.Title)
	if len(title) == 0 {
		return nil, cerror.WrapError(400, cerror.ErrCantBeEmpty)
	}

	row := model.Movie{
		Title:       payload.Title,
		Description: payload.Description,
		Rating:      payload.Rating,
		Image:       payload.Image,
	}

	result, err := u.movieRepo.InsertMovie(row)
	if err != nil {
		logger.Error().
			Err(err).
			Str("context", "usecase.movie").
			Send()
		return nil, cerror.WrapError(500, fmt.Errorf("internal server error"))
	}

	return u.ParseMovieResponse(result), nil
}

func (u *movieUseCase) ParseMovieResponse(m *model.Movie) *dto.Movie {
	return &dto.Movie{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Rating:      m.Rating,
		Image:       m.Image,
		Timestamp:   dto.ParseTimestampResponse(m.TimestampColumn),
	}
}

func (u *movieUseCase) fetchMovie(ctx *gin.Context, id int) (*model.Movie, *cerror.CustomError) {
	row, err := u.movieRepo.SelectMovieById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.WrapError(404, cerror.ErrCantBeEmpty)
		}

		logger.Error().
			Err(err).
			Str("context", "usecase.movie").
			Send()
		return nil, cerror.WrapError(500, fmt.Errorf("internal server error"))
	}
	return row, nil
}

func (u *movieUseCase) DetailMovie(ctx *gin.Context, param dto.MovieDetailParam) (*dto.Movie, *cerror.CustomError) {
	row, err := u.fetchMovie(ctx, param.ID)
	if err != nil {
		return nil, err
	}

	return u.ParseMovieResponse(row), nil
}
