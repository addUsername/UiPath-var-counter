package main

import (
	"io/ioutil"
	"strings"
)

func countVars(xalm Xalm, path string) Xalm {

	XAMLContent, err := ioutil.ReadFile(path)

	if err != nil {
		Perror(err)
	}

	text := string(XAMLContent)
	for i, scope := range xalm.Scopes {

		for j, v := range scope.Vars {

			xalm.Scopes[i].Vars[j].Count = strings.Count(text, "\""+v.Name+"\"") +
				strings.Count(text, ">"+v.Name+"<") +
				strings.Count(text, "["+v.Name+"]") - 1
		}
	}

	if members {
		text := string(XAMLContent)
		for i, a := range xalm.Arguments.Arguments {

			xalm.Arguments.Arguments[i].Count = strings.Count(text, "\""+a.Name+"\"") +
				strings.Count(text, ">"+a.Name+"<") +
				strings.Count(text, "["+a.Name+"]") - 1
		}

	}
	return xalm
}
