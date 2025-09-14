package repository

import "errors"

var (
	// Common database errors
	ErrNotFound = errors.New("not found")
	ErrInternal = errors.New("internal error")

	// User creation errors
	ErrUserEmailAlreadyExists = errors.New("email already exists")
	ErrUserCreationFailed     = errors.New("failed to create user")

	// Agent creation errors
	ErrAgentCnpjAlreadyExists     = errors.New("CNPJ already exists")
	ErrAgentCceeCodeAlreadyExists = errors.New("CCEE code already exists")

	// User authentication errors
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrIncorrectPassword  = errors.New("incorrect password")
)
