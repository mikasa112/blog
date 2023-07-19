package util

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"time"
	"v1/pkg"
)

func MD5Encode(data string) (string, error) {
	h := md5.New()
	_, err := io.WriteString(h, data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// GenerateToken 生成token
func GenerateToken(appKey, appSecret string) (token string, err error) {
	now := time.Now()
	expireTime := now.Add(pkg.Js.Expire)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":       expireTime,
		"iss":       pkg.Js.Issuer,
		"appKey":    MD5Encode(appKey),
		"appSecret": MD5Encode(appSecret),
	})
	return claims.SignedString([]byte(pkg.Js.Secret))
}

// ParseToken 解析token
func ParseToken(token string) {

}
