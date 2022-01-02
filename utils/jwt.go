package utils

import (
	"bytes"
	"ecommerce-v1/auth/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

func GenerateRefreshToken() (string, error) {
	expires, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRE"))
	secret := os.Getenv("REFRESH_TOKEN_SECRET")
	claims := jwt.MapClaims{}

	claims["jti"] = jti()
	claims["exp"] = time.Now().Add((time.Hour * 24) * time.Duration(expires)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func DecodeRefreshToken(refreshToken string) (*models.RefreshTokenMetadata, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &models.RefreshTokenMetadata{
			Jti:     claims["jti"].(string),
			Expires: int64(claims["exp"].(float64)),
			Iat:     int64(claims["iat"].(float64)),
			Sub:     claims["sub"].(string),
			Nbf:     int64(claims["nbf"].(float64)),
		}, nil
	}

	return nil, err
}

func jti() string {
	t := time.Now()
	entropy := ulid.Monotonic(bytes.NewReader([]byte(uuid.New().String())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
