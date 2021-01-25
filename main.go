package main

import (
	"fmt"

	"github.com/Hybrid-Performance-Method/interval/interval"
)

func main() {
	// get inputs
	nb := interval.GetNotebook()
	out := interval.GetOutputNotebook()
	reqs := interval.GetRequirements()
	params := interval.GetParams()
	paramsFile := interval.GetParamsFile()
	interval.ReadSecrets()

	// parameters can either be string or parameters.yml file
	if paramsFile != "" {
		params = paramsFile
	}

	// run interval
	interval.CreateEnv(reqs)
	interval.RunNotebook(nb, out, params)
	fmt.Println("✔️ Finished ✔️")
}
