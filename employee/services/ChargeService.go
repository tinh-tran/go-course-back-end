package services

import (
	. "employee/models"
	. "employee/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ChargeService struct {
	ChargeAPI ChargeAction
}

func (service ChargeService) InsertCharge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start inserting")
	charge := Charge{}
	errPs := json.NewDecoder(r.Body).Decode(&charge)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	t := time.Now()
	result := strings.Split(charge.Resource, "_")
	charge.PaymentDate = t.Format("2006.01.02 15:04:05")
	charge.UserId, _ = strconv.Atoi(result[0])
	service.ChargeAPI.InsertCharge(charge)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
