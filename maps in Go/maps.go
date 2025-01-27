package main

import (
	"fmt"
)

/*
Question
Write a Go program that allows a teacher to track the grades of students using a map. The program should:

Prompt the user to enter the number of students.
Allow the user to input each student's name and grade (out of 100).
Store the data in a map where the key is the student's name and the value is their grade.
Provide the following functionalities:
Print all students and their grades.
Calculate and display the average grade of all students.
Find and display the name of the student with the highest grade.
Allow the teacher to update a specific student's grade.
Use appropriate error handling for invalid inputs or non-existent students during the update operation.
*/
func main() {
	var numOfStudents int
	fmt.Println("Enter the number of Students\n")
	fmt.Scan(&numOfStudents)
	var studentMap = make(map[string]int)
	var sumOfGrades int
	average := 0.0
	var hightStudent = ""
	currentHighestGrade := 0
	for i := 0; i < numOfStudents; i++ {
		studentName := ""
		studentGrade := 0

		fmt.Println("Enter studnets name\n")
		fmt.Scan(&studentName)
		fmt.Println("Enter studnets grade\n")
		fmt.Scan(&studentGrade)
		if studentGrade < 0 || studentGrade > 100 {
			fmt.Printf("Please enter a grade between 0 and 100\n")
			fmt.Scan(&studentGrade)
			studentMap[studentName] = studentGrade
		} else {
			fmt.Println("Information stored\n")
			studentMap[studentName] = studentGrade
			sumOfGrades += studentGrade
			if studentGrade < currentHighestGrade {
				hightStudent = studentName
			}
		}
	}

	fmt.Println(studentMap)
	average = float64(sumOfGrades / numOfStudents)
	fmt.Printf("Avergae is %f\n", average)
	fmt.Printf("Student with the highest grade is %s \n", hightStudent)

	var choice string
	fmt.Printf("Do you want to change a students score?\n")
	fmt.Scan(&choice)
	switch choice {
	case "Yes", "Y", "y":
		var newGrade int
		name := ""
		fmt.Printf("Enter student name")
		fmt.Scan(&name)
		fmt.Printf("Enter new grade for %s", name)
		fmt.Scan(&newGrade)
		studentMap[name] = newGrade
		fmt.Printf("%s grade has been changed to %d", name, newGrade)
	case "No", "N", "n":
		break
	default:
		fmt.Printf("Wrong choice please enter Yes or No")
	}
}
