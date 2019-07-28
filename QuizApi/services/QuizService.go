package services

import (
	. "QuizApi/models"
	. "QuizApi/repository"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type QuizService struct {
	Api DBaction
}

func (service QuizService) GetAllQuizHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAllQuizHandler Services------------->")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	E := service.Api.GetAllQuiz()
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

func (service QuizService) GetQuizByCourseAndChapter(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetQuizByCourseAndChapter Services------------->")
	log.Debug("Go to GetQuizByCourseAndChapter")
	quiz := Quiz{}
	json.NewDecoder(r.Body).Decode(&quiz)
	E := service.Api.GetQuizByCourseAndChater(quiz.CourseId)
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

func (service QuizService) InsertQuizHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside InsertQuizHandler Services------------->")
	fmt.Println("start inserting")
	quiz := Quiz{}
	errPs := json.NewDecoder(r.Body).Decode(&quiz)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.Api.InsertQuiz(quiz)
}

func (service QuizService) UpdateQuizHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside UpdateQuizHandler Services------------->")
	fmt.Println("start inserting")
	quiz := Quiz{}
	errPs := json.NewDecoder(r.Body).Decode(&quiz)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.Api.UpdateQuiz(quiz)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
