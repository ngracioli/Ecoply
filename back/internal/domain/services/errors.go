package services

import "errors"

var (
	// Common
	ErrInternal = errors.New("unexpected internal error")

	// User
	ErrUserNotFound           = errors.New("user not found")
	ErrUserEmailAlreadyExists = errors.New("email already exists")
	ErrUserCreationFailed     = errors.New("failed to create user")

	// Submarket
	ErrSubmarketNotFound = errors.New("submarket not found")

	// UserType
	ErrUserTypeNotFound = errors.New("user type not found")

	// Agent
	ErrAgentCnpjAlreadyExists     = errors.New("CNPJ already exists")
	ErrAgentCceeCodeAlreadyExists = errors.New("CCEE code already exists")

	// User
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrIncorrectPassword  = errors.New("incorrect password")

	// Availability
	ErrInvalidAvailabilityType = errors.New("invalid availability type")

	// JWT
	ErrFailedToGenerateToken = errors.New("failed to generate token")
)
