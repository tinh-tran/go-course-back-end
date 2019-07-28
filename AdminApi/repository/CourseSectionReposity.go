package repository

import (
	. "AdminApi/models"
	"database/sql"
	"fmt"
	"log"
)

type CourseSectionAction interface {
	InsertCourseChapter(chapter Chapter)
	InsertCourseSection(section Section)
	InsertResource(video *string, pdf *string, sectionId int)
}

type SectionAPI struct {
	DBcon *sql.DB
}

func (secApi SectionAPI) InsertCourseSection(section Section) {
	log.Println("<-----------------Inside InsertCourseSection Resposiroty------------->")
	insForm, err := secApi.DBcon.Prepare("INSERT INTO CourseSection(CourseId, ChapterId, SectionName, SectionOrder) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("INSERT: CourseId: ", section.CourseId, " | SectionName: ", section.SectionName)
	insForm.Exec(section.CourseId, section.ChapterId, section.SectionName, section.SectionOrder)
}

func (secApi SectionAPI) InsertCourseChapter(chapter Chapter) {
	log.Println("<-----------------Inside InsertCourseChapter Resposiroty------------->")
	insForm, err := secApi.DBcon.Prepare("INSERT INTO CourseChapter(CourseId, ChapterName , ChapterOrder) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("INSERT: CourseId: ", chapter.CourseId, " | ChapterName: ", chapter.ChapterName)
	insForm.Exec(chapter.CourseId, chapter.ChapterName, chapter.ChapterOrder)
}
func (secApi SectionAPI) InsertResource(video *string, pdf *string, sectionId int) {
	if video != nil {
		insForm, err := secApi.DBcon.Prepare("UPDATE CourseSection SET SectionVideo= ? WHERE CourseSectionId = ?")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("UPDATE: Video: ", video, " | SectionId: ", sectionId, "video Yes")
		insForm.Exec(video, sectionId)
	} else {
		insForm, err := secApi.DBcon.Prepare("UPDATE CourseSection SET SectionPdf= ? WHERE CourseSectionId = ?")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("UPDATE: Pdf: ", pdf, " | SectionId: ", sectionId, "Pdf Yes")
		insForm.Exec(pdf, sectionId)
	}
}
