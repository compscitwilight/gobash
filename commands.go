package main

import (
	"fmt"
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
	command, args := parse_command(cmd)
	/**
	command := strings.Split(cmd, " ")[0]
	args := strings.Split(strings.TrimPrefix(cmd, command), " ")
	*/
	is_file_reference := char_0 == "." || char_0 == "/"
	if is_file_reference {
		wd, _ := get_working_directory()
		exec.Command(filepath.Join(wd, cmd)).Run()
		return nil, nil, nil
	} else {
		cmd_path := strings.Join([]string{"/usr/bin", command}, "/")
		var command_exec *exec.Cmd
		//log.Println(len(args))
		if len(args) == 1 {
			command_exec = exec.Command(cmd_path)
		} else {
			command_exec = exec.Command(cmd_path, args...)
		}
		command_exec.Dir = WORKING_DIRECTORY
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
func shell_exec(cmd string) {
	if string(cmd[0]) == SHELL_COMMAND_PREFIX {
		cmd = strings.TrimPrefix(cmd, SHELL_COMMAND_PREFIX)
	}

	command, args := parse_command(cmd)

	//log.Println(cmd)
	if command == "help" {
		show_shell_menu()
	} else if command == "exit" {
		os.Exit(1)
	} else if command == "cd" {
		//os.Open(args[0])
		path := filepath.Join(WORKING_DIRECTORY, args[0])
		if err := os.Chdir(path); err != nil {
			log.Println(err)
			return
		}
		//new_path := filepath.Join([]string{WORKING_DIRECTORY, "../", args[0]}...)
		//log.Println(new_path, args[0])
		//WORKING_DIRECTORY = new_path
	} else if command == "mk_config" {
		if _, stat_err := os.Stat(CONFIG_PATH); stat_err != nil {
			// mode 0755 represents a directory
			if _, config_creation_error := os.Create(CONFIG_PATH); config_creation_error != nil {
				log.Printf(config_creation_error.Error())
			} else {
				log.Printf(color.GreenString("Created config file."))
			}
		}
	} else if strings.Split(command, ".")[0] == "config" {
		log.Println("yo")
	} else if strings.Split(command, ".")[0] == "script" {
		script_name := strings.Split(command, ".")[1]
		lua_log(fmt.Sprintf("Running lua script %v", script_name))
		if err := run_lua_script(script_name); err != nil {
			log.Println(err)
		}
	}
}

func parse_command(cmd string) (string, []string) {
	command := strings.Split(cmd, " ")[0]
	args := strings.Split(strings.TrimPrefix(cmd, command), " ")
	return command, args
}

func log_help_manual() {
	log.Println(color.CyanString("Welcome to Gobash"))
}

func string_trim(str string) {

}
