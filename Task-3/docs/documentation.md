# Library Management System Documentation

## Overview
This is a console-based library management system implemented in Go. It allows users to manage books and track borrowing activities.

## Features
1. **Book Management**
   - Add new books to the library
   - Remove books from the library
   - List all available books

2. **Member Management**
   - Track books borrowed by members
   - Handle book borrowing and returning

## Technical Details
- **Data Structures**:
  - Books are stored in a map with book ID as the key
  - Members are stored in a map with member ID as the key
  - Borrowed books are tracked in a slice within each member

- **Error Handling**:
  - Book not found
  - Member not found
  - Book already borrowed
  - Member hasn't borrowed the book being returned

## Usage
Run the program and follow the menu prompts to perform various operations.

## Future Enhancements
1. Add member registration functionality
2. Implement persistent storage (database/file)
3. Add due dates and late return tracking
4. Implement search functionality
5. Add user authentication