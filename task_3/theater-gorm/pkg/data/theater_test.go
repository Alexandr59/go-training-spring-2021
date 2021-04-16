package data

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	//"gorm.io/gorm"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

func TestTheaterData_ReadAllTickets(t *testing.T) {
	//assert := assert.New(t)
	//db, mock := NewMock()
	//data := NewTheaterData(db.)

}
