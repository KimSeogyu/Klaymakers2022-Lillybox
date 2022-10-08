package database

import (
	"time"

	"gorm.io/gorm/clause"
)

// Comments ...
type Comments struct {
	VidsID        uint             `json:"vids_id"`
	Nickname      string           `gorm:"not_null" json:"nickname"`
	Description   string           `gorm:"not_null" json:"description"`
	Version       uint             `json:"version"`
	CommentReport []CommentReports `json:"comment_reports"`
	CreatedAt     time.Time        `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt     time.Time        `gorm:"autoUpdateTime:nano" json:"updated_at"`
	ID            uint             `gorm:"primaryKey;authoIncrement;index;"`
}

// CreateCommentDto : Dto for create comment
type CreateCommentDto struct {
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
}

// UpdateCommentDto : Dto for update comment
type UpdateCommentDto struct {
	Description string `json:"description"`
}

// ReadComment ...
func (db *Database) ReadComment(id uint) ([]Comments, error) {
	var result Vids
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Table(db.Vids.TableName()).Where(&Vids{ID: id}).Preload(clause.Associations).Find(&result); err.Error != nil {
		return nil, err.Error
	}
	return result.Comment, tx.Commit().Error
}

// InsertComment ...
func (db *Database) InsertComment(dto CreateCommentDto, id uint) (bool, error) {
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false, tx.Error
	}
	if err := tx.Create(&Comments{
		VidsID:      id,
		Nickname:    dto.Nickname,
		Description: dto.Description,
		Version:     0,
	}); err.Error != nil {
		tx.Rollback()
		return false, err.Error
	}
	return true, tx.Commit().Error
}

// UpdateComment ...
func (db *Database) UpdateComment(dto UpdateCommentDto, id uint) (bool, error) {
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false, tx.Error
	}
	if err := tx.Table(db.Comments.TableName()).Where(&Comments{ID: id}).Update("description", dto.Description); err.Error != nil {
		tx.Rollback()
		return false, tx.Error
	}
	return true, tx.Commit().Error
}

// DeleteComment ...
func (db *Database) DeleteComment(id uint) (bool, error) {
	tx := db.Client.Begin()
	if err := tx.Delete(&Comments{}, id); err.Error != nil {
		return false, err.Error
	}
	return true, tx.Commit().Error
}
