# Library Management System Documentation

## Overview

This is a minimal console-based Library Management System written in Go.  
The application uses an in-memory store (maps) to manage books and members.

## Module

- Module name: `library_management`
- Go version: `1.25.4`

## Project Structure

- `main.go`: application entry point and seed data setup.
- `models/book.go`: `Book` domain model.
- `models/member.go`: `Member` domain model.
- `services/library_services.go`: business logic and service interface.
- `controllers/library_controllers.go`: console menu and input/output handlers.
- `docs/documentation.go`: docs package placeholder.

## Data Models

### Book

Defined in `models/book.go`:

- `ID int`
- `Title string`
- `Author string`
- `Status string` (`"Available"` or `"Borrowed"` in current logic)

### Member

Defined in `models/member.go`:

- `ID int`
- `Name string`
- `BorrowedBooks []Book`

## Service Layer

The `LibraryManager` interface in `services/library_services.go` defines:

- `AddBook(book models.Book) error`
- `RemoveBook(bookID int) error`
- `BorrowBook(bookID int, memberID int) error`
- `ReturnBook(bookID int, memberID int) error`
- `ListAvailableBooks() []models.Book`
- `ListBorrowedBooks(memberID int) []models.Book`
- `AddMember(member models.Member) error`

`Library` is the concrete implementation using:

- `Books map[int]models.Book`
- `Members map[int]models.Member`

### Behavior Summary

- `AddBook`: adds a new book, rejects duplicate IDs, sets status to `"Available"`.
- `RemoveBook`: removes by ID, errors if book does not exist.
- `BorrowBook`: validates book/member existence and availability, marks as `"Borrowed"`, appends to member borrowed list.
- `ReturnBook`: ensures member exists and owns the book, removes from borrowed list, marks book as `"Available"`.
- `ListAvailableBooks`: returns books with status `"Available"`.
- `ListBorrowedBooks`: returns borrowed books for a member if found, otherwise `nil`.

## Controller Layer (Console UI)

`LibraryController.MainMenu()` provides:

1. Add a Book
2. Remove a Book
3. Borrow a Book
4. Return a Book
5. List Available Books
6. List Borrowed by Member
7. Exit

Each option maps to a handler that collects input via `fmt.Scanln` and calls the service.

## Application Startup

In `main.go`:

1. Create service with `services.NewLibrary()`.
2. Seed initial data:
   - Member: `ID=1`, `Name="John Doe"`
   - Book: `ID=1`, `Title="The Go Programming Language"`, `Author="Donovan & Kernighan"`
3. Create controller with the service.
4. Start interactive menu loop.

## How to Run

From `task-3_library_management`:

```bash
go run .
```

Or:

```bash
go build ./...
./library_management
```

## Current Limitations

- Data is in-memory only; restarting the app clears all data.
- No concurrency protection (not safe for parallel access).
- Input handling uses `Scanln`, so multi-word titles/authors may not be captured as expected.
- `handleListAvailableBooks` prints `Count: %d` without passing a value, so output formatting is inconsistent.
- `RemoveBook` does not currently prevent removing books that are already borrowed.

## Suggested Next Improvements

- Add persistence (file or database).
- Add validation for empty titles/authors and invalid IDs.
- Improve input parsing for full-line text.
- Track quantities/copies explicitly if needed.
- Add unit tests for service methods and edge cases.
