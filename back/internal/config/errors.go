package config

import "errors"

var (
	ErrFailedToLoadConfig  = makeError("failed to load environment config")
	ErrFailedToLoadEnvFile = makeError("failed to load env file(s)")
	ErrConfigNotLoaded     = makeError("environment is not loaded")
)

func makeError(message string) error {
	const prefix string = "Config -> "

	return errors.New(prefix + message)
}
