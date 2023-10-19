package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

const SHELL_COMMAND_PREFIX = "#"

// Method for executing operating system commands using the os package
func sys_exec(cmd string, channel ...chan any) (*exec.Cmd, []byte, error) {
	char_0 := string(cmd[0])
	command, args := parse_command(cmd)
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
		//command_exec.Wait()
		output, err := command_exec.Output()
		if cmd == "ls" {
			new_output := ""
			output_string := string(output)
			for _, ln := range strings.Split(output_string, "\n") {
				path := filepath.Join(WORKING_DIRECTORY, ln)
				if _, isNotDir := os.ReadDir(path); isNotDir == nil {
					new_output = fmt.Sprintf("%v\n[dir] %v", new_output, color.HiMagentaString(ln))
				} else if _, isNotFile := os.ReadFile(path); isNotFile == nil {
					new_output = fmt.Sprintf("%v\n[file] %v", new_output, color.MagentaString(ln))
				}
			}
			output = []byte(new_output)
		}
		return command_exec, output, err
	}
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
		//path := filepath.Join(WORKING_DIRECTORY, args[1])
		//log.Println(args[1])
		if err := os.Chdir(args[1]); err != nil {
			log.Println(err)
			return
		}
		change_working_directory(args[1])
		log.Println(fmt.Sprintf("#cd: %v", args[1]))
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
		valid_size := len(strings.Split(command, ".")) != 0
		if valid_size {
			script_name := strings.Split(command, ".")[1]
			lua_log(fmt.Sprintf("Running lua script %v", script_name))
			if err := run_lua_script(script_name); err != nil {
				log.Println(err)
			}
		}
	} else if command == "tree" {
		tree := create_tree()
		tree.content = "Root Node"
		directory := create_tree()
		directory.content = "Directory"
		sub_directory := create_tree()
		sub_directory.content = "Subdirectory"

		tree = append_child_node(tree, &sub_directory)
		tree = append_child_node(tree, &directory)
		log_tree(tree)
	} else if command == "cr" {
		if len(args) < 3 {
			log.Println("Coroutines must have more than 3 arguments.")
			return
		}

		name := args[1]
		task1 := args[2]
		task2 := args[3]

		channel := *create_global_channel(name)
		go sys_exec(task1, channel)
		go sys_exec(task2, channel)
	} else if command == "img" {
		img, err := decodeImage("/home/shiny/Documents/code/gobash/twilight.png")
		if err != nil {
			log.Println(err)
		}
		log.Println(img.ColorModel())
	}
}

func parse_command(cmd string) (string, []string) {
	command := strings.Split(cmd, " ")[0]
	args := strings.Split(strings.TrimPrefix(cmd, command), " ")
	return command, args
}

func register_command(name string) {
	_, file, _, ok := runtime.Caller(0)
	if ok {
		log.Println(file)
	}
}

func log_help_manual() {
	log.Println(color.CyanString("Welcome to Gobash"))
}

func string_trim(str string) {

}
