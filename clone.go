package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("rsync")
	if lookErr != nil {
		fmt.Println("\033[31mRsync Not Found. \033[0m")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("\033[31mInvaild Argument\nType --help For Usage.\033[0m")
		os.Exit(1)
	}

	if os.Args[1] == "--help " {
		fmt.Println("Usage:- clone [Source] [Distination]\n")
		fmt.Println(os.Args[0], "Version 1.0")
		os.Exit(1)

	}

	fmt.Println("\033[32m            Progress\033[0m\033[35m    Speed \033[0m\033[36m       Time \033[0m")
	args := []string{"rsync", "-ah", "--info=progress2", os.Args[1], os.Args[2]}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
