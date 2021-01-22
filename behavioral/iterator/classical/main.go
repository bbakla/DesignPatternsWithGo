package main

import "fmt"

func main() {
	books := []*book{
		&book{
			Name:   "Ölüler Evi",
			Author: "Ali Veli",
		},

		&book{
			Name:   "Canlilar kaciyor",
			Author: "Selim Can",
		},

		&book{
			Name:   "Kacmayanlar ölüyor zaten",
			Author: "Karacaahmet",
		},
	}

	bookStore := &BookStore{Books: books}
	booksIterator := bookStore.getIterator()

	fmt.Printf("%15s%40s\n", "Name", "Author")

	for booksIterator.hasNext() {
		book := booksIterator.next()
		fmt.Printf("%15s%40s\n", book.Name, book.Author)
	}
}
