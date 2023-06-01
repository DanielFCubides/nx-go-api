package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"nx-go-api/app"
	"os"
	"time"
)

func init() {
	err := app.Injector.Provide(NewMySQLConnection)
	if err != nil {
		log.Println("Error providing  MySQL connection:", err)
	}
}

// getURL retrieves the URL to connection to SQL database.
func getURL(params ...string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", params[0], params[1], params[2], params[3], params[4])
}

type MySQLConnection struct {
	db *gorm.DB
}

// NewMySQLConnection retrieves a sql connection to MySQL server
func NewMySQLConnection() (Connection, error) {
	dbUsername := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_USER_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := getURL(dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	log.Println(url)
	db, err := gorm.Open("mysql", url)
	i := 0
	for {
		if err != nil {
			db, err = gorm.Open("mysql", url)
		}
		if i >= 10 || err == nil {
			break
		}
		log.Printf("Iteration %d\n", i+1)
		time.Sleep(1 * time.Second)
		i++
	}

	db.LogMode(true)

	if err != nil {
		return nil, err
	}

	return &MySQLConnection{db: db}, nil
}

func (c *MySQLConnection) GetDatabase() *gorm.DB {
	return c.db
}
