package model

type Course struct {
	CourseId    string `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Time
}

type CourseFilter struct {
	Title  *string
	Limit  *int
	Offset *int
}
