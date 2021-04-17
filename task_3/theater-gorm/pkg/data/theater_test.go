package data

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	data *TheaterData
}

//func (s *Suite) AfterTest(_, _ string) {
//	require.NoError(s.T(), s.mock.ExpectationsWereMet())
//}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.data = NewTheaterData(s.DB)
}

var testTicket = &SelectTicket{
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

var testPoster = &SelectPoster{
	Id:                  2,
	PerformanceName:     "The Dragon",
	GenreName:           "a musical",
	PerformanceDuration: "0000-01-01T04:00:00Z",
	DateTime:            "2021-04-13T16:00:00Z",
	HallName:            "Middle",
	HallCapacity:        1500,
	LocationAddress:     "Gaidara_6",
	LocationPhoneNumber: "+375443564987",
	Comment:             "We invite you! It will be cool!!!",
}

var testUser = &SelectUser{
	Id:                  1,
	FirstName:           "Charles",
	LastName:            "Dean",
	Role:                "Actor",
	LocationAddress:     "Gaidara_6",
	LocationPhoneNumber: "+375443564987",
	PhoneNumber:         "+375445239375",
}

var testAccount = &Account{
	Id:          1,
	FirstName:   "Alexander",
	LastName:    "Antoshkov",
	PhoneNumber: "+37544*******",
	Email:       "aaaa@gmail.com",
}

var testGenre = &Genre{
	Id:   1,
	Name: "a musical",
}

func (s *Suite) TestTheaterData_ReadAllTickets() {
	rows := sqlmock.NewRows([]string{"tickets.id", "performance.name", "genres.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "sectors.name", "places.name", "price.price",
		"tickets.date_of_issue", "tickets.paid", "tickets.reservation", "tickets.destroyed"}).
		AddRow(testTicket.Id, testTicket.PerformanceName, testTicket.GenreName, testTicket.PerformanceDuration,
			testTicket.DateTime, testTicket.HallName, testTicket.HallCapacity, testTicket.LocationAddress, testTicket.LocationPhoneNumber,
			testTicket.SectorName, testTicket.Place, testTicket.Price, testTicket.DateOfIssue, testTicket.Paid, testTicket.Reservation,
			testTicket.Destroyed)
	s.mock.ExpectQuery(`SELECT tickets.id, performance.name, genres.name, performance.duration, schedule.date, halls.name, halls.capacity, 
locations.address, locations.phone_number, sectors.name, places.name, price.price, tickets.date_of_issue, tickets.paid, tickets.reservation, 
tickets.destroyed FROM "tickets" 
JOIN schedule on schedule.id = tickets.schedule_id 
JOIN performance on schedule.performance_id = performance.id 
JOIN genres on performance.genre_id = genres.id JOIN halls on schedule.hall_id = halls.id 
JOIN locations on halls.location_id = locations.id JOIN places on tickets.place_id = places.id 
JOIN sectors on places.sector_id = sectors.id 
JOIN price on performance.id = price.performance_id and sectors.id = price.sector_id`).
		WillReturnRows(rows)
	res, err := s.data.ReadAllTickets()
	fmt.Println(res)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testTicket, &res[0]))
}

func (s *Suite) TestTheaterData_ReadAllTicketsErr() {
	s.mock.ExpectQuery(`SELECT tickets.id, performance.name, genres.name, performance.duration, schedule.date, halls.name, halls.capacity, 
locations.address, locations.phone_number, sectors.name, places.name, price.price, tickets.date_of_issue, tickets.paid, tickets.reservation, 
tickets.destroyed FROM "tickets" 
JOIN schedule on schedule.id = tickets.schedule_id 
JOIN performance on schedule.performance_id = performance.id 
JOIN genres on performance.genre_id = genres.id JOIN halls on schedule.hall_id = halls.id 
JOIN locations on halls.location_id = locations.id JOIN places on tickets.place_id = places.id 
JOIN sectors on places.sector_id = sectors.id 
JOIN price on performance.id = price.performance_id and sectors.id = price.sector_id`).
		WillReturnError(errors.New("something went wrong..."))
	users, err := s.data.ReadAllTickets()
	require.Error(s.T(), err)
	require.Empty(s.T(), users)
}

func (s *Suite) TestTheaterData_ReadAllPosters() {
	rows := sqlmock.NewRows([]string{"poster.id", "performance.name", "genres.name", "performance.duration", "schedule.date",
		"halls.name", "halls.capacity", "locations.address", "locations.phone_number", "poster.comment"}).
		AddRow(testPoster.Id, testPoster.PerformanceName, testPoster.GenreName, testPoster.PerformanceDuration,
			testPoster.DateTime, testPoster.HallName, testPoster.HallCapacity, testPoster.LocationAddress,
			testPoster.LocationPhoneNumber, testPoster.Comment)
	s.mock.ExpectQuery(`SELECT poster.id, performance.name, genres.name, performance.duration, schedule.date, halls.name, 
halls.capacity, locations.address, locations.phone_number, poster.comment FROM "poster" 
JOIN schedule on schedule.id = poster.schedule_id 
JOIN performance on schedule.performance_id = performance.id 
JOIN genres on performance.genre_id = genres.id 
JOIN halls on schedule.hall_id = halls.id 
JOIN locations on halls.location_id = locations.id`).
		WillReturnRows(rows)
	res, err := s.data.ReadAllPosters()
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testPoster, &res[0]))
}

func (s *Suite) TestTheaterData_ReadAllPostersErr() {
	s.mock.ExpectQuery(`SELECT poster.id, performance.name, genres.name, performance.duration, schedule.date, halls.name, 
halls.capacity, locations.address, locations.phone_number, poster.comment FROM "poster" 
JOIN schedule on schedule.id = poster.schedule_id 
JOIN performance on schedule.performance_id = performance.id 
JOIN genres on performance.genre_id = genres.id 
JOIN halls on schedule.hall_id = halls.id 
JOIN locations on halls.location_id = locations.id`).
		WillReturnError(errors.New("something went wrong..."))
	res, err := s.data.ReadAllPosters()
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_ReadAllUsers() {
	rows := sqlmock.NewRows([]string{"users.id", "users.first_name", "users.last_name", "roles.name", "locations.address",
		"locations.phone_number", "users.phone_number"}).
		AddRow(testUser.Id, testUser.FirstName, testUser.LastName, testUser.Role,
			testUser.LocationAddress, testUser.LocationPhoneNumber, testUser.PhoneNumber)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT users.id, users.first_name, users.last_name, roles.name, locations.address, 
locations.phone_number, users.phone_number FROM "users" 
JOIN roles on users.role_id = roles.id 
JOIN locations on locations.id = users.account_id 
WHERE (users.account_id = $1)`)).
		WithArgs(Account{Id: 1}.Id).
		WillReturnRows(rows)
	res, err := s.data.ReadAllUsers(Account{Id: 1})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testUser, &res[0]))
}

func (s *Suite) TestTheaterData_ReadAllUsersErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT users.id, users.first_name, users.last_name, roles.name, locations.address, 
locations.phone_number, users.phone_number FROM "users" 
JOIN roles on users.role_id = roles.id 
JOIN locations on locations.id = users.account_id 
WHERE (users.account_id = $1)`)).
		WithArgs(Account{Id: 1}.Id).
		WillReturnError(errors.New("something went wrong..."))
	res, err := s.data.ReadAllUsers(Account{Id: 1})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdAccount() {
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "phone_number", "email"}).
		AddRow(testAccount.Id, testAccount.FirstName, testAccount.LastName, testAccount.PhoneNumber, testAccount.Email)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" ASC LIMIT 1`)).
		WithArgs(testAccount.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdAccount(Account{Id: 1})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testAccount, &res))
}

func (s *Suite) TestTheaterData_FindByIdAccountErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" ASC LIMIT 1`)).
		WithArgs(testAccount.Id).
		WillReturnError(errors.New("something went wrong..."))
	res, err := s.data.FindByIdAccount(Account{Id: 1})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

func (s *Suite) TestTheaterData_FindByIdGenre() {
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(testGenre.Id, testGenre.Name)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "genres" WHERE "genres"."id" = $1 ORDER BY "genres"."id" ASC LIMIT 1`)).
		WithArgs(testGenre.Id).
		WillReturnRows(rows)
	res, err := s.data.FindByIdGenre(Genre{Id: 1})
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(testGenre, &res))
}

func (s *Suite) TestTheaterData_FindByIdGenreErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "genres" WHERE "genres"."id" = $1 ORDER BY "genres"."id" ASC LIMIT 1`)).
		WithArgs(testAccount.Id).
		WillReturnError(errors.New("something went wrong..."))
	res, err := s.data.FindByIdAccount(Account{Id: 1})
	require.Error(s.T(), err)
	require.Empty(s.T(), res)
}

//func (s *Suite) TestTheaterData_AddAccount() {
//	testAccount := &Account{
//		FirstName:   "Dim",
//		LastName:    "Ivanov",
//		PhoneNumber: "+375296574897",
//		Email:       "dimaivanov@gmail.com",
//	}
//	//s.mock.ExpectBegin()
//	//s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "accounts" ("first_name", "last_name", "phone_number", "email")
//	//VALUES ($1,$2,$3,$4) RETURNING "accounts"."id"`)).
//		s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "accounts" ("first_name", "last_name", "phone_number", "email")
//	VALUES ($1,$2,$3,$4)`)).
//		//WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
//		//s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "accounts" ("first_name", "last_name", "phone_number", "email")
//		//VALUES ($1,$2,$3,$4)`)).
//		//	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO accounts (first_name, last_name, phone_number, email)
//		//VALUES ($1,$2,$3,$4) RETURNING accounts.id`)).
//		//s.mock.ExpectExec(`INSERT INTO "accounts" ("first_name", "last_name", "phone_number", "email")
//		//VALUES ($1,$2,$3,$4)`).
//		//	s.mock.ExpectExec(`INSERT INTO "accounts" ("first_name","last_name","phone_number","email") VALUES ($1,$2,$3,$4) RETURNING "accounts"."id"`).
//		WithArgs(testAccount.FirstName, testAccount.LastName, testAccount.PhoneNumber, testAccount.Email).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testAccount.Id))
//
//	id, err := s.data.AddAccount(*testAccount)
//	fmt.Println(id)
//	require.NoError(s.T(), err)
//}
