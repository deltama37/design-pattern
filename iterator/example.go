package iterator

import "fmt"

type Example struct{}

func (_ Example) Run() {
	books := NewBooks()
	books.Append(&Book{
		Id:   1,
		Name: "harry potter",
	})
	books.Append(&Book{
		Id:   2,
		Name: "road of the ring",
	})

	for books.HasNext() {
		book := books.Next()
		fmt.Println(book)
	}
}
