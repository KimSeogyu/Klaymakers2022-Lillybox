package database

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Tokens ...
type Tokens struct {
	TokenID   string    `gorm:"primaryKey;autoIncrement:false;"`
	TokenURI  string    `gorm:"not_null"`
	CreatedAt time.Time `gorm:"not_null"`
	UpdatedAt time.Time `gorm:"not_null"`
}

// ReadToken ...
func (db *Database) ReadToken(tokenID string) (*Tokens, error) {
	var result Tokens
	if err := db.Client.Table(db.Tokens.TableName()).Where(&Tokens{TokenID: tokenID}, "token_id").First(&result); err.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &result, nil
}

// InsertToken ...
func (db *Database) InsertToken(tokenID string, tokenURI string) (bool, error) {
	token := Tokens{TokenID: tokenID, TokenURI: tokenURI}
	if err := db.Client.Create(&token); err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

// UpdateToken ...
func (db *Database) UpdateToken(addr string) (bool, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return false, fiber.ErrInternalServerError
	}
	if err := db.Client.Model(db.Users).Where(&Users{Address: addr}).Update("LoginRequestID", uid); err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

// DeleteToken ...
func (db *Database) DeleteToken(addr string) (bool, error) {
	if err := db.Client.Delete(&Users{Address: addr}); err.Error != nil {
		return false, err.Error
	}
	return true, nil
}
