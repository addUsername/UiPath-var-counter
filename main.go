package main

import (
	"flag"
	"os"
	"strconv"
)

/*
TODO:
- parse arguments too
- put colors on that table
- export .csv
*/

var ver = "0.2.5"

// path to uipath project
var path string

// single file
var s bool

// verbose
var v bool

// parse args
func init() {

	flag.StringVar(&path, "path", ".", "Path to step folder")
	flag.BoolVar(&s, "s", false, "Single file, no recursive search")
	flag.BoolVar(&v, "verbose", false, "Print log lines")

	var help bool
	flag.BoolVar(&help, "help", false, "Show help")

	var version bool
	flag.BoolVar(&version, "version", false, "Show version")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if version {
		v = true
		Pinfo(ver)
		os.Exit(0)
	}

}

func main() {

	Pwarn("Reading files..")
	files, err := getAllPaths(path, s)

	if err != nil {
		Perror(err)
		os.Exit(-1)
	}

	Pwarn("Parsing files..")
	var allVars []Xalm

	for _, file := range files {

		vars := parseFile(file)
		vars = countVars(vars, file)
		allVars = append(allVars, vars)
	}

	Pwarn("Rendering table..")
	for i, va := range allVars {

		if len(va.Scopes) == 0 && !v {
			continue
		}
		//[16 - 17]    ProcessTestCase.xaml
		//┌───────────────────────────┬─────────────────────────────┬──────┐
		Ptitle(strconv.Itoa(i+1), strconv.Itoa(len(allVars)), va.Filename)
		for _, scope := range va.Scopes {

			for k, va := range scope.Vars {
				//  TransactionItem           │   QueueItem                 │   2
				Prow(va.Name, va.Class, strconv.Itoa(va.Count), k%2 == 0)
			}
			//──────────────────────────────────────────────────────────────────ProcessTestCase
			PfooterScope(scope.Name)
		}
		//└───────────────────────────┴─────────────────────────────┴──────┘
		Pfooter()
	}
	os.Exit(0)
}
