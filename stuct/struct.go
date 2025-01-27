package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

type Student struct {
	name  string
	age   int
	grade string
}

type Account struct {
	Name    string
	Balance float64
}

func (book Book) getDatails() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", book.Title, book.Author, book.Pages)
}

func (st1 Student) copareGrades(st2 Student) {
	if st1.grade > st2.grade {
		fmt.Println(st1.name, "has a higher grade than", st2.name)
		return
	}
	fmt.Println(st2.name, "has a higher grade than", st1.name)
}

func upgradeBalance(account *Account, amount float64) {
	account.Balance += amount

}
func main() {

	book1 := Book{
		Title:  "The Alchemist",
		Author: "Paulo Coelho",
		Pages:  163}

	fmt.Println(book1)

	student1 := Student{name: "John", age: 25, grade: "A"}
	fmt.Println(student1)
	student2 := Student{name: "Ama", age: 25, grade: "A"}
	student1.copareGrades((student2))
}
