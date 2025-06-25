package uploader

import (
	"encoding/json"
	"strings"
)

// AnonymizeData 用于匿名化数据的结构体
type AnonymizeData struct {
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
	UserName  string `json:"user_name"`
}

// Anonymize 将敏感数据匿名化
func Anonymize(data AnonymizeData) AnonymizeData {
	return AnonymizeData{
		UserID:    "anonymous",
		UserEmail: "anonymous@example.com",
		UserName:  "Anonymous User",
	}
}

// AnonymizeJSON 将JSON格式的数据进行匿名化
func AnonymizeJSON(jsonData string) (string, error) {
	var data AnonymizeData
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return "", err
	}

	anonymizedData := Anonymize(data)
	anonymizedJSON, err := json.Marshal(anonymizedData)
	if err != nil {
		return "", err
	}

	return string(anonymizedJSON), nil
}

// AnonymizeString 将字符串中的敏感信息进行匿名化
func AnonymizeString(input string) string {
	// 替换敏感信息
	anonymized := strings.ReplaceAll(input, "user_id", "anonymous")
	anonymized = strings.ReplaceAll(anonymized, "user_email", "anonymous@example.com")
	anonymized = strings.ReplaceAll(anonymized, "user_name", "Anonymous User")
	return anonymized
}
