package database

import "time"

// CommentReports ...
type CommentReports struct {
	CommentsID uint
	Nickname   string
	CreatedAt  time.Time `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	ID         uint      `gorm:"primaryKey;authoIncrement;index;"`
}

// VidReports ...
type VidReports struct {
	VidsID    uint      `json:"vids_id"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	ID        uint      `gorm:"primaryKey;authoIncrement;index;"`
}

// ReportVidDto ...
type ReportVidDto struct {
	Nickname string `json:"nickname"`
}

// ReportCommentDto ...
type ReportCommentDto struct {
	Nickname string `json:"nickname"`
}

// ReportVid ...
func (db *Database) ReportVid(dto ReportVidDto, id uint) (bool, error) {
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false, tx.Error
	}
	if err := tx.Create(&VidReports{VidsID: id, Nickname: dto.Nickname}); err.Error != nil {
		tx.Rollback()
		return false, err.Error
	}
	return true, tx.Commit().Error
}

// ReportComment ...
func (db *Database) ReportComment(dto ReportCommentDto, id uint) (bool, error) {
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return false, tx.Error
	}
	if err := tx.Create(&CommentReports{CommentsID: id, Nickname: dto.Nickname}); err.Error != nil {
		tx.Rollback()
		return false, err.Error
	}
	return true, tx.Commit().Error
}

// ReadReportVid ...
func (db *Database) ReadReportVid(id uint) ([]VidReports, error) {
	var result []VidReports
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Table(db.VidReports.TableName()).Where(&VidReports{VidsID: id}).Find(&result); err.Error != nil {
		return nil, err.Error
	}
	return result, tx.Commit().Error
}

// ReadReportComment ...
func (db *Database) ReadReportComment(id uint) (*CommentReports, error) {
	var result CommentReports
	tx := db.Client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Table(db.CommentReports.TableName()).Where(&CommentReports{CommentsID: id}).Find(&result); err.Error != nil {
		return nil, err.Error
	}
	return &result, tx.Commit().Error
}
