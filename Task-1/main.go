package main

import (
	"fmt"
	"math"
)

func convertGrade(c int) (string) {
	switch  {
		case c >= 90:
			return "A"
		case c >= 80:
			return "B"
		case c >= 70:
			return "C"
		case c >= 60:
			return "D"
		default:
			return "F"
	}
}

func calculateGrade(grades map[string]string) (string) {
	var total int

	for _, grade := range grades {
		switch grade {
			case "A":
				total += 4
			case "B":
				total += 3
			case "C":
				total += 2
			case "D":
				total += 1
			case "F":
				total += 0
		}
	}

	if len(grades) == 0{
		return "Please enter valid data"
	}

	average := total  / len(grades)
	average = int(math.Round(float64(average)))

	switch average{
		case 4:
			return "A"
		case 3:
			return "B"
		case 2:
			return "C"
		case 1:
			return "D"
		default:
			return "F"
	}
}

func isValidGrade(grade int) bool {
	return grade >= 0 && grade <= 100
}

func main() {
	var name string
	var noOfCourse int
	courseGrades := make(map[string]string)

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter the number of courses: ")
	fmt.Scan(&noOfCourse)

	if noOfCourse <= 0 {
		fmt.Println("Invalid number of courses. Please enter a positive integer.")
		return
	}

	fmt.Println("Enter the grades for each course:")
	for i := 0; i < noOfCourse; i++ {
		var courseName string
		var grade int

		fmt.Print("Enter course name: ")
		fmt.Scan(&courseName)

		fmt.Print("Enter grade for ", courseName, ": ")
		fmt.Scan(&grade)

		if !isValidGrade(grade) {
			fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
			i-- // Decrement i to repeat this iteration
			continue
		}

		courseGrades[courseName] = convertGrade(grade)
	}

	fmt.Println("\nYour grade despaly as follows:")
	fmt.Println("Name:", name)
	fmt.Printf("\nName of course\t Grade\n")
	for course,grades := range(courseGrades){
		fmt.Printf("%s\t\t %s\n", course, grades)
	}

	fmt.Println("Your final grade is:\t", calculateGrade(courseGrades))
}