package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Articles - articles models in database
type Articles struct {
	ID      uint64     `gorm:"column:id; primaryKey; type:int unsigned auto_increment; not null" json:"id"`
	Author  string     `gorm:"column:author; type:text; not null" json:"author"`
	Body    string     `gorm:"column:body; type:text; not null" json:"body"`
	Created *time.Time `gorm:"created; type:datetime; not null; default:current_timestamp" json:"created"`
}

// TableName -
func (Articles) TableName() string {
	return "kumparan.articles"
}

// InsertArticles -
func InsertArticles(db *gorm.DB, article *Articles) error {
	return db.Create(article).Error
}

// GetDataArticles -
func GetDataArticles(db *gorm.DB, arr *Articles, ID uint64) error {
	if err := db.Raw(`
		SELECT * from kumparan.articles WHERE id = ? 
	`, ID).Scan(arr).Error; err != nil {
		return err
	}

	return nil
}
