package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type URL struct {
	gorm.Model
	ShortCode   string `gorm:"uniqueIndex;not null"`
	OriginalURL string `gorm:"not null"`
	Clicks      int    `gorm:"default:0;not null"`
}

func Connect() error {
	dsn := "host=localhost user=godev password=godevpass dbname=urlshortener port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connect to db: %w", err)
	}
	DB = db
	return nil
}
func SaveURL(shortCode, originalURL string) (URL, error) {
	url := URL{
		ShortCode:   shortCode,
		OriginalURL: originalURL,
	}
	result := DB.Create(&url)
	return url, result.Error
}

func GetURL(shortCode string) (URL, error) {
	var url URL
	result := DB.Where("short_code = ? ", shortCode).First(&url)
	return url, result.Error
}

func IncrementClicks(shortCode string) error {
	result := DB.
		Model(&URL{}).
		Where("short_code = ?", shortCode).
		Update("clicks", gorm.Expr("clicks + 1"))
	return result.Error
}

func GetStats(shortCode string) (URL, error) {
    return GetURL(shortCode)
}
