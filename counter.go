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
	return xalm
}
