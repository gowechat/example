package main

import "github.com/gowechat/example"

func main() {
	err := example.Run()
	if err != nil {
		panic(err)
	}
}
