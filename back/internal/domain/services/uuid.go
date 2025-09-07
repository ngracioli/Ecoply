package services

import (
	"ecoply/internal/mlog"

	"github.com/google/uuid"
)

func NewUuidV7() uuid.UUID {
	uuid, err := uuid.NewV7()
	if err != nil {
		mlog.Log(err.Error() + ": failed to generate uuid")
		panic(err)
	}
	return uuid
}

func NewUuidV7String() string {
	return NewUuidV7().String()
}
