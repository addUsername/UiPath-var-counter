package main

import (
	"errors"
	"os"
)

func getAllPaths(dir string, s bool) ([]string, error) {

	var toReturn []string

	if s {
		if len(dir) < 6 || dir[len(dir)-5:] != ".xaml" {
			return nil, errors.New("invalid path has not .xalm extension")
		}

		toReturn = append(toReturn, dir)
		return toReturn, nil
	}

	dir = processPath(dir)

	return lookForXalms(toReturn, dir)

}

func processPath(dir string) string {

	if dir == "." {

		root, err := os.Getwd()

		if err != nil {
			Perror(err)
			os.Exit(-2)
		}
		return root //+ "\\Main.xalm"
	}
	return dir

}

func lookForXalms(toReturn []string, dir string) ([]string, error) {

	files, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	var name string
	for _, file := range files {

		name = dir + "\\" + file.Name()

		if len(name) > 6 && name[len(name)-5:] == ".xaml" {

			Pdlog(name)
			toReturn = append(toReturn, name)

		} else if isFolder(name) {

			Pinfo(name)
			toReturn, err = lookForXalms(toReturn, name)
			if err != nil {
				return nil, err
			}

		}

	}
	return toReturn, nil

}

func isFolder(path string) bool {

	fileInfo, _ := os.Stat(path)
	return fileInfo.IsDir()
}
