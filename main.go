package main

import (
	"flag"
	"os"
	"strconv"
)

/*
TODO:

- some model management
- Show results..

*/
var ver = "0.2"

//var root string
var path string
var s bool
var v bool

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
	}

	Pwarn("Parsing files..")

	var allVars []Xalm
	for _, file := range files {

		vars := parseFile(file)
		vars = countVars(vars, file)
		allVars = append(allVars, vars)
	}
	// Generate Xalm
	Pwarn("Rendering table..")
	for i, v := range allVars {

		//Ptitle(strconv.Itoa(i+1) + " / " + strconv.Itoa(len(allVars)) + v.Filename)
		Ptitle(strconv.Itoa(i+1), strconv.Itoa(len(allVars)), v.Filename)
		for _, scope := range v.Scopes {

			for k, va := range scope.Vars {

				//x := va.Name + "\t\t" + va.Class + "\t\t" + strconv.Itoa(va.Count)
				Prow(va.Name, va.Class, strconv.Itoa(va.Count), k%2 == 0)
				//	fmt.Printf("%-6s | % 6s | %6s", va.Name, va.Class, strconv.Itoa(va.Count))
				//	fmt.Println()
			}
			PfooterScope(scope.Name)
		}
		Pfooter()
	}

	os.Exit(0)
	Pwarn("\nRendering table..")

	/*

		// Fill Xalm with count


		fmt.Println(vars.Scopes[0].Vars[1].Count)

		fmt.Println(vars.Print())
	*/
}
