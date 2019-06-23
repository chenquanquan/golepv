package controllers

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/chenquanquan/golepv/models"
	"github.com/gorilla/websocket"
)

type MonitorController struct {
	beego.Controller
}

var (
	mclients = make(map[string]models.MonitorClient) // Client list
)

func mergeMethod(method, add []string) []string {
	list := []string{}

	for i := range add {
		flag := 0
		for j := range method {
			if method[j] == add[i] {
				flag = 1
				break
			}
		}
		if flag == 0 {
			list = append(list, add[i])
		}
	}

	return append(method, list...)
}

func monitorHandle(client *websocket.Conn) {
	for {
		_, r, err := client.ReadMessage()
		if err != nil {
			client.Close()
			return
		}

		input := make(map[string]string)
		err = json.Unmarshal(r, &input)
		if err != nil {
			log.Println("Error input:")
			log.Println(input)
			client.Close()
			return
		}

		server := input["server"]
		method := strings.Split(input["method"], ",")

		for s, m := range mclients { // server, monitor
			if len(m.Conn) == 0 { // If server in list but no webclient, rebuild the server struct
				delete(mclients, s)
				break
			}

			if s == server { // If server in list with webclient, add websocket to connect list
				var wc models.WebConn
				wc.Connected = true
				m.Conn[client] = &wc
				m.Method = mergeMethod(m.Method, method)
				return
			}
		}

		// Create a new monitor
		var monitor models.MonitorClient
		monitor.Conn = make(map[*websocket.Conn]*models.WebConn)
		var wc models.WebConn
		wc.Connected = true
		monitor.Conn[client] = &wc

		monitor.Server = input["server"]
		monitor.Method = strings.Split(input["method"], ",")
		monitor.LepdTimer = time.NewTimer(time.Second)
		monitor.WebTimer = time.NewTimer(time.Second)
		mclients[server] = monitor

		go models.WebServer(monitor)
		go models.LepdServer(monitor)
	}
}

func (c *MonitorController) Get() {
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	go monitorHandle(ws)

	c.Ctx.Output.SetStatus(200)
}
