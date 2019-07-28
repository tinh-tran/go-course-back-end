package models

import (
	"time"
)

type Course struct {
	CourseId          int
	CourseName        string
	CategoryID        int
	CourseDescription string
	CoursePrice       float64
	CourseImage       string
	StartDate         time.Time
	EndDate           time.Time
	SlotAvailable     int
	SlotRegistered    int
	CategoryName      string
	Page              int
	StudentJoined     *int
	CreateId          int
	Status            bool
}
