package repository

import (
	. "QuizApi/models"
	"database/sql"
	"fmt"
	"log"
)

type DBaction interface {
	GetAllQuiz() []Quiz
	GetQuizByCourseAndChater(courseId int) Quiz
	InsertQuiz(quiz Quiz)
	UpdateQuiz(quiz Quiz)
}

type QuizApi struct {
	DBcon *sql.DB
}

func (api QuizApi) GetAllQuiz() []Quiz {
	log.Println("<-----------------Inside GetAllQuiz Resposiroty------------->")
	selDB, err := api.DBcon.Query("SELECT * FROM quizs")
	if err != nil {
		panic(err.Error())
	}
	var listQuiz []Quiz
	for selDB.Next() {
		quiz := Quiz{}
		var id, course, chapter int
		var content string
		err = selDB.Scan(&id, &content, &course, &chapter)
		if err != nil {
			panic(err.Error())
		}
		quiz.QuizId = id
		quiz.QuizContent = content
		quiz.CourseId = course
		quiz.ChapterId = chapter
		listQuiz = append(listQuiz, quiz)
	}
	return listQuiz
}

func (api QuizApi) GetQuizByCourseAndChater(courseId int) Quiz {
	log.Println("<-----------------Inside GetQuizByCourseAndChater Resposiroty------------->")
	selDB, err := api.DBcon.Query("SELECT * FROM quizs WHERE CourseId=?", courseId)
	if err != nil {
		panic(err.Error())
	}
	quiz := Quiz{}
	for selDB.Next() {
		var id, course, chapter int
		var content string
		err = selDB.Scan(&id, &content, &course, &chapter)
		if err != nil {
			panic(err.Error())
		}
		quiz.QuizId = id
		quiz.QuizContent = content
		quiz.CourseId = course
		quiz.ChapterId = chapter
	}
	return quiz
}

func (api QuizApi) InsertQuiz(quiz Quiz) {
	log.Println("<-----------------Inside InsertCourse Resposiroty------------->")
	insForm, err := api.DBcon.Prepare("INSERT INTO quizs(QuizContent, CourseId, ChapterId) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("INSERT: CourseName: "+cour.CourseName+" | CourseDes: "+cour.CourseDescription, "| ImageUrl: "+cour.CourseImage)
	insForm.Exec(quiz.QuizContent, quiz.CourseId, quiz.ChapterId)
}

func (api QuizApi) UpdateQuiz(quiz Quiz) {
	log.Println("<-----------------Inside UpdateQuiz Resposiroty------------->")
	insForm, err := api.DBcon.Prepare("Update quizs SET QuizContent =?, CourseId=?, ChapterId=? WHERE QuizId =?")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(quiz.QuizContent)
	//fmt.Println("INSERT: CourseName: "+cour.CourseName+" | CourseDes: "+cour.CourseDescription, "| ImageUrl: "+cour.CourseImage)
	insForm.Exec(quiz.QuizContent, quiz.CourseId, quiz.ChapterId, quiz.QuizId)
}
