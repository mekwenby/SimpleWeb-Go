package databases

/*
数据库工具函数存储路径
*/

import (
	"SimpleWeb/configs"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// StringToSHA256 字符串SHA256
func StringToSHA256(input string) string {
	// 创建一个新的 SHA-256 哈希对象
	hash := sha256.New()
	// 写入数据
	input = input + configs.PasswdKey
	hash.Write([]byte(input))
	// 获取哈希结果并转换为十六进制字符串
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

// GetCurrentTime 格式化时间
func GetCurrentTime() string {
	// 获取当前时间
	currentTime := time.Now()
	// 格式化为指定格式
	return currentTime.Format("2006-01-02 15:04:05")
}

// Base64UrlEncode 对数据进行 Base64 URL 编码
func Base64UrlEncode(input []byte) string {
	// 使用 Base64 URL Safe 字符集
	enc := base64.URLEncoding.EncodeToString(input)
	// 移除最后的 `=` 填充字符
	enc = string([]byte(enc)[:len(enc)-len(input)%3])
	return enc
}

// GenerateToken 生成JWT令牌
func GenerateToken(username string) (string, error) {
	// 1. 构建 Header
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}

	// 2. 构建 Payload
	payload := map[string]interface{}{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // 设置过期时间为 3 天
	}

	// 3. 将 Header 和 Payload 转为 JSON 字符串并编码为 Base64
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("header JSON 编码失败: %v", err)
	}
	encodedHeader := Base64UrlEncode(headerJSON)

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("payload JSON 编码失败: %v", err)
	}
	encodedPayload := Base64UrlEncode(payloadJSON)

	// 4. 创建签名
	dataToSign := encodedHeader + "." + encodedPayload
	signature := createSignature(dataToSign, configs.PasswdKey)

	// 5. 拼接 JWT
	token := encodedHeader + "." + encodedPayload + "." + signature

	return token, nil
}

// createSignature 使用 HMAC SHA-256 对数据进行签名
func createSignature(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	signature := h.Sum(nil)
	return Base64UrlEncode(signature)
}
