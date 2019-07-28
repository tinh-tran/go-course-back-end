package repository

import (
	. "CourseApi/models"
	"database/sql"
	"fmt"
	"log"
)

type CourseOrderAction interface {
	GetOrderByID(studentID string, courseID string) Order
	GetAllOrders() []Order
	InsertOrder(order Order)
	UpdateOrder(cour Order)
	DeleteOrder(nID string)
	UpdateOrderStatus(nID string, status string)
}

type CourseOrderAPI struct {
	DBcon *sql.DB
}

func (courApi CourseOrderAPI) GetOrderByID(studentID string, courseID string) Order {
	log.Println("<-----------------Inside GetCourseByID Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT * FROM Orders WHERE studentID=? and courseID=?", studentID, courseID)
	if err != nil {
		panic(err.Error())
	}
	cour := Order{}
	for selDB.Next() {
		var orderId, courseId, studentID, status *int
		var orderDate, orderNote *string
		err = selDB.Scan(&orderId, &courseId, &studentID, &orderDate, &orderNote, &status)
		if err != nil {
			panic(err.Error())
		}
		cour.OrderId = orderId
		cour.CourseId = courseId
		cour.StudentID = studentID
		cour.OrderDate = orderDate
		cour.OrderNote = orderNote
		cour.Status = status

	}
	return cour
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
		var orderId, courseId, studentID, status *int
		var orderDate, orderNote *string
		err = selDB.Scan(&orderId, &courseId, &studentID, &orderDate, &orderNote, &status)
		if err != nil {
			panic(err.Error())
		}
		cour.OrderId = orderId
		cour.CourseId = courseId
		cour.StudentID = studentID
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
	insForm, err := courApi.DBcon.Prepare("INSERT INTO orders(CourseId, StudentID,orderDate, orderNote, status) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(order.CourseId, order.StudentID, order.OrderDate, order.OrderNote, order.Status)
	insForm.Exec(order.CourseId, order.StudentID, order.OrderDate, order.OrderNote, order.Status)
}

//-UpdateOrder
func (courApi CourseOrderAPI) UpdateOrder(order Order) {
	log.Println("<-----------------Inside UpdateCategory Resposiroty------------->")
	insForm, err := courApi.DBcon.Prepare("Update orders SET CourseId=?,StudentID=?,orderDate=?,orderNote=?,status=? WHERE orderID=?")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("INSERT: order", order)
	insForm.Exec(order.CourseId, order.StudentID, order.OrderDate, order.OrderNote, order.Status, order.OrderId)
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
func (courApi CourseOrderAPI) UpdateOrderStatus(orderID string, status string) {
	log.Println("<-----------------Inside UpdateCategory Resposiroty------------->")
	insForm, err := courApi.DBcon.Prepare("Update orders SET status=? WHERE orderID=?")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Update staus: order" + orderID + " status : " + status)
	insForm.Exec(status, orderID)
}
