package io

import (
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"testing"
)

type Person struct {
	XMLName Name    `xml:"person"`
	Name    Name    `xml:"name"`
	Emails  []Email `xml:"email"`
}
type Name struct {
	First string `xml:"first"`
	Last  string `xml:"last"`
}
type Email struct {
	Kind    string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

var person = Person{
	Name: Name{First: "Ууганбаяр", Last: "Сүхбаатар"},
	Emails: []Email{{Kind: "хувийн", Address: "ub@gmail.com"},
		{Kind: "ажлын", Address: "ub@hotmail.com"}}}

func TestGob(t *testing.T) {
	outFile, _ := os.Create("person.gob")
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err := encoder.Encode(person)
	if err != nil {
		t.Error(err)
	}
}

func TestJSON(t *testing.T) {
	outFile, _ := os.Create("person.json")
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)
	err := encoder.Encode(person)
	if err != nil {
		t.Error(err)
	}
}

func TestXML(t *testing.T) {
	str := `<?xml version="1.0" encoding="utf-8"?>
	<person>
	<name>
	<first>Ууганбаяр</first>
	<last>Сүхбаатар</last>
	</name>
	<email type="хувийн">ub@gmail.com</email>
	<email type="ажлын">ub@hotmail.com</email>
	</person>`

	var person Person
	err := xml.Unmarshal([]byte(str), &person)
	if err != nil {
		t.Error(err)
	}

	// обектыг ашиглах
	fmt.Println("Нэр: \"" + person.Name.First + "\"")
	fmt.Println("Э-мэйл 2: \"" + person.Emails[1].Address + "\"")
}
