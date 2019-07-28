package repository

import (
	"database/sql"
	. "employee/models"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

type ChargeAction interface {
	InsertCharge(charge Charge)
}

type ChargeAPI struct {
	DBcon *sql.DB
}

func (chargeApi ChargeAPI) InsertCharge(charge Charge) {
	insForm, err := chargeApi.DBcon.Prepare("INSERT INTO payments(UserId, Brand, PaymentDate, Amount, Exp_Month, Exp_Year, Last4, Country) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println("INSERT: Name: ", charge)
	insForm.Exec(charge.UserId, charge.Brand, charge.PaymentDate, charge.Amount, charge.Exp_Month, charge.Exp_Year, charge.Last4, charge.Country)
}
