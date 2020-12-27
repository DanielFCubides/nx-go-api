package gateway

import (
	"nx-go-api/app/account"
)

type AccountRepository interface {
	Create(account account.Account) account.Account
	FindAll() []account.Account
	FindByEmail(email string) account.Account
	Edit(account account.Account) account.Account
}
