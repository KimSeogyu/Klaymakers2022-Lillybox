package batch

import (
	"context"
	"fmt"
	"io"
	"lillybox-backend/internal/database"
	"log"
	"net/http"
	"strings"
)

// GetHTTPRequest returns a http response from the given request
func GetHTTPRequest(method string, url string, body io.Reader, header map[string]string, auth map[string]string) (*http.Request, error) {
	request, err := http.NewRequestWithContext(context.Background(), method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		request.Header.Set(k, v)
	}
	for k, v := range auth {
		request.SetBasicAuth(k, v)
	}
	if err != nil {
		return nil, err
	}
	return request, err
}

// PushTokenInfo into given channel
func PushTokenInfo(channel chan *TokenInfo, context *TokenInfo) {
	channel <- context
}

// PushMetadata into given channel
func PushMetadata(channel chan MappedData, context MappedData) {
	channel <- context
}

// ParseMetadataVideoDto ...
func ParseMetadataVideoDto(metadata Metadata, owner string) *database.CreateVideoDto {
	ok, cid := ParseContentID(metadata.Properties.VideoURI)
	if !ok {
		return nil
	}
	video := &database.CreateVideoDto{
		CID:          cid,
		Nickname:     owner,
		VideoName:    metadata.Properties.Name,
		Description:  metadata.Properties.Description,
		Category:     ParseCategory(metadata.Properties.Categories),
		ThumbnailURI: metadata.Properties.ThumbnailURI,
		VideoURI:     metadata.Properties.VideoURI,
	}
	return video
}

// ParseMetadataAdsDto ...
func ParseMetadataAdsDto(metadata Metadata, owner string) *database.CreateAdsDto {
	ok, cid := ParseContentID(metadata.Properties.VideoURI)
	if !ok {
		return nil
	}
	video := &database.CreateAdsDto{
		CID:      cid,
		Nickname: owner,
		VideoURI: metadata.Properties.VideoURI,
	}
	return video
}

// ParseContentID ...
func ParseContentID(tokenURI string) (bool, string) {
	_, after, ok := strings.Cut(tokenURI, "https://lillybox.infura-ipfs.io/ipfs/")
	if !ok {
		return false, ""
	}
	return true, after
}

// ParseCategory ...
func ParseCategory(category []interface{}) []string {
	result := make([]string, len(category))
	for i, v := range category {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}

func FatalWithError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
