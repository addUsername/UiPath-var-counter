package main

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
)

func parseFile(path string) Xalm {

	XAMLContent, err := ioutil.ReadFile(path)

	if err != nil {
		Perror(err)
	}

	strContent := string(XAMLContent)
	strContent = clean(strContent)
	yeah := getVariablesAstext(strContent)

	final := "<root>\n" + strings.Join(yeah, "\n") + "</root>\n"

	var coverFile Xalm

	err = xml.Unmarshal([]byte(final), &coverFile)
	if err != nil {
		Perror(err)
	}

	filename := strings.Split(path, "\\")
	coverFile.Filename = filename[len(filename)-1]
	return coverFile

}

func clean(text string) string {

	flag := "</TextExpression.ReferencesForImplementation>"
	loc := strings.Index(text, flag)
	return text[loc+len(flag):]
}

func getVariablesAstext(text string) []string {

	flag := ".Variables>"

	count := strings.Count(text, flag) / 2
	toReturn := make([]string, count)

	for i := 0; i < count; i++ {

		loc1 := strings.Index(text, flag)
		text = strings.Replace(text, flag, "####", 1)

		loc2 := strings.Index(text, flag)

		x := strings.Split(text[loc1+6:loc2], "\n")
		x[len(x)-1] = "</variables>\n"
		toReturn[i] = "<variables>\n" + strings.Join(x, "\n")
		text = text[loc2+len(flag)+1:]

	}
	return toReturn
}
