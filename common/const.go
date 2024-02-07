package common

import (
	"errors"
	"log"
)

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

const (
	CurrentUser = "user"
)

var (
	RecordNotFound = errors.New("record not found")
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error", err)
	}
}
