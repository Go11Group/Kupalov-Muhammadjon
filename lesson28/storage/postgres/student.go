package postgres

import (
	"database/sql"
	"module/model"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{db}
}

// Create
func (s *StudentRepo) CreateStudent(name string, age int) {
	_, err := s.db.Exec(`Insert into student(name, age) values($1, $2)`, name, age)
	if err != nil {
		panic(err)
	}
}

// Read
func (s *StudentRepo) GetStudents() (*[]model.Student, error) {
	rows, err := s.db.Query(`select * from student`)
	if err != nil {
		return nil, err
	}
	students := []model.Student{}
	student := model.Student{}

	for rows.Next() {
		err = rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}

func (s *StudentRepo) GetCoursesOfStudent(id string) (*[]string, error) {
	query := `
	select c.name
	from
		student_course as sc 
	join 
		course as c
	on
		sc.course_id = c.id and $1 = sc.student_id 
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	courses := []string{}
	for rows.Next() {
		var course string
		err := rows.Scan(&course)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return &courses, nil

}

func (s *StudentRepo) GetGrades(id string) (*[]model.Grade, error) {
	query := `
	select c.name, g.grade
	from
		student_course as sc 
	join 
		course as c
	on
		sc.course_id = c.id and $1 = sc.student_id 
	join 
		grade as g
	on
		g.student_course_id = sc.id
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	grades := []model.Grade{}
	grade := model.Grade{}
	for rows.Next() {
		err := rows.Scan(&grade.CourseName, &grade.Points)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return &grades, nil
}

// Update
func (s *StudentRepo) UpdateStudent(student model.Student) error {
	query := `
	update
		student 
	set 
		name=$1, 
		age=$2
	where
		id=$3
	`
	_, err := s.db.Exec(query, student.Name, student.Age, student.Id)
	return err
}

func (s *StudentRepo) AssignCourse(student_id, course_id string) error {
	query := `
	insert into
		student_course(student_id, course_id)
	values(
		$1, $2
	)
	`
	_, err := s.db.Exec(query, student_id, course_id)
	return err
}

// Delete
func (s *StudentRepo) DeleteStudent(id string) error{
	query := `
	delete from
		student
	where
		id=$1
	`
	_, err := s.db.Exec(query, id)
	return err
}

// func (s *StudentRepo) DeleteCourseOfStudent(id string) {

// }
