package constants

import "time"

const (
	JWTTimeout = time.Hour * 2

	ContextUserEmailKey       = "user_email"
	ContextIsAuthenticatedKey = "is_authenticated"
	ContextIsAdminKey         = "is_admin"
	ContextIsFirstLoginKey    = "is_first_login"
	ContextGoogleJWTKey       = "google_jwt"
)
