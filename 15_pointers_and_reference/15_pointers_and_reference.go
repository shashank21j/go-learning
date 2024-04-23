package main

import "fmt"

type Book struct {
	title  string
	author string
}

func changeTitle(book *Book) {
	book.title = "New Title"
	// (*book).title = "New Title"
}

func (b *Book) changeAuthor() {
	b.author = "New Author"
	// (*book).author = "New Author"
}

func change(s []int) {
	s[0] = 100
}

func change_int(num *int) {
	*num = 20
}

func main() {
	x := 2
	x_ptr := &x
	fmt.Println("x: ", x, "x_ptr: ", x_ptr)

	var y_ptr *int = &x
	fmt.Println("y_ptr: ", y_ptr, "*y_ptr: ", *y_ptr)

	s := []int{1, 2, 3, 4, 5}
	change(s)
	fmt.Println("s: ", s)

	num := 10
	change_int(&num)
	fmt.Println("num: ", num)

	x_1 := []string{"a", "b", "c"}
	x_2 := &x_1
	x_3 := &x_2
	fmt.Println("x_1: ", x_1, "*x_2: ", *x_2, "**x_3: ", **x_3)

	book := Book{"Go Programming", "John Doe"}
	book_ptr := &book
	fmt.Println("book: ", book, "book_ptr: ", book_ptr)
	changeTitle(book_ptr)
	fmt.Println("book: ", book, "book_ptr: ", book_ptr)

	book.changeAuthor()
	fmt.Println("book: ", book, "book_ptr: ", book_ptr)

	p1 := 0
	p2 := 1
	p_slice := []*int{&p1, &p2}
	fmt.Println("p_slice: ", p_slice)
}
