package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAddJsonContact(t *testing.T) {
	contact := `
	{
	
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
	  }
	`

	var jsonAddressBook = NewAddressBookJson()
	id, err := jsonAddressBook.addContact(contact)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}
