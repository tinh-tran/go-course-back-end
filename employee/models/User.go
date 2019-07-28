package models

type User struct {
	UserId     int
	UserName   string
	Password   string
	Role       int
	CustomerId *int
}

type Jwt struct {
	UserId     int
	Token      string
	Role       string
	UserName   string
	CustomerId *int
}
