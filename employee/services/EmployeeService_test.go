package services

import (
	. "employee/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var EmployeeMock = Employee{1, "thinh", "ninh phat"}

type EmployeeAPIMock struct {
	mock.Mock
}

func (m *EmployeeAPIMock) GetEmployeeByID(nID string) Employee {
	return EmployeeMock
}
func TestSum(t *testing.T) {
	assert.Equal(t, 1, 1)
	req, err := http.NewRequest("GET", "/api/Employee/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	api := new(EmployeeAPIMock)

	api.On("GetEmployeeByID", 1).Return(true)

	// next we want to define the service we wish to test
	myService := EmployeeService{api}

	// and call said method
	myService.GetEmployeeByIDHandler(w, req)

	emp := Employee{}
	errPs := json.NewDecoder(w.Body).Decode(&emp)
	assert.NoError(t, errPs)
	assert.Equal(t, emp, EmployeeMock)
}
