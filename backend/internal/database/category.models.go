package database

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Category ...
type Category struct {
	Name      string    `gorm:"not_null;index;unique"`
	Videos    []*Vids   `gorm:"many2many:category_videos;"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	ID        uint      `gorm:"primaryKey;authoIncrement;index;"`
}

// ReadCategories ...
func (db *Database) ReadCategories(data []string) ([]*Category, error) {
	var categories []*Category
	if err := db.Client.Table(db.Category.TableName()).Where("Name IN ?", data).Find(&categories); err.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return categories, nil
}

// InsertCategories ...
func (db *Database) InsertCategories(name string) (bool, error) {
	if err := db.Client.Create(&Category{Name: name}); err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

// BatchInsertCategories ...
func (db *Database) BatchInsertCategories(data []string) (bool, error) {
	var categories []Category
	for _, v := range data {
		temp := Category{Name: v}
		categories = append(categories, temp)
	}
	if err := db.Client.CreateInBatches(categories, 100); err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

// UpdateCategories ...
func (db *Database) UpdateCategories(name string) (bool, error) {
	if err := db.Client.Model(db.Category).Where(&Category{Name: name}).Update("Name", name); err != nil {
		return false, fiber.ErrInternalServerError
	}
	return true, nil
}

// DeleteCategories ...
func (db *Database) DeleteCategories(name string) (bool, error) {
	if err := db.Client.Delete(&Category{Name: name}); err != nil {
		return false, fiber.ErrInternalServerError
	}
	return true, nil
}
