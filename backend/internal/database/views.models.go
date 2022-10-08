package database

import (
	"time"

	"gorm.io/gorm"
)

// VidViews ...
type VidViews struct {
	VidsID    uint      `json:"vids_id"`
	Count     uint      `gorm:"not_null;default:0" json:"count"`
	Version   uint      `gorm:"not_null;default:0" json:"version"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	ID        uint      `gorm:"primaryKey;authoIncrement;index;"`
}

// AdsViews ...
type AdsViews struct {
	Count     uint      `gorm:"not_null" json:"count"`
	Version   uint      `gorm:"not_null" json:"version"`
	AdsID     string    `gorm:"not_null" json:"ads_id"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	ID        uint      `gorm:"primaryKey;authoIncrement;index;"`
}

// ReadVidView ...
func (db *Database) ReadVidView(id uint) (*VidViews, error) {
	var result VidViews
	if err := db.Client.Table(db.VidViews.TableName()).Where(&VidViews{VidsID: id}).First(&result); err.Error != nil {
		return nil, err.Error
	}
	return &result, nil
}

// ReadAdsView ...
func (db *Database) ReadAdsView(cid string) (*AdsViews, error) {
	var result AdsViews
	if err := db.Client.Table(db.AdsViews.TableName()).Where(&AdsViews{AdsID: cid}).First(&result); err.Error != nil {
		return nil, err.Error
	}
	return &result, nil
}

// InsertVidView ...
func (db *Database) InsertVidView(vid Vids) (bool, error) {
	data := VidViews{
		Count:   0,
		Version: 0,
		VidsID:  vid.ID,
	}
	if err := db.Client.Table(db.VidViews.TableName()).Create(&data); err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

// InsertAdsView ...
func (db *Database) InsertAdsView(ads Ads) (bool, error) {
	data := AdsViews{
		Count:   0,
		Version: 0,
		AdsID:   ads.ContentID,
	}
	if err := db.Client.Table(db.AdsViews.TableName()).Create(&data); err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

// UpdateVidView ...
func (db *Database) UpdateVidView(id uint) (bool, error) {
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}
	var v VidViews
	var AtFirst uint
	var AtLast uint
	if err := tx.Table(db.VidViews.TableName()).Where(&VidViews{VidsID: id}).First(&v); err.Error != nil {
		tx.Rollback()
		return false, err.Error
	}
	AtFirst = v.Version
	if err := tx.Table(db.VidViews.TableName()).Where(&VidViews{VidsID: id}).Updates(&VidViews{Count: v.Count + 1, Version: v.Version + 1}); err.Error != nil {
		tx.Rollback()
		return false, err.Error
	}
	if err := tx.Table(db.VidViews.TableName()).Where(&VidViews{VidsID: id}).First(&v); err.Error != nil {
		tx.Rollback()
		return false, err.Error
	}
	AtLast = v.Version
	if AtFirst+1 != AtLast {
		tx.Rollback()
		return false, gorm.ErrInvalidTransaction
	}
	return true, tx.Commit().Error
}

// UpdateAdsView ...
func (db *Database) UpdateAdsView(cid string) (bool, error) {
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}
	var v AdsViews
	if err := db.Client.Table(db.AdsViews.TableName()).Where(&AdsViews{AdsID: cid}).First(&v); err.Error != nil {
		tx.Rollback()
		return false, err.Error
	}
	if err := tx.Model(db.AdsViews).Updates(&AdsViews{Count: v.Count + 1, Version: v.Version + 1}); err.Error != nil {
		tx.Rollback()
		return false, err.Error
	}
	return true, tx.Commit().Error
}

// DeleteVidView ...
func (db *Database) DeleteVidView() {
}
