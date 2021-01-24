package interval

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

func CreateEnv(requirements string) {
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
	if requirements != "" {
		req := ParseRequirements(requirements)
		for _, r := range req {
			cmd3 := exec.Command("pip", "install", r)
			if err := cmd3.Run(); err != nil {
				log.Fatal("pip install error ", err)
			}
		}
	}
}

func RunNotebook(notebook, params string) {
	fmt.Println("Taking notes with Papermill ⏳")
	// Start kernel
	cmdKernel := exec.Command("python", "-m", "ipykernel", "install", "--user", "--name", "interval")
	if err := cmdKernel.Run(); err != nil {
		log.Fatal("kernel error", err)
	}
	var cmdRun *exec.Cmd
	if params == "" {
		cmdRun = exec.Command("papermill", notebook, "success.ipynb")
	} else {
		cmdRun = exec.Command("papermill", notebook, "success.ipynb", "-y", params)
	}

	cmdRun.Stdin = os.Stdin
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	if err := cmdRun.Run(); err != nil {
		log.Fatal("Interval Error: ", err)
	}
}

// Helpers
func ParseRequirements(val string) []string {
	r := strings.Split(val, " ")
	for _, v := range r {
		if strings.Contains(v, "==") {
			githubactions.Warningf("Requirement %v does not have pinned version", val)
		}
	}
	return r
}

func GetRequirements() string {
	r := githubactions.GetInput("requirements")
	fmt.Println("Reading Requirements:", r)
	return r
}

func GetNotebook() string {
	nb := githubactions.GetInput("notebook")
	fmt.Println("Reading Notebook:", nb)
	return nb
}

// parameters are formatted in yaml string
func GetParams() string {
	p := githubactions.GetInput("parameters")
	fmt.Println("Reading in Paramters:", p)
	return p
}
