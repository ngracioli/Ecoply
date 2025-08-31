package database

import "errors"

var (
	ErrFailedToOpenConnectionPostgres = makeError("Failed to open connection with postgres")
)

func makeError(message string) error {
	const prefix string = "Database -> "

	return errors.New(prefix + message)
}
