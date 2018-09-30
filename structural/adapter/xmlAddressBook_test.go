package adapter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAddContact(t *testing.T) {
	data := `
	<contact>
		<Id></Id>
		<name>
			<first>BürgerSteig</first>
			<surname>IstImmerVernünftig</surname>
		</name>
	</contact>`

	var addressBook = NewAddressBookInXml()
	id, err := addressBook.addContact(data)

	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}

func TestCanPrintAllContacts(t *testing.T) {
	data := `
	<contact>
		<Id></Id>
		<name>
			<first>BürgerSteig</first>
			<last>IstImmerVernünftig</last>
		</name>
	</contact>`

	var addressBook = NewAddressBookInXml()
	_, err := addressBook.addContact(data)
	assert.Nil(t, err)

	err = addressBook.print()
	assert.Nil(t, err)
}

func TestCanGetContact(t *testing.T) {
	data := `
	<contact>
		<Id></Id>
		<name>
			<first>BürgerSteig</first>
			<last>IstImmerVernünftig</last>
		</name>
	</contact>`

	var addressBook = NewAddressBookInXml()
	id, err := addressBook.addContact(data)
	assert.Nil(t, err)

	savedContact, err := addressBook.get(id)

	fmt.Println(savedContact)

	assert.Nil(t, err)
	assert.NotEqual(t, "", savedContact)
}

func TestCanGetAllContacts(t *testing.T) {
	data := `
	<contact>
		<Id></Id>
		<name>
			<first>BürgerSteig</first>
			<last>IstImmerVernünftig</last>
		</name>
	</contact>`

	data2 := `
	<contact>
		<Id></Id>
		<name>
			<first>Laufsteg</first>
			<last>IchhabemneinesHandyaufgehoben</last>
		</name>
	</contact>`

	var addressBook = NewAddressBookInXml()
	addressBook.addContact(data)
	addressBook.addContact(data2)

	contacts, err := addressBook.getAll()

	fmt.Println(contacts)

	assert.Nil(t, err)
	assert.True(t, strings.Contains(contacts, "Laufsteg"))

}
