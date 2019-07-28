package repository

import (
	. "QuizApi/models"
	"database/sql"
	"log"
)

type AnswerAction interface {
	GetAllAnswer() []Answer
	GetAnswerByCourseAndChater(courseId int) []Answer
	DeleteAnswer(nId string)
	InsertAnswer(answer Answer)
	UpdateAnswer(answer Answer)
}

type AnswerApi struct {
	DBcon *sql.DB
}

func (api AnswerApi) GetAllAnswer() []Answer {
	log.Println("<-----------------Inside GetAllAnswer Resposiroty------------->")
	selDB, err := api.DBcon.Query("SELECT * FROM answers")
	if err != nil {
		panic(err.Error())
	}
	var listAnswer []Answer
	for selDB.Next() {
		answer := Answer{}
		var answerid, quizid, status, course, chapter int
		var content string
		err = selDB.Scan(&answerid, &content, &quizid, &status, &course, &chapter)
		if err != nil {
			panic(err.Error())
		}
		answer.AnswerId = answerid
		answer.AnswerContent = content
		answer.QuizId = quizid
		answer.Status = status
		answer.CourseId = course
		answer.ChapterId = chapter
		listAnswer = append(listAnswer, answer)
	}
	return listAnswer
}

func (api AnswerApi) GetAnswerByCourseAndChater(courseId int) []Answer {
	log.Println("<-----------------Inside GetAnswerByCourseAndChater Resposiroty------------->")
	selDB, err := api.DBcon.Query("SELECT * FROM answers WHERE CourseId=?", courseId)
	if err != nil {
		panic(err.Error())
	}
	var listAnswer []Answer
	for selDB.Next() {
		answer := Answer{}
		var answerid, quizid, status, course, chapter int
		var content string
		err = selDB.Scan(&answerid, &content, &quizid, &status, &course, &chapter)
		if err != nil {
			panic(err.Error())
		}
		answer.AnswerId = answerid
		answer.AnswerContent = content
		answer.QuizId = quizid
		answer.Status = status
		answer.CourseId = course
		answer.ChapterId = chapter
		listAnswer = append(listAnswer, answer)
	}
	return listAnswer
}

func (api AnswerApi) DeleteAnswer(nID string) {
	log.Println("<-----------------Inside DeleteCourse Resposiroty------------->")
	delForm, err := api.DBcon.Prepare("DELETE FROM answers WHERE AnswerId=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nID)
}

func (api AnswerApi) InsertAnswer(answer Answer) {
	log.Println("<-----------------Inside InsertAnswer Resposiroty------------->")
	insForm, err := api.DBcon.Prepare("INSERT INTO answers(AnswerContent, QuizId, Status, CourseId, ChapterId) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("INSERT: CourseName: "+cour.CourseName+" | CourseDes: "+cour.CourseDescription, "| ImageUrl: "+cour.CourseImage)
	insForm.Exec(answer.AnswerContent, answer.QuizId, answer.Status, answer.CourseId, answer.ChapterId)
}

func (api AnswerApi) UpdateAnswer(answer Answer) {
	log.Println("<-----------------Inside InsertAnswer Resposiroty------------->")
	insForm, err := api.DBcon.Prepare("Update answers SET AnswerContent=?, QuizId =?, Status=?, CourseId=?, ChapterId=? WHERE AnswerId =?")
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("INSERT: CourseName: "+cour.CourseName+" | CourseDes: "+cour.CourseDescription, "| ImageUrl: "+cour.CourseImage)
	insForm.Exec(answer.AnswerContent, answer.QuizId, answer.Status, answer.CourseId, answer.ChapterId, answer.AnswerId)
}
