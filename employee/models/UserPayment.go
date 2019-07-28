package models

type Charge struct {
	Resource    string
	UserId      int
	Brand       string
	PaymentDate string
	Amount      int64
	Exp_Month   int64
	Exp_Year    int64
	Last4       string
	Country     string
}
