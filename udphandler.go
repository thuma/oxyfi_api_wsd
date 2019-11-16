package main

import (
    "net"
    "fmt"
)

func udpListen() {
    message := make([]byte, 1500)
    pc, err := net.ListenPacket("udp", udp_port)
    if err != nil {
            fmt.Println("Unable To Open UDP port")
            return
    }
    for {
        _, _, err := pc.ReadFrom(message)
        if err != nil {
            fmt.Println("Unable To Read Message")
            return
        }
        for k := range listeners {
            if !send(listeners[k], message) {
                delete(listeners, k)
                }
            }
        }
    }
