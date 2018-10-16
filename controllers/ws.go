package controllers

import (
	"github.com/astaxie/beego"
    "hello/models"
    "log"
    "fmt"
    "github.com/gorilla/websocket"
	"github.com/astaxie/beego/session"	
    // "time"
)

var upgrader = websocket.Upgrader{}
var GlobalSessions *session.Manager

func init() {
	initSession()
}

type WsController struct {
	beego.Controller

}

func (c *WsController) Get() {

    sess, _ := GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)


	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
    if err != nil {
        log.Fatal(err)
    }
    //  defer ws.Close()

    clients[ws] = true
    
     // Read in a new message as JSON and map it to a Message object
  
	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)

		if err != nil {
			name := sess.Get("loginUser")
		    msg:= models.Message{Type : "quit", Client_name:name.(string), Client_img:"" , Content:"退出了聊天室"}   
		    broadcast <- msg
		    sess.Delete("loginuser") 
		    delete(clients, ws)
		    break

		} else {
            switch msg.Type{
               case "login": 
               msg:= models.Message{Type : "login", Client_name:msg.Client_name , Client_img:"" , Content:"进入了聊天室"}  
               
               sess.Set("loginUser", msg.Client_name)
          
               broadcast <- msg
               break
               case "say": 
               msg:= models.Message{Type : "say", Client_name:msg.Client_name , Client_img:msg.Client_img , Content:msg.Content} 
               broadcast <- msg
               break
               case "quit": 
               msg:= models.Message{Type : "quit", Client_name:msg.Client_name , Client_img:"" , Content:"退出了聊天室"}   
               broadcast <- msg
               break
            }
            
      
			name := sess.Get("loginUser")
	        fmt.Println("session2:",name.(string))
		    fmt.Println("类型 名称",msg)
		}

      
    }
}

func initSession() {
    sessionConfig := &session.ManagerConfig{
    CookieName:"gosessionid", 
    EnableSetCookie: true, 
    Gclifetime:3600,
    Maxlifetime: 3600, 
    Secure: false,
    CookieLifeTime: 3600,
    ProviderConfig: "data/session",
    }
    GlobalSessions, _ = session.NewManager("file",sessionConfig)
    go GlobalSessions.GC()
}
