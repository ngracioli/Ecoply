package config

import "errors"

var (
	ErrFailedToLoadConfig  = makeError("Failed to load environment config")
	ErrFailedToLoadEnvFile = makeError("Failed to load env file(s)")
	ErrConfigNotLoaded     = makeError("Environment is not loaded")
)

func makeError(message string) error {
	const prefix string = "Config : "

	return errors.New(prefix + message)
}
