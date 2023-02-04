package database

import (
	"github.com/AnishriM/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and create comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return nil
	}
	return nil
}
