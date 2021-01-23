package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	CreateEnv()
	RunNotebook()
	fmt.Println("✔️ Finished ✔️")
}

func CreateEnv() {
	fmt.Println("⏱️ Starting Interval ⏱️")
	// create virtual environment
	// create virtual environment
	cmd1 := exec.Command("python", "-m", "venv", "venv")
	if err := cmd1.Run(); err != nil {
		log.Fatal("python venv Error: ")
	}
	// activate by running a bash command
	cmd2 := exec.Command("bash", "-c", "source venv/bin/activate")
	if err := cmd2.Run(); err != nil {
		log.Fatal("bash source error: ", err)
	}
	// install requirements
	cmd3 := exec.Command("pip", "install", "-r", "requirements.txt")
	if err := cmd3.Run(); err != nil {
		log.Fatal("pip install error ", err)
	}
}

func RunNotebook() {
	cmdKernel := exec.Command("python", "-m", "ipykernel", "install", "--user", "--name", "interval")
	if err := cmdKernel.Run(); err != nil {
		log.Fatal("kernel error", err)
	}
	fmt.Println("Taking Notes ⏳")
	cmdRun := exec.Command("papermill", "sample.ipynb", "success.ipynb")
	cmdRun.Stdin = os.Stdin
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	if err := cmdRun.Run(); err != nil {
		log.Fatal("Interval Error: ", err)
	}
}
