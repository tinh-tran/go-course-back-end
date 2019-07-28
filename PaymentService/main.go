package main

import (
	. "PaymentService/common"
	. "PaymentService/services"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer, middleware.DefaultCompress)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to our go API"))
	})

	r.Post("/debits", DebitsHandler)
	port := GetServicePort()
	log.Println("Service started successfully and now Listening on port ", port)
	fmt.Println("Start service successfully... ")
	http.ListenAndServe(":"+port, r)
}
