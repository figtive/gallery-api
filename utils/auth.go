package utils

import (
	"context"
	"time"

	"google.golang.org/api/idtoken"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
)

func ValidateGoogleJWT(jwtString string) (valid bool, claims dtos.GoogleJWTClaim, err error) {
	var token *idtoken.Payload
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if token, err = idtoken.Validate(ctx, jwtString, configs.AppConfig.GoogleClientID); err != nil {
		return false, dtos.GoogleJWTClaim{}, err
	}

	claims = dtos.GoogleJWTClaim{
		Name:     token.Claims["name"].(string),
		Email:    token.Claims["email"].(string),
		Subject:  token.Subject,
		Expiry:   token.Expires,
		IssuedAt: token.IssuedAt,
	}

	return true, claims, nil
}
