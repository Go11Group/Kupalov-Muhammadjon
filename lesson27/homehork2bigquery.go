package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "study"
	password = "root"
)

type TopStudents struct{
	CourseName string
	Grade float64
	students []byte
}

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	q := `
	with students_avg_grade as (
		select 
			s.id as student_id, sc.course_id as course_id, round(avg(g.grade)::numeric, 2) as avarage_grade
		from student as s
		right join 
			student_course as sc
		on 
			s.id = sc.student_id
		right join 
			grade as g
		on 
			sc.id = g.student_course_id
		group by 
			sc.course_id, s.id
	),
	max_scores as (
		select course_id, max(avarage_grade) as max_grade
		from students_avg_grade
		group by course_id
	)
	
	select c.name, m.max_grade, array_agg(s.name)
	from course as c
	join max_scores as m
	on m.course_id = c.id
	join students_avg_grade as sag
	on sag.course_id = c.id and m.max_grade = sag.avarage_grade
	join student as s
	on sag.student_id = s.id 
	group by c.name, m.max_grade;
	`
	rows, err := db.Query(q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	topStudents := []TopStudents{}

	for rows.Next(){
		gp := TopStudents{}
		err = rows.Scan(&gp.CourseName, &gp.Grade, &gp.students)
		if err != nil {
		panic(err)
		}
		topStudents = append(topStudents, gp)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	for _, grp := range topStudents{
		fmt.Println(grp.CourseName, grp.Grade, string(grp.students))
	}
}
