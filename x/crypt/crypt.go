package crypt

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var (
	PWD_INVALID   = errors.New("password invalid")
	TOKEN_INVALID = errors.New("token invalid")
	TOKEN_TIMEOUT = errors.New("token timeout")
	TOKEN_ILLEGAL = errors.New("token illegal")
)

// JWT由三部分组成：
// Header：头部，表明类型和加密算法
// Claims(Payload)：声明，即载荷（承载的内容）
// Signature：签名，这一部分是将header和claims进行base64转码后，并用header中声明的加密算法加密(secre)后构成

func NewToken(key string, m map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for i, v := range m {
		claims[i] = v
	}
	token.Claims = claims
	tokenStr, _ := token.SignedString([]byte(key))
	return tokenStr
}

func ParseToken(tokenStr string, key string) (interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	}
	return nil, err
}

// 判断token是否合法
func ValidateToken(tokenStr string, key string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return err
	}

	s, err := token.SignedString([]byte(key))
	if s != tokenStr {
		return TOKEN_ILLEGAL
	}
	return nil
}

func EncryptPwd(pwd string, salt string) string {
	h := md5.New()
	h.Write([]byte(pwd + salt))
	return hex.EncodeToString(h.Sum(nil))
}
