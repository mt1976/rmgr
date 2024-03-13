package support

import (
	"log"

	"github.com/google/uuid"
)

func NewUUID() string {
	uuid, err := uuid.NewV6()
	if err != nil {
		log.Panicln(err.Error())
	}
	return uuid.String()
}
