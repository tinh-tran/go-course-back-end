package models

type OneQuiz struct {
	QuizId      int
	QuizContent string
	ChapterId   int
	CourseId    int
	Answer      []Answer
}
