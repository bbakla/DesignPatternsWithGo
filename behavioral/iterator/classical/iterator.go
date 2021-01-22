package main

type iterator interface {
	next() *book
	hasNext() bool
}

type BookIterator struct {
	books []*book
	index int
}

func (i *BookIterator) next() *book {
	b := i.books[i.index]
	i.index++
	return b
}

func (i *BookIterator) hasNext() bool {
	return i.index < len(i.books)
}
