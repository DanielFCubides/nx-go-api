package usecases

import "nx-go-api/app/account"

type AccountUseCase interface {
	Create(account account.Account) account.Account
	FindAll() []account.Account
	FindByEmail(email string) account.Account
	Edit(account account.Account) account.Account
}
