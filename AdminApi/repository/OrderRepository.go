package repository

import (
	. "AdminApi/models"
	"database/sql"
	"fmt"
	"log"
)

type CourseOrderAction interface {
	GetOrderByID(customerId int, courseID int) []Order
	GetAllOrders() []Order
	GetAllCartById(customerId int) []Order
	GetAllClassById(customerId int) []Order
	InsertOrder(order Order)
	UpdateOrder(cour Order)
	DeleteOrder(nID string)
	UpdateOrderStatus(orderID int, status int)
}

type CourseOrderAPI struct {
	DBcon *sql.DB
}

func (courApi CourseOrderAPI) GetOrderByID(customerId int, courseID int) []Order {
	log.Println("<-----------------Inside GetCourseByID Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT * FROM Orders WHERE CustomerId=? and courseID=?", customerId, courseID)
	if err != nil {
		panic(err.Error())
	}
	var listOrder = make([]Order, 0)
	for selDB.Next() {
		order := Order{}
		var orderId, courseId, customerId, status int
		var orderNote *string
		var orderDate string
		err = selDB.Scan(&orderId, &courseId, &customerId, &orderDate, &orderNote, &status)
		if err != nil {
			panic(err.Error())
		}
		order.OrderId = orderId
		order.CourseId = courseId
		order.CustomerId = customerId
		order.OrderDate = orderDate
		order.OrderNote = orderNote
		order.Status = status
		listOrder = append(listOrder, order)
	}
	return listOrder
}

func (courApi CourseOrderAPI) GetAllOrders() []Order {
	log.Println("<-----------------Inside GetAllOrder Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT * FROM Orders")
	if err != nil {
		panic(err.Error())
	}
	var listOrder = make([]Order, 0)
	for selDB.Next() {
		cour := Order{}
		var orderId, courseId, customerId, status int
		var orderNote *string
		var orderDate string
		err = selDB.Scan(&orderId, &courseId, &customerId, &orderDate, &orderNote, &status)
		if err != nil {
			panic(err.Error())
		}
		cour.OrderId = orderId
		cour.CourseId = courseId
		cour.CustomerId = customerId
		cour.OrderDate = orderDate
		cour.OrderNote = orderNote
		cour.Status = status
		listOrder = append(listOrder, cour)
	}
	return listOrder
}

//--insertOrder
func (courApi CourseOrderAPI) InsertOrder(order Order) {
	log.Println("<-----------------Inside Insertorder Resposiroty------------->")
	insForm, err := courApi.DBcon.Prepare("INSERT INTO Orders(CourseId, CustomerId ,orderDate, orderNote, status) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(order.CourseId, order.CustomerId, order.OrderDate, order.OrderNote, order.Status)
	insForm.Exec(order.CourseId, order.CustomerId, order.OrderDate, order.OrderNote, order.Status)
}

//-UpdateOrder
func (courApi CourseOrderAPI) UpdateOrder(order Order) {
	log.Println("<-----------------Inside UpdateCategory Resposiroty------------->")
	insForm, err := courApi.DBcon.Prepare("Update orders SET CourseId=?,CustomerId=?,orderDate=?,orderNote=?,status=? WHERE orderID=?")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Update: order", order)
	insForm.Exec(order.CourseId, order.CustomerId, order.OrderDate, order.OrderNote, order.Status, order.OrderId)
}

// DeleteOrder
func (courApi CourseOrderAPI) DeleteOrder(nID string) {
	log.Println("<-----------------Inside DeleteCategory Resposiroty------------->")
	delForm, err := courApi.DBcon.Prepare("DELETE FROM orders WHERE orderID=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nID)
}

//-UpdateOrderStatus
func (courApi CourseOrderAPI) UpdateOrderStatus(orderID int, status int) {
	log.Println("<-----------------Inside UpdateCategory Resposiroty------------->")
	insForm, err := courApi.DBcon.Prepare("Update orders SET status=? WHERE orderID=?")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Update staus: order", orderID, " status : ", status)
	insForm.Exec(status, orderID)
}

func (courApi CourseOrderAPI) GetAllCartById(customerId int) []Order {
	log.Println("<-----------------Inside GetAllCartById Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT O.*, C.CourseName ,C.CoursePrice, C.CourseImage FROM Orders O JOIN Courses C ON C.CourseId= O.CourseId  WHERE O.Status = 2 and O.CustomerId = ?", customerId)
	if err != nil {
		panic(err.Error())
	}
	var listOrder = make([]Order, 0)
	for selDB.Next() {
		cour := Order{}
		var orderId, courseId, customerId, status int
		var orderNote *string
		var orderDate, courseName, courseImage string
		var coursePrice float64
		err = selDB.Scan(&orderId, &courseId, &customerId, &orderDate, &orderNote, &status, &courseName, &coursePrice, &courseImage)
		if err != nil {
			panic(err.Error())
		}
		cour.OrderId = orderId
		cour.CourseId = courseId
		cour.CustomerId = customerId
		cour.OrderDate = orderDate
		cour.OrderNote = orderNote
		cour.Status = status
		cour.CourseName = courseName
		cour.CoursePrice = coursePrice
		cour.CourseImage = courseImage
		listOrder = append(listOrder, cour)
	}
	return listOrder
}
func (courApi CourseOrderAPI) GetAllClassById(customerId int) []Order {
	log.Println("<-----------------Inside GetAllCartById Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT O.*, C.CourseName ,C.CoursePrice, C.CourseImage FROM Orders O JOIN Courses C ON C.CourseId= O.CourseId  WHERE O.Status = 1 and O.CustomerId = ?", customerId)
	if err != nil {
		panic(err.Error())
	}
	var listOrder = make([]Order, 0)
	for selDB.Next() {
		cour := Order{}
		var orderId, courseId, customerId, status int
		var orderNote *string
		var orderDate, courseName, courseImage string
		var coursePrice float64
		err = selDB.Scan(&orderId, &courseId, &customerId, &orderDate, &orderNote, &status, &courseName, &coursePrice, &courseImage)
		if err != nil {
			panic(err.Error())
		}
		cour.OrderId = orderId
		cour.CourseId = courseId
		cour.CustomerId = customerId
		cour.OrderDate = orderDate
		cour.OrderNote = orderNote
		cour.Status = status
		cour.CourseName = courseName
		cour.CoursePrice = coursePrice
		cour.CourseImage = courseImage
		listOrder = append(listOrder, cour)
	}
	return listOrder
}
