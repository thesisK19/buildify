package util

import (
	uuid "github.com/satori/go.uuid"
)

func GenUUID(size int) string {
	// create a new UUID
	u := uuid.NewV4()
	return u.String()[0:size]
}
