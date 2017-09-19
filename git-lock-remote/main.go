package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: git lock <remote>")
	}

	out, err := exec.Command("git", "config", fmt.Sprintf("remote.%s.pushurl", os.Args[1]), "no_push").Output()
	if err != nil {
		stderr := string(err.(*exec.ExitError).Stderr)
		if len(stderr) != 0 {
			log.Fatal(stderr)
		}
		log.Fatal("error locking remote '", os.Args[1], "' - may not exist?")
	}

	fmt.Printf("locked remote '%s'\n", os.Args[1])
	if len(out) > 1 {
		fmt.Println(string(out))
	}
}
