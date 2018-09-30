package adapter

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"strings"
)

type xmlProcessor interface {
	print()
	addContact() error
	get(id int) string
	getAll() string
}

type Contacts struct {
	XmlName  xml.Name  `xml:"contacts"`
	Contacts []Contact `xml:"contact"`
}
type Contact struct {
	XMLName xml.Name `xml:"contact"`
	Id      int      `xml:id, attr`
	Name
}

type Name struct {
	FirstName string `xml:"name>first"`
	LastName  string `xml:"name>last"`
}
type addressBookInXml struct {
	people map[int]Contact
}

func NewAddressBookInXml() addressBookInXml {
	var addressBook addressBookInXml
	addressBook.people = make(map[int]Contact)

	return addressBook
}

func (xmlAddressBook addressBookInXml) print() error {
	for key, value := range xmlAddressBook.people {
		fmt.Printf("Id = %d\t name = %s\n", key, value.FirstName+" "+value.LastName)
	}

	return nil
}

func (xmlAddressBook addressBookInXml) addContact(contact string) (int, error) {
	var entity Contact
	var id int

	err := xml.Unmarshal([]byte(contact), &entity)
	if err != nil {
		return id, err
	}

	id = rand.Intn(10000)
	entity.Id = id

	xmlAddressBook.people[id] = entity

	return id, nil
}

func (xmlAddressBook addressBookInXml) get(id int) (string, error) {
	contact := xmlAddressBook.people[id]

	contactInString, err := xml.MarshalIndent(contact, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err

	}
	return string(contactInString), nil
}

func (xmlAddressBook addressBookInXml) getAll() (string, error) {

	var contactsAsString strings.Builder
	// contactsAsString.WriteString("<contacts>")

	for key, _ := range xmlAddressBook.people {
		contact, _ := xmlAddressBook.get(key)
		contactsAsString.WriteString(contact)
		contactsAsString.WriteString("\n")
	}

	// contactsAsString.WriteString("</contacts>")

	return contactsAsString.String(), nil

}

func (xmlAddressBook addressBookInXml) getAll2() (string, error) {

	// var contactsAsString strings.Builder
	// // contactsAsString.WriteString("<contacts>")
	var contacts Contacts
	contacts.Contacts = make([]Contact, 1)

	for key, _ := range xmlAddressBook.people {
		contact, _ := xmlAddressBook.get(key)
		contactsAsString.WriteString(contact)
		contactsAsString.WriteString("\n")
	}

	// contactsAsString.WriteString("</contacts>")

	return contactsAsString.String(), nil

}
