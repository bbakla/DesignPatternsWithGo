package main

type Iterable interface {
	getIterator() iterator
}

type book struct {
	_      struct{}
	Name   string
	Author string
}

type BookStore struct {
	Books []*book
}

func (store *BookStore) getIterator() iterator {

	bookIterator := &BookIterator{
		books: store.Books,
		index: 0,
	}
	return bookIterator
}
