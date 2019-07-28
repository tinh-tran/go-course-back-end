package repository

import (
	. "QuizApi/models"
	"database/sql"
	"log"
)

type OneQuizAction interface {
	GetAllOneQuiz() []OneQuiz
}
type OneQuizAPI struct {
	DBcon *sql.DB
}

func (oneAPI OneQuizAPI) GetAllOneQuiz() []OneQuiz {
	log.Println("<-----------------Inside GetAllQuiz Resposiroty------------->")
	selQuiz, err := oneAPI.DBcon.Query("SELECT QuizId, QuizContent, ChapterId, CourseId FROM Quizs")
	if err != nil {
		panic(err.Error())
	}
	var listQuiz = []OneQuiz{}
	for selQuiz.Next() {
		oneQuiz := OneQuiz{}
		err = selQuiz.Scan(&oneQuiz.QuizId, &oneQuiz.QuizContent, &oneQuiz.ChapterId, &oneQuiz.CourseId)
		if err != nil {
			panic(err.Error())
		}
		selDB, err := oneAPI.DBcon.Query("SELECT a.AnswerId, a.AnswerContent, a.CourseId, a.ChapterId, a.Status, a.QuizId FROM Quizs q JOIN Answers a ON q.QuizId = a.QuizId WHERE q.CourseId = 2 AND q.ChapterId = 20")
		for selDB.Next() {
			ans := Answer{}
			err = selDB.Scan(&ans.AnswerId, &ans.AnswerContent, &ans.CourseId, &ans.ChapterId, &ans.Status, &ans.QuizId)
			if err != nil {
				panic(err.Error())
			}
			if oneQuiz.QuizId == ans.QuizId {
				oneQuiz.Answer = append(oneQuiz.Answer, ans)
			} else {
				oneQuiz.Answer = nil
			}

		}
		listQuiz = append(listQuiz, oneQuiz)

	}
	return listQuiz
}
