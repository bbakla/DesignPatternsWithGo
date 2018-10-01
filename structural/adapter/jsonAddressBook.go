package adapter

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type jsonProcessor interface {
	print()
	addContact(jsonContact string) (int, error)
	get(id int) string
	getAll() string
	addAddress(address Address) error
}

type jsonContactList struct {
	contactList map[int]jsonContact
}

type jsonContact struct {
	Addresses []Address `json:"addresses"`
	jsonName  `json:"name"`
	Id        int `json:"id"`
}

type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type jsonName struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

func NewAddressBookJson() jsonContactList {
	var contacts jsonContactList
	contacts.contactList = make(map[int]jsonContact)

	return contacts
}

func (contactList jsonContactList) addContact(contactInString string) (int, error) {

	var contact jsonContact
	var id int
	err := json.Unmarshal([]byte(contactInString), &contact)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	id = rand.Intn(10000)
	contact.Id = id

	contactList.contactList[id] = contact

	return id, err
}

func (contactList jsonContactList) print() {

}
func (contactList jsonContactList) get(id int) string {

}
func (contactList jsonContactList) getAll() string {

}
func (contactList jsonContactList) addAddress(address Address) error {

}
