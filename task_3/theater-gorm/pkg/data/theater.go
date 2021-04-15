package data

import (
	"fmt"
	"gorm.io/gorm"
)

type Account struct {
	Id          int    `gorm:"primaryKey"`
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	PhoneNumber string `gorm:"phone_number"`
	Email       string `gorm:"email"`
}

type SelectTicket struct { /////?????
	Id                  int    `gorm:"primaryKey"`
	PerformanceName     string `gorm:"p.name"`
	GenreName           string `gorm:"g.name"`
	PerformanceDuration string `gorm:"p.duration"`
	DateTime            string `gorm:"s.date"`
	HallName            string `gorm:"h.name"`
	HallCapacity        int    `gorm:"h.capacity"`
	LocationAddress     string `gorm:"l.address"`
	LocationPhoneNumber string `gorm:"l.phone_number"`
	SectorName          string `gorm:"s2.name"`
	Place               int    `gorm:"p2.name"`
	Price               int    `gorm:"p3.price"`
	DateOfIssue         string `gorm:"date_of_issue"`
	Paid                bool   `gorm:"paid"`
	Reservation         bool   `gorm:"reservation"`
	Destroyed           bool   `gorm:"destroyed"`
}

//type SelectTicket struct { /////?????
//	Id                  int    `gorm:"primaryKey"`
//	PerformanceName     string `gorm:"name"`
//	GenreName           string `gorm:"name"`
//	PerformanceDuration string `gorm:"duration"`
//	DateTime            string `gorm:"date"`
//	HallName            string `gorm:"name"`
//	HallCapacity        int    `gorm:"capacity"`
//	LocationAddress     string `gorm:"address"`
//	LocationPhoneNumber string `gorm:"phone_number"`
//	SectorName          string `gorm:"name"`
//	Place               int    `gorm:"name"`
//	Price               int    `gorm:"price"`
//	DateOfIssue         string `gorm:"date_of_issue"`
//	Paid                bool   `gorm:"paid"`
//	Reservation         bool   `gorm:"reservation"`
//	Destroyed           bool   `gorm:"destroyed"`
//}

type Ticket struct {
	Id          int    `gorm:"primaryKey"`
	AccountId   int    `gorm:"account_id"`
	ScheduleId  int    `gorm:"schedule_id"`
	PlaceId     int    `gorm:"place_id"`
	DateOfIssue string `gorm:"date_of_issue"`
	Paid        bool   `gorm:"paid"`
	Reservation bool   `gorm:"reservation"`
	Destroyed   bool   `gorm:"destroyed"`
}

type SelectPoster struct { /////////?????
	Id                  int
	PerformanceName     string
	GenreName           string
	PerformanceDuration string
	DateTime            string
	HallName            string
	HallCapacity        int
	LocationAddress     string
	LocationPhoneNumber string
	Comment             string
}

type SelectUser struct { ////////////////?????
	Id                  int    `gorm:"primaryKey"`
	FirstName           string `gorm:"first_name"`
	LastName            string `gorm:"last_name"`
	Role                string `gorm:"role"`
	LocationAddress     string `gorm:"location_address"`
	LocationPhoneNumber string `gorm:"location"`
	PhoneNumber         string `gorm:"phone_number"`
}

type User struct {
	Id          int    `gorm:"primaryKey"`
	AccountId   int    `gorm:"account_id"`
	FirstName   string `gorm:"first_name"`
	LastName    string `gorm:"last_name"`
	RoleId      int    `gorm:"role_id"`
	LocationId  int    `gorm:"location_id"`
	PhoneNumber string `gorm:"phone_number"`
}

type Hall struct {
	Id         int    `gorm:"primaryKey"`
	AccountId  int    `gorm:"account_id"`
	Name       string `gorm:"name"`
	Capacity   int    `gorm:"capacity"`
	LocationId int    `gorm:"location_id"`
}

type Location struct {
	Id          int    `gorm:"primaryKey"`
	AccountId   int    `gorm:"account_id"`
	Address     string `gorm:"address"`
	PhoneNumber string `gorm:"phone_number"`
}

type Performance struct {
	Id        int    `gorm:"primaryKey"`
	AccountId int    `gorm:"account_id"`
	Name      string `gorm:"name"`
	GenreId   int    `gorm:"genre_id"`
	Duration  string `gorm:"duration"`
}

type Place struct {
	id       int    `gorm:"primaryKey"`
	SectorId int    `gorm:"sector_id"`
	Name     string `gorm:"name"`
}

type Poster struct {
	Id         int    `gorm:"primaryKey"`
	AccountId  int    `gorm:"account_id"`
	ScheduleId int    `gorm:"schedule_id"`
	Comment    string `gorm:"comment"`
}

type Price struct {
	Id            int `gorm:"primaryKey"`
	AccountId     int `gorm:"account_id"`
	SectorId      int `gorm:"sector_id"`
	PerformanceId int `gorm:"performance_id"`
	Price         int `gorm:"price"`
}

type Role struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

type Schedule struct {
	Id            int    `gorm:"primaryKey"`
	AccountId     int    `gorm:"account_id"`
	PerformanceId int    `gorm:"performance_id"`
	Date          string `gorm:"date"`
	HallId        int    `gorm:"hall_id"`
}

type Sector struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

type Genre struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"name"`
}

type TheaterData struct {
	db *gorm.DB
}

func NewTheaterData(db *gorm.DB) *TheaterData {
	return &TheaterData{db: db}
}

func (u TheaterData) ReadAllTickets() ([]Ticket, error) {
	var tickets []Ticket
	//result := u.db.Model(&Ticket{}).Select("tickets.id, p.name, g.name, p.duration, s.date, h.name, h.capacity, l.address, " +
	//	"l.phone_number, s2.name, p2.name, p3.price, tickets.date_of_issue, tickets.paid, tickets.reservation, " +
	//	"tickets.destroyed").Joins("JOIN schedule s on s.id = tickets.schedule_id").Joins("JOIN " +
	//	"performance p on s.performance_id = p.id ").Joins("JOIN genre " +
	//	"g on p.genre_id = g.id ").Joins("JOIN halls h on s.hall_id = h.id").Joins("JOIN " +
	//	"locations l on h.location_id = l.id").Joins("JOIN places p2 on tickets.place_id = " +
	//	"p2.id").Joins("JOIN sectors s2 on p2.sector_id = s2.id").Joins("JOIN price p3 on p.id = " +
	//	"p3.performance_id and s2.id = p3.sector_id").Scan(&tickets)

	//u.db = u.db.Joins("JOIN schedule s on s.id = tickets.schedule_id")
	//u.db = u.db.Joins("JOIN performance p on s.performance_id = p.id")
	//u.db = u.db.Joins("JOIN genre g on p.genre_id = g.id")
	//u.db = u.db.Joins("JOIN halls h on s.hall_id = h.id")
	//u.db = u.db.Joins("JOIN locations l on h.location_id = l.id")
	//u.db = u.db.Joins("JOIN places p2 on tickets.place_id = p2.id")
	//u.db = u.db.Joins("JOIN sectors s2 on p2.sector_id = s2.id")
	//u.db = u.db.Joins("JOIN price p3 on p.id = p3.performance_id and s2.id = p3.sector_id")
	//u.db = u.db.Select("tickets.id, p.name, g.name, p.duration, s.date, h.name, h.capacity, l.address, l.phone_number, s2.name, p2.name, p3.price, tickets.date_of_issue, tickets.paid, tickets.reservation, tickets.destroyed")
	//u.db = u.db.Table("tickets").Find(&tickets)

	result := u.db.Find(&tickets)

	//result := u.db.Preload("performance").Find(&tickets)

	if result.Error != nil {
		return nil, fmt.Errorf("can't read users from database, error: %w", result.Error)
	}
	return tickets, nil
}

//func (u TheaterData) ReadAllPosters() ([]Poster, error) {
//	var posters []Poster
//	rows, err := u.db.Query(readAllPostersQuery)
//	if err != nil {
//		return nil, fmt.Errorf("can't get posters from database, error:%w", err)
//	}
//	for rows.Next() {
//		var temp Poster
//		err = rows.Scan(&temp.Id, &temp.PerformanceName, &temp.GenreName, &temp.PerformanceDuration,
//			&temp.DateTime, &temp.HallName, &temp.HallCapacity, &temp.LocationAddress, &temp.LocationPhoneNumber,
//			&temp.Comment)
//		if err != nil {
//			return nil, fmt.Errorf("can't scan posters from database, error:%w", err)
//		}
//		posters = append(posters, temp)
//	}
//	return posters, nil
//}
//
//func (u TheaterData) ReadAllUsers(account Account) ([]User, error) {
//	var users []User
//	rows, err := u.db.Query(readAllUsersQuery, account.Id)
//	if err != nil {
//		return nil, fmt.Errorf("can't get users from database, error:%w", err)
//	}
//	for rows.Next() {
//		var temp User
//		err = rows.Scan(&temp.Id, &temp.FirstName, &temp.LastName, &temp.Role,
//			&temp.LocationAddress, &temp.LocationPhoneNumber, &temp.PhoneNumber)
//		if err != nil {
//			return nil, fmt.Errorf("can't scan users from database, error:%w", err)
//		}
//		users = append(users, temp)
//	}
//	return users, nil
//}

func (u TheaterData) AddAccount(account Account) error {
	result := u.db.Create(&account)
	if result.Error != nil {
		return fmt.Errorf("can't inser account to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddGenre(genre Genre) error {
	result := u.db.Create(&genre)
	if result.Error != nil {
		return fmt.Errorf("can't inser genre to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddHall(hall Hall) error {
	result := u.db.Create(&hall)
	if result.Error != nil {
		return fmt.Errorf("can't inser hall to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddLocation(location Location) error {
	result := u.db.Create(&location)
	if result.Error != nil {
		return fmt.Errorf("can't inser location to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddPerformance(performance Performance) error {
	result := u.db.Create(&performance)
	if result.Error != nil {
		return fmt.Errorf("can't inser Performance to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddPlace(place Place) error {
	result := u.db.Create(&place)
	if result.Error != nil {
		return fmt.Errorf("can't inser Place to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddPoster(poster Poster) error {
	result := u.db.Create(&poster)
	if result.Error != nil {
		return fmt.Errorf("can't inser Poster to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddPrice(price Price) error {
	result := u.db.Create(&price)
	if result.Error != nil {
		return fmt.Errorf("can't inser Price to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddRole(role Role) error {
	result := u.db.Create(&role)
	if result.Error != nil {
		return fmt.Errorf("can't inser Role to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddSchedule(schedule Schedule) error {
	result := u.db.Create(&schedule)
	if result.Error != nil {
		return fmt.Errorf("can't inser Schedule to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddSector(sector Sector) error {
	result := u.db.Create(&sector)
	if result.Error != nil {
		return fmt.Errorf("can't inser Sector to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddTicket(ticket Ticket) error {
	result := u.db.Create(&ticket)
	if result.Error != nil {
		return fmt.Errorf("can't inser Ticket to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) AddUser(user User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("can't inser User to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteAccount(entry Account) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Account to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteGenre(entry Genre) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Genre to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteHall(entry Hall) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Hall to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteLocation(entry Location) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Location to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeletePerformance(entry Performance) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Performance to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeletePlace(entry Place) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Place to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeletePoster(entry Poster) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Poster to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeletePrice(entry Price) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Price to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteRole(entry Role) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Role to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteSchedule(entry Schedule) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Schedule to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteSector(entry Sector) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Sector to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteTicket(entry Ticket) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete Ticket to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) DeleteUser(entry User) error {
	result := u.db.Delete(&entry)
	if result.Error != nil {
		return fmt.Errorf("can't Delete User to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateAccount(account Account) error {
	result := u.db.Model(&account).Updates(account)
	if result.Error != nil {
		return fmt.Errorf("can't update account to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateGenre(genre Genre) error {
	result := u.db.Model(&genre).Updates(genre)
	if result.Error != nil {
		return fmt.Errorf("can't update genre to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateHall(hall Hall) error {
	result := u.db.Model(&hall).Updates(hall)
	if result.Error != nil {
		return fmt.Errorf("can't update hall to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateLocation(location Location) error {
	result := u.db.Model(&location).Updates(location)
	if result.Error != nil {
		return fmt.Errorf("can't update location to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdatePerformance(performance Performance) error {
	result := u.db.Model(&performance).Updates(performance)
	if result.Error != nil {
		return fmt.Errorf("can't update Performance to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdatePlace(place Place) error {
	result := u.db.Model(&place).Updates(place)
	if result.Error != nil {
		return fmt.Errorf("can't update Place to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdatePoster(poster Poster) error {
	result := u.db.Model(&poster).Updates(poster)
	if result.Error != nil {
		return fmt.Errorf("can't update Poster to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdatePrice(price Price) error {
	result := u.db.Model(&price).Updates(price)
	if result.Error != nil {
		return fmt.Errorf("can't update Price to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateRole(role Role) error {
	result := u.db.Model(&role).Updates(role)
	if result.Error != nil {
		return fmt.Errorf("can't update Role to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateSchedule(schedule Schedule) error {
	result := u.db.Model(&schedule).Updates(schedule)
	if result.Error != nil {
		return fmt.Errorf("can't update Schedule to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateSector(sector Sector) error {
	result := u.db.Model(&sector).Updates(sector)
	if result.Error != nil {
		return fmt.Errorf("can't update Sector to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateTicket(ticket Ticket) error {
	result := u.db.Model(&ticket).Updates(ticket)
	if result.Error != nil {
		return fmt.Errorf("can't update Ticket to database, error: %w", result.Error)
	}
	return nil
}

func (u TheaterData) UpdateUser(user User) error {
	result := u.db.Model(&user).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("can't update User to database, error: %w", result.Error)
	}
	return nil
}
