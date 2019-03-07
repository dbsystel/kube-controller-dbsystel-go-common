package flag

import (
	log "github.com/dbsystel/kube-controller-dbsystel-go-common/log"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func AddFlags(app *kingpin.Application, config *log.Config) {
	app.Flag("log-level", "desired log level, one of: [debug, info, warn, error]").Default("info").StringVar(&config.LogLevel)
	app.Flag("log-format", "desired log format, one of: [json, logfmt]").Default("json").StringVar(&config.LogFormat)
}
