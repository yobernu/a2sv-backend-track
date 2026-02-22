package main

import (
	"fmt"
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	libraryService := services.NewLibrary()
	seedData(libraryService)
	controller := controllers.LibraryController{
		Service: libraryService,
	}
	fmt.Println("Welcome to the A2SV Library System")
	controller.MainMenu()
}
func seedData(s *services.Library) {
	s.Members[1] = models.Member{
		ID:   1,
		Name: "John Doe",
	}
	s.AddBook(models.Book{
		ID:     1,
		Title:  "The Go Programming Language",
		Author: "Donovan & Kernighan",
	})
}
