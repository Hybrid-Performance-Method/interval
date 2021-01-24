package main

import (
	"fmt"

	"github.com/Hybrid-Performance-Method/interval/interval"
)

func main() {
	// get inputs
	nb := interval.GetNotebook()
	reqs := interval.GetRequirements()
	params := interval.GetParams()
	paramsFile := interval.GetParamsFile()
	interval.GetSecrets()

	// parameters can either be string or parameters.yml file
	if paramsFile != "" {
		params = paramsFile
	}

	// run interval
	interval.CreateEnv(reqs)
	interval.RunNotebook(nb, params)
	fmt.Println("✔️ Finished ✔️")
}
