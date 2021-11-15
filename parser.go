package main

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
)

// TODO: get args definition too

// Search for all vars definition and parse them
func ParseFile(path string) Xalm {

	XAMLContent, err := ioutil.ReadFile(path)

	if err != nil {
		Perror(err)
	}

	strContent := string(XAMLContent)
	var aux = ""
	if members {
		aux = getMembers(strContent)

	}

	strContent = clean(strContent)
	yeah := getVariablesAstext(strContent)

	final := "<root>" + "\n" + aux + "\n" + strings.Join(yeah, "\n") + "</root>"

	var coverFile Xalm

	err = xml.Unmarshal([]byte(final), &coverFile)
	if err != nil {
		Perror(err)
	}

	filename := strings.Split(path, "\\")
	coverFile.Filename = filename[len(filename)-1]
	return coverFile

}

// Find arguments, they call them Members
func getMembers(text string) string {

	flag := "<x:Members>"
	flagEnd := "</x:Members>"
	count := strings.Count(text, flag)
	if count < 1 {
		return ""
	}
	loc1 := strings.Index(text, flag)
	loc2 := strings.Index(text, flagEnd)

	return text[loc1 : loc2+len(flagEnd)]

}

// Find variabes definitions and buid a valid .xml with all the info
func getVariablesAstext(text string) []string {

	flag := ".Variables>"

	count := strings.Count(text, flag) / 2
	toReturn := make([]string, count)

	for i := 0; i < count; i++ {

		loc1 := strings.Index(text, flag)
		displayName := getDisplayname(text, loc1)
		text = strings.Replace(text, flag, "####", 1)

		loc2 := strings.Index(text, flag)

		x := strings.Split(text[loc1+6:loc2], "\n")
		x[len(x)-1] = "</variables>\n"

		toReturn[i] = "<variables DisplayName=\"" + displayName + "\">\n" + strings.Join(x, "\n")

		text = text[loc2+len(flag):]

	}
	return toReturn
}

// Get rid of not necesary text
func clean(text string) string {

	flag := "</TextExpression.ReferencesForImplementation>"
	loc := strings.Index(text, flag)
	return text[loc+len(flag):]
}

//Get scope's name
func getDisplayname(text string, loc1 int) string {

	if loc1 < 200 {

		text = text[0:]
	} else {
		text = text[loc1-200 : loc1]
	}
	s := strings.Split(text, "DisplayName=\"")
	if len(s) > 1 {

		text = strings.Split(s[1], "\" ")[0]
	} else {
		s = strings.Split(text, "IdRef=\"")
		if len(s) < 2 {
			text = "StateMachine"
		} else {
			text = strings.Split(s[1], "\">")[0]

		}

	}

	return text

}
