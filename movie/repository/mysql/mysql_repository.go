package mysql

import (
	"github.com/mqnoy/movie-rest-api/domain"
	"github.com/mqnoy/movie-rest-api/dto"
	"github.com/mqnoy/movie-rest-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (m mysqlMovieRepository) SelectAndCountUnit(param dto.ListParam[dto.FilterMovieParams]) (*dto.SelectAndCount[model.Movie], error) {
	var rows []*model.Movie
	var count int64
	var result *gorm.DB

	filters := param.Filters
	orders := param.Orders
	pagination := param.Pagination
	whereClause := clause.Where{}
	mDB := m.DB

	if filters.Title != "" {
		whereClause.Exprs = append(whereClause.Exprs, clause.Where{
			Exprs: []clause.Expression{
				clause.Like{
					Column: clause.Column{
						Name: "title",
					},
					Value: "%" + filters.Title + "%",
				},
			},
		})
	}

	if len(whereClause.Exprs) > 0 {
		mDB = m.DB.Clauses(whereClause)
	}

	m.DB.Model(&model.Movie{}).Count(&count)

	result = mDB.
		Limit(pagination.Limit).Offset(pagination.Offset).
		Order(orders).
		Find(&rows)

	if result.Error != nil {
		return &dto.SelectAndCount[model.Movie]{
			Rows:  rows,
			Count: count,
		}, result.Error
	}

	return &dto.SelectAndCount[model.Movie]{
		Rows:  rows,
		Count: count,
	}, nil
}
