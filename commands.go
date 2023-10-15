package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

const SHELL_COMMAND_PREFIX = "#"

// Method for executing operating system commands using the os package
func sys_exec(cmd string) (*exec.Cmd, []byte, error) {
	//USER_CONTROL = false
	char_0 := string(cmd[0])
	command := strings.Split(cmd, " ")[0]
	args := strings.Split(strings.TrimPrefix(cmd, command), " ")

	is_file_reference := char_0 == "." || char_0 == "/"
	if is_file_reference {
		wd, _ := get_working_directory()
		exec.Command(filepath.Join(wd, cmd)).Run()
		return nil, nil, nil
	} else {
		cmd_path := strings.Join([]string{"/usr/bin", command}, "/")
		var command_exec *exec.Cmd
		log.Println(len(args))
		if len(args) == 1 {
			command_exec = exec.Command(cmd_path)
		} else {
			command_exec = exec.Command(cmd_path, args...)
		}

		output, err := command_exec.Output()
		return command_exec, output, err
	}
	/**
	command := exec.Command(cmd)
	output, err := command.Output()
	*/
	//USER_CONTROL = true
	//return command, output, err
}

// For handling internal shell commands that manipulate the
// shell only.
func shell_exec(cmd string, args ...string) {
	if string(cmd[0]) == SHELL_COMMAND_PREFIX {
		cmd = strings.TrimPrefix(cmd, SHELL_COMMAND_PREFIX)
	}
	log.Println(cmd)
	if cmd == "help" {
		log_help_manual()
	} else if cmd == "exit" {
		os.Exit(1)
	} else if cmd == "menu" {
		show_shell_menu()
	}
}

func log_help_manual() {
	log.Println(color.CyanString("Welcome to Gobash"))
}

func string_trim(str string) {

}
