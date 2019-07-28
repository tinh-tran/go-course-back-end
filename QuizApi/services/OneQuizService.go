package services

import (
	. "QuizApi/repository"
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type OneQuizService struct {
	OneQuizAPI OneQuizAction
}

func (service OneQuizService) GetAllOneQuiz(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAllOneQuiz Services------------->")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	E := service.OneQuizAPI.GetAllOneQuiz()
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
