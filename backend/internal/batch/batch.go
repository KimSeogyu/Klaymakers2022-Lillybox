package batch

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// BatchLoop executes a batch operation for every 5 seconds
func (c *Client) BatchLoop() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(fmt.Sprint("Unexpected Errors Occurred :", r))
		}
	}()
	go c.PollingToken()
	for {
		select {
		case tokenInfo := <-c.Channel.TokenChan:
			if tokenInfo != nil && tokenInfo.TokenID != "0x0" {
				_metadata, err := c.GetMetadata(tokenInfo.TokenURI)
				switch {
				case errors.Is(err, nil):
					go PushMetadata(c.Channel.MetadataChan, MappedData{Metadata: *_metadata, TokenInfo: *tokenInfo})
				default:
					log.Println(fmt.Sprint("Unexpected Error Occurred!", err))
				}
			}
		case data := <-c.Channel.MetadataChan:
			owner, _ := c.GetTokenOwner(data.TokenInfo.TokenID)
			nickname, err := c.GetUserNickname(owner)
			if err != nil {
				continue
			}
			if data.Metadata.Properties.Type == "videos" {
				dto := ParseMetadataVideoDto(data.Metadata, nickname)
				c.Database.InsertVOD(*dto)
				c.Database.InsertToken(data.TokenInfo.TokenID, data.TokenInfo.TokenURI)
			} else {
				dto := ParseMetadataAdsDto(data.Metadata, nickname)
				c.Database.InsertAds(*dto)
			}
		}
	}
}

// PollingToken push new token to TokenChan
func (c *Client) PollingToken() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(fmt.Sprint("Unexpected Errors Occurred :", r))
		}
	}()
	t := time.Tick(time.Minute)
	var cursor any = ""
	for range t {
		ok, _body, err := c.GetToken(cursor)
		if err != nil || !ok {
			continue
		}
		for _, v := range _body["items"].([]interface{}) {
			items := v.(map[string]interface{})
			_tokenInfo, _ := c.GetTokenInfo(items)
			if _tokenInfo.TokenID != "0x0" && _tokenInfo.TokenID != "0x1" {
				_, err := c.Database.ReadToken(_tokenInfo.TokenID)
				if errors.Is(err, gorm.ErrRecordNotFound) {
					go PushTokenInfo(c.Channel.TokenChan, _tokenInfo)
				}
			}
		}
		if _body["cursor"] != "" {
			cursor = _body["cursor"]
		}
	}
}
