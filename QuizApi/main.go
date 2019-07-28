package main

import (
	. "QuizApi/common"
	. "QuizApi/repository"
	. "QuizApi/services"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	var DBcon = GetConnection()
	var quizAPI = QuizApi{DBcon}
	var quizService = QuizService{quizAPI}
	var answerAPI = AnswerApi{DBcon}
	var answerService = AnswerService{answerAPI}
	var oneQuizAPI = OneQuizAPI{DBcon}
	var oneQuizService = OneQuizService{oneQuizAPI}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer, middleware.DefaultCompress)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to our go API"))
	})

	r.Route("/api/quiz", func(r chi.Router) {
		r.Get("/", quizService.GetAllQuizHandler)
		r.Post("/bycourseandchapter", quizService.GetQuizByCourseAndChapter)
		r.Post("/", quizService.InsertQuizHandler)
		r.Put("/byid", quizService.UpdateQuizHandler)
	})

	r.Route("/api/answer", func(r chi.Router) {
		r.Get("/", answerService.GetAllAnswerHandler)
		r.Post("/bycourseandchapter", answerService.GetAnswerByCourseAndChapter)
		r.Post("/", answerService.InsertAnswerHandler)
		r.Put("/byid", answerService.UpdateAnswerHandler)
		r.Delete("/{uid}", answerService.DeleteAnswer)
	})
	r.Route("/api/onequiz", func(r chi.Router) {
		r.Post("/getonequiz", oneQuizService.GetAllOneQuiz)
	})

	port := GetServicePort()
	log.Println("Service started successfully and now Listening on port ", port)
	fmt.Println("Start service successfully... ")
	http.ListenAndServe(":"+port, r)
}
