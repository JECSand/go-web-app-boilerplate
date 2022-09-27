package main

import "github.com/JECSand/go-web-app-boilerplate/cmd"

//go:generate rice embed-go

func main() {
	a := cmd.App{}
	a.Initialize("production")
	a.Run()
}
