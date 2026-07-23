package main

import (
	"log"
	"os"
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
	_, err := os.ReadFile("out/go.mod")
	if err != nil {
		initProject()
	}

	cmd := exec.Command("go", "get", name);
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}