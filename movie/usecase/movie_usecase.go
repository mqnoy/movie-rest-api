package usecase

import (
	"errors"
	"fmt"
	"math"
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
			return nil, cerror.WrapError(404, cerror.ErrResourceNotFound)
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

func (u *movieUseCase) ListMovies(ctx *gin.Context, param dto.ListParam[dto.FilterMovieParams]) (*dto.ListResponse[dto.Movie], *cerror.CustomError) {
	pagination := param.Pagination
	param.Pagination.Offset = (pagination.Page - 1) * pagination.Limit

	rows, err := u.movieRepo.SelectAndCountUnit(param)
	if err != nil {
		logger.Error().
			Err(err).
			Str("context", "usecase.movie").
			Send()

		return nil, cerror.WrapError(500, fmt.Errorf("internal server error"))
	}

	totalItems := rows.Count
	// Create pagination metadata
	totalPages := int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))

	return &dto.ListResponse[dto.Movie]{
		Rows: u.ParseMovieListResponse(rows.Rows),
		MetaData: dto.Pagination{
			Page:       pagination.Page,
			Limit:      pagination.Limit,
			TotalPages: totalPages,
			TotalItems: totalItems,
		},
	}, nil
}

func (u *movieUseCase) ParseMovieListResponse(m []*model.Movie) []*dto.Movie {
	result := make([]*dto.Movie, len(m))
	for idx, el := range m {
		result[idx] = u.ParseMovieResponse(el)
	}

	return result
}

func (u *movieUseCase) RemoveMovie(ctx *gin.Context, param dto.MovieDetailParam) *cerror.CustomError {
	// validate resource is exists
	row, err := u.fetchMovie(ctx, param.ID)
	if err != nil {
		return err
	}

	// persist delete movie
	if err := u.movieRepo.DeleteMovieById(row.ID); err != nil {
		logger.Error().
			Err(err).
			Str("context", "usecase.movie").
			Send()

		return cerror.WrapError(500, fmt.Errorf("internal server error"))
	}

	return nil
}

func (u *movieUseCase) UpdateMovie(ctx *gin.Context, param dto.MovieUpdateParam) (*dto.Movie, *cerror.CustomError) {
	// validate resource is exists
	id := param.Payload.ID
	updateValue := param.Payload

	row, err := u.fetchMovie(ctx, id)
	if err != nil {
		return nil, err
	}

	values := map[string]interface{}{}
	if updateValue.Title != "" {
		values["title"] = updateValue.Title
	}

	if updateValue.Description != "" {
		values["description"] = updateValue.Description
	}

	rating := updateValue.Rating
	if rating != nil {
		values["rating"] = *updateValue.Rating
	}

	if updateValue.Image != "" {
		values["image"] = updateValue.Image
	}

	// persist update
	if err := u.movieRepo.UpdateMovieById(row.ID, values); err != nil {

		logger.Error().
			Err(err).
			Str("context", "usecase.movie").
			Send()

		return nil, cerror.WrapError(500, fmt.Errorf("internal server error"))
	}

	// select updated row
	updatedRow, err := u.fetchMovie(ctx, id)
	if err != nil {
		return nil, err
	}

	return u.ParseMovieResponse(updatedRow), nil
}
