package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// TableName ...
func (*CommentReports) TableName() string {
	return "comment_reports"
}

// TableName ...
func (*VidReports) TableName() string {
	return "vid_reports"
}

// TableName ...
func (*CategoryVideos) TableName() string {
	return "category_videos"
}

// TableName ...
func (*Ads) TableName() string {
	return "ads"
}

// TableName ...
func (*Users) TableName() string {
	return "users"
}

// TableName ...
func (*Vids) TableName() string {
	return "vids"
}

// TableName ...
func (*Comments) TableName() string {
	return "comments"
}

// TableName ...
func (*VidViews) TableName() string {
	return "vid_views"
}

// TableName ...
func (*AdsViews) TableName() string {
	return "ads_views"
}

// TableName ...
func (*Tokens) TableName() string {
	return "tokens"
}

// TableName ...
func (*Category) TableName() string {
	return "categories"
}

// ConnectDatabase ...
func (db *Database) ConnectDatabase() {

	db.Users = &Users{}
	db.Comments = &Comments{}
	db.Category = &Category{}
	db.Tokens = &Tokens{}
	db.Vids = &Vids{}
	db.CommentReports = &CommentReports{}
	db.VidReports = &VidReports{}
	db.Ads = &Ads{}
	db.VidViews = &VidViews{}
	db.AdsViews = &AdsViews{}

	err := godotenv.Load("./.env")
	FatalWithError(err)

	pgUser := os.Getenv("PG_USER")
	pgPw := os.Getenv("PG_PASSWORD")
	pgDBName := os.Getenv("PG_DATABASE")
	pgPort := os.Getenv("PG_PORT")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul", pgUser, pgPw, pgDBName, pgPort)
	sqldb, err := sql.Open("postgres", dsn)
	FatalWithError(err)
	err = sqldb.Ping()
	FatalWithError(err)
	client, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 sqldb,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: newLogger,
	})
	FatalWithError(err)
	pl("Postgresql connection done successfully")

	client.Migrator().CreateTable(db.Users)
	client.Migrator().CreateTable(db.Tokens)
	client.Migrator().CreateTable(db.Vids)
	client.Migrator().CreateTable(db.Comments)
	client.Migrator().CreateTable(db.Category)
	client.Migrator().CreateTable(db.Ads)
	client.Migrator().CreateTable(db.AdsViews)
	client.Migrator().CreateTable(db.VidViews)
	client.Migrator().CreateTable(db.CommentReports)
	client.Migrator().CreateTable(db.VidReports)
	client.Migrator().CreateTable(db.CategoryVideos)
	pl("CreateTable called successfully")

	Genres := []string{
		"All",
		"Music",
		"Sports",
		"Gaming",
		"News",
		"Entertainment",
		"Education",
		"Science & Technology",
		"Travel",
		"Other",
	}

	db.Client = client
	db.BatchInsertCategories(Genres)
}
