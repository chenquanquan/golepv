package models

import (
	"container/list"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type LepdFunc func(client string) map[string]interface{}

type WebConn struct {
	sync.Mutex
	Connected bool
}

type MonitorClient struct {
	Server    string
	LepdTimer *time.Timer //time.NewTicker(d time.Duration)
	WebTimer  *time.Timer //time.NewTicker(d time.Duration)
	Conn      map[*websocket.Conn]*WebConn
	Method    []string
}

var (
	listResult *list.List
	listMutex  sync.Mutex
)

func init() {
	listResult = list.New()
}

func getLepd(method, server string, f LepdFunc) {
	result := make(map[string]interface{})

	result[method] = f(server)
	//log.Printf("get :%s, len :%d", method, len(method))

	listMutex.Lock()
	listResult.PushBack(result)
	listMutex.Unlock()
}

func WebServer(monitor MonitorClient) {
	for {
		<-monitor.WebTimer.C

		//log.Printf("Web server len: %d", listResult.Len())

		if len(monitor.Conn) == 0 {
			log.Println("Web connect break")
			return
		} else {
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

			for client, wc := range monitor.Conn {
				wc.Lock()
				err := client.WriteJSON(result)
				wc.Unlock()
				if err != nil {
					client.Close()
					wc := monitor.Conn[client]
					wc.Connected = false
					delete(monitor.Conn, client)

					continue
				}
			}
		}

		monitor.WebTimer.Reset(time.Second)
	}
}

func LepdServer(monitor MonitorClient) {
	counter := 0

	for {
		<-monitor.LepdTimer.C

		if len(monitor.Conn) == 0 {
			log.Println("Lepd connect break")
			return
		} else {
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

				if counter%i != 0 {
					continue
				}

				if keylist[key] != nil {
					continue
				}
				keylist[key] = true

				switch key {
				case "cpu.stat":
					go getLepd(method, server, CpuStat)
				case "cpu.softirq":
					go getLepd(method, server, CpuSoftirq)
				case "cpu.avgload":
					go getLepd(method, server, CpuAvgload)
				case "cpu.top":
					go getLepd(method, server, CpuTopOutput)
				case "memory.status":
					go getLepd(method, server, MemStatus)
				case "memory.procrank":
					go getLepd(method, server, MemProcrank)
				case "io.status":
					go getLepd(method, server, IoStatus)
				case "io.top":
					go getLepd(method, server, IoTop)
				case "io.jnet":
					go getLepd(method, server, JnetTop)
				case "perf.cpuclock":
					go getLepd(method, server, PerfCpuClock)
				case "perf.flame":
					go getLepd(method, server, PerfFlame)
				default:
				}
			}

		}

		counter++
		monitor.LepdTimer.Reset(time.Second)
	}
}
