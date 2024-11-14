package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	ctx := context.TODO()

	path, err := exec.LookPath("dagger")
	if err != nil {
		fmt.Println("dagger is not found. Please install using https://docs.dagger.io/install")
		os.Exit(1)
	}

	cmd := exec.CommandContext(ctx, path, []string{
		"call",
		"-m=github.com/rajatjindal/daggerverse/wasi@main",
		"build-env",
		"--source=.",
		"terminal",
	}...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		fmt.Printf("failed to run cmd %v\n", err)
		os.Exit(1)
	}

	fmt.Println("closed the terminal succcessfully")
}
