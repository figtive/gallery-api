package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/constants"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/dtos"
)

func (m *module) AuthParseGoogleJWT(jwtString string) (claims dtos.GoogleJWTClaim, err error) {
	var token *jwt.Token
	parser := jwt.Parser{}
	if token, _, err = parser.ParseUnverified(jwtString, &claims); err != nil || !token.Valid {
		return
	}
	// TODO: JWK
	return
}

func (m *module) AuthGenerateJWT(userInfo dtos.User) (token string, err error) {
	now := time.Now()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, dtos.GalleryJWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(constants.JWTTimeout)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "gallery-api",
			Subject:   userInfo.ID,
		},
		Name:  userInfo.Name,
		Email: userInfo.Email,
	})

	return jwtToken.SignedString([]byte(configs.AppConfig.Secret))
}
