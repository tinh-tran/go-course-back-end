package models

type Section struct {
	CourseId        int     `json:"CourseId,omitempty" `
	ChapterId       int     `json:"ChapterId"`
	SectionName     string  `json:"SectionName"`
	SectionVideo    *string `json:"SectionVideo"`
	SectionPdf      *string `json:"SectionPdf"`
	CourseSectionId int     `json:"CourseSectionId"`
	SectionOrder    int     `json:"SectionOrder"`
}
