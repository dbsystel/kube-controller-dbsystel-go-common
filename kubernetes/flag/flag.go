package flag

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func AddFlags(app *kingpin.Application, config *bool) {
	app.Flag("run-outside-cluster", "set if you want to run outside from k8s").Default("false").BoolVar(config)
}
