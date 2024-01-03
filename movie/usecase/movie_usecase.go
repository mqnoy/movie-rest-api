package usecase

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mqnoy/movie-rest-api/domain"
	"github.com/mqnoy/movie-rest-api/dto"
	"github.com/mqnoy/movie-rest-api/model"
	"github.com/mqnoy/movie-rest-api/pkg/cerror"
	"github.com/mqnoy/movie-rest-api/pkg/logger"
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
