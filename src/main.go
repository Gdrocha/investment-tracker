package main

import (
	"investment-tracker/src/core/interfaces"
	"investment-tracker/src/core/registry"
	_ "investment-tracker/src/selic"
)

func main() {
	println("---------------------")

	reporters, err := registry.GetRegistry().GetAll((*interfaces.Reporter)(nil))

	if err != nil {
		println("Error while trying to get all reporters in the registry. Terminating application.")
	}

	for _, f := range reporters {
		if reporterImpl, ok := f.(interfaces.Reporter); ok {
			reporterImpl.Report()
		}
	}
}
