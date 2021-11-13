package main

// Groups variables
type Scope struct {
	Vars []Variable `xml:"Variable"`
	Name string     `xml:"DisplayName,attr"`
}
