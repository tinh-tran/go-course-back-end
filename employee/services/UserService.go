package services

import (
	. "employee/models"
	. "employee/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService struct {
	Api UserDBaction
}

//GetEmployeeByIDHandler Handler
func (service UserService) RequestLogin(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("credentials", user.UserName)
	jwt := Jwt{}

	userDB, er := service.Api.GetUser(user.UserName, user.Password)

	if er == nil {
		jwt.Token = "true"
		fmt.Println(userDB.Role)
		if userDB.Role == 1 {
			jwt.Role = "admin"
		}
		if userDB.Role == 0 {
			jwt.Role = "user"
		}
		if userDB.Role == 2 {
			jwt.Role = "student"
		}
		jwt.UserName = userDB.UserName
		jwt.CustomerId = userDB.CustomerId
		jwt.UserId = userDB.UserId
		b, err := json.Marshal(jwt)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Write([]byte(b))
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (service UserService) InsertUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start inserting")
	user := User{}
	errPs := json.NewDecoder(r.Body).Decode(&user)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.Api.RegisterUser(user.UserName, user.Password, user.Role)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
