package repository

import (
	"database/sql"
	. "employee/models"
	"errors"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type UserDBaction interface {
	GetUser(userName string, passWord string) (User, error)
	RegisterUser(userName string, passWord string, role int)
}

type UserAPI struct {
	DBcon *sql.DB
}

func (api UserAPI) GetUser(userName string, passWord string) (User, error) {
	log.Println("<-----------------Inside GetUser Resposiroty------------->")
	selDB, err := api.DBcon.Query("SELECT * FROM Users WHERE userName=? and passWord=?", userName, passWord)
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	if selDB.Next() {
		var role, id int
		var customerId *int
		var usersc, passwordsc string
		err = selDB.Scan(&id, &usersc, &passwordsc, &customerId, &role)
		if err != nil {
			panic(err.Error())
		}
		user.UserId = id
		user.UserName = usersc
		user.Password = passwordsc
		user.Role = role
		user.CustomerId = customerId

	} else {
		return user, errors.New("can not get user : " + userName)
	}

	return user, nil
}

func (api UserAPI) RegisterUser(userName string, passWord string, role int) {
	res, err := api.DBcon.Exec("INSERT INTO customers (CustomerName, Phone, Email, Address) VALUES ('', '','','')")
	if err != nil {
		println("Exec err:", err.Error())
	} else {
		id, err := res.LastInsertId()
		insUser, err := api.DBcon.Prepare("INSERT INTO users(username, password, role, customerid) VALUES(?,?,?,?)")
		if err != nil {
			log.Error(err.Error())
		}
		fmt.Println("INSERT: Name: ", userName, passWord, role)
		insUser.Exec(userName, passWord, role, id)
	}
}
