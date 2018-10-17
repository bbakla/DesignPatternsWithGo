package adapter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

var jsonContacts jsonContactList

func NewJsonAddressBook() {
	jsonContacts.contactList = make(map[int]jsonContact)
}

func convertXmlContactToJsonContact(xmlContact string) Contact {
	var entity Contact

	err := xml.Unmarshal([]byte(xmlContact), &entity)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	return entity
}

func addXmlContactToJsonContactList(xmlContact string) {
	var jsonContact jsonContact

	contactInXml := convertXmlContactToJsonContact(xmlContact)

	jsonContact.Id = contactInXml.Id
	jsonContact.FirstName = contactInXml.FirstName
	jsonContact.LastName = contactInXml.LastName

	stringFormat, err := json.Marshal(jsonContact)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	jsonContacts.addContact(string(stringFormat))

}
