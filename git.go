package main

import (
	"fmt"
	"os/exec"
)

func clone(repo string) {
	cmd := exec.Command("git", "clone", repo)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("clone err: %v", err)
	}
}
