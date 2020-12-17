package implementations

import (
	"fmt"
	"nx-go-api/app"
	"nx-go-api/app/account"
	"nx-go-api/app/account/usecases"
	"time"
)

type AccountUseCase struct{}

func NewAccountUseCase() usecases.AccountUseCase {
	return &AccountUseCase{}
}

func init() {
	err := app.Injector.Provide(NewAccountUseCase)
	if err != nil {
		fmt.Println("Error providing AccountUseCase :", err)
		panic(err)
	}
}

func (AccountUseCase) Create(account account.Account) account.Account {
	return account
}

func (AccountUseCase) FindAll() []account.Account {
	return []account.Account{{"", "", "", "", time.Now()}}
}

func (AccountUseCase) Find(email string) account.Account {
	return account.Account{
		Email: email,
	}
}

func (AccountUseCase) Edit(account account.Account) account.Account {
	return account
}
