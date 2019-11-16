package main

import (
    "net/http"
    "time"
    "github.com/gorilla/websocket"
    "log"
)

var listeners = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func close(conn *websocket.Conn){
    conn.WriteControl(
        websocket.CloseMessage,
        websocket.FormatCloseMessage(
            websocket.CloseGoingAway,
            ""), 
        time.Now().Add(time.Second))
}

func send(conn *websocket.Conn, message []byte) (bool){
    err := conn.WriteMessage(websocket.TextMessage, message)
    if err != nil {
        return false
    }
    return true
}

func readLoop(c *websocket.Conn) {
    for {
        if _, _, err := c.NextReader(); err != nil {
            c.Close()
            break
        }
    }
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    if key == "" {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("401 - Unauthorized, key query parameter not valid. "))
        return
    }
    if get_key(key).Active == false {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write([]byte("401 - Unauthorized, key is not active. "))
        return
    }
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    go readLoop(conn)
    if _, ok := listeners[key]; ok {
        close(listeners[key])
    }
    listeners[key] = conn
}