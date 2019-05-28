package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type MainController struct {
	beego.Controller
}

func initControllerData(c *MainController) {
	langs := make(map[string]string)

        langs["Summary"] = "概况"
        langs["CPU"] = "处理器"
        langs["Memory"] = "内存"
        langs["IO"] = "磁盘"
        langs["Perf"] = "Perf"

        langs["SystemSummary"] = "系统概况"
        langs["TotalDiskSpace"] = "磁盘总空间"
        langs["FreeDiskSpace"] = "空闲磁盘空间"
        langs["Settings"] = "设置"
        langs["Server"] = "服务器"
        langs["Port"] = "端口"
        langs["Connect"] = "连接"
        langs["Configurations"] = "配置"
        langs["Language"] = "语言"

        langs["perfTableTitle"] = "基于Symbol的时间分布"
        langs["perfTableTitleFull"] = "基于Symbol的时间分布 (perf top)"

        langs["ioChartTitle"] = "I/O吞吐量"
        langs["ioTopTableTitle"] = "I/O TOP"

        langs["memoryConsumptionChartTitle"] = "内存消耗"
        langs["ramChartTitle"] = "消耗的内存、Page Cahe(Buffers, Cached)和空闲的内存"
        langs["memoryPssDonutChartTitle"] = "应用程序内存消耗比例分布(基于PSS)"
        langs["memoryPssAgainstTotalChartTitle"] = "应用程序耗费内存占总内存比例"

        langs["averageLoadChartTitle"] = "Average Load"
        langs["averageLoadChartTitleFull"] = "Average Load: CPU的平均负载; 当达到0.7*核数时，负载较重; 当达到1.0*核数时，CPU是性能瓶颈"

        langs["cpuUserGroupChartTitle"] = "CPU Stat: User+Sys+Nice"
        langs["cpuUserGroupChartTitleFull"] = "CPU Stat: User+Sys+Nice; 进程上下文占据的时间比例，如果是多核，可由此观察CPU的负载均衡"

        langs["cpuIdleGroupChartTitle"] = "CPU Stat: Idle"
        langs["cpuIdleGroupChartTitleFull"] = "CPU Stat: Idle; 系统空闲占据的时间比例，如果是多核，可由此观察CPU的负载均衡"

        langs["cpuIrqGroupChartTitle"] = "CPU Stat: IRQ + SoftIRQ"
        langs["cpuIrqGroupChartTitleFull"] = "CPU Stat: IRQ + SoftIRQ; 中断/软中断占据的时间比例，如果是多核，可由此观察中断/软中断的负载均衡"

        langs["cpuIrqChartTitle"] = "CPU Stat: IRQ"
        langs["cpuIrqChartTitleFull"] = "CPU Stat: IRQ; 中断/软中断占据的时间比例，如果是多核，可由此观察中断的负载均衡"

        langs["cpuNettxIrqChartTitle"] = "CPU Stat: SoftIRQ - NET_TX"
        langs["cpuNettxIrqChartTitleFull"] = "CPU Stat: SoftIRQ - NET_TX"

        langs["cpuNetrxIrqChartTitle"] = "CPU Stat: SoftIRQ - NET_RX"
        langs["cpuNetrxIrqChartTitleFull"] = "CPU Stat: SoftIRQ - NET_RX"

        langs["cpuTaskletIrqChartTitle"] = "CPU Stat: SoftIRQ - TASKLET"
        langs["cpuTaskletIrqChartTitleFull"] = "CPU Stat: SoftIRQ - TASKLET"

        langs["cpuHrtimerIrqChartTitle"] = "CPU Stat: SoftIRQ - HRTIMER"
        langs["cpuHrtimerIrqChartTitleFull"] = "CPU Stat: SoftIRQ - HRTIMER"

        langs["cpuStatRatioChartTitle"] = "CPU Stat"
        langs["cpuStatRatioChartTitleFull"] = "CPU Stat; 各种上下文占据的时间"

	c.Data["languages"] = langs

	widgets := make(map[string]string)
	widgets["input"] = "127.0.0.0"
	c.Data["widgets"] = widgets

	c.Data["watch"] = ""
}


func addClient(in string)(out string){
	log.Println("addClient:" + in)
	out =  in + "world"
	return
}

func initFuncMap() {
	beego.AddFuncMap("addClient", addClient)
}

func init() {
	initFuncMap()

	beego.SetStaticPath("/static","static")
	beego.SetStaticPath("/socket.io", "static/vendors/socket-io")
}

func (c *MainController) Get() {
	initControllerData(c)

	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	log.Println("Post")
	//initControllerData(c)

	//c.TplName = "index.tpl"
}
