package database

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Users ...
type Users struct {
	Address        string           `gorm:"unique;not_null;index;" json:"address"`
	Nickname       string           `gorm:"unique;not_null;index;" json:"nickname"`
	LoginRequestID string           `gorm:"not_null" json:"login_request_id"`
	Videos         []Vids           `gorm:"foreignKey:Nickname;references:Nickname;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"videos"`
	Comments       []Comments       `gorm:"foreignKey:Nickname;references:Nickname;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
	VidReports     []VidReports     `gorm:"foreignKey:Nickname;references:Nickname;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"vid_reports"`
	CommentReports []CommentReports `gorm:"foreignKey:Nickname;references:Nickname;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comment_reports"`
	Ads            []Ads            `gorm:"foreignKey:Nickname;references:Nickname" json:"ads"`
	CreatedAt      time.Time        `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt      time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
	ID             uint             `gorm:"primaryKey;authoIncrement;index;"`
}

// ReadUser ...
func (db *Database) ReadUser(addr string) (*Users, error) {
	var result Users
	if err := db.Client.Table(db.Users.TableName()).Where(&Users{Address: addr}).First(&result); err.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &result, nil
}

// ReadUserByNickname ...
func (db *Database) ReadUserByNickname(nickname string) (*Users, error) {
	var result Users
	if err := db.Client.Where(&Users{Nickname: nickname}).First(&result); err.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &result, nil
}

// InsertUser ...
func (db *Database) InsertUser(addr string, nick string) (bool, error) {
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		tx.Rollback()
		return false, tx.Error
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		return false, fiber.ErrInternalServerError
	}
	user := Users{Address: addr, Nickname: nick, LoginRequestID: uid.String()}
	if err := tx.Create(&user); err.Error != nil {
		return false, gorm.ErrInvalidTransaction
	}
	return true, tx.Commit().Error
}

// UpdateUser ...
func (db *Database) UpdateUser(addr string) (bool, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return false, fiber.ErrInternalServerError
	}
	if err := db.Client.Model(db.Users).Where(&Users{Address: addr}).Update("LoginRequestID", uid); err.Error != nil {
		return false, fiber.ErrInternalServerError
	}
	return true, nil
}

// DeleteUser ...
func (db *Database) DeleteUser(addr string) (bool, error) {
	if err := db.Client.Delete(&Users{Address: addr}); err.Error != nil {
		return false, fiber.ErrInternalServerError
	}
	return true, nil
}
