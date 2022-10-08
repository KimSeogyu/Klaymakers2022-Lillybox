package batch

import (
	"lillybox-backend/internal/database"
	"net/http"

	"github.com/klaytn/klaytn/common"
	"go.uber.org/zap"
)

// Client ...
type Client struct {
	Environ      Environ
	ContractAddr common.Address
	httpClient   *http.Client
	Channel      PollingChannel
	Database     database.Database
	AccessLogger *zap.Logger
	ErrorLogger  *zap.Logger
}

// PollingChannel ...
type PollingChannel struct {
	TokenChan    chan *TokenInfo
	MetadataChan chan MappedData
}

// MappedData ...
type MappedData struct {
	TokenInfo TokenInfo
	Metadata  Metadata
}

// Environ ...
type Environ struct {
	KasAccessKey   string
	KasSecretKey   string
	LivepeerAPIKey string
}

// TokenInfo ...
type TokenInfo struct {
	TokenID  string `json:"token_id"`
	TokenURI string `json:"token_uri"`
}

// Property represents properties in metadata
type Property struct {
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	Description  string        `json:"description"`
	Categories   []interface{} `json:"categories"`
	VideoURI     string        `json:"video_uri"`
	ThumbnailURI string        `json:"thumbnail_uri"`
	CreatedAt    string        `json:"created_at"`
}

// Metadata represents metadata about the specified token
type Metadata struct {
	Title      string   `json:"title"`
	Type       string   `json:"type"`
	Properties Property `json:"properties"`
}

// TokenOwnerInfo ...
type TokenOwnerInfo struct {
	Items  []Items
	Cursor string
}

// Items represents ...
type Items struct {
	TokenID         string  `json:"token_id"`
	Owner           string  `json:"owner"`
	TokenAddress    string  `json:"token_address"`
	TokenURI        string  `json:"token_uri"`
	Balance         string  `json:"balance"`
	TransactionHash string  `json:"transaction_hash"`
	TransferFrom    string  `json:"transfer_from"`
	TransferTo      string  `json:"transfer_to"`
	UpdatedAt       float64 `json:"updated_at"`
}

// LivepeerRequest represents the request to livepeer
type LivepeerRequest struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// UploadLivepeerResp represents the response from UploadLivepeerResp
type UploadLivepeerResp struct {
	Asset Asset `json:"asset"`
	Task  Task  `json:"task"`
}

// Asset represents the asset associated with UploadResponse
type Asset struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Status    AssetStatus `json:"status"`
	UserID    string      `json:"user_id"`
	CreatedAt string      `json:"created_at"`
}

// AssetStatus represents the status of uploading on Livepeer
type AssetStatus struct {
	Phase     string `json:"phase"`
	UpdatedAt string `json:"updated_at"`
}

// Task represents the id of the task on Livepeer
type Task struct {
	ID string `json:"id"`
}

// CheckUploadStatusResp represents the response from CheckUploadStatus
type CheckUploadStatusResp struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Hash        VODHash      `json:"hash"`
	Status      UploadStatus `json:"status"`
	UserID      string       `json:"user_id"`
	Size        uint64       `json:"size"`
	CreatedAt   float64      `json:"created_at"`
	PlaybackID  string       `json:"playback_id"`
	PlaybackURL string       `json:"playback_url"`
	DownloadURL string       `json:"download_url"`
	VideoSpec   VideoSpec    `json:"video_spec"`
}

// VideoSpec ...
type VideoSpec struct {
	Format   string        `json:"format"`
	Tracks   []interface{} `json:"tracks"`
	Duration float64       `json:"duration"`
}

// LivepeerHash ...
type LivepeerHash struct {
	Hash      string `json:"hash"`
	Algorithm string `json:"algorithm"`
}

// UploadStatus represents the status of uploading on Livepeer as same as AssetStatus
type UploadStatus struct {
	Phase     string  `json:"phase"`
	Progress  float64 `json:"progress"`
	UpdatedAt float64 `json:"updated_at"`
}

// VODHash ...
type VODHash []LivepeerHash
