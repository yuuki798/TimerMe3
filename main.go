package main

import "github.com/yuuki798/TimerMe3/cmd/server"

func main() {
	//cmd.Execute()
	server.SetUp()
	server.Run()
}
