package repository

import (
	. "AdminApi/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type CourseAction interface {
	GetCourseByID(nID int) []Course
	DeleteCourse(nID string)
	UpdateCourse(cour Course)
	InsertCourse(cour Course)
	GetChapterByCourseId(nID int) []Chapter
	GetAllCourse() []Course
	GetSectionByCourseId(nID int) []Section
	GetCourseByCategoryId(nID int, page int) []Course
	GetCourseNum(CategoryID int) []Number
	AddSeeCourse(userId int, courseId int, pdf string, video string, courseSectionId int)
	GetProcessCourse(courseId int, userId int) []SeeCourse
	GetALlCourseByUserId(createId int) []Course
}

type CourseAPI struct {
	DBcon *sql.DB
}

func (courApi CourseAPI) GetCourseByID(nID int) []Course {
	log.Println("<-----------------Inside GetCourseByID Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT * FROM Courses WHERE CourseId=?", nID)
	fmt.Println(nID)
	if err != nil {
		panic(err.Error())
	}
	var listCours = make([]Course, 0)
	for selDB.Next() {
		cour := Course{}
		var id, catid, slotvai, slotregis, status, createId int
		var courPrice float64
		var startdate, enddate time.Time
		var name, description, imageUrl string
		err = selDB.Scan(&id, &name, &catid, &description, &courPrice, &imageUrl, &startdate, &enddate, &slotvai, &slotregis, &createId, &status)
		if err != nil {
			panic(err.Error())
		}
		cour.CourseId = id
		cour.CourseName = name
		cour.CourseDescription = description
		cour.CategoryID = catid
		cour.CoursePrice = courPrice
		cour.CourseImage = imageUrl
		cour.EndDate = enddate
		cour.StartDate = startdate
		cour.SlotAvailable = slotvai
		cour.SlotRegistered = slotregis
		cour.CreateId = createId
		if status == 1 {
			cour.Status = true
		} else {
			cour.Status = false
		}
		listCours = append(listCours, cour)
	}
	return listCours
}

//--insertCate
func (courApi CourseAPI) InsertCourse(cour Course) {
	log.Println("<-----------------Inside InsertCourse Resposiroty------------->")
	insForm, err := courApi.DBcon.Prepare("INSERT INTO Courses(CourseName, CategoryID, CourseDescription, CoursePrice, CourseImage, StartDate, EndDate, SlotAvailable, SlotRegistered, CreateID, Status) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("INSERT: CourseName: "+cour.CourseName+" | CourseDes: "+cour.CourseDescription, "| ImageUrl: "+cour.CourseImage)
	insForm.Exec(cour.CourseName, cour.CategoryID, cour.CourseDescription, cour.CoursePrice, cour.CourseImage, cour.StartDate, cour.EndDate, cour.SlotAvailable, cour.SlotRegistered, cour.CreateId, cour.Status)
}

//-UpdateCate
func (courApi CourseAPI) UpdateCourse(cour Course) {
	log.Println("<-----------------Inside UpdateCouses Resposiroty------------->")
	insForm, err := courApi.DBcon.Prepare("Update Courses SET CourseName =?, CategoryID =?, CourseDescription =?, CoursePrice =?, CourseImage =?, StartDate =?, EndDate =?, SlotAvailable =?, SlotRegistered =?, CreateId=?, Status= ? WHERE CourseId =?")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("UPDATE:  CourseName: "+cour.CourseName+" | CourseDes: "+cour.CourseDescription, "| ImageUrl: "+cour.CourseImage, "|WHERE ID =", cour.CourseId)
	insForm.Exec(cour.CourseName, cour.CourseDescription, cour.CategoryID, cour.CoursePrice, cour.CourseImage, cour.EndDate, cour.StartDate, cour.SlotAvailable, cour.SlotRegistered, cour.CreateId, cour.Status, cour.CourseId)
	//fmt.Println(insForm.Exec(cour.CourseName, cour.CategoryID, cour.CourseDescription, cour.CoursePrice, cour.CourseImage, cour.EndDate, cour.StartDate, cour.SlotAvailable, cour.SlotRegistered, cour.CreateId, cour.CourseId))
}

// Delete
func (courApi CourseAPI) DeleteCourse(nID string) {
	log.Println("<-----------------Inside DeleteCourse Resposiroty------------->")
	delForm, err := courApi.DBcon.Prepare("DELETE FROM Courses WHERE CourseId=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nID)
}

func (courApi CourseAPI) GetAllCourse() []Course {
	log.Println("<-----------------Inside GetAllCourse Resposiroty------------->")
	query := "SELECT c.*,ca.CategoryName, A.StudentJoined FROM Courses c JOIN Categories ca ON c.CategoryID= ca.CategoryID left join ( SELECT courseid,count(CustomerId) as StudentJoined FROM Orders group by courseid ) A on C.courseid = A.courseid"
	selDB, err := courApi.DBcon.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var listCour = make([]Course, 0)
	for selDB.Next() {
		cour := Course{}
		var id, catid, slotvai, slotregis, createId int
		var studentJoined *int
		var status int
		var courPrice float64
		var startDate, endDate time.Time
		var name, description, imageUrl, cateName string
		err = selDB.Scan(&id, &name, &catid, &description, &courPrice, &imageUrl, &startDate, &endDate, &slotvai, &slotregis, &createId, &status, &cateName, &studentJoined)
		if err != nil {
			panic(err.Error())
		}
		cour.CourseId = id
		cour.CourseName = name
		cour.CourseDescription = description
		cour.CategoryID = catid
		cour.CoursePrice = courPrice
		cour.CourseImage = imageUrl
		cour.EndDate = endDate
		cour.StartDate = startDate
		cour.SlotAvailable = slotvai
		cour.SlotRegistered = slotregis
		cour.CreateId = createId
		cour.CategoryName = cateName
		cour.StudentJoined = studentJoined
		if status == 1 {
			cour.Status = true
		} else {
			cour.Status = false
		}

		listCour = append(listCour, cour)
	}
	return listCour

}

func (courApi CourseAPI) GetChapterByCourseId(nID int) []Chapter {
	log.Println("<-----------------Inside GetChapterByCourseId Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT CourseId,ChapterName,ChapterId,ChapterOrder FROM CourseChapter WHERE CourseId = ? ORDER BY ChapterOrder", nID)
	if err != nil {
		panic(err.Error())
	}

	var listChap = make([]Chapter, 0)
	for selDB.Next() {
		chap := Chapter{}
		var ChapterId, ChapterOrder, CourseId int
		var ChapterName string
		err = selDB.Scan(&CourseId, &ChapterName, &ChapterId, &ChapterOrder)
		if err != nil {
			panic(err.Error())
		}
		chap.CourseId = CourseId
		chap.ChapterId = ChapterId
		chap.ChapterName = ChapterName
		chap.ChapterOrder = ChapterOrder
		listChap = append(listChap, chap)
	}
	return listChap
}
func (courApi CourseAPI) GetSectionByCourseId(nID int) []Section {
	log.Println("<-----------------Inside GetSectionByCourseId Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT * FROM CourseSection WHERE CourseId = ? ORDER BY ChapterId AND SectionOrder", nID)
	if err != nil {
		panic(err.Error())
	}
	var listSec = make([]Section, 0)
	for selDB.Next() {
		sec := Section{}
		var ChapterId, SectionOrder, CourseSectionId, CourseId int
		var SectionName string
		var SectionVideo, SectionPdf *string
		err = selDB.Scan(&CourseId, &ChapterId, &SectionName, &SectionVideo, &SectionPdf, &CourseSectionId, &SectionOrder)
		if err != nil {
			panic(err.Error())
		}
		sec.CourseId = CourseId
		sec.ChapterId = ChapterId
		sec.SectionName = SectionName
		sec.SectionPdf = SectionPdf
		sec.SectionVideo = SectionVideo
		sec.SectionOrder = SectionOrder
		sec.CourseSectionId = CourseSectionId
		listSec = append(listSec, sec)
	}
	return listSec
}
func (courApi CourseAPI) GetCourseByCategoryId(nID int, page int) []Course {
	log.Println("<-----------------Inside GetCourseByCategoryId Resposiroty------------->")
	var listCourse = make([]Course, 0)
	if nID == 0 {
		pagenum := ((page - 1) * 4)
		selDB, err := courApi.DBcon.Query("SELECT * FROM Courses LIMIT ?, 6", pagenum)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(page, pagenum)
		for selDB.Next() {
			cour := Course{}
			var id, catid, slotvai, slotregis, createId, status int
			var courPrice float64
			var startDate, endDate time.Time
			var name, description, imageUrl string
			err = selDB.Scan(&id, &name, &catid, &description, &courPrice, &imageUrl, &startDate, &endDate, &slotvai, &slotregis, &createId, &status)
			if err != nil {
				panic(err.Error())
			}
			cour.CourseId = id
			cour.CourseName = name
			cour.CourseDescription = description
			cour.CategoryID = catid
			cour.CoursePrice = courPrice
			cour.CourseImage = imageUrl
			cour.EndDate = endDate
			cour.StartDate = startDate
			cour.SlotAvailable = slotvai
			cour.SlotRegistered = slotregis
			cour.CreateId = createId
			if status == 1 {
				cour.Status = true
			} else {
				cour.Status = false
			}
			listCourse = append(listCourse, cour)
		}
	} else {
		pagenum := ((page - 1) * 4)
		selDB, err := courApi.DBcon.Query("SELECT * FROM Courses WHERE CategoryId = ? LIMIT ?, 6", nID, pagenum)
		if err != nil {
			panic(err.Error())
		}
		for selDB.Next() {
			cour := Course{}
			var id, catid, slotvai, slotregis, createId, status int
			var courPrice float64
			var startDate, endDate time.Time
			var name, description, imageUrl string
			err = selDB.Scan(&id, &name, &catid, &description, &courPrice, &imageUrl, &startDate, &endDate, &slotvai, &slotregis, &createId, &status)
			if err != nil {
				panic(err.Error())
			}
			cour.CourseId = id
			cour.CourseName = name
			cour.CourseDescription = description
			cour.CategoryID = catid
			cour.CoursePrice = courPrice
			cour.CourseImage = imageUrl
			cour.EndDate = endDate
			cour.StartDate = startDate
			cour.SlotAvailable = slotvai
			cour.SlotRegistered = slotregis
			listCourse = append(listCourse, cour)
		}
	}
	return listCourse
}
func (courApi CourseAPI) GetCourseNum(CategoryID int) []Number {
	log.Println("<-----------------Inside GetCourseNum Resposiroty------------->")
	var listNum = make([]Number, 0)
	if CategoryID == 0 {
		selDB, err := courApi.DBcon.Query("SELECT COUNT(*), CategoryID FROM Courses")
		if err != nil {
			panic(err.Error())
		}
		for selDB.Next() {
			num := Number{}
			var count, id int
			err = selDB.Scan(&count, &id)
			if err != nil {
				panic(err.Error())
			}
			num.Num = count
			num.CategoryID = id
			listNum = append(listNum, num)
		}
	} else {
		selDB, err := courApi.DBcon.Query("SELECT COUNT(*),CategoryID FROM Courses WHERE CategoryID=?", CategoryID)
		if err != nil {
			panic(err.Error())
		}
		for selDB.Next() {
			num := Number{}
			var count, id int
			err = selDB.Scan(&count, &id)
			if err != nil {
				panic(err.Error())
			}
			num.Num = count
			num.CategoryID = id
			listNum = append(listNum, num)
		}
	}
	return listNum
}
func (courApi CourseAPI) AddSeeCourse(userId int, courseId int, pdf string, video string, courseSectionId int) {
	log.Println("<-----------------Inside AddSeeCourse Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT COUNT(*) FROM SeeCourse WHERE CourseId = ? AND SectionId = ? AND UserId = ?", courseId, courseSectionId, userId)
	if err != nil {
		panic(err.Error())
	}
	var count int
	for selDB.Next() {
		if err := selDB.Scan(&count); err != nil {
			log.Fatal(err)
		} else {
			if count > 0 {
				if video != "" {
					insForm, err := courApi.DBcon.Prepare("UPDATE SeeCourse SET Video='Yes' WHERE CourseId= ? AND SectionId = ? AND UserId= ?")
					if err != nil {
						panic(err.Error())
					}
					fmt.Println("UPDATE: SeeCourse: ", userId, " | CourseId: ", courseId, "video Yes")
					insForm.Exec(courseId, courseSectionId, userId)
				} else {
					insForm, err := courApi.DBcon.Prepare("UPDATE SeeCourse SET Pdf='Yes' WHERE CourseId= ? AND SectionId = ? AND UserId= ?")
					if err != nil {
						panic(err.Error())
					}
					fmt.Println("UPDATE: SeeCourse: ", userId, " | CourseId: ", courseId, "Pdf Yes")
					insForm.Exec(courseId, courseSectionId, userId)
				}
			} else {
				if video != "" {
					insForm, err := courApi.DBcon.Prepare("INSERT INTO SeeCourse(CourseId , SectionId , Video , Pdf , UserId) VALUES (?,?,?,?,?)")
					if err != nil {
						panic(err.Error())
					}
					fmt.Println("INSERT: SeeCourse: ", userId, " | CourseId: ", courseId, "Video Yes")
					insForm.Exec(courseId, courseSectionId, "Yes", "No", userId)
				} else {
					insForm, err := courApi.DBcon.Prepare("INSERT INTO SeeCourse(CourseId , SectionId , Video , Pdf , UserId) VALUES (?,?,?,?,?)")
					if err != nil {
						panic(err.Error())
					}
					fmt.Println("INSERT: SeeCourse: ", userId, " | CourseId: ", courseId)
					insForm.Exec(courseId, courseSectionId, "No", "Yes", userId)
				}
			}
		}
	}
}
func (courApi CourseAPI) GetProcessCourse(courseId int, userId int) []SeeCourse {
	log.Println("<-----------------Inside GetProcessCourse Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT * FROM SeeCourse WHERE CourseId= ? AND UserId= ?", courseId, userId)
	if err != nil {
		panic(err.Error())
	}
	var seeCourse = make([]SeeCourse, 0)
	for selDB.Next() {
		see := SeeCourse{}
		var courseId, Id, sectionId, userId int
		var video, pdf string
		err = selDB.Scan(&courseId, &Id, &sectionId, &video, &pdf, &userId)
		if err != nil {
			panic(err.Error())
		}
		see.CourseId = courseId
		see.Id = Id
		see.SectionId = sectionId
		see.Pdf = pdf
		see.Video = video
		see.UserId = userId
		seeCourse = append(seeCourse, see)
	}
	return seeCourse
}
func (courApi CourseAPI) GetALlCourseByUserId(createId int) []Course {
	log.Println("<-----------------Inside GetALlCourseByUserId Resposiroty------------->")
	selDB, err := courApi.DBcon.Query("SELECT C.*, O.StudentJoined  FROM Courses C LEFT JOIN (SELECT CourseId, count(CustomerId) as StudentJoined FROM Orders group by CourseId)  O on C.CourseId = O.CourseId  WHERE CreateId = ?", createId)
	if err != nil {
		panic(err.Error())
	}
	var listCour = make([]Course, 0)
	for selDB.Next() {
		cour := Course{}
		var id, catid, slotvai, slotregis, createId int
		var status int
		var studentJoined *int
		var courPrice float64
		var startDate, endDate time.Time
		var name, description, imageUrl string
		err = selDB.Scan(&id, &name, &catid, &description, &courPrice, &imageUrl, &startDate, &endDate, &slotvai, &slotregis, &createId, &status, &studentJoined)
		if err != nil {
			panic(err.Error())
		}
		cour.CourseId = id
		cour.CourseName = name
		cour.CourseDescription = description
		cour.CategoryID = catid
		cour.CoursePrice = courPrice
		cour.CourseImage = imageUrl
		cour.EndDate = endDate
		cour.StartDate = startDate
		cour.SlotAvailable = slotvai
		cour.SlotRegistered = slotregis
		cour.CreateId = createId
		if cour.StudentJoined == nil {
			a := 0
			cour.StudentJoined = &a
		} else {
			cour.StudentJoined = studentJoined
		}
		if status == 1 {
			cour.Status = true
		} else {
			cour.Status = false
		}

		listCour = append(listCour, cour)
	}
	return listCour
}
