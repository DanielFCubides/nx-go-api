package gateway

import (
	"nx-go-api/app"
	sql "nx-go-api/infrastructure/datasources"
	"time"
)

type Account struct {
	ID        uint `gorm:"primary_key;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Email     string     `gorm:"unique;not null"`
	Password  string
	Status    string
	Username  string `gorm:"unique;not null"`
}

func Migrate() {
	var conn sql.Connection
	invokeFunc := func(connection sql.Connection) {
		conn = connection
	}
	err := app.Injector.Invoke(invokeFunc)
	if err != nil {
		panic(err)
	}

	db := conn.GetDatabase()

	db.AutoMigrate(Account{})
}
