package models

import "time"

//User is the user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreateAt    time.Time
	UpdateAt    time.Time
}

//Room is the room model
type Room struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Restriction is the restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreateAt        time.Time
	UpdateAt        time.Time
}

//Reservation is the reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreateAt  time.Time
	UpdateAt  time.Time
	Room      Room
}

//RoomRestriction is the rooms restriction model
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreateAt      time.Time
	UpdateAt      time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}