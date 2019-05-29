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
	Conn   *websocket.Conn
	Method []string
}

var (
	mclients = make(map[*websocket.Conn]MonitorClient) // Client list
	//mbroadcast = make(chan bool)                         // New message flag
)

func monitorServer(monitor MonitorClient) {
	for {
		<-monitor.Timer.C

		result := make(map[string]interface{})
		for _, m := range monitor.Method {
			switch m {
			case "cpu.stat":
				result[m] = models.CpuStat(monitor.Server)
			case "cpu.softirq":
				result[m] = models.CpuSoftirq(monitor.Server)
			case "cpu.avgload":
				result[m] = models.CpuAvgload(monitor.Server)
			case "cpu.top":
				result[m] = models.CpuTopOutput(monitor.Server)
			case "memory.status":
				result[m] = models.MemStatus(monitor.Server)
			case "memory.procrank":
				result[m] = models.MemProcrank(monitor.Server)
			case "io.status":
				result[m] = models.IoStatus(monitor.Server)
			case "io.top":
				result[m] = models.IoTop(monitor.Server)
			case "perf.cpuclock":
				result[m] = models.PerfCpuClock(monitor.Server)
			case "perf.flame":
				result[m] = models.PerfFlame(monitor.Server)
			default:
			}
		}

		result["time"] = time.Now().UnixNano()

		client := monitor.Conn
		err := client.WriteJSON(result)
		if err != nil {
			client.Close()
			delete(mclients, client)

			return
		}

		monitor.Timer.Reset(2 * time.Second)
	}
}

func monitorHandle(client *websocket.Conn, monitor MonitorClient) {
	for {
		_, r, err := client.ReadMessage()
		if err != nil {
			client.Close()
			delete(mclients, client)
			return
		}

		input := make(map[string]string)
		err = json.Unmarshal(r, &input)
		if err != nil {
			log.Println("Error input:")
			log.Println(input)
			client.Close()
			delete(mclients, client)
		}

		monitor.Conn = client
		monitor.Server = input["server"]
		monitor.Method = strings.Split(input["method"], ",")

		if monitor.Timer == nil {
			sec, err := strconv.Atoi(input["interval"])
			if err == nil {
				monitor.Timer = time.NewTimer(time.Second *
					time.Duration(sec))

				go monitorServer(monitor)
			}

		}
	}
}

func (c *MonitorController) Get() {
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	/* add websocket to list */
	mclients[ws] = MonitorClient{Timer: nil}
	//mbroadcast <- true

	go monitorHandle(ws, mclients[ws])

	//c.ServeJSON() /* Return empty */
	c.Ctx.Output.SetStatus(200)
}
