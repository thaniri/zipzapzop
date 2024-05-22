package main

import (
	"fmt"
	"runtime/debug"
)

var version = "unknown"

func main() {
	// Print the version set during build time
	fmt.Println("Version:", version)

	// Get build info
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Failed to read build info")
		return
	}

	fmt.Println("Version2:", buildInfo.Main.Version)

	// Print build settings
	fmt.Println("Build Info:")
	for _, setting := range buildInfo.Settings {
		fmt.Printf("%s: %s\n", setting.Key, setting.Value)
	}
}
