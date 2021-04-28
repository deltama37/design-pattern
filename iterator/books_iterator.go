package iterator

type BooksIterator struct {
	Books *Books
	Idx   int64
}

func (iterator *BooksIterator) HasNext() bool {
	if iterator.Idx < iterator.Books.GetSize() {
		return true
	}
	return false
}

func (iterator *BooksIterator) Next() *Book {
	book := iterator.Books.GetItemAt(iterator.Idx)
	iterator.Idx += 1
	return book
}
