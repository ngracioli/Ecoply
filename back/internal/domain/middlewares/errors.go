package middlewares

import "errors"

var (
	ErrJwtAuthorizationHeaderRequired = errors.New("authorization header required")
	ErrJwtBearerTokenRequired         = errors.New("bearer token required")
	ErrJwtInvalidToken                = errors.New("invalid token")
	ErrMissingClaim                   = errors.New("missing claim")
	ErrUserIsNotSupplier              = errors.New("user is not supplier")
)
