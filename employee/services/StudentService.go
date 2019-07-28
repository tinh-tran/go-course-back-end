package services

import (
	. "employee/models"
	. "employee/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type StudentService struct {
	Api DBaction
}

// GetAllStudentHandler Handler
func (service StudentService) GetAllStudentHandler(w http.ResponseWriter, r *http.Request) {
	E := service.Api.GetAllStudent()
	b, err := json.Marshal(E)
	if err != nil {
		w.Write([]byte("error roi "))
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(b))
	}
}

// GetStudentByCourseHandler Handler
func (service StudentService) GetStudentsByCourseIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/Student/Course/"))
	E := service.Api.GetStudentByCourseID(id)
	b, err := json.Marshal(E)
	if err != nil {
		w.Write([]byte("error roi "))
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(b))
	}
}

//GetStudentByIDHandler Handler
func (service StudentService) GetStudentByIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetStudentByIDHandler")
	log.Debug("Go to GetStudentByIDHandler")
	id := strings.TrimPrefix(r.URL.Path, "/api/Student/")
	E := service.Api.GetStudentByID(id)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetStudentByIDHandler", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(b))
	}

}

func (service StudentService) InsertStudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start inserting")
	student := Student{}
	errPs := json.NewDecoder(r.Body).Decode(&student)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.Api.InsertStudent(student)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

// Delete dd
func (service StudentService) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/Student/")
	service.Api.DeleteStudent(id)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
