package usecase

import "github.com/mqnoy/movie-rest-api/domain"

type movieUseCase struct {
	movieRepo domain.MovieRepository
}

func New(movieRepo domain.MovieRepository) domain.MovieUseCase {
	return &movieUseCase{
		movieRepo: movieRepo,
	}
}

// TODO: implement movie useCase
