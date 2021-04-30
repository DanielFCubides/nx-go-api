package implementations

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
	"nx-go-api/app"
	"nx-go-api/app/account"
	"nx-go-api/app/account/repositories"
	sql "nx-go-api/infrastructure/datasources"
	"time"
)

type DBAccountRepository struct {
	db *gorm.DB
}

func NewDBAccountRepository(conn sql.Connection) repositories.AccountRepository {
	db := conn.GetDatabase()
	return &DBAccountRepository{db: db}
}

func init() {
	err := app.Injector.Provide(NewDBAccountRepository)
	if err != nil {
		fmt.Println("Error providing AccountRepository :", err)
		panic(err)
	}
}

func (repo *DBAccountRepository) Create(accountToCreate account.Account) account.Account {
	acc := toEntity(accountToCreate)
	err := repo.db.Save(&acc)
	if err.Error != nil {
		return account.Account{}
	}
	return accountToCreate
}

func (repo *DBAccountRepository) FindAll() []account.Account {
	var accList []repositories.Account
	result := repo.db.Find(&accList)
	if result.Error != nil {
		return []account.Account{}
	}
	return funk.Map(accList, toDomain).([]account.Account)
}

func (repo *DBAccountRepository) FindByEmail(email string) account.Account {
	var acc repositories.Account
	err := repo.db.Where("email = ?", email).First(&acc).Error
	if err != nil {
		return account.Account{}
	}
	return toDomain(acc)
}

func (repo *DBAccountRepository) Edit(accountToEdit account.Account) account.Account {
	var accToUpdate repositories.Account
	if err := repo.db.Where("email = ?", accountToEdit.Email).First(&accToUpdate).Error; err != nil {
		return account.Account{}
	}
	acc := toEntity(accountToEdit)
	acc.ID = accToUpdate.ID
	if err := repo.db.Save(&acc); err.Error != nil {
		return account.Account{}
	}
	return toDomain(acc)

}

func toDomain(acc repositories.Account) account.Account {
	return account.Account{
		ID:           acc.ID,
		Email:        acc.Email,
		Password:     acc.Password,
		Status:       acc.Status,
		Username:     acc.Username,
		CreationDate: acc.CreatedAt,
	}
}

func toEntity(a account.Account) repositories.Account {
	return repositories.Account{
		ID:        a.ID,
		CreatedAt: a.CreationDate,
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		Email:     a.Email,
		Password:  a.Password,
		Status:    a.Status,
		Username:  a.Username,
	}
}
