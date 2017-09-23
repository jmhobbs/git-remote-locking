package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: git unlock-remote <remote>")
	}

	out, err := exec.Command("git", "config", "--unset", fmt.Sprintf("remote.%s.pushurl", os.Args[1])).Output()
	if err != nil {
		stderr := string(err.(*exec.ExitError).Stderr)
		if len(stderr) != 0 {
			log.Fatal(stderr)
		}
		log.Fatal("error unlocking remote '", os.Args[1], "' - may not exist?")
	}

	fmt.Printf("unlocked remote '%s'\n", os.Args[1])
	if len(out) > 1 {
		fmt.Println(string(out))
	}
}
