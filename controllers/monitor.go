package controllers

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"strings"
	"time"
	"container/list"

	"github.com/astaxie/beego"
	"github.com/chenquanquan/golepv/models"
	"github.com/gorilla/websocket"
)

type MonitorController struct {
	beego.Controller
}

type webConn struct {
	Connected bool
	Mux sync.Mutex
}

type MonitorClient struct {
	Server string
	LepdTimer  *time.Timer //time.NewTicker(d time.Duration)
	WebTimer  *time.Timer //time.NewTicker(d time.Duration)
	Conn   map[*websocket.Conn]webConn
	Method []string
}

type LepdFunc func (client string) map[string]interface{}

var (
	mclients = make(map[string]MonitorClient) // Client list
	//dataResult = make(chan map[string]interface{})
	resultFlag = make(chan bool)
	listResult *list.List
	listMutex sync.Mutex
)

func init() {
	listResult = list.New()
}

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

func getLepd(method, server string, f LepdFunc) {
	result := make(map[string]interface{})


	result[method] = f(server)
	//log.Printf("get :%s, len :%d", method, len(method))

	listMutex.Lock()
	listResult.PushBack(result)
	listMutex.Unlock()
}

func monitorWebServer(monitor MonitorClient) {
	for {
		<-monitor.WebTimer.C

		//log.Printf("Web server len: %d", listResult.Len())

		result := make(map[string]interface{})

		listMutex.Lock()
		for listResult.Len() > 0 {
			e := listResult.Front() // First element

			item := e.Value.(map[string]interface{})
			/* merge result */
			for k, v := range item {
				result[k] = v
			}

			listResult.Remove(e) // Dequeue
		}
		listMutex.Unlock()


		result["time"] = time.Now().UnixNano()
		result["client"] = monitor.Server

		for client, wc:= range monitor.Conn {
			wc.Mux.Lock()
			err := client.WriteJSON(result)
			wc.Mux.Unlock()
			if err != nil {
				client.Close()
				wc := monitor.Conn[client]
				wc.Connected = false
				delete(monitor.Conn, client)

				continue
			}
		}

		if len(monitor.Conn) == 0 {
			log.Println("Web connect break")
			return
		}

		monitor.WebTimer.Reset(time.Second)
	}
}

func monitorLepdServer(monitor MonitorClient) {
	var counter int

	for {
		<-monitor.LepdTimer.C

		server := monitor.Server
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
				go getLepd(method, server, models.CpuStat)
			case "cpu.softirq":
				go getLepd(method, server, models.CpuSoftirq)
			case "cpu.avgload":
				go getLepd(method, server, models.CpuAvgload)
			case "cpu.top":
				go getLepd(method, server, models.CpuTopOutput)
			case "memory.status":
				go getLepd(method, server, models.MemStatus)
			case "memory.procrank":
				go getLepd(method, server, models.MemProcrank)
			case "io.status":
				go getLepd(method, server, models.IoStatus)
			case "io.top":
				go getLepd(method, server, models.IoTop)
			case "io.jnet":
				go getLepd(method, server, models.JnetTop)
			case "perf.cpuclock":
				go getLepd(method, server, models.PerfCpuClock)
			case "perf.flame":
				go getLepd(method, server, models.PerfFlame)
			default:
			}
		}

		if len(monitor.Conn) == 0 {
			log.Println("Lepd connect break")
			return
		}

		counter++
		monitor.LepdTimer.Reset(time.Second)
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
				wc := m.Conn[client]
				wc.Connected = true
				m.Method = mergeMethod(m.Method, method)
				return
			}
		}


		// Create a new monitor
		var monitor MonitorClient
		monitor.Conn = make(map[*websocket.Conn]webConn)
		var wc webConn
		//wc := monitor.Conn[client]
		wc.Connected = true
		monitor.Conn[client]=wc

		monitor.Server = input["server"]
		monitor.Method = strings.Split(input["method"], ",")
		monitor.LepdTimer = time.NewTimer(time.Second)
		monitor.WebTimer = time.NewTimer(time.Second)
		mclients[server] = monitor

		go monitorWebServer(monitor)
		go monitorLepdServer(monitor)
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
