package account

import "time"

const StatusActive = "ACTIVE"
const StatusInactive = "INACTIVE"

type Account struct {
	ID           uint
	Email        string
	Password     string
	Status       string
	Username     string
	CreationDate time.Time
}
