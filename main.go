package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal(("No filename provided"))
	}

	if !strings.HasSuffix(args[0], ".el") {
		log.Fatal("File is not a Enlang (.el) file")
	}

	file, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	files, dependencies := transpile(string(file))

	os.Mkdir("out", 0755)
	removeTempFiles()
	initProject()

	for _, dependencie := range dependencies {
		installDependencie(dependencie)
		fmt.Printf("Installed %s", dependencie)
	}


	for _, file := range files {
		err = os.WriteFile(fmt.Sprintf("out/%s", file.name), []byte(file.content), 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created %s", file.name)
	}

	tidyCommand()

	cmd := exec.Command("go", "build", "-o", "binary", ".");
	cmd.Dir = "out/"
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("go build failed: %v\n%s", err, out)
	}
	removeTempFiles()
}

func removeTempFiles() {
	os.Remove("out/main.go")
	os.Remove("out/go.sum")
	os.Remove("out/go.mod")
}