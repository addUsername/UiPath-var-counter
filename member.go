package main

// Represents a var
type Member struct {
	Class string `xml:"Type,attr"`
	Name  string `xml:"Name,attr"`
	Count int
}
