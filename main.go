package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/fatih/color"
)

const (
	_VERSION = 1.0
)

var USER_CONTROL = true
var USER_INPUT = ""
var KEY_INPUT []byte
var HOSTNAME string
var USERNAME string
var SOURCE_DIR string

var CONFIG_PATH string
var CLI_PREFIX = "#"

func main() {
	if runtime.GOOS == "windows" {
		log.Println("GoBash is not supported on Windows.")
		log.Println("Please see https://github.com/devrusty/gobash for more details")
		os.Exit(0)
	}

	log.SetFlags(0)
	log.SetPrefix(color.HiMagentaString("/)"))
	log.Println(strings.Join([]string{
		fmt.Sprintf("version: %g.0", _VERSION),
	}, " "))
	show_shell_menu()

	//log.SetPrefix(color.GreenString("out -> "))

	if dir, wd_err := os.Getwd(); wd_err != nil {
		SOURCE_DIR = dir
	}

	hostname, err := os.Hostname()
	username, err := exec.Command("whoami").Output()
	if err != nil {
		log.Println(err)
	}

	HOSTNAME = hostname
	USERNAME = strings.TrimSuffix(string(username), "\n")
	CONFIG_PATH = fmt.Sprintf("/home/%v/.config/gobash", USERNAME)

	config_data, parse_err := parse_config_file()
	if parse_err != nil {
		log.Println(parse_err)
	}

	if val, ok := config_data["CLI_PREFIX"]; ok {
		CLI_PREFIX = val
	}

	if val, ok := config_data["USERNAME"]; ok {
		USERNAME = val
	}

	in := bufio.NewScanner(os.Stdin)
	// Input loop
	for {
		wd, _ := get_working_directory()
		wd_subs := strings.Split(wd, "/")
		fmt.Printf("%v@%v [%v/%v]", USERNAME, HOSTNAME, color.BlueString(fmt.Sprint(len(wd_subs)-1)), wd_subs[len(wd_subs)-1])
		fmt.Printf(color.MagentaString(fmt.Sprintf("%v ", CLI_PREFIX)))

		in.Scan()
		USER_INPUT = in.Text()
		KEY_INPUT = in.Bytes()
		//fmt.Scanf("%s", &USER_INPUT)
		//USER_INPUT = strings.Join([]string{USER_INPUT}, " ")
		//log.Println(len(USER_INPUT))
		//USER_LOG[len(USER_LOG)] = CommandLog{USER_INPUT, int(time.Now().Unix())}
		history_append(USER_INPUT, int(time.Now().Unix()))

		var is_shell_command bool = false
		var is_valid_input bool = len(USER_INPUT) != 0
		if is_valid_input {
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
}

func show_shell_menu() {
	man_box := box.New(box.Config{Px: 15, Type: "Single", Color: "Cyan"})
	options := []string{"Docs", "About", "Configuration", "Scripts"}

	get_options_string := func() string {
		res := ""
		for index, opt := range options {
			res = fmt.Sprintf("%v\n%d - %v", res, index+1, opt)
		}
		return res
	}
	man_box.TopRight = "1"
	man_box.TopLeft = "0"
	man_box.BottomLeft = "1"
	man_box.BottomRight = "0"
	man_box.Print("Startup", get_options_string())
}

func parse_config_file() (map[string]string, error) {
	data := make(map[string]string)
	content, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		return nil, err
	}

	content_string := string(content)
	lines := strings.Split(content_string, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			sep := strings.Split(line, "=")
			key := sep[0]
			value := sep[1]
			data[key] = value
		}
	}
	return data, nil
}
