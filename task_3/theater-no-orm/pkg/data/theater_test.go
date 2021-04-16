package data

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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

var testTicket = &Ticket{
	Id:                  21,
	PerformanceName:     "The Dragon",
	GenreName:           "a musical",
	PerformanceDuration: "0000-01-01T04:00:00Z",
	DateTime:            "2021-04-13T16:00:00Z",
	HallName:            "Middle",
	HallCapacity:        1500,
	LocationAddress:     "Gaidara_6",
	LocationPhoneNumber: "+375443564987",
	SectorName:          "A",
	Place:               1,
	Price:               40,
	DateOfIssue:         "2021-04-12T22:48:15.344148Z",
	Paid:                false,
	Reservation:         false,
	Destroyed:           false,
}

func TestTheaterData_ReadAllTickets(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	rows := sqlmock.NewRows([]string{"tickets.id", "performance.name", "genre.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "sectors.name", "places.name", "price.price",
		"tickets.date_of_issue", "tickets.paid", "tickets.reservation", "tickets.destroyed"}).
		AddRow(testTicket.Id, testTicket.PerformanceName, testTicket.GenreName, testTicket.PerformanceDuration,
			testTicket.DateTime, testTicket.HallName, testTicket.HallCapacity, testTicket.LocationAddress, testTicket.LocationPhoneNumber,
			testTicket.SectorName, testTicket.Place, testTicket.Price, testTicket.DateOfIssue, testTicket.Paid, testTicket.Reservation,
			testTicket.Destroyed)
	mock.ExpectQuery(readAllTicketsQuery).WillReturnRows(rows)
	users, err := data.ReadAllTickets()
	assert.NoError(err)
	assert.NotEmpty(users)
	assert.Equal(users[0], *testTicket)
	assert.Len(users, 1)
}

func TestTheaterData_ReadAllTicketsErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	mock.ExpectQuery(readAllTicketsQuery).WillReturnError(errors.New("something went wrong..."))
	users, err := data.ReadAllTickets()
	assert.Error(err)
	assert.Empty(users)
}

func TestTheaterData_ReadAllPosters(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	data := NewTheaterData(db)
	rows := sqlmock.NewRows([]string{"tickets.id", "performance.name", "genre.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "sectors.name", "places.name", "price.price",
		"tickets.date_of_issue", "tickets.paid", "tickets.reservation", "tickets.destroyed"}).
		AddRow(testTicket.Id, testTicket.PerformanceName, testTicket.GenreName, testTicket.PerformanceDuration,
			testTicket.DateTime, testTicket.HallName, testTicket.HallCapacity, testTicket.LocationAddress, testTicket.LocationPhoneNumber,
			testTicket.SectorName, testTicket.Place, testTicket.Price, testTicket.DateOfIssue, testTicket.Paid, testTicket.Reservation,
			testTicket.Destroyed)
	mock.ExpectQuery(readAllTicketsQuery).WillReturnRows(rows)
	users, err := data.ReadAllTickets()
	assert.NoError(err)
	assert.NotEmpty(users)
	assert.Equal(users[0], *testTicket)
	assert.Len(users, 1)
}
