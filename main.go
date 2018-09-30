package main

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"os"
)

type Address struct {
	City, State string
}
type Person struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}

type Contact struct {
	Name      xml.Name `xml:person`
	Id        int      `xml:id`
	FirstName string   `xml:name`
	LastName  string   `xml:surname`
}

func main() {

	// v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	// v.Comment = " Need more details. "
	// v.Address = Address{"Hanga Roa", "Easter Island"}
	data := `
	<Contact>
		<Id></Id>
		<LastName>BürgerSteig</LastName>
		<LastName>IstImmerVernünftig</LastName>
	</Contact>`

	var entity Contact
	var id int

	err := xml.Unmarshal([]byte(data), &entity)
	if err != nil {
		fmt.Println(err)
	}
	id = rand.Intn(10000)
	entity.Id = id

	output, err := xml.MarshalIndent(entity, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}
