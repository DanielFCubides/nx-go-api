package implementations

import (
	"fmt"
	"nx-go-api/app"
	"nx-go-api/app/account"
	"nx-go-api/app/account/gateway"
	"nx-go-api/app/account/usecases"
)

type AccountUseCase struct {
	AccountRepository gateway.AccountRepository
}

func NewAccountUseCase(AccountRepository gateway.AccountRepository) usecases.AccountUseCase {
	return &AccountUseCase{AccountRepository: AccountRepository}
}

func init() {
	err := app.Injector.Provide(NewAccountUseCase)
	if err != nil {
		fmt.Println("Error providing AccountUseCase :", err)
		panic(err)

	}
}

func (uc *AccountUseCase) Create(account account.Account) account.Account {
	return uc.AccountRepository.Create(account)
}

func (uc *AccountUseCase) FindAll() []account.Account {
	return uc.AccountRepository.FindAll()
}

func (uc *AccountUseCase) FindByEmail(email string) account.Account {
	return uc.AccountRepository.FindByEmail(email)
}

func (uc *AccountUseCase) Edit(acc account.Account) account.Account {
	return uc.AccountRepository.Edit(acc)
}
