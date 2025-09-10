package pkg

import (
	"billing-app/internal/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DBAddr string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func NewDatabase(addr, port, user, pass, name string) *Database {
	return &Database{
		DBAddr: addr,
		DBPort: port,
		DBUser: user,
		DBPass: pass,
		DBName: name,
	}
}

func (db *Database) Connect() *gorm.DB {
	address := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		db.DBUser, db.DBPass, db.DBAddr, db.DBPort, db.DBName)

	database, err := gorm.Open(postgres.Open(address), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(model.Orders{})
	if err != nil {
		panic(err)
	}

	return database

}
