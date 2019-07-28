package main

import (
	. "employee/common"
	. "employee/repository"
	. "employee/services"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	var DBcon = GetConnection()
	var studentAPI = StudentAPI{DBcon}
	var studentService = StudentService{studentAPI}
	log.Println("Start service ... ")
	var userAPI = UserAPI{DBcon}
	var userService = UserService{userAPI}

	var chargeAPI = ChargeAPI{DBcon}
	var chargeService = ChargeService{chargeAPI}

	var customerAPI = CustomerAPI{DBcon}
	var customerService = CustomerService{customerAPI}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer, middleware.DefaultCompress)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to our go API"))
	})

	r.Route("/api/Student", func(r chi.Router) {
		r.Get("/{uid}", studentService.GetStudentByIDHandler)
		r.Get("/", studentService.GetAllStudentHandler)
		r.Post("/", studentService.InsertStudentHandler)
		r.Delete("/{uid}", studentService.DeleteStudentHandler)
		r.Get("/Course/{uid}", studentService.GetStudentsByCourseIDHandler)
	})

	r.Route("/api/User", func(r chi.Router) {
		r.Post("/login", userService.RequestLogin)
		r.Post("/register", userService.InsertUser)
		r.Post("/updateinfo", customerService.UpdateUserInfo)
	})
	r.Route("/api/charge", func(r chi.Router) {
		r.Post("/savecharge", chargeService.InsertCharge)
	})
	port := GetServicePort()
	log.Println("Service started successfully and now Listening on port ", port)
	fmt.Println("Start service successfully... ")
	http.ListenAndServe(":"+port, r)
}
