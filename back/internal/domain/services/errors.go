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
	ErrAgentAlreadyExists = errors.New("agent already exists")

	// User
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrIncorrectPassword  = errors.New("incorrect password")
	ErrUserAlreadyExists  = errors.New("user already exists")

	// Availability
	ErrInvalidAvailabilityType = errors.New("invalid availability type")

	// JWT
	ErrFailedToGenerateToken = errors.New("failed to generate token")

	// EnergyType
	ErrInvalidEnergyType = errors.New("invalid energy type")

	// Offer
	ErrInvalidPeriodStart        = errors.New("invalid period start")
	ErrInvalidPeriodEnd          = errors.New("invalid period end")
	ErrInvalidPeriod             = errors.New("invalid period")
	ErrOfferNotFound             = errors.New("offer not found")
	ErrInvalidPrice              = errors.New("invalid price")
	ErrInvalidQuantity           = errors.New("invalid quantity")
	ErrUserIsNotTheOfferOwner    = errors.New("user is not the offer owner")
	ErrCannotDeleteOffer         = errors.New("offer can't be deleted")
	ErrInsufficientOfferQuantity = errors.New("insufficient offer quantity")
	ErrCannotPurchaseOwnOffer    = errors.New("cannot purchase own offer")
	ErrCannotUpdateOffer         = errors.New("offer can't be updated")
	ErrOfferHasEnded             = errors.New("offer has ended")

	// Purchase
	ErrUserIsNotThePurchaseOwner = errors.New("user is not the purchase owner")
	ErrPurchaseNotFound          = errors.New("purchase not found")
	ErrPurchaseCannotBeCancelled = errors.New("purchase can not be cancelled")

	// Contract
	ErrUserIsNotContractMember = errors.New("user is not a member of the contract")
	ErrPurchaseIsNotCompleted  = errors.New("purchase is not completed")
)
