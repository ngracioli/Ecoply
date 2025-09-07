package auth

import "errors"

var (
	ErrFailedToGenerateToken = errors.New("failed to generate token")

	ErrFailedToLogin = errors.New("failed to login")
)
