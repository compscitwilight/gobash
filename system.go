package main

import (
	"os"
)

var DIRECTORY string = ""

func get_working_directory() (string, error) {
	var wd string
	if DIRECTORY == "" {
		wd, _ = os.Getwd()
		return wd, nil
	}
	return DIRECTORY, nil
}
