package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("got an error when tried to make connection with database:%w", err)
	}
	//connection.AutoMigrate(data.Account{}, data.TicketInsert{}, data.UserInsert{}, data.Hall{}, data.Location{},
	//	data.Performance{}, data.Place{}, data.PosterInsert{}, data.Price{}, data.Role{}, data.Schedule{},
	//	data.Sector{}, data.Genre{}, data.Ticket{}, data.Poster{}, data.User{})
	return connection, nil
}
