package main

import (
	. "CourseApi/common"
	. "CourseApi/repository"
	. "CourseApi/services"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	var DBcon = GetConnection()
	var orderAPI = CourseOrderAPI{DBcon}
	var orderService = CourseOrderService{orderAPI}
	log.Println("Start service ... ")

	r := chi.NewRouter()
	r.Use(middleware.Recoverer, middleware.DefaultCompress)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to our go API"))
	})

	r.Route("/api/order", func(r chi.Router) {
		r.Get("/{Studentid}/{Courseid}", orderService.GetOrderIDHandler)
		r.Get("/", orderService.GetAllOrderHandler)
		r.Post("/", orderService.UpsertOrderHandler)
		r.Delete("/{uid}", orderService.DeleteOrderHandler)
		r.Put("/{orderId}/{Status}", orderService.UpDateOrderStatusHandler)
	})

	port := GetServicePort()
	log.Println("Service started successfully and now Listening on port ", port)
	fmt.Println("Start service successfully... ")
	http.ListenAndServe(":"+port, r)
}
