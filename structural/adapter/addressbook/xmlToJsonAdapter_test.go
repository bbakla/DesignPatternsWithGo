package addressbook

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const xmlContact = `
<contact>
	<id>325</id>
	<name>
		<first>BürgerSteig</first>
		<last>IstImmerVernünftig</last>
	</name>
</contact>`

func TestCanConvertAnXmlToJsonContact(t *testing.T) {
	jsonContact := convertXmlContactToJsonContact(xmlContact)

	assert.NotEqual(t, "", jsonContact.FirstName)
	assert.NotEqual(t, "", jsonContact.LastName)
	assert.NotEqual(t, 0, jsonContact.Id)
}

func TestCanAddXmlContactToJsonContactList(t *testing.T) {
	NewJsonAddressBook()
	addXmlContactToJsonContactList(xmlContact)

	assert.Equal(t, 1, len(jsonContacts.contactList))
}
