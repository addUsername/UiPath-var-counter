package main

// This struct represents the .xaml file as a whole
type Xalm struct {
	Scopes    []Scope `xml:"variables"`
	Filename  string
	Arguments Members `xml:"Members"`
}

// Test function..
func (xa Xalm) Print() string {
	return "booh"
}
