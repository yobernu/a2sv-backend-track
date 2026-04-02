package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book

	//
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) error {
	if _, exists := l.Books[book.ID]; exists {
		return errors.New("book ID already exists")
	}
	book.Status = "Available"
	l.Books[book.ID] = book
	return nil
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.Books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bExists := l.Books[bookID]
	member, mExists := l.Members[memberID]

	if !bExists {
		return errors.New("book not found")
	}
	if !mExists {
		return errors.New("member not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, mExists := l.Members[memberID]
	if !mExists {
		return errors.New("member not found")
	}

	found := false
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("member does not have this book")
	}

	book := l.Books[bookID]
	book.Status = "Available"
	l.Books[bookID] = book
	l.Members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var avail []models.Book
	for _, b := range l.Books {
		if b.Status == "Available" {
			avail = append(avail, b)
		}
	}
	return avail
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	if m, exists := l.Members[memberID]; exists {
		return m.BorrowedBooks
	}
	return nil
}

func (l *Library) AddMember(member models.Member) error {
	if _, exists := l.Members[member.ID]; exists {
		return errors.New("member ID already exists")
	}
	l.Members[member.ID] = member
	return nil
}
