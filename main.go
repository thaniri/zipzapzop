package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println("hello")
	info, ok := debug.ReadBuildInfo()
	//info.Main.Version = "v0.0.1"
	if ok {
		fmt.Println(info)
		fmt.Println("Module:", info.Main.Path)
		fmt.Println("Version:", info.Main.Version)
	}
}
