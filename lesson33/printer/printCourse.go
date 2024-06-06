package printer

import (
	"fmt"
	"module/model"
)

func PrintCourses(courses *[]model.Course){
	for _, course := range *courses {
		fmt.Println(course.Name)
	}
	fmt.Println()
}

func PrintRatingOfCourses(courses *[]model.Course){
	for _, course := range *courses {
		fmt.Println(course.Name, course.AvgGrade)
	}
	fmt.Println()
}