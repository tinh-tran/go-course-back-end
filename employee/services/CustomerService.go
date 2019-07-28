package services

import (
	. "employee/models"
	. "employee/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomerService struct {
	CustomerAPI CustomerAction
}

func (service CustomerService) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start inserting")
	cus := Customer{}
	errPs := json.NewDecoder(r.Body).Decode(&cus)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.CustomerAPI.UpdateUserInfo(cus)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
