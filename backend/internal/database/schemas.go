package database

import (
	"gorm.io/gorm"
)

// Table ...
type Table interface {
	TableName() string
}

// Database ...
type Database struct {
	Users          *Users
	Vids           *Vids
	Category       *Category
	Comments       *Comments
	Tokens         *Tokens
	VidViews       *VidViews
	AdsViews       *AdsViews
	CommentReports *CommentReports
	VidReports     *VidReports
	Ads            *Ads
	CategoryVideos *CategoryVideos
	Client         *gorm.DB
}
