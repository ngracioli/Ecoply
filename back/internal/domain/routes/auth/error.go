package auth

import "errors"

var (
	ErrFailedToGenerateToken = errors.New("failed to generate token")

	ErrFailedToLogin = errors.New("failed to login")

	ErrInvalidAvailabilityType = errors.New("invalid availability type. Must be 'email', 'cpf' or 'cnpj'")
)
