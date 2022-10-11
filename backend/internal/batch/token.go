package batch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetMetadata returns metadata about the specified token
func (c *Client) GetMetadata(TokenURI string) (*Metadata, error) {
	req, err := GetHTTPRequest("GET", "https://ipfs.io/ipfs/"+TokenURI+"?filename=metadata.json", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var body map[string]Metadata
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, err
	}
	if _, ok := body["metadata"]; !ok {
		return nil, fiber.UnknownKeyError{Key: "metadata"}
	}
	metadata := body["metadata"]
	return &metadata, nil
}

// GetTokenInfo returns information about the specified token
func (c *Client) GetTokenInfo(items map[string]interface{}) (*TokenInfo, error) {
	if _, ok := items["tokenUri"]; !ok {
		return nil, fiber.UnknownKeyError{Key: "tokenUri"}
	}
	if _, ok := items["tokenId"]; !ok {
		return nil, fiber.UnknownKeyError{Key: "tokenId"}
	}
	tokenURI := items["tokenUri"].(string)
	tokenID := items["tokenId"].(string)
	return &TokenInfo{TokenID: tokenID, TokenURI: tokenURI}, nil
}

// GetToken ...
func (c *Client) GetToken(cursor any) (bool, map[string]any, error) {
	url := fmt.Sprintf("https://th-api.klaytnapi.com/v2/contract/mt/%v/token?size=50&cursor=%v", c.ContractAddr.String(), cursor)
	req, err := GetHTTPRequest("GET", url, nil, map[string]string{"x-chain-id": "1001", "Content-Type": "application/json"}, map[string]string{c.Environ.KasAccessKey: c.Environ.KasSecretKey})
	if err != nil {
		return false, nil, err
	}
	log.Println(fmt.Sprintf("GET ACCESS TO %s", url))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, nil, err
	}
	var body map[string]any
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return false, nil, err
	}
	if _, ok := body["cursor"]; !ok {
		return false, nil, err
	}
	if _, ok := body["items"]; !ok {
		return false, nil, err
	}
	return true, body, nil
}

// GetTokenOwner ...
func (c *Client) GetTokenOwner(tokenID string) (string, error) {
	url := fmt.Sprintf("https://th-api.klaytnapi.com/v2/contract/mt/%v/token/%v?size=1", c.ContractAddr.String(), tokenID)
	req, err := GetHTTPRequest("GET", url, nil,
		map[string]string{"x-chain-id": "1001", "Content-Type": "application/json"},
		map[string]string{c.Environ.KasAccessKey: c.Environ.KasSecretKey})
	if err != nil {
		return "", err
	}
	log.Println(fmt.Sprintf("GET ACCESS TO %s", url))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var res TokenOwnerInfo
	err = json.Unmarshal(body, &res)
	return res.Items[0].Owner, nil
}

// GetUserNickname ...
func (c *Client) GetUserNickname(owner string) (string, error) {
	user, err := c.Database.ReadUser(owner)
	switch {
	case errors.Is(err, nil):
		return user.Nickname, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return "", gorm.ErrRecordNotFound
	default:
		return "", gorm.ErrInvalidTransaction
	}
}
