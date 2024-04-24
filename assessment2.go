package main

// Write your code here.
type Book struct {
	id       int
	title    string
	author   string
	quantity int
}

type Library struct {
	books []*Book
}

func (l *Library) CheckoutBook(id int) (*Book, bool) {
	for _, book := range l.books {
		if book.id == id && book.quantity > 0 {
			book.quantity--
			return book, true
		}
	}
	return nil, false
}

func (l *Library) ReturnBook(id int) bool {
	for _, book := range l.books {
		if book.id == id {
			book.quantity++
			return true
		}
	}
	return false
}
