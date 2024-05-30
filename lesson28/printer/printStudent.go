package printer

import (
	"fmt"
	"module/model"
)

func PrintStudents(students *[]model.Student){
	for _, student := range *students{
		fmt.Println(student.Name, student.Age)
	}
	fmt.Println()
}

func PrintCoursesOfStudent(courses *[]string){
	for _, course := range *courses{
		fmt.Println(course)
	}
	fmt.Println()
}
func PrintGradesOfStudent(gardes *[]model.Grade){
	for _, garde := range *gardes{
		fmt.Println(garde.CourseName, garde.Points)
	}
	fmt.Println()
}
