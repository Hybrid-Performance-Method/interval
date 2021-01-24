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

	// run interval
	interval.CreateEnv(reqs)
	interval.RunNotebook(nb, params)
	fmt.Println("✔️ Finished ✔️")
}
