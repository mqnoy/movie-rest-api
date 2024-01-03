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

func (m mysqlMovieRepository) SelectMovieById(id int) (*model.Movie, error) {
	var row *model.Movie
	if err := m.DB.Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}

	return row, nil
}
