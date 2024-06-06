package postgres

import (
	"database/sql"
	"module/model"
)

type CoursesRepo struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CoursesRepo {
	return &CoursesRepo{db: db}
}

// Create
func (s *CoursesRepo) CreateCourse(name string) error {
	_, err := s.db.Exec(`Insert into course(name) values($1)`, name)

	return err
}

// Read
func (s *CoursesRepo) GetCourses() (*[]model.Course, error) {
	rows, err := s.db.Query(`select * from course`)
	if err != nil {
		return nil, err
	}
	courses := []model.Course{}
	course := model.Course{}

	for rows.Next() {
		err = rows.Scan(&course.Id, &course.Name)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return &courses, nil
}

func (s *CoursesRepo) ShowRatingOfCourses() (*[]model.Course, error) {
	query := `
	select c.name, round(avg(g.grade)::numeric, 2) as avarage_grade
	from course as c
	right join student_course as sc
	on sc.course_id = c.id
	right join grade as g
	on sc.id = g.student_course_id
	group by c.name
	order by avarage_grade desc
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	courses := []model.Course{}
	course := model.Course{}

	for rows.Next() {
		err = rows.Scan(&course.Name, &course.AvgGrade)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return &courses, nil
}

// func (s *CoursesRepo) GetStudentsOfCourse(id string) {

// }

// Update
func (s *CoursesRepo) UpdateCourseName(course model.Course) error {
	query := `
	update
		student 
	set 
		name=$1, 
	where
		id=$2
	`
	_, err := s.db.Exec(query, course.Name, course.Id)
	return err
}

// Delete
func (s *CoursesRepo) DeleteCourse(id string) error {
	query := `
	delete from
		course
	where
		id=$1
	`
	_, err := s.db.Exec(query, id)
	return err
}

// func (s *CoursesRepo) DeleteStudentOfCourse(id string){

// }
