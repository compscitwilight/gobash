package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/Shopify/go-lua"
	"github.com/fatih/color"
)

func run_lua_script(alias string) error {
	if !strings.HasSuffix(alias, ".lua") {
		alias = fmt.Sprintf("%v.lua", alias)
	}
	l := lua.NewState()
	wd, _ := get_working_directory()
	path := filepath.Join(wd, alias)
	lua.OpenLibraries(l)
	if err := lua.DoFile(l, path); err != nil {
		return err
	}
	return nil
}

func lua_log(msg string) {
	log.Println(fmt.Sprintf("%v%v", color.CyanString("Lua Interpreter: "), msg))
}
