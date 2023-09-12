package usecase

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"nx-go-api/app"
	"nx-go-api/app/account/repositories"
	"nx-go-api/app/authentication/token"
	"os"
)

const activeStatus = "active"

type DBAuthUseCase struct {
	accountRepo repositories.AccountRepository
}

func (c DBAuthUseCase) GenerateToken(email string, status string) (string, error) {
	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Errorf("{secret: NOT SET }")
		return "", errors.New("no secret")
	}
	log.Debugf("{secret=%s }", secret)
	return token.GenerateToken(email, status, secret)
}

func (c DBAuthUseCase) Authenticate(email string, password string) bool {
	account := c.accountRepo.FindByEmail(email)
	log.Debugf("{email: %s password:%s}", account.Email, account.Password)
	if account.Email != email || account.Password != password {
		log.Infof("{authentication: unsucessfull }")
		return false
	}

	log.Debugf("{\"status\": \"%s\" }", account.Status)
	if account.Status != activeStatus {
		log.Infof("{\"authentication\": \"unsucessfull\" }")
		return false
	}
	log.Infof("{\"authentication\": \"sucessfull\" }")
	return true
}

func New(repository repositories.AccountRepository) AuthUseCase {
	return &DBAuthUseCase{
		accountRepo: repository,
	}
}

func init() {
	err := app.Injector.Provide(New)
	if err != nil {
		fmt.Println("Error providing DBAuthUseCase :", err)
		panic(err)

	}
}
