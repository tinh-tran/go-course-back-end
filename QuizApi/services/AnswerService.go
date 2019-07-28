package services

import (
	. "QuizApi/models"
	. "QuizApi/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type AnswerService struct {
	Api AnswerAction
}

func (service AnswerService) GetAllAnswerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAllAnswerHandler Services------------->")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	E := service.Api.GetAllAnswer()
	b, err := json.Marshal(E)
	if err != nil {
		w.Write([]byte("error!! "))
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Write([]byte(b))
	}
}

func (service AnswerService) GetAnswerByCourseAndChapter(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAnswerByCourseAndChapter Services------------->")
	log.Debug("Go to GetQuizByCourseAndChapter")
	answer := Answer{}
	json.NewDecoder(r.Body).Decode(&answer)
	E := service.Api.GetAnswerByCourseAndChater(answer.CourseId)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetQuizByCourseAndChapter", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Write([]byte(b))
	}
}

func (service AnswerService) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside DeleteAnswer Services------------->")
	fmt.Println("start deleting")
	id := strings.TrimPrefix(r.URL.Path, "/api/answer/")
	service.Api.DeleteAnswer(id)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully", "object": id})
}

func (service AnswerService) InsertAnswerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside InsertAnswerHandler Services------------->")
	fmt.Println("start inserting")
	answer := Answer{}
	errPs := json.NewDecoder(r.Body).Decode(&answer)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.Api.InsertAnswer(answer)
}

func (service AnswerService) UpdateAnswerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside UpdateAnswerHandler Services------------->")
	fmt.Println("start inserting")
	answer := Answer{}
	errPs := json.NewDecoder(r.Body).Decode(&answer)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.Api.UpdateAnswer(answer)
}
