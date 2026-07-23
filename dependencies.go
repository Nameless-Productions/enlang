package main

import (
	"log"
	"os/exec"
)

func initProject() {
	cmd := exec.Command("go", "mod", "init", "project");
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}

func installDependencie(name string) {
	cmd := exec.Command("go", "get", name);
	cmd.Dir = "out/"
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}