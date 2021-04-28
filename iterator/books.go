package iterator

type Books struct {
	items    []*Book
	iterator *BooksIterator
}

func NewBooks() *Books {
	books := &Books{
		iterator: &BooksIterator{},
	}
	books.iterator.Books = books
	return books
}

func (books *Books) GetItemAt(idx int64) *Book {
	if books.GetSize() <= idx {
		return nil
	}
	return books.items[idx]
}

func (items *Books) GetSize() int64 {
	return int64(len(items.items))
}

func (books *Books) HasNext() bool {
	return books.iterator.HasNext()
}

func (books *Books) Next() *Book {
	return books.iterator.Next()
}

func (books *Books) Append(book *Book) {
	books.items = append(books.items, book)
}
