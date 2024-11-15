package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "buildx" {
		buildx()
		return
	}

	buildEnv()
}

func buildEnv() {
	ctx := context.TODO()
	bin := checkDagger()
	cmd := exec.CommandContext(ctx, bin, []string{
		"call",
		"-m=github.com/rajatjindal/daggerverse/wasi@main",
		"build-env",
		"--source=.",
		"terminal",
	}...)

	run(cmd)
}

func buildx() {
	ctx := context.TODO()

	bin, err := exec.LookPath("dagger")
	if err != nil {
		fmt.Println("dagger is not found. Please install using https://docs.dagger.io/install")
		os.Exit(1)
	}

	cmd := exec.CommandContext(ctx, bin, []string{
		"call",
		"-m=github.com/rajatjindal/daggerverse/wasi@main",
		"build",
		"--source=.",
		"export",
		"--path=.",
	}...)

	run(cmd)
}

func run(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("failed to run cmd %v\n", err)
		os.Exit(1)
	}
}

func checkDagger() string {
	path, err := exec.LookPath("dagger")
	if err != nil {
		fmt.Println("dagger is not found. Please install using https://docs.dagger.io/install")
		os.Exit(1)
	}

	return path
}
