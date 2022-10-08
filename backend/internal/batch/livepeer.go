package batch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// UploadLivepeer ...
func (c *Client) UploadLivepeer(metadata *Metadata) (bool, *UploadLivepeerResp, error) {
	url := "https://livepeer.studio/api/asset/import"
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(LivepeerRequest{Name: metadata.Properties.Name, URL: metadata.Properties.VideoURI})
	if err != nil {
		return false, nil, err
	}
	resp, err := GetHTTPRequest("POST", url, &buf, map[string]string{"Content-Type": "application/json", "Authorization": fmt.Sprintf("Bearer %v", c.Environ.LivepeerAPIKey)}, nil)
	if err != nil {
		return false, nil, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return false, nil, err
	}
	var result UploadLivepeerResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, nil, err
	}
	return true, &result, nil
}

// CheckUploadStatus ...
func (c *Client) CheckUploadStatus(data *UploadLivepeerResp) (bool, *CheckUploadStatusResp, error) {
	url := fmt.Sprintf("https://livepeer.studio/api/asset/%v?details=true", data.Asset.ID)
	resp, err := GetHTTPRequest("GET", url, nil, map[string]string{"Authorization": fmt.Sprintf("Bearer %v", c.Environ.LivepeerAPIKey), "Content-Type": "application/json"}, nil)
	if err != nil {
		return false, nil, err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return false, nil, err
	}
	var result CheckUploadStatusResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, nil, err
	}
	return true, &result, nil
}
