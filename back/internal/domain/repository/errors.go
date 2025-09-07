package repository

import "errors"

var (
	// Common database errors
	ErrNotFound = errors.New("not found")
	ErrInternal = errors.New("internal error")

	// User creation errors
	ErrUserEmailAlreadyExists = errors.New("email already exists")
	ErrUserCnpjAlreadyExists  = errors.New("CNPJ already exists")
	ErrUserCreationFailed     = errors.New("failed to create user")

	// User authentication errors
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrIncorrectPassword  = errors.New("incorrect password")
)
