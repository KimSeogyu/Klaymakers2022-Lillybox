package database

import (
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Vids ...
type Vids struct {
	ContentID    string       `gorm:"not_null;" json:"content_id"`
	Nickname     string       `gorm:"not_null" json:"nickname"`
	VideoName    string       `gorm:"not_null" json:"video_name"`
	Description  string       `json:"description"`
	Version      uint         `gorm:"not_null" json:"version"`
	VideoURI     string       `gorm:"not_null" json:"video_uri"`
	ThumbnailURI string       `gorm:"not_null" json:"thumbnail_uri"`
	VidViews     VidViews     `gorm:"foreignKey:VidsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"vid_views"`
	Category     []*Category  `gorm:"many2many:category_videos;" json:"category"`
	Comment      []Comments   `gorm:"foreignKey:VidsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment"`
	VidReports   []VidReports `gorm:"foreignKey:VidsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"vid_reports"`
	CreatedAt    time.Time    `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt    time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	ID           uint         `gorm:"primaryKey;authoIncrement;index;"`
}

// CategoryVideos ...
type CategoryVideos struct {
	VidsID     uint `json:"vids_id"`
	CategoryID uint `json:"category_id"`
}

// CreateVideoDto ...
type CreateVideoDto struct {
	CID          string   `json:"cid"`
	Nickname     string   `json:"nickname"`
	VideoName    string   `json:"video_name"`
	Description  string   `json:"description"`
	Category     []string `json:"category"`
	ThumbnailURI string   `json:"thumbnail_uri"`
	VideoURI     string   `json:"video_uri"`
}

// UpdateVideoDto : Dto for updating comment
type UpdateVideoDto struct {
	VideoName   string   `json:"video_name"`
	Description string   `json:"description"`
	Category    []string `json:"category"`
}

// ReadVOD ...
func (db *Database) ReadVOD(id uint) (*Vids, error) {
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return nil, tx.Error
	}
	var result Vids
	if err := tx.Table(db.Vids.TableName()).Where(&Vids{ID: id}).Preload(clause.Associations).Find(&result); err.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &result, tx.Commit().Error
}

// ReadManyVOD ...
func (db *Database) ReadManyVOD(offset string, category string) ([]*Vids, error) {
	tx := db.Client.Begin()
	var result []*Vids
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return nil, tx.Error
	}
	o, err := strconv.Atoi(offset)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if category == "1" {
		if err := db.Client.Table(db.Vids.TableName()).Offset(o * 20).Limit(20).Preload(clause.Associations).Find(&result); err.Error != nil {
			tx.Rollback()
			return nil, err.Error
		}
		return result, nil
	}
	if err := tx.Raw("SELECT * FROM vids WHERE id IN (SELECT vids_id FROM category_videos WHERE category_id = ?) LIMIT 20 OFFSET ?", category, 20*o).Preload(clause.Associations).Find(&result); err.Error != nil {
		tx.Rollback()
		return nil, err.Error
	}
	if len(result) != 0 {
		return result, nil
	}
	return nil, gorm.ErrRecordNotFound
}

// InsertVOD ...
func (db *Database) InsertVOD(dto CreateVideoDto) (bool, error) {
	tx := db.Client.Begin()
	tx2 := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}
	if err := tx2.Error; err != nil {
		return false, err
	}
	categories, err := db.ReadCategories(dto.Category)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	data := Vids{
		Nickname:     dto.Nickname,
		ContentID:    dto.CID,
		VideoName:    dto.VideoName,
		Description:  dto.Description,
		VideoURI:     dto.VideoURI,
		ThumbnailURI: dto.ThumbnailURI,
		Category:     categories,
		Version:      0,
		VidViews:     VidViews{Count: 0, Version: 0},
	}
	if err := tx2.Create(&data); err.Error != nil {
		tx.Rollback()
		return false, err.Error
	}
	tx2.Commit()
	if _, err := db.InsertVidView(data); err != nil {
		tx.Rollback()
	}
	return true, tx.Commit().Error
}

// UpdateVOD ...
func (db *Database) UpdateVOD(dto UpdateVideoDto, id uint) (bool, error) {
	tx := db.Client.Begin()
	tx2 := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			tx2.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}
	if err := tx2.Error; err != nil {
		return false, err
	}
	categories, err := db.ReadCategories(dto.Category)
	if err != nil {
		return false, err
	}
	var vid Vids
	var AtFirst uint
	var AtLast uint
	if err := tx2.Table(db.Vids.TableName()).Where(&Vids{ID: id}, "content_id").First(&vid); err.Error != nil {
		tx2.Rollback()
		tx.Rollback()
		return false, err.Error
	}
	AtFirst = vid.Version
	if err := tx2.Table(db.Vids.TableName()).Where(&Vids{ID: id}).Updates(&Vids{
		VideoName:   dto.VideoName,
		Description: dto.Description,
		Version:     vid.Version + 1,
	}); err.Error != nil {
		tx2.Rollback()
		tx.Rollback()
		return false, err.Error
	}
	if err := tx2.Model(&vid).Association("Category").Clear(); err != nil {
		tx2.Rollback()
		tx.Rollback()
		return false, err
	}
	if err := tx2.Model(&vid).Association("Category").Replace(categories); err != nil {
		tx2.Rollback()
		tx.Rollback()
		return false, err
	}
	if err := tx2.Table(db.Vids.TableName()).Where(&Vids{ID: id}, "content_id").First(&vid); err.Error != nil {
		tx2.Rollback()
		tx.Rollback()
		return false, err.Error
	}
	tx2.Commit()
	AtLast = vid.Version
	if AtFirst+1 != AtLast {
		tx2.Rollback()
		tx.Rollback()
		return false, gorm.ErrInvalidTransaction
	}
	return true, tx.Commit().Error
}

// DeleteVOD ...
func (db *Database) DeleteVOD(id uint) (bool, error) {
	if err := db.Client.Delete(&Vids{ID: id}); err.Error != nil {
		return false, err.Error
	}
	return true, nil
}
