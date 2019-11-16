package main

import (
    "github.com/docopt/docopt-go"
)

var inifile string

func args_init() {
    usage := `Oxyfi GPS websocket API server 

Usage:
  oxify_gps_wsd [--cfg=<inifile>]

Options:
  --cfg=<inifile>    Settings file path [default: oxyfigpsws.ini]

`
    arguments, _ := docopt.ParseDoc(usage)
    inifile = arguments["--cfg"].(string)
}