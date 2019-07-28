package models

type Order struct {
	OrderId     int
	CourseId    int
	CustomerId  int
	OrderDate   string
	OrderNote   *string
	Status      int
	CoursePrice float64
	CourseName  string
	CourseImage string
}
