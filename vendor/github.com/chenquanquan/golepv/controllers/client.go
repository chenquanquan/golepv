package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
	"strconv"
	"github.com/chenquanquan/golepv/models"
)

type ClientController struct {
	beego.Controller
}

var (
	clients   = make(map[*websocket.Conn]bool) // Client list
	broadcast = make(chan bool) // New message flag
	upgrader = websocket.Upgrader{}
)

func handleMessages() {
        for {
                // Wait new message
                <-broadcast
                // Anaylsis command
                for client := range clients {
			_, r, err := client.ReadMessage()
			if err != nil {
				client.Close()
				delete(clients, client)
				continue
			}

			input := make(map[string]string)
			err = json.Unmarshal(r, &input)
			if err != nil {
				log.Println("Error input:")
				log.Println(input)
				client.Close()
				delete(clients, client)
			}

			for cmd, v := range input {
				result := make(map[string]string)
				switch cmd {
				case "addClient":
					models.AddClient(v)

					for i, c := range models.ClientList {
						key := "addr" + strconv.Itoa(i)
						result[key] = c.Addr
					}
				default:
					result["error"] = "No this command"
				}
				js, _ := json.Marshal(result)
				_ = client.WriteJSON(string(js))
			}

			client.Close()
			delete(clients, client)
		}
	}
}

func init() {
	go handleMessages()
}

func (c *ClientController) Get() {
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	/* add websocket to list */
	clients[ws] = true
	broadcast <- true

	c.ServeJSON() /* Return empty */
}

