package models

type Student struct {
	CustomerId   int
	CustomerName string
	Phone        string
	Email        string
	Address      string
	CourseJoined *int
}

type StudentCourse struct {
	CustomerId   int
	CustomerName string
	Phone        string
	Email        string
	Address      string
	DisplayName  string
	CourseId     int
	Status       int
}
