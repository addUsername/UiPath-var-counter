package main

import (
	"flag"
	"os"
	"strconv"
)

var ver = "0.3.5"

// path to uipath project
var path string

// single file
var s bool

// verbose
var v bool

// show default variables, it triggers auto when -verbose
var d bool

// parse and show arguments at the end of each table
var members bool

// path to output.json
var toJson string

func parseArgs() {

	flag.StringVar(&path, "path", ".", "Path to step folder")
	flag.BoolVar(&s, "s", false, "Single file, no recursive search")
	flag.BoolVar(&v, "verbose", false, "Print log lines")
	flag.BoolVar(&members, "args", false, "parse and show arguments too")
	flag.BoolVar(&d, "default", false, "show default values too")

	var help bool
	flag.BoolVar(&help, "help", false, "Show help")

	var version bool
	flag.BoolVar(&version, "version", false, "Show version")

	var install string
	flag.StringVar(&install, "install", "", "(No fully implemented) Copy binary to specified folder and add to %PATH%")

	flag.StringVar(&toJson, "json", "", "Path to export output.json")

	var noColors bool
	flag.BoolVar(&noColors, "coloroff", false, "Turn off colors")

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

	if len(install) > 0 {
		v = true
		err := installMe(install)
		if err != nil {
			Perror(err)
			Pinfo("Please run this command as administrator :P")
			os.Exit(-2)
		}
		os.Exit(0)
	}

	if v {
		d = true
	}

	if noColors {
		primary = "ffffff"
		second = "ffffff"
		alternative = "ffffff"
		ints = "ffffff"
	}

}

func main() {

	parseArgs()
	Pwarn("Reading files..")
	files, err := getAllPaths(path, s)

	if err != nil {
		Perror(err)
		os.Exit(-1)
	}

	Pwarn("Parsing files..")
	var allVars []Xalm

	for _, file := range files {

		vars := ParseFile(file)
		vars = countVars(vars, file)
		allVars = append(allVars, vars)
	}

	// if -json flag
	if len(toJson) > 0 {

		Pwarn("Generating json")
		text := xamlToJson(allVars)
		err = stringToFile(text, toJson)
		if err != nil {
			Perror(err)
			os.Exit(-1)
		} else {
			Pwarn("File saved -> " + processPath(toJson))
		}
		os.Exit(0)
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
				Prow(va, k%2 == 0)
			}
			//──────────────────────────────────────────────────────────────────ProcessTestCase
			PfooterScope(scope.Name)
		}
		//└───────────────────────────┴─────────────────────────────┴──────┘
		Pfooter()
		// Print args

		if members && len(va.Arguments.Arguments) > 0 {

			//Ptitle(strconv.Itoa(i+1), strconv.Itoa(len(allVars)), va.Filename)
			for k, me := range va.Arguments.Arguments {
				ProwArguments(me, k%2 == 0)

			}
			PfooterArguments()
		}
	}
	os.Exit(0)
}
