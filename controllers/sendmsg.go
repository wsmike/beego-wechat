package controllers

import (
    "hello/models"
    "fmt"
    "log"

    "github.com/gorilla/websocket"
)

var (
    clients   = make(map[*websocket.Conn]bool)
    broadcast = make(chan models.Message)
)

func init() {
    go handleMessages()
}

//广播发送至页面
func handleMessages() {
    for {
        msg := <-broadcast

        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                client.Close()
                delete(clients, client)
            }
        }  
        
    }
}
