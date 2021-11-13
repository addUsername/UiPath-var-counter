package main

type Scope struct {
	Vars []Variable `xml:"Variable"`
	Name string     `xml:"DisplayName,attr"`
}
