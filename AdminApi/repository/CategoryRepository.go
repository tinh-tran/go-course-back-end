package repository

import (
	. "AdminApi/common"
	. "AdminApi/models"
	"database/sql"
	"fmt"
	"log"
)

type DBaction interface {
	GetCategoryByID(nID string) Category
	DeleteCategory(nID string)
	UpdateCategory(cat Category)
	InsertCategory(cat Category)
	GetAllCategory() []Category
}

type CategoryAPI struct {
	DBcon *sql.DB
}

func (api CategoryAPI) GetCategoryByID(nID string) Category {
	log.Println("<-----------------Inside GetCategoryByID Resposiroty------------->")
	selDB, err := api.DBcon.Query("SELECT * FROM Categories WHERE CategoryID=?", nID)
	if err != nil {
		panic(err.Error())
	}
	cat := Category{}
	for selDB.Next() {
		var id int
		var name, description, imageUrl string
		err = selDB.Scan(&id, &name, &description, &imageUrl)
		if err != nil {
			panic(err.Error())
		}
		cat.CategoryID = id
		cat.CategoryName = name
		cat.CategoryDescription = description
		cat.CategoryImage = imageUrl
	}
	return cat
}

//--insertCate
func (api CategoryAPI) InsertCategory(cat Category) {
	log.Println("<-----------------Inside InsertCategory Resposiroty------------->")
	insForm, err := api.DBcon.Prepare("INSERT INTO Categories(CategoryName, CategoryDescription,CategoryImage) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("INSERT: CategoryName: "+cat.CategoryName+" | CategoryDes: "+cat.CategoryDescription, "| ImageUrl: "+cat.CategoryImage)
	insForm.Exec(cat.CategoryName, cat.CategoryDescription, cat.CategoryImage)
}

//-UpdateCate
func (api CategoryAPI) UpdateCategory(cat Category) {
	log.Println("<-----------------Inside UpdateCategory Resposiroty------------->")
	insForm, err := api.DBcon.Prepare("Update Categories SET CategoryName=?,CategoryDescription=?,CategoryImage = ? WHERE CategoryId=?")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("UPDATE: CategoryName: "+cat.CategoryName+" | CategoryDes: "+cat.CategoryDescription, "| ImageUrl"+cat.CategoryImage, "|WHERE ID =", cat.CategoryID)
	insForm.Exec(cat.CategoryName, cat.CategoryDescription, cat.CategoryImage, cat.CategoryID)
}

// Delete DeleteEmployee
func (api CategoryAPI) DeleteCategory(nID string) {
	log.Println("<-----------------Inside DeleteCategory Resposiroty------------->")
	delForm, err := api.DBcon.Prepare("DELETE FROM Categories WHERE CategoryID=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nID)

}

// Get All Employee with gorm
func (api CategoryAPI) GetAllCategoryWithGorm() {
	var cat []Category
	db := GetdbConnforgorm()
	db.Find(&cat)
	fmt.Println("cat", cat)
}

func (api CategoryAPI) GetAllCategory() []Category {
	log.Println("<-----------------Inside GetAllCategory Resposiroty------------->")
	selDB, err := api.DBcon.Query("SELECT * FROM Categories")
	if err != nil {
		panic(err.Error())
	}
	var listCategory []Category
	for selDB.Next() {
		cat := Category{}
		var id int
		var name, des, imgUrl string
		err = selDB.Scan(&id, &name, &des, &imgUrl)
		if err != nil {
			panic(err.Error())
		}
		cat.CategoryID = id
		cat.CategoryName = name
		cat.CategoryDescription = des
		cat.CategoryImage = imgUrl
		listCategory = append(listCategory, cat)
	}
	return listCategory
}
