package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
)

type LibraryController struct {
	libraryService services.LibraryManager
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{libraryService: service}
}

func acceptIntInput(prompt string, scanner *bufio.Scanner) int {
	fmt.Print(prompt)
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())
	return id
}
func acceptStrInput(prompt string, scanner *bufio.Scanner) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func (lc *LibraryController) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n-----------------------------------")
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Remove Book")
		fmt.Println("4. Borrow Book")
		fmt.Println("5. Return Book")
		fmt.Println("6. List Available Books")
		fmt.Println("7. List Borrowed Books by Member")
		fmt.Println("0. Exit")
		fmt.Print("-----------------------------------\n")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			id := acceptIntInput("Enter book ID: ", scanner)
			title := acceptStrInput("Enter book title: ", scanner)
			author := acceptStrInput("Enter book author: ", scanner)

			book := models.Book{
				ID:     id,
				Title:  title,
				Author: author,
			}

			lc.libraryService.AddBook(book)
			fmt.Println("Book added successfully!")

		case 2:
			id := acceptIntInput("Enter member ID: ", scanner)
			name := acceptStrInput("Enter member name: ", scanner)

			member := models.Member{
				ID:   id,
				Name: name,
			}

			lc.libraryService.AddMember(member)
			fmt.Println("Member added successfully!")
			
		case 3:
			id := acceptIntInput("Enter book ID: ", scanner)

			err := lc.libraryService.RemoveBook(id)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			fmt.Println("Book removed successfully!")

		case 4:
			bookId := acceptIntInput("Enter book ID: ", scanner)
			memberId := acceptIntInput("Enter member ID: ", scanner)

			err := lc.libraryService.BorrowBook(bookId, memberId)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			fmt.Println("Book borrowed successfully!")

		case 5:
			bookId := acceptIntInput("Enter book ID: ", scanner)
			memberId := acceptIntInput("Enter member ID: ", scanner)

			err := lc.libraryService.ReturnBook(bookId, memberId)
			if err != nil {
				fmt.Println("Error: ", err)
			}else{
				fmt.Println("Book returned successfully!")
			}

		case 6:
			books := lc.libraryService.ListAvailableBooks()

			if len(books) == 0 {
				fmt.Println("No available books in the library.")
			}else{
				fmt.Println("Avaliable Books!")
				for _, book := range books {
					fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
				}
			}			

		case 7:
			id := acceptIntInput("Enter member ID: ", scanner)

			books, err := lc.libraryService.ListBorrowedBooks(id)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			if len(books) == 0 {
				fmt.Println("This member hasn't borrowed any books.")
			}else{
				fmt.Println("Borrowed books by member ", id)
				for _, book := range books {
					fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
				}
			}			

		case 0:
			fmt.Println("Exiting the system...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
