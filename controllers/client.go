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
	upgrader = websocket.Upgrader{}
)

func writeClientList(client *websocket.Conn) {
	result := make(map[string]string)

	for i, c := range models.ClientList {
		key := "addr" + strconv.Itoa(i)
		result[key] = c.Addr
	}

	err := client.WriteJSON(result)
	if err != nil {
		client.Close()
		delete(clients, client)
	}
}

func clientHandler(client *websocket.Conn) {
	for {
		_, r, err := client.ReadMessage()
		if err != nil {
			client.Close()
			delete(clients, client)
			return
		}

		input := make(map[string]string)
		err = json.Unmarshal(r, &input)
		if err != nil {
			continue
		}

		for cmd, v := range input {
			result := make(map[string]string)
			switch cmd {
			case "addClient":
				models.AddClient(v)

				for c := range clients {
					writeClientList(c)
				}

			default:
				result["error"] = "No this command"
			}
		}
	}
}

func (c *ClientController) Get() {
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	/* add websocket to list */
	clients[ws] = true
	writeClientList(ws)
	go clientHandler(ws)

	c.Ctx.Output.SetStatus(200)
}

