package main

import (
	"fmt"
	"sort"
)

var studentMap = make(map[string]int)

func validateGrade(grade int) bool {
	return grade >= 0 && grade <= 100
}

func AddStudent() {
	var name string
	var grade int

	fmt.Println("Enter student name:")
	fmt.Scan(&name)

	fmt.Println("Enter student grade (0-100):")
	fmt.Scan(&grade)

	if !validateGrade(grade) {
		fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
		return
	}

	studentMap[name] = grade
	fmt.Println("Student details stored successfully.")
}

func getStudentDetails() {
	if len(studentMap) == 0 {
		fmt.Println("No students in the list.")
		return
	}

	var sumOfGrades int
	highestStudent := ""
	highestGrade := 0

	for name, grade := range studentMap {
		sumOfGrades += grade
		if grade > highestGrade {
			highestGrade = grade
			highestStudent = name
		}
	}

	average := float64(sumOfGrades) / float64(len(studentMap))
	fmt.Printf("Total students: %d\n", len(studentMap))
	fmt.Printf("Average grade: %.2f\n", average)
	fmt.Printf("Student with the highest grade: %s (%d)\n", highestStudent, highestGrade)
}

func updateGrade() {
	var name string
	var grade int

	fmt.Println("Enter the name of the student to update:")
	fmt.Scan(&name)

	if _, exists := studentMap[name]; !exists {
		fmt.Printf("No student found with the name '%s'.\n", name)
		return
	}

	fmt.Println("Enter the new grade (0-100):")
	fmt.Scan(&grade)

	if !validateGrade(grade) {
		fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
		return
	}

	studentMap[name] = grade
	fmt.Printf("%s's grade has been updated to %d.\n", name, grade)
}

func setRange() {
	var min, max int

	fmt.Println("Enter minimum grade:")
	fmt.Scan(&min)
	fmt.Println("Enter maximum grade:")
	fmt.Scan(&max)

	if min > max {
		fmt.Println("Invalid range: Minimum should not exceed Maximum.")
		return
	}

	filteredStudents := make(map[string]int)
	for name, grade := range studentMap {
		if grade >= min && grade <= max {
			filteredStudents[name] = grade
		}
	}

	if len(filteredStudents) == 0 {
		fmt.Println("No students found in the specified range.")
	} else {
		fmt.Println("Students within the specified range:", filteredStudents)
	}
}

func removeStudent() {
	var name string

	fmt.Println("Enter the name of the student to remove:")
	fmt.Scan(&name)

	if _, exists := studentMap[name]; !exists {
		fmt.Printf("No student found with the name '%s'.\n", name)
		return
	}

	delete(studentMap, name)
	fmt.Printf("Student '%s' has been removed.\n", name)
}

func sortStudentList() {
	if len(studentMap) == 0 {
		fmt.Println("No students to sort.")
		return
	}

	sortedNames := make([]string, 0, len(studentMap))
	for name := range studentMap {
		sortedNames = append(sortedNames, name)
	}

	sort.Strings(sortedNames)
	fmt.Println("Sorted student list:", sortedNames)
}

func main() {
	for {
		fmt.Println(`
Welcome to the university admin system. What do you want to do?

[1] Add Student
[2] Get Students Stats
[3] Update Grade
[4] Filter by Range
[5] Remove Student
[6] Sort List
[7] Exit

Enter your choice:`)

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			AddStudent()
		case 2:
			getStudentDetails()
		case 3:
			updateGrade()
		case 4:
			setRange()
		case 5:
			removeStudent()
		case 6:
			sortStudentList()
		case 7:
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
