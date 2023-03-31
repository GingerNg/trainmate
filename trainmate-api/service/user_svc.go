package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

var (
	Secret     = "dong_tech_1" // 加盐
	ExpireTime = 3600          // token有效期
)

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID      int      `json:"user_id"`
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	FullName    string   `json:"full_name"`
	Permissions []string `json:"permissions"`
}

func VerifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	fmt.Println("verify")
	return claims, nil
}

func GetToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	return signedToken, nil
}

func Login(username string, password string) (string, error) {
	// check whether password is correct
	// if correct, generate token
	// if incorrect, return error

	claims := &JWTClaims{
		UserID:      1,
		Username:    username,
		Password:    password,
		FullName:    username,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, err := GetToken(claims)
	return signedToken, err
}
