package controllers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/chenquanquan/golepv/models"
	"github.com/gorilla/websocket"
)

type MonitorController struct {
	beego.Controller
}

type MonitorClient struct {
	Server string
	Timer  *time.Timer //time.NewTicker(d time.Duration)
	Conn   map[*websocket.Conn]bool
	Method []string
}

var (
	mclients = make(map[string]MonitorClient) // Client list
)

func mergeMethod(method, add []string) []string {
	list := []string{}

	for i := range add {
		flag := 0
		for j := range method{
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

func monitorServer(monitor MonitorClient) {
	var counter int

	for {
		<-monitor.Timer.C

		result := make(map[string]interface{})
		keylist := make(map[string]interface{})

		for _, method := range monitor.Method {
			params := strings.Split(method, "@")
			if len(params) < 2 {
				continue
			}

			key := params[0]
			interval := params[1]

			i, err := strconv.Atoi(interval)
			if err != nil {
				continue
			}

			if counter % i != 0 {
				continue
			}

			if keylist[key] != nil {
				continue
			}
			keylist[key] = true

			switch key {
			case "cpu.stat":
				result[method] = models.CpuStat(monitor.Server)
			case "cpu.softirq":
				result[method] = models.CpuSoftirq(monitor.Server)
			case "cpu.avgload":
				result[method] = models.CpuAvgload(monitor.Server)
			case "cpu.top":
				result[method] = models.CpuTopOutput(monitor.Server)
			case "memory.status":
				result[method] = models.MemStatus(monitor.Server)
			case "memory.procrank":
				result[method] = models.MemProcrank(monitor.Server)
			case "io.status":
				result[method] = models.IoStatus(monitor.Server)
			case "io.top":
				result[method] = models.IoTop(monitor.Server)
			case "perf.cpuclock":
				result[method] = models.PerfCpuClock(monitor.Server)
			case "perf.flame":
				result[method] = models.PerfFlame(monitor.Server)
			default:
			}
		}

		result["time"] = time.Now().UnixNano()
		result["client"] = monitor.Server

		for client := range monitor.Conn {
			err := client.WriteJSON(result)
			if err != nil {
				client.Close()
				monitor.Conn[client] = false
				delete(monitor.Conn, client)

				continue
			}
		}

		if len(monitor.Conn) == 0 {
			return
		}

		counter++
		monitor.Timer.Reset(2 * time.Second)
	}
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
			if s == server {
				m.Conn[client] = true
				m.Method = mergeMethod(m.Method, method)
				return
			}
		}


		// Create a new monitor
		var monitor MonitorClient
		monitor.Conn = make(map[*websocket.Conn]bool)
		monitor.Conn[client] = true
		monitor.Server = input["server"]
		monitor.Method = strings.Split(input["method"], ",")
		monitor.Timer = time.NewTimer(time.Second)
		mclients[server] = monitor

		go monitorServer(monitor)
	}
}

func (c *MonitorController) Get() {
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	go monitorHandle(ws)

	//c.ServeJSON() /* Return empty */
	c.Ctx.Output.SetStatus(200)
}
