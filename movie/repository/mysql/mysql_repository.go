package mysql

import (
	"github.com/mqnoy/movie-rest-api/domain"
	"gorm.io/gorm"
)

type mysqlMovieRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) domain.MovieRepository {
	return &mysqlMovieRepository{
		DB: db,
	}
}

// TODO: implement movie repository
