package main

import (
	"errors"
	"fmt"
)

var USER_LOG_LIMIT = 100
var USER_LOG = [100]CommandLog{}

type CommandLog struct {
	context   string
	timestamp int
}

func get_history_len() int {
	var count = 0
	for _, item := range USER_LOG {
		exists := &item == &CommandLog{}
		if exists {
			count++
		}
	}
	return count
}

func history_append(context string, timestamp int) error {
	length := get_history_len()
	var new_index = 0
	if length > 0 {
		new_index = length + 1
	}

	if new_index == USER_LOG_LIMIT {
		return errors.New("History limit reached.")
	}

	USER_LOG[new_index] = CommandLog{context, timestamp}
	return nil
}

func history_remove(c *CommandLog) error {
	var removed bool = false
	for index, log := range USER_LOG {
		check := c == &log
		if check {
			USER_LOG[index] = CommandLog{}
			removed = true
		}
	}

	if !removed {
		return errors.New("Couldn't remove history log because it does not exist.")
	}

	return nil
}

func log_history() {
	for _, content := range USER_LOG {
		exists := content != CommandLog{}
		if exists {
			fmt.Println(content.context, "-", content.timestamp)
		}
	}
}
