package main

type Variable struct {
	Class   string `xml:"TypeArguments,attr"`
	Name    string `xml:"Name,attr"`
	Default string `xml:"Variable.Default>Literal"`
	Count   int
}
