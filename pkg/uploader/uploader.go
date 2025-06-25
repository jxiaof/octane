package uploader

import (
    "bytes"
    "encoding/json"
    "net/http"
    "time"
)

// Uploader 负责处理数据上报
type Uploader struct {
    ServerURL string
    Client    *http.Client
}

// NewUploader 创建一个新的 Uploader 实例
func NewUploader(serverURL string) *Uploader {
    return &Uploader{
        ServerURL: serverURL,
        Client:    &http.Client{Timeout: 10 * time.Second},
    }
}

// UploadData 上传数据到服务器
func (u *Uploader) UploadData(data interface{}) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", u.ServerURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := u.Client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return &httpError{StatusCode: resp.StatusCode}
    }

    return nil
}

// httpError 用于处理 HTTP 错误
type httpError struct {
    StatusCode int
}

func (e *httpError) Error() string {
    return http.StatusText(e.StatusCode)
}