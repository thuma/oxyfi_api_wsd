package main

import (
    "log"
    "net/http"
)

func main() {
    args_init()
    read_cfg()
    init_db()
    http.HandleFunc("/v1/apikeys/apis/oxygps/profiles", rest_auth(profilesHandler))
    http.HandleFunc("/v1/apikeys/apis/oxygps/keys", rest_auth(adminKeysHandler))
    http.HandleFunc("/v1/apikeys/keys/", rest_auth(keysHandler))
    http.HandleFunc("/v0/status/", rest_auth(statusHandler))
    http.HandleFunc("/ws", wsHandler)
    go udpListen()
    log.Fatal(http.ListenAndServe(http_server_port, nil))
}