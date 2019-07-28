package services

import (
	. "CourseApi/models"
	. "CourseApi/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type CourseOrderService struct {
	CourseOrderAPI CourseOrderAction
}

//GetCategoryIDHandler Handler
func (service CourseOrderService) GetOrderIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetOrderIDHandler Services------------->")
	log.Debug("Go to GetCourseIDHandler")
	ids := strings.TrimPrefix(r.URL.Path, "/api/order/")
	idar := strings.Split(ids, "/")
	E := service.CourseOrderAPI.GetOrderByID(idar[0], idar[1])
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetCourseIDHandler", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(b))
	}
}

// GetAllCategoryHandler Handler
func (service CourseOrderService) GetAllOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAllOrderHandler Services------------->")
	E := service.CourseOrderAPI.GetAllOrders()
	b, err := json.Marshal(E)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("error!! "))
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
	fmt.Println(order.OrderId)

	if order.OrderId == nil {
		service.CourseOrderAPI.InsertOrder(order)
	} else {
		service.CourseOrderAPI.UpdateOrder(order)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

// Delete dd
func (service CourseOrderService) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside DeleteOrder Services------------->")
	fmt.Println("start deleting")
	id := strings.TrimPrefix(r.URL.Path, "/api/order/")
	service.CourseOrderAPI.DeleteOrder(id)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

// Update status Hanlder
func (service CourseOrderService) UpDateOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside upsertOrder Services------------->")
	fmt.Println("start UpDateOrderStatusHandler")
	ids := strings.TrimPrefix(r.URL.Path, "/api/order/")
	idar := strings.Split(ids, "/")
	service.CourseOrderAPI.UpdateOrderStatus(idar[0], idar[1])
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
