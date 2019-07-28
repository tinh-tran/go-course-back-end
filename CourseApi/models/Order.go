package models

type Order struct {
	OrderId   *int
	CourseId  *int
	StudentID *int
	OrderDate *string
	OrderNote *string
	Status    *int
}
