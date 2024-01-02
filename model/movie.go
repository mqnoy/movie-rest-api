package model

type Movie struct {
	ID          int
	Title       string  `gorm:"column:title;type:varchar(255);index"`
	Description string  `gorm:"column:description"`
	Rating      float32 `gorm:"column:rating;type:decimal(2,1);default:0"`
	Image       string  `gorm:"column:image"`
	TimestampColumn
}
