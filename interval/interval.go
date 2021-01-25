package interval

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/sethvargo/go-githubactions"
)

func CreateEnv(requirements string) {
	fmt.Println("⏱️ Starting Interval ⏱️")

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

func RunNotebook(notebook, output, params string) {
	fmt.Println("Taking notes with Papermill ⏳")
	// Start kernel
	cmdKernel := exec.Command("python", "-m", "ipykernel", "install", "--user", "--name", "interval")
	if err := cmdKernel.Run(); err != nil {
		log.Fatal("kernel error", err)
	}

	if readDate() == "true" && output != "output.ipynb" {
		output = AddDateString(output, time.Now())
	}

	var cmdRun *exec.Cmd
	switch params {
	case "":
		cmdRun = exec.Command("papermill", notebook, output)
	case "parameters.yml":
		cmdRun = exec.Command("papermill", notebook, output, "-f", params)
	default:
		cmdRun = exec.Command("papermill", notebook, output, "-y", params)
	}

	cmdRun.Stdin = os.Stdin
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	if err := cmdRun.Run(); err != nil {
		githubactions.Fatalf("Interval Error: %v", err)
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
	fmt.Println("Reading requirements:", r)
	return r
}

func GetNotebook() string {
	nb := githubactions.GetInput("notebook")
	if !strings.HasSuffix(nb, "ipynb") {
		githubactions.Fatalf("Notebook Error: %v is not a valid", nb)
	}
	fmt.Println("Reading Notebook:", nb)
	return nb
}

// parameters are formatted in yaml string
func GetParams() string {
	p := githubactions.GetInput("parameters")
	fmt.Println("Reading paramter string:", p)
	return p
}

func GetParamsFile() string {
	yml := githubactions.GetInput("parameterFile")
	fmt.Println("Reading Parameters file:", yml)
	if yml != "" && !strings.HasPrefix(yml, "parameters") {
		githubactions.Fatalf("Parameters Error: %v is not a valid", yml)
	}
	return yml
}

// ReadSecrets finds a secrets string and assigns it to the environmnent variable SECRETS
func ReadSecrets() {
	fmt.Println("Looking for secrets...")
	secrets := githubactions.GetInput("secrets")
	if secrets != "" {
		fmt.Println("⚡ Found Secrets ⚡")
	}
}

func GetOutputNotebook() string {
	o := githubactions.GetInput("outputNotebook")
	fmt.Println("Reading output pat:", o)
	if !strings.HasSuffix(o, "ipynb") {
		githubactions.Fatalf("Notebook Error: %v is not a valid", o)
	}
	return o
}

func readDate() string {
	return githubactions.GetInput("hasDate")
}

func AddDateString(file string, d time.Time) string {
	f := strings.Split(file, ".")
	output := fmt.Sprint(f[0] + "-" + d.Format("2006-01-02") + ".ipynb")
	return output
}
