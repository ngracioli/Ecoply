package middlewares

import "errors"

var (
	ErrJwtAuthorizationHeaderRequired = errors.New("authorization header required")
	ErrJwtBearerTokenRequired         = errors.New("bearer token required")
	ErrJwtInvalidToken                = errors.New("invalid token")
)
