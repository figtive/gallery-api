package dtos

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type GoogleJWTClaim struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Subject  string `json:"sub"`
	Expiry   int64  `json:"exp"`
	IssuedAt int64  `json:"iat"`
}

func (c GoogleJWTClaim) Valid() (err error) {
	now := time.Now()
	if c.Name == "" {
		return errors.New("empty name")
	}
	if c.Email == "" {
		return errors.New("empty email")
	}
	if now.After(time.Unix(c.Expiry, 0)) {
		return errors.New("token expired")
	}
	if now.Before(time.Unix(c.IssuedAt, 0)) {
		return errors.New("token used before issued")
	}
	return
}

type GalleryJWTClaim struct {
	jwt.RegisteredClaims
	Name  string
	Email string
}

func (c GalleryJWTClaim) Valid() (err error) {
	if c.Name == "" {
		return errors.New("empty name")
	}
	if c.Email == "" {
		return errors.New("empty email")
	}
	return c.RegisteredClaims.Valid()
}
