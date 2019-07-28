package services

import (
	. "AdminApi/models"
	. "AdminApi/repository"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type CourseSectionService struct {
	SecApi CourseSectionAction
}

func (service CourseSectionService) InsertCourseChapter(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside InsertCourseChapter Services------------->")
	fmt.Println("start inserting")
	chapter := Chapter{}
	errPs := json.NewDecoder(r.Body).Decode(&chapter)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.SecApi.InsertCourseChapter(chapter)
	//respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func (service CourseSectionService) InsertCourseSection(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside InsertCourseSection Services------------->")
	fmt.Println("start inserting")
	section := Section{}
	errPs := json.NewDecoder(r.Body).Decode(&section)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.SecApi.InsertCourseSection(section)
	//respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func (service CourseSectionService) InsertResource(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside InsertResource Services------------->")
	fmt.Println("start inserting")
	section := Section{}
	errPs := json.NewDecoder(r.Body).Decode(&section)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.SecApi.InsertResource(section.SectionVideo, section.SectionPdf, section.CourseSectionId)
	//respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}
