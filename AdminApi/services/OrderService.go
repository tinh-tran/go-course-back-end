package services

import (
	. "AdminApi/models"
	. "AdminApi/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

type CourseOrderService struct {
	CourseOrderAPI CourseOrderAction
}

//GetCategoryIDHandler Handler
func (service CourseOrderService) GetOrderIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetOrderIDHandler Services------------->")
	log.Debug("Go to GetCourseIDHandler")
	order := Order{}
	errPs := json.NewDecoder(r.Body).Decode(&order)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	E := service.CourseOrderAPI.GetOrderByID(order.CustomerId, order.CourseId)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetCourseIDHandler", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(b))
	}
}

// GetAllCategoryHandler Handler
func (service CourseOrderService) GetAllOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAllOrderHandler Services------------->")
	E := service.CourseOrderAPI.GetAllOrders()
	b, err := json.Marshal(E)
	if err != nil {
		w.Write([]byte("error!! "))
	} else {
		w.Write([]byte(b))

	}
}

// Delete dd
func (service CourseOrderService) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside DeleteOrder Services------------->")
	fmt.Println("start deleting")
	id := strings.TrimPrefix(r.URL.Path, "/api/order/")
	service.CourseOrderAPI.DeleteOrder(id)
}

// Update status Hanlder
func (service CourseOrderService) UpDateOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside upsertOrder Services------------->")
	fmt.Println("start UpDateOrderStatusHandler")
	order := Order{}
	errPs := json.NewDecoder(r.Body).Decode(&order)
	if errPs != nil {
		fmt.Println("parse input failed", errPs)
	}
	service.CourseOrderAPI.UpdateOrderStatus(order.OrderId, order.Status)
	E := service.CourseOrderAPI.GetAllCartById(order.CustomerId)
	b, _ := json.Marshal(E)
	w.Write([]byte(b))
}

func (service CourseOrderService) GetAllCartById(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAllCartById Services------------->")
	order := Order{}
	errPs := json.NewDecoder(r.Body).Decode(&order)
	if errPs != nil {
		fmt.Println("parse input failed", errPs)
	}
	E := service.CourseOrderAPI.GetAllCartById(order.CustomerId)
	b, err := json.Marshal(E)
	if err != nil {
		w.Write([]byte("error!! "))
	} else {
		w.Write([]byte(b))

	}
}
func (service CourseOrderService) GetAllClassById(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAllClassById Services------------->")
	order := Order{}
	errPs := json.NewDecoder(r.Body).Decode(&order)
	if errPs != nil {
		fmt.Println("parse input failed", errPs)
	}
	E := service.CourseOrderAPI.GetAllClassById(order.CustomerId)
	b, err := json.Marshal(E)
	if err != nil {
		w.Write([]byte("error!! "))
	} else {
		w.Write([]byte(b))

	}
}
func (service CourseOrderService) UpsertOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside upsertOrder Services------------->")
	fmt.Println("start inserting")
	order := Order{}
	errPs := json.NewDecoder(r.Body).Decode(&order)
	if errPs != nil {
		fmt.Println("parse input failed", errPs)
	}
	t := time.Now()
	order.OrderDate = t.Format("2006.01.02 15:04:05")
	if order.OrderId == 0 {
		service.CourseOrderAPI.InsertOrder(order)
	} else {
		service.CourseOrderAPI.UpdateOrder(order)
	}
	E := service.CourseOrderAPI.GetOrderByID(order.CustomerId, order.CourseId)
	b, _ := json.Marshal(E)
	w.Write([]byte(b))
}
