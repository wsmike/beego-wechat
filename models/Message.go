package models

type Message struct {
    Type string `json:"type"`
    Client_name string `json:"client_name"`
    Client_img string `json:"client_img"`
    Content string `json:"content"`
}
