package database

import (
	"time"

	"gorm.io/gorm"
)

// Ads ...
type Ads struct {
	ContentID string    `gorm:"not_null; index; unique;" json:"content_id"`
	Nickname  string    `gorm:"not_null; index;" json:"nickname"`
	VideoURI  string    `gorm:"not_null;" json:"video_uri"`
	Version   uint      `gorm:"not_null;" json:"version"`
	AdsViews  AdsViews  `gorm:"foreignKey:AdsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ads_view"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	ID        uint      `gorm:"primaryKey;authoIncrement;index;"`
}

// CreateAdsDto ...
type CreateAdsDto struct {
	CID      string `json:"cid"`
	Nickname string `json:"nickname"`
	VideoURI string `json:"video_uri"`
}

// ReadAds ...
func (db *Database) ReadAds(cid string) (*Ads, error) {
	var result Ads
	if err := db.Client.Where(Ads{ContentID: cid}).First(&result); err.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	return &result, nil
}

// InsertAds ...
func (db *Database) InsertAds(dto CreateAdsDto) (bool, error) {
	if err := db.Client.Transaction(func(tx *gorm.DB) error {
		data := &Ads{
			ContentID: dto.CID,
			Nickname:  dto.Nickname,
			VideoURI:  dto.VideoURI,
			Version:   0,
			AdsViews: AdsViews{
				Count:   0,
				Version: 0,
			},
		}
		db.Client.Transaction(func(tx2 *gorm.DB) error {
			if err := tx2.Create(data); err != nil {
				return err.Error
			}
			return tx2.Commit().Error
		})
		if _, err := db.InsertAdsView(*data); err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit().Error
	}); err != nil {
		return false, err
	}
	return true, nil
}

// UpdateAds ...
func (db *Database) UpdateAds() (bool, error) {
	// Not Implemented yet
	if err := db.Client.Transaction(func(tx *gorm.DB) error {
		return tx.Commit().Error
	}); err != nil {
		return false, err
	}
	return true, nil
}

// DeleteAds ...
func (db *Database) DeleteAds(cid string) (bool, error) {
	if err := db.Client.Delete(&Ads{ContentID: cid}); err.Error != nil {
		return false, err.Error
	}
	return true, nil

}
