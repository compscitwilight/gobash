package main

import (
	"os"
)

var WORKING_DIRECTORY string

func get_working_directory() (string, error) {
	var wd string
	if WORKING_DIRECTORY == "" {
		wd, _ = os.Getwd()
		WORKING_DIRECTORY = wd
		return wd, nil
	}
	return WORKING_DIRECTORY, nil
}
