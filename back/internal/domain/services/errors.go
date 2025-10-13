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
	ErrInvalidSubmarket = errors.New("invalid submarket")

	// UserType
	ErrInvalidUserType = errors.New("invalid user type")

	// Agent
	ErrAgentAlreadyExists = errors.New("Agent already exists")

	// User
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrIncorrectPassword  = errors.New("incorrect password")

	// Availability
	ErrInvalidAvailabilityType = errors.New("invalid availability type")

	// JWT
	ErrFailedToGenerateToken = errors.New("failed to generate token")

	// EnergyType
	ErrInvalidEnergyType = errors.New("invalid energy type")

	// Offer
	ErrInvalidPeriodStart = errors.New("invalid period start")
	ErrInvalidPeriodEnd   = errors.New("invalid period end")
	ErrInvalidPeriod      = errors.New("invalid period")
	ErrOfferNotFound      = errors.New("offer not found")
)
