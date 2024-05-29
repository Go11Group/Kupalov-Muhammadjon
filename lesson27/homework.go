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

type GroupAvarage struct {
	Name         string
	AvarageGrade float64
}

type TopStudents struct{
	CourseName string
	Grade float64
	students []string
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
	query := `
	select c.name, round(avg(g.grade)::numeric, 2) as avarage_grade
	from course as c
	right join student_course as sc
	on sc.course_id = c.id
	right join grade as g
	on sc.id = g.student_course_id
	group by c.name
	order by avarage_grade desc;`
	row := db.QueryRow(query)
	ga := GroupAvarage{}
	err = row.Scan(&ga.Name, &ga.AvarageGrade)
	if err != nil {
		panic(err)
	}
	fmt.Println(ga)
	
	/*
	QueryRow faqat bitta ma'lumot olib kela oladi. Agar biz ga ko'plab malumot kelsa bir necha qatorli Query
	ishlatiladi. Ya'ni query qatorlarni qaytaradi va biz for orqali ularni birma bir o'zlashtira olamiz
	*/
	groups := []GroupAvarage{}
	rows, err := db.Query(query)
	if err != nil{
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		g := GroupAvarage{}

		err = rows.Scan(&g.Name, &g.AvarageGrade)
		if err != nil {
			panic(err)
		}

		groups = append(groups, g)
	}
	err = row.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println(groups)
}
