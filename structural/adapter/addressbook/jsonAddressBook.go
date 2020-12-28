package addressbook

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

	for key, value := range contactList.contactList {
		fmt.Printf("Id = %d\t name = %s\n", key, value.FirstName+" "+value.LastName)
	}

}
func (contactList jsonContactList) get(id int) string {
	contact := contactList.contactList[id]

	stringFormat, err := json.Marshal(contact)

	if err != nil {
		return ""
	}

	return string(stringFormat)
}

func (contactList jsonContactList) getAll() string {

	list := make([]jsonContact, 0)

	for _, value := range contactList.contactList {
		list = append(list, value)
	}

	result, err := json.Marshal(list)
	if err != nil {
		return ""
	}

	return string(result)
}
func (contactList jsonContactList) addAddress(id int, address Address) error {

	addresses := contactList.contactList[id].Addresses
	addresses = append(addresses, address)

	return nil
}
