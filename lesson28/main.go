package main

import (
	"log"
	"module/printer"
	"module/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	
	st := postgres.NewStudentRepo(db)
	students, err := st.GetStudents()
	if err != nil {
		log.Println("Can not get students")
	}

	(*students)[1].Age = 16
	st.UpdateStudent((*students)[1])

	printer.PrintStudents(students)
	// st.CreateStudent("Javlon", 16)
	// students, err = st.GetStudents()
	// if err != nil {
	// 	log.Println("Can not get students")
	// }
	// printer.PrintStudents(students)

	// courses, err := st.GetCoursesOfStudent((*students)[0].Id)
	// if err != nil {
	// 	log.Println("Can not get courses of student")
	// }
	// printer.PrintCoursesOfStudent(courses)	st.UpdateStudent((*students)[11])


	// grades, err := st.GetGrades((*students)[0].Id)
	// if err != nil {
	// 	log.Println("Cannnot get grades")
	// }
	// printer.PrintGradesOfStudent(grades)
	c := postgres.NewCourseRepo(db)
	courses, err := c.GetCourses()
	if err != nil {
		log.Println("Can not get courses")
	}
	printer.PrintCourses(courses)
	courses, err = c.ShowRatingOfCourses()
	if err != nil {
		log.Println("Can not get rating of courses")
	}
	printer.PrintRatingOfCourses(courses)

}
