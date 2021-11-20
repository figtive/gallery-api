package constants

import "time"

const (
	JWTTimeout = time.Hour * 2

	ContextUserEmailKey       = "user_email"
	ContextIsAuthenticatedKey = "is_authenticated"
	ContextIsAdminKey         = "is_admin"
	ContextGoogleJWTKey       = "google_jwt"
)
