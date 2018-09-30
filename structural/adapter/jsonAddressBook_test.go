package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAddJsonContact(t *testing.T) {
	contact := `
	{
		"adress": [
		  {
			"country": "GErmany",
			"city": "Nurnberg"
		  },
		   {
			"country": "Turkey",
			"city": "Istanbul"
		  }
		],
		"id": "",
		"name": {
		  "first": "b",
		  "last": "d"
		}
	  }
	`

	var jsonAddressBook = NewAddreeBookJson()
	id, err := jsonAddressBook.addContact(contact)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}
