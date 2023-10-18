package main

var IPC_CHANNELS = make(map[string]chan any)

func create_global_channel(name string) *chan any {
	channel := make(chan any)
	IPC_CHANNELS[name] = channel
	return &channel
}
