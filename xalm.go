package main

type Xalm struct {
	Scopes   []Scope `xml:"variables"`
	Filename string
}

func (xa Xalm) Print() string {
	return "booh"
}
