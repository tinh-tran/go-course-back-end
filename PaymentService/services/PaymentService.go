package services

import (
	. "PaymentService/common"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
)

type PaymentDetails struct {
	CustomerID  *string
	Tokenid     *string
	Amount      *int64
	Description *string
	OrderId     int
	CourseId    int
	CustomerId  int
	Status      int
}

func createOneTimeDebit(token string, amount int64, description string) (*stripe.Charge, error) {
	stripe.Key = "sk_test_FgJPbuo5CoNjYRMABa7Yqess"

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String(description),
	}
	params.SetSource(token)
	ch, err := charge.New(params)

	if err != nil {
		return nil, err
	}
	return ch, nil
}

func createSaveDebit(token string, amount int64, description string, email string) (*stripe.Charge, error) {
	stripe.Key = "sk_test_FgJPbuo5CoNjYRMABa7Yqess"

	customerParams := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	customerParams.SetSource(token)
	cus, _ := customer.New(customerParams)

	chargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Customer: stripe.String(string(cus.ID)),
	}
	ch, err := charge.New(chargeParams)
	fmt.Println("-------cus--------", cus)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func createExistDebit(token string, amount int64, description string, email string) (*stripe.Charge, error) {
	stripe.Key = "sk_test_FgJPbuo5CoNjYRMABa7Yqess"

	chargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Customer: stripe.String(string("cus_DyGuJbA5WabVkX")),
	}
	ch, err := charge.New(chargeParams)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func saveToDB(ch *stripe.Charge) {
	url := GetUrlService()
	jsonData := map[string]interface{}{"Resource": ch.Description, "Amount": ch.Amount, "Last4": ch.Source.Card.Last4, "Brand": ch.Source.Card.Brand, "Country": ch.Source.Card.Country, "Exp_Month": ch.Source.Card.ExpMonth, "Exp_Year": ch.Source.Card.ExpYear}
	jsonValue, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("POST", url+"/api/charge/savecharge", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
func DebitsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("-----------------DebitsHandler-------------")
	paymentDetails := PaymentDetails{}
	errPs := json.NewDecoder(r.Body).Decode(&paymentDetails)
	if errPs != nil {
		log.Error("ssssssssssssss")
		log.Error(errPs.Error())
		log.Error("parse input failed", errPs.Error())
	}
	ch, er := createOneTimeDebit(*paymentDetails.Tokenid, *paymentDetails.Amount, *paymentDetails.Description)
	if er != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Pay Error at " + er.Error()))
	} else {
		saveToDB(ch)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Success Payment"))
	}
}
