package account

import "time"

type Account struct {
	ID           uint
	Email        string
	Password     string
	Status       string
	Username     string
	CreationDate time.Time
}
