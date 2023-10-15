package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/fatih/color"
)

const _VERSION = float32(1.0)

var USER_CONTROL = true
var USER_INPUT = ""
var HOSTNAME string
var USERNAME string

func main() {
	log.SetFlags(0)
	log.Println(color.CyanString("GoBash version:"), _VERSION)
	show_shell_menu()

	//log.SetPrefix(color.GreenString("out -> "))

	hostname, err := os.Hostname()
	username, err := exec.Command("whoami").Output()
	if err != nil {
		log.Println(err)
	}

	HOSTNAME = hostname
	USERNAME = strings.TrimSuffix(string(username), "\n")

	in := bufio.NewScanner(os.Stdin)
	// Input loop
	for {
		wd, _ := get_working_directory()
		wd_subs := strings.Split(wd, "/")
		fmt.Printf("%v@%v [%v]", USERNAME, HOSTNAME, wd_subs[len(wd_subs)-1])
		fmt.Printf(color.MagentaString("# "))

		in.Scan()
		USER_INPUT = in.Text()
		//fmt.Scanf("%s", &USER_INPUT)
		//USER_INPUT = strings.Join([]string{USER_INPUT}, " ")
		//log.Println(len(USER_INPUT))
		//USER_LOG[len(USER_LOG)] = CommandLog{USER_INPUT, int(time.Now().Unix())}
		history_append(USER_INPUT, int(time.Now().Unix()))

		var is_shell_command bool = false
		if string(USER_INPUT[0]) == SHELL_COMMAND_PREFIX {
			is_shell_command = true
		}

		if is_shell_command {
			shell_exec(USER_INPUT)
		} else {
			_, out, err := sys_exec(USER_INPUT)
			if err != nil {
				log.Println(err)
			} else {
				log.Printf(string(out))
			}
		}
	}
}

func show_shell_menu() {
	log.Println("")
	man_box := box.New(box.Config{Px: 5, Py: 0, Type: "Single", Color: "Cyan"})
	man_box.Print("Man", "")
}
