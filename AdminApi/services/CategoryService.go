package services

import (
	. "AdminApi/models"
	. "AdminApi/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type CategoryService struct {
	Api DBaction
}

func (service CategoryService) InsertCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside InsertCategory Services------------->")
	fmt.Println("start inserting")
	cat := Category{}
	errPs := json.NewDecoder(r.Body).Decode(&cat)
	if errPs != nil {
		fmt.Println("parse input failed")
	}
	if cat.CategoryID <= 0 {
		service.Api.InsertCategory(cat)
	} else {
		service.Api.UpdateCategory(cat)
	}

}

//
func (service CategoryService) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside UpdateCategory Services------------->")
	fmt.Println("start updating")
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/category/"))
	cat := Category{CategoryID: id}
	errPs := json.NewDecoder(r.Body).Decode(&cat)
	if errPs != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}
	service.Api.UpdateCategory(cat)
}

// Delete dd
func (service CategoryService) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside DeleteCategory Services------------->")
	fmt.Println("start deleting")
	id := strings.TrimPrefix(r.URL.Path, "/api/category/")
	service.Api.DeleteCategory(id)
	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully", "object": id})
}

//GetCategoryIDHandler Handler
func (service CategoryService) GetCategoryIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetCategoryIDHandler Services------------->")
	log.Debug("Go to GetEmployeeByIDHandler")
	id := strings.TrimPrefix(r.URL.Path, "/api/category/")
	E := service.Api.GetCategoryByID(id)
	log.Debug(E)
	b, err := json.Marshal(E)
	if err != nil {
		log.Error("Fail to GetCategoryByIDHandler", err)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(b))
	}

}

// GetAllCategoryHandler Handler
func (service CategoryService) GetAllCategoryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("<-----------------Inside GetAllCategoryHandler Services------------->")
	E := service.Api.GetAllCategory()
	b, err := json.Marshal(E)
	if err != nil {
		w.Write([]byte("error!! "))
	} else {
		w.Write([]byte(b))
	}
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
