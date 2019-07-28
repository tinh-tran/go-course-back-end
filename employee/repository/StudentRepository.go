package repository

import (
	"database/sql"
	. "employee/models"
	"fmt" //. "employee/common"

	log "github.com/Sirupsen/logrus"
)

type DBaction interface {
	GetStudentByID(nID string) Student
	GetAllStudent() []Student
	GetStudentByCourseID(nID int) []StudentCourse
	InsertStudent(student Student)
	DeleteStudent(nID string)
}

type StudentAPI struct {
	DBcon *sql.DB
}

func (api StudentAPI) GetAllStudent() []Student {
	query := "select s.*,A.joinedCourses from customers s left join ( " +
		"SELECT customerid,count(courseid) as joinedCourses FROM " +
		" orders group by studentid) A" +
		" on  s.customerid = A.customerid "
	selDB, err := api.DBcon.Query(query)
	if err != nil {
		log.Error(err.Error())
	}

	var listStudent = make([]Student, 0)
	for selDB.Next() {
		student := Student{}
		var studentid int
		var studentname, phone, email, address string
		var joinedcourses *int
		err = selDB.Scan(&studentid, &studentname, &phone, &email, &address, &joinedcourses)
		if err != nil {
			panic(err.Error())
		}
		student.CustomerId = studentid
		student.CustomerName = studentname
		student.Phone = phone
		student.Email = email
		student.Address = address
		student.CourseJoined = joinedcourses
		listStudent = append(listStudent, student)
	}
	return listStudent
}

func (api StudentAPI) GetStudentByCourseID(nID int) []StudentCourse {
	log.Println("<-----------------Inside GetStudentByCourseID Resposiroty------------->")
	selDB, err := api.DBcon.Query("SELECT s.*,o.CourseId,o.Status FROM Customers s JOIN Orders o ON s.CustomerId=o.CustomerId  WHERE o.CourseId = ? and Status = 1", nID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(nID)
	var listStudent = make([]StudentCourse, 0)
	for selDB.Next() {
		student := StudentCourse{}
		var studentid, courseId, status int
		var studentname, phone, email, address, displayName string
		err = selDB.Scan(&studentid, &studentname, &phone, &email, &address, &courseId, &status, &displayName)
		if err != nil {
			log.Error(err.Error())
		}
		student.CustomerId = studentid
		student.CustomerName = studentname
		student.Phone = phone
		student.Email = email
		student.Address = address
		student.CourseId = courseId
		student.Status = status
		student.DisplayName = displayName
		listStudent = append(listStudent, student)
	}
	return listStudent
}

func (api StudentAPI) GetStudentByID(nID string) Student {
	selDB, err := api.DBcon.Query("SELECT * FROM customers WHERE customerid=?", nID)
	if err != nil {
		log.Error(err.Error())
	}
	student := Student{}
	for selDB.Next() {
		var studentid int
		var studentname, phone, email, address string
		err = selDB.Scan(&studentid, &studentname, &phone, &email, &address)
		if err != nil {
			log.Error(err.Error())
		}
		student.CustomerId = studentid
		student.CustomerName = studentname
		student.Phone = phone
		student.Email = email
		student.Address = address
	}
	return student
}

func (api StudentAPI) InsertStudent(student Student) {
	insForm, err := api.DBcon.Prepare("INSERT INTO customers(customername, phone, email, address) VALUES(?,?,?,?)")
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println("INSERT: Name: ", student)
	insForm.Exec(student.CustomerName, student.CustomerId, student.Email, student.Address)
}

// Delete DeleteStudent
func (api StudentAPI) DeleteStudent(nID string) {
	delForm, err := api.DBcon.Prepare("DELETE FROM Students WHERE studentid=?")
	if err != nil {
		log.Error(err.Error())
	}
	delForm.Exec(nID)
}
