package repository

import (
	"database/sql"
	. "employee/models"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type CustomerAction interface {
	UpdateUserInfo(customer Customer)
}

type CustomerAPI struct {
	DBcon *sql.DB
}

func (cusApi CustomerAPI) UpdateUserInfo(customer Customer) {
	insForm, err := cusApi.DBcon.Prepare("UPDATE SET CustomerName = ?, Phone = ? ,Email = ? , Address = ? , DisplayName = ? , CustomerPhoto = ?  WHERE CustomerId = ? ")
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println("UPDATE: Info: ", customer)
	insForm.Exec(customer.CustomerName, customer.Phone, customer.Email, customer.Address, customer.DisplayName, customer.CustomerPhoto, customer.CustomerId)
}
