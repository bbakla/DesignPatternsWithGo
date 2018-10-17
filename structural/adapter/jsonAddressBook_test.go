package adapter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const contact = `{
	
	"id": 0,
	"name": {
		"first": "b",
		"last": "d"
	},
	"addresses": [
		{
		"country": "GErmany",
		"city": "Nurnberg"
		},
		 {
		"country": "Turkey",
		"city": "Istanbul"
		}
	]
	}`

const contact2 = `	{
	
		"id": 1,
		"name": {
			"first": "dieFreude",
			"last": "db"
		},
		"addresses": [
			{
			"country": "GErmany",
			"city": "Nurnberg"
			},
			 {
			"country": "Turkey",
			"city": "Istanbul"
			}
		]
		}`

func TestCanAddJsonContact(t *testing.T) {
	var jsonAddressBook = NewAddressBookJson()
	id, err := jsonAddressBook.addContact(contact)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}

func TestCanAddPrintJsonContact(t *testing.T) {

	var jsonAddressBook = NewAddressBookJson()
	jsonAddressBook.addContact(contact)
	jsonAddressBook.addContact(contact2)

	jsonAddressBook.print()
}

func TestCanGetTheContactSavedBeforeToJsonContactList(t *testing.T) {
	var jsonAddressBook = NewAddressBookJson()
	id, _ := jsonAddressBook.addContact(contact)

	contactInString := jsonAddressBook.get(id)
	fmt.Println(contactInString)

	assert.NotEqual(t, "", contactInString)
}

func TestCanGetAllJsonContacts(t *testing.T) {

	var jsonAddressBook = NewAddressBookJson()
	jsonAddressBook.addContact(contact)
	jsonAddressBook.addContact(contact2)

	contactsInString := jsonAddressBook.getAll()

	assert.True(t, strings.Contains(contactsInString, "dieFreude"))
}
