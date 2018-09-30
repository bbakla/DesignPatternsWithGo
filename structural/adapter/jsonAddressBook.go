package adapter

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
	addresses []Address `json:"addresses"`
	jsonName  `json:"name"`
	Id        string `json:"id"`
}

type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type jsonName struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

func NewAddreeBookJson() jsonContactList {
	var contacts jsonContactList
	contacts.contactList = make(map[int]jsonContact)

	return contacts
}

func (contacts jsonContactList) addContact(jsonContact string) (int, error) {

	return 0, nil

}
