package http

import (
	"errors"
	"fmt"
	"lillybox-backend/internal/database"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/livepeer/go-api-client"
	"go.uber.org/zap"
)

// EXPORT
const (
	ADDR           = "account"
	SIGNATURE      = "signature"
	NICK           = "nickname"
	lillySignature = "Lillybox login message\nLogin request id is "
)

var err = errors.New
var toString = fmt.Sprint
var pl = log.Println

// Handlers ...
type Handlers struct {
	LivepeerClient *api.Client
	Database       *database.Database
	AccessLogger   *zap.Logger
	ErrorLogger    *zap.Logger
}

// GetLoginRequestIDBody ...
type GetLoginRequestIDBody struct {
	Account string `json:"account"`
}

// SignUpBody ...
type SignUpBody struct {
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
}

// GetUserInfoBody ...
type GetUserInfoBody struct {
	Account string `json:"account"`
}

// GetNicknameBody ...
type GetNicknameBody struct {
	Nickname string `json:"nickname"`
}

// ViewEvent ...
type ViewEvent struct {
	Type      interface{}
	ContentID string
}

// DefaultResponse ...
type DefaultResponse struct {
	Message string `json:"message"`
}

// LoginResponse ...
type LoginResponse struct {
	RequestID string `json:"request_id"`
}

// BoolResponse ...
type BoolResponse struct {
	Success bool `json:"success"`
}

// GetOnDemandResponse ...
type GetOnDemandResponse struct {
	Success bool         `json:"success"`
	Result  []LillyVideo `json:"result"`
}

// GetOnDemandByIDResponse ...
type GetOnDemandByIDResponse struct {
	Success bool       `json:"success"`
	Result  LillyVideo `json:"result"`
}

// DefaultDataResponse ...
type DefaultDataResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

// LillyVideo ...
type LillyVideo struct {
	ID           uint      `json:"id"`
	CID          string    `json:"cid"`
	Nickname     string    `json:"nickname"`
	Account      string    `json:"account"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Type         string    `json:"type"`
	Categories   []string  `json:"categories"`
	VideoURI     string    `json:"video_uri"`
	ThumbnailURI string    `json:"thumbnail_uri"`
	Views        int       `json:"views"`
	CreatedAt    time.Time `json:"created_at"`
}

// LillyComment ...
type LillyComment struct {
	ID          uint      `json:"id"`
	Nickname    string    `json:"nickname"`
	Description string    `json:"description"`
	CretaedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

// Server ...
type Server struct {
	App      *fiber.App
	Handlers Handlers
	Config   AppConfig
}

// AppConfig ...
type AppConfig struct {
	Port     int
	Livepeer struct {
		Host   string
		APIKey string
	}
}
