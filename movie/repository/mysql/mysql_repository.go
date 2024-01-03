package mysql

import (
	"github.com/mqnoy/movie-rest-api/domain"
	"github.com/mqnoy/movie-rest-api/model"
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

func (m mysqlMovieRepository) InsertMovie(row model.Movie) (*model.Movie, error) {
	err := m.DB.Create(&row).Error

	return &row, err
}
