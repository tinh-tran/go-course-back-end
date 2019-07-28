package main

import (
	. "AdminApi/common"
	. "AdminApi/repository"
	. "AdminApi/services"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	var DBcon = GetConnection()
	var categoryAPI = CategoryAPI{DBcon}
	var categoryService = CategoryService{categoryAPI}
	var courseAPI = CourseAPI{DBcon}
	var courseService = CourseService{&courseAPI}
	var orderAPI = CourseOrderAPI{DBcon}
	var orderService = CourseOrderService{orderAPI}
	var sectionAPI = SectionAPI{DBcon}
	var courseSectionService = CourseSectionService{sectionAPI}
	log.Println("Start service ... ")

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Use(middleware.Recoverer, middleware.DefaultCompress)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to our go API"))
	})

	r.Route("/api/category", func(r chi.Router) {
		r.Get("/{uid}", categoryService.GetCategoryIDHandler)
		r.Get("/", categoryService.GetAllCategoryHandler)
		r.Post("/", categoryService.InsertCategory)
		r.Delete("/{uid}", categoryService.DeleteCategory)
	})

	r.Route("/api/course", func(r chi.Router) {
		r.Post("/byid", courseService.GetCourseIDHandler)
		r.Get("/", courseService.GetAllCourseHandler)
		r.Post("/", courseService.InsertCourse)
		r.Delete("/{uid}", courseService.DeleteCourse)
		r.Post("/chapter/byid", courseService.GetChapterCourseByCourseId)
		r.Post("/section/byid", courseService.GetSectionCourseByCourseId)
		r.Post("/bycat", courseService.GetCourseByCategoryId)
		r.Post("/number", courseService.GetCourseNum)
		r.Post("/seeclass", courseService.AddSeeCourse)
		r.Post("/seeone", courseService.GetProcessCourse)
		r.Post("/createid", courseService.GetALlCourseByUserId)
		r.Post("/chapter", courseSectionService.InsertCourseChapter)
		r.Post("/section", courseSectionService.InsertCourseSection)
		r.Post("/section/addresource", courseSectionService.InsertResource)
	})
	r.Route("/api/order", func(r chi.Router) {
		r.Post("/byuserid", orderService.GetOrderIDHandler)
		r.Get("/", orderService.GetAllOrderHandler)
		r.Post("/", orderService.UpsertOrderHandler)
		r.Delete("/{uid}", orderService.DeleteOrderHandler)
		r.Post("/update", orderService.UpDateOrderStatusHandler)
		r.Post("/getcartbyid", orderService.GetAllCartById)
		r.Post("/getmyclassbyid", orderService.GetAllClassById)
	})
	port := GetServicePort()
	log.Println("Service started successfully and now Listening on port ", port)
	fmt.Println("Start service successfully... ")
	http.ListenAndServe(":"+port, r)
}
