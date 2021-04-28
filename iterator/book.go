package iterator

import "fmt"

type Book struct {
	Id   int64
	Name string
}

func (book Book) String() string {
	return fmt.Sprintf("ID: %d, Name: %s", book.Id, book.Name)
}
