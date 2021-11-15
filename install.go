package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

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
