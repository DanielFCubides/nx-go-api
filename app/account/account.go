package account

import "time"

type Account struct {
	Email        string
	Password     string
	Status       string
	Username     string
	CreationDate time.Time
}
