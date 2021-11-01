package constants

import "time"

const (
	JWTTimeout = time.Hour * 2

	ContextUserEmailKey       = "user_email"
	ContextIsAuthenticatedKey = "is_authenticated"
	ContextGoogleJWTKey       = "google_jwt"
)
