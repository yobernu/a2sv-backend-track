package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	Service services.LibraryManager
}

func (c *LibraryController) MainMenu() {
	for {
		fmt.Println("\n--- Library Management System ---")
		fmt.Println("1. Add a Book")
		fmt.Println("2. Remove a Book") // Added to menu
		fmt.Println("3. Borrow a Book")
		fmt.Println("4. Return a Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed by Member")
		fmt.Println("7. Exit")
		fmt.Print("Select an option: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			c.handleAddBook()
		case 2:
			c.handleRemoveBook()
		case 3:
			c.handleBorrowBook()
		case 4:
			c.handleReturnBook()
		case 5:
			c.handleListAvailableBooks()
		case 6:
			c.handleListBorrowedBooks()
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Option not recognized. Please choose between 1-7")
		}
	}
}

func (c *LibraryController) handleAddBook() {
	var id int
	var title, author string
	fmt.Print("Enter Book ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter Title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter Author: ")
	fmt.Scanln(&author)

	newBook := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}

	err := c.Service.AddBook(newBook)
	if err != nil {
		fmt.Printf("Failed to add book: %v\n", err)
	} else {
		fmt.Println("Book added successfully!")
	}
}

func (c *LibraryController) handleListBorrowedBooks() {
	var memberID int
	fmt.Print("Enter Member ID: ")
	fmt.Scanln(&memberID)
	books := c.Service.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No books borrowed by this member.")
		return
	}
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) handleRemoveBook() {
	var bookID int
	fmt.Print("Enter Book ID to remove: ")
	fmt.Scanln(&bookID)
	err := c.Service.RemoveBook(bookID)
	if err != nil {
		fmt.Printf("Failed to remove book: %v\n", err)
	} else {
		fmt.Println("Book removed successfully!")
	}
}

func (c *LibraryController) handleBorrowBook() {
	var bookID, memberID int
	fmt.Print("Enter Book ID to borrow: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter Member ID: ")
	fmt.Scanln(&memberID)

	err := c.Service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Failed to borrow book: %v\n", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func (c *LibraryController) handleReturnBook() {
	var bookID, memberID int
	fmt.Print("Enter Book ID to return: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter Member ID: ")
	fmt.Scanln(&memberID)

	err := c.Service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Failed to return book: %v\n", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func (c *LibraryController) handleListAvailableBooks() {
	books := c.Service.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books at the moment.")
		return
	}
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Count: %d\n", book.ID, book.Title, book.Author)
	}
}
