package services

import (
	. "AdminApi/models"
	. "AdminApi/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type CourseService struct {
	CourApi CourseAction
}

func (service *CourseService) InsertCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside InsertCourse Services------------->")
	fmt.Println("start inserting")
	cour := Course{}
	errPs := json.NewDecoder(r.Body).Decode(&cour)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	if cour.CourseId <= 0 {
		service.CourApi.InsertCourse(cour)
	} else {
		service.CourApi.UpdateCourse(cour)
	}
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Delete dd
func (service *CourseService) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside DeleteCourse Services------------->")
	fmt.Println("start deleting")
	id := strings.TrimPrefix(r.URL.Path, "/api/course/")
	service.CourApi.DeleteCourse(id)
	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully", "object": id})
}

//GetCategoryIDHandler Handler
func (service *CourseService) GetCourseIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetCourseIDHandler Services------------->")
	log.Debug("Go to GetCourseIDHandler")
	cour := Course{}
	errPs := json.NewDecoder(r.Body).Decode(&cour)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	E := service.CourApi.GetCourseByID(cour.CourseId)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetCourseIDHandler", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(b))
	}
}

//GetCategoryIDHandler Handler
func (service *CourseService) GetChapterCourseByCourseId(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetChapterCourseByCourseId Services------------->")
	log.Debug("Go to GetChapterCourseByCourseId")
	chap := Chapter{}
	errPs := json.NewDecoder(r.Body).Decode(&chap)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	E := service.CourApi.GetChapterByCourseId(chap.CourseId)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetChapterCourseByCourseId", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(b))
	}

}

func (service *CourseService) GetSectionCourseByCourseId(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetSectionCourseByCourseId Services------------->")
	log.Debug("Go to GetSectionCourseByCourseId")
	sec := Section{}
	errPs := json.NewDecoder(r.Body).Decode(&sec)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	E := service.CourApi.GetSectionByCourseId(sec.CourseId)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetChapterCourseByCourseId", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(b))
	}

}

// GetAllCategoryHandler Handler
func (service *CourseService) GetAllCourseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAllCourseHandler Services------------->")
	E := service.CourApi.GetAllCourse()
	b, err := json.Marshal(E)
	if err != nil {
		w.Write([]byte("error!! "))
	} else {
		w.Write([]byte(b))
	}
}

func (service *CourseService) GetCourseByCategoryId(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetCourseByCategoryId Services------------->")
	log.Debug("Go to GetCourseByCategoryId")
	cour := Course{}
	json.NewDecoder(r.Body).Decode(&cour)
	E := service.CourApi.GetCourseByCategoryId(cour.CategoryID, cour.Page)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetCourseByCategoryId", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(b))
	}
}
func (service *CourseService) GetCourseNum(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetCourseNum Services------------->")
	log.Debug("Go to GetCourseNum")
	num := Number{}
	json.NewDecoder(r.Body).Decode(&num)
	E := service.CourApi.GetCourseNum(num.CategoryID)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetCourseNum", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(b))
	}
}
func (service *CourseService) AddSeeCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside AddSeeCourse Services------------->")
	fmt.Println("start AddSeeCourse")
	seeCour := SeeCourse{}
	errPs := json.NewDecoder(r.Body).Decode(&seeCour)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	service.CourApi.AddSeeCourse(seeCour.UserId, seeCour.CourseId, seeCour.Pdf, seeCour.Video, seeCour.SectionId)
	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func (service *CourseService) GetProcessCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetProcessCourse Services------------->")
	seeCour := SeeCourse{}
	json.NewDecoder(r.Body).Decode(&seeCour)
	E := service.CourApi.GetProcessCourse(seeCour.CourseId, seeCour.UserId)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetProcessCourse", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(b))
	}
}

func (service *CourseService) GetALlCourseByUserId(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetALlCourseByUserId Services------------->")
	course := Course{}
	json.NewDecoder(r.Body).Decode(&course)
	E := service.CourApi.GetALlCourseByUserId(course.CreateId)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetALlCourseByUserId", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(b))
	}
}
