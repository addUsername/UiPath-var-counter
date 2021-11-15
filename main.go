package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
TODO:
- parse arguments too
- put colors on that table
- export .csv
- install (add to path)
- 0.7.0-beta
- test
- 0.8.0-pre
*/

var ver = "0.3.1-alpha"

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

func parseArgs() {

	flag.StringVar(&path, "path", ".", "Path to step folder")
	flag.BoolVar(&s, "s", false, "Single file, no recursive search")
	flag.BoolVar(&v, "verbose", false, "Print log lines")
	flag.BoolVar(&members, "args", false, "parse and show arguments too")
	flag.BoolVar(&d, "default", false, "show default values too")

	//os.Exit(0)
	var help bool
	flag.BoolVar(&help, "help", false, "Show help")

	var version bool
	flag.BoolVar(&version, "version", false, "Show version")

	var install string
	flag.StringVar(&install, "install", "", "(No fully implemented) Copy binary to specified folder and add to %PATH%")

	flag.Parse()

	if v {
		d = true
	}

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
			Pinfo("Please tun this command as administrator :P")
			os.Exit(-2)
		}
		os.Exit(0)
	}

}

func installMe(install string) error {

	Pinfo("Do you want to install myvars to " + install + "?\n[y/n]: ")

	reader := bufio.NewReader(os.Stdin)
	var ans string
	ans, _ = reader.ReadString('\n')
	ans = strings.TrimSpace(ans)

	var err error
	if ans == "yes" || ans == "y" {
		Pwarn("Installing my vars..")
		err = doTheInstall(install)
	} else {
		Pwarn(":(")
	}
	return err

}

func doTheInstall(install string) error {

	//Get exe path
	pathToExe, err := os.Executable()
	if err != nil {
		Perror(err)
		return err
	}
	name := strings.Split(pathToExe, "\\")

	/*
		path := savePath(install)

		if len(path)+len(install)+20 > 1023 {
			Pwarn("[WARNING] lenght path is bigger than 1024, add this route to the PATH manually: ")
			Pinfo(install + "\\" + name[len(name)-1])
		}

		out, err := addToPath(path, install)
		if err != nil {
			Perror(err)
			return err
		}
		Pinfo(out)
		os.Exit(0)
	*/
	Pinfo("Reading file: " + path)

	input, err := ioutil.ReadFile(path)
	if err != nil {
		Perror(err)
		return err
	}

	Pinfo("Copy file: " + install + "\\" + name[len(name)-1])
	err = ioutil.WriteFile(install+"\\"+name[len(name)-1], input, 0644)
	if err != nil {
		Perror(err)
		return err
	}

	Pinfo("Pls maually Add \"" + install + "\\" + name[len(name)-1] + "\" to %PATH% ")
	Pinfo("Now you can use the command \"myvars\" anywhere in your Powershell")
	panic("add to path it is not implemented yet, srry")
	return nil
}

func addToPath(path string, install string) (string, error) {

	path = path + ";" + install
	cmd := exec.Command("powershell", "setx /M PATH "+path)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		fmt.Println("error??")
		Perror(err)
		return "", nil
	}
	output := out.String()
	return output, nil
}

func savePath(install string) string {
	cmd := exec.Command("sh", "$Env:Path")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		Perror(err)
	}
	path := out.String()
	return path
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

		/*
			if members && len(va.Arguments.Arguments) > 1 {

				Ptitle(strconv.Itoa(i+1), strconv.Itoa(len(allVars)), va.Filename)
				for k, me := range va.Arguments.Arguments {
					//Prow(me.Name, me.Class, strconv.Itoa(me.Count), k%2 == 0, d)
					//Prow(va, k%2 == 0, d)
				}
				Pfooter()
			}
		*/
	}
	os.Exit(0)
}
