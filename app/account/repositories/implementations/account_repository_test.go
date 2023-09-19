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

		gomocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"accounts\" (\"id\",\"created_at\",\"updated_at\",\"deleted_at\",\"email\",\"password\",\"status\",\"username\") VALUES (?,?,?,?,?,?,?,?)")
		result := repository.Create(accountToCreate)
		if result.ID == 0 {
			t.Errorf("Failed to create account")
			return
		}
	})
}

func TestDBAccountRepository_CreateAccount_Failure(t *testing.T) {
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

		gomocket.Catcher.Reset().
			NewMock().
			WithQuery("INSERT INTO \"accounts\" (\"id\",\"created_at\",\"updated_at\",\"deleted_at\",\"email\",\"password\",\"status\",\"username\") VALUES (?,?,?,?,?,?,?,?)").WithExecException()
		result := repository.Create(accountToCreate)
		if result.ID != 0 {
			t.Errorf("Failed to send the correct exception")
			return
		}
	})
}
