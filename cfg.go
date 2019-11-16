package main

import (
    "log"
    "gopkg.in/ini.v1"
)

var db_path, admin_username, admin_password, http_server_port, udp_port string

func read_cfg() {
    cfg, err := ini.Load(inifile)
    if err != nil {
        log.Fatal("Unable to read ini file.")
    }
    db_path = cfg.Section("").Key("db_path").String()
    admin_username = cfg.Section("").Key("admin_username").String()
    admin_password = cfg.Section("").Key("admin_password").String()
    http_server_port = cfg.Section("").Key("http_server_port").String()
    udp_port = cfg.Section("").Key("udp_port").String()
}

