package main

import (
	"errors"
	"os"
)

// Check for errors and returns all valid file paths (.xaml)
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

// Adds dot "." functionality
func processPath(dir string) string {

	if dir == "." {

		root, err := os.Getwd()

		if err != nil {
			Perror(err)
			os.Exit(-2)
		}
		return root
	}
	return dir

}

// Recursive function, do the actual search
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

// Cheks if arguments is a folder
func isFolder(path string) bool {

	fileInfo, _ := os.Stat(path)
	return fileInfo.IsDir()
}
