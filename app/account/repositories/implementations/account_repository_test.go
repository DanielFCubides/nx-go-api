package implementations

import (
	gomocket "github.com/selvatico/go-mocket"
	"nx-go-api/app/account"
	sql "nx-go-api/app/configuration/datasources"
	"testing"
	"time"
)

func TestDBAccountRepository_CreateAccount(t *testing.T) {
	repository := DBAccountRepository{db: sql.NewMockConnection().GetDatabase()}
	t.Run("Create an account", func(t *testing.T) {
		accountToCreate := account.Account{
			ID:           1,
			Email:        "a.b@c.com",
			Password:     "123345",
			Status:       "ACTIVE",
			Username:     "a.b",
			CreationDate: time.Time{},
		}
		gomocket.Catcher.Logging = true

		gomocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"accounts\" ")
		result := repository.Create(accountToCreate)
		if result.ID == 0 {
			t.Errorf("Failed to create account")
			return
		}
	})
}

func TestDBAccountRepository_CreateAccount_Failure(t *testing.T) {
	repository := DBAccountRepository{db: sql.NewMockConnection().GetDatabase()}
	t.Run("Fail to create an account", func(t *testing.T) {
		accountToCreate := account.Account{
			ID:           1,
			Email:        "a.b@c.com",
			Password:     "123345",
			Status:       "ACTIVE",
			Username:     "a.b",
			CreationDate: time.Time{},
		}

		gomocket.Catcher.Reset().
			NewMock().
			WithQuery("INSERT INTO \"accounts\" ").WithExecException()
		result := repository.Create(accountToCreate)
		if result.ID != 0 {
			t.Errorf("Failed to send the correct exception")
			return
		}
	})
}

func TestDBAccountRepository_FindAll_Success(t *testing.T) {
	repository := DBAccountRepository{db: sql.NewMockConnection().GetDatabase()}
	t.Run("Get all the accounts", func(t *testing.T) {
		result := []map[string]interface{}{{"id": 1}}
		gomocket.Catcher.Reset().NewMock().WithQuery("SELECT * FROM \"accounts\"").WithReply(result)
		accounts := repository.FindAll()
		if len(accounts) != 1 {
			t.Errorf("List lenght is %d and is different from 1", len(accounts))
		}
	})
}

func TestDBAccountRepository_FindAll_Failure(t *testing.T) {
	repository := DBAccountRepository{db: sql.NewMockConnection().GetDatabase()}
	t.Run("Get all the accounts", func(t *testing.T) {
		gomocket.Catcher.Reset().NewMock().WithQuery("SELECT * FROM \"accounts\"").WithExecException()
		accounts := repository.FindAll()
		if len(accounts) != 0 {
			t.Errorf("List lenght is %d and is different from 0", len(accounts))
		}
	})
}

func TestDBAccountRepository_FindByEmail(t *testing.T) {

}
