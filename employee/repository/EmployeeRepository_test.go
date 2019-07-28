package repository

import (
	. "employee/models"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var EmployeeMock = Employee{1, "thinh", "ninh phat"}

func TestSum(t *testing.T) {
	db, mock, mErr := sqlmock.New()
	assert.Equal(t, 1, 1)
	if mErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mErr)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Id", "Name", "City"}).
		AddRow(1, "thinh", "ninh phat")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	Eapi := EmployeeAPI{db}
	emp := Eapi.GetEmployeeByID("1")
	assert.Equal(t, emp, EmployeeMock)
}
