package account

import "time"

type Account struct {
	id           int64
	email        string
	password     string
	status       string
	creationDate time.Time
}
