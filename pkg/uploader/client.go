package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client 用于处理数据上报的客户端
type Client struct {
	ServerURL string
	Timeout   time.Duration
}

// NewClient 创建一个新的上报客户端
func NewClient(serverURL string, timeout time.Duration) *Client {
	return &Client{
		ServerURL: serverURL,
		Timeout:   timeout,
	}
}

// UploadData 上传数据到服务器
func (c *Client) UploadData(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.ServerURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: c.Timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload data, status code: %d", resp.StatusCode)
	}

	return nil
}
