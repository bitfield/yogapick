package main

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"

	"github.com/bitfield/yogapick"
)

func main() {
	path, err := xdg.ConfigFile("yogapick/poses.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	poses, err := yogapick.LoadPoses(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for p := range yogapick.Suggest(poses, 3) {
		fmt.Println(p)
	}
}
