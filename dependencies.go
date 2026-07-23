package main

import (
	"log"
	"os/exec"
)

func initProject() {
	cmd := exec.Command("go", "mod", "init", "project");
	cmd.Dir = "out/"
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("go mod init failed: %v\n%s", err, out)
	}
}

func installDependencie(name string) {
	cmd := exec.Command("go", "get", name);
	cmd.Dir = "out/"
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("go get %s failed: %v\n%s", name, err, out)
	}
}