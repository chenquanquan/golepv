package models

import (
	"log"
	"regexp"
	. "strconv"
	"strings"
)

func CpuStat(client string) map[string]interface{} {
	lepd_command := "GetCmdMpstat"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	if len(response_lines) < 3 {
		return nil
	}

	response_lines = response_lines[3:] // discard the first three lines
	irq_ret := make(map[string]interface{})
	cpu_name := make(map[string]interface{})

	for _, line := range response_lines {
		if strings.Trim(line, " ") == "" {
			break
		}

		values := strings.Fields(line)

		if len(values) < 12 {
			break
		}

		irq_stat := make(map[string]float64)

		irq_stat["idle"], _ = ParseFloat(values[11], 32)
		irq_stat["gnice"], _ = ParseFloat(values[10], 32)
		irq_stat["guest"], _ = ParseFloat(values[9], 32)
		irq_stat["steal"], _ = ParseFloat(values[8], 32)
		irq_stat["soft"], _ = ParseFloat(values[7], 32)
		irq_stat["irq"], _ = ParseFloat(values[6], 32)
		irq_stat["iowait"], _ = ParseFloat(values[5], 32)
		irq_stat["system"], _ = ParseFloat(values[4], 32)
		irq_stat["nice"], _ = ParseFloat(values[3], 32)
		irq_stat["user"], _ = ParseFloat(values[2], 32)

		name := values[1]
		cpu_name[name] = irq_stat
	}
	irq_ret["data"] = cpu_name

	return irq_ret
}

func CpuSoftirq(client string) map[string]interface{} {
	lepd_command := "GetCmdMpstat-I"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	if len(response_lines) < 3 {
		return nil
	}

	response_lines = response_lines[2:] // discard the first 2 lines
	softirq_ret := make(map[string]interface{})
	var softirq_resp []string

	startIndex := 0
	for _, line := range response_lines {
		if strings.Trim(line, " ") == "" {
			startIndex++
		}

		if startIndex < 2 {
			continue
		} else if startIndex > 2 {
			break
		}

		softirq_resp = append(softirq_resp, line)
	}

	softirq_resp = softirq_resp[2:] // discard the first 2 lines
	cpu_name := make(map[string]interface{})
	for _, line := range softirq_resp {
		values := strings.Fields(line)
		irq_stat := make(map[string]float64)

		irq_stat["HRTIMER"], _ = ParseFloat(values[9], 32)
		irq_stat["TASKLET"], _ = ParseFloat(values[7], 32)
		irq_stat["NET_RX"], _ = ParseFloat(values[4], 32)
		irq_stat["NET_TX"], _ = ParseFloat(values[3], 32)

		name := values[1]
		cpu_name[name] = irq_stat
	}
	softirq_ret["data"] = cpu_name

	return softirq_ret
}

func CpuAvgload(client string) map[string]interface{} {
	lepd_command := "GetProcLoadavg"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	avg_ret := make(map[string]interface{})

	for _, line := range response_lines {
		values := strings.Fields(line)

		if len(values) < 3 {
			continue
		}

		avg_stat := make(map[string]interface{})

		avg_stat["last1"] = values[0]
		avg_stat["last5"] = values[1]
		avg_stat["last15"] = values[2]

		avg_ret["data"] = avg_stat
	}

	return avg_ret
}

func CpuTopOutput(client string) map[string]interface{} {
	lepd_command := "GetCmdTop"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	var validTitle = regexp.MustCompile(`\W*PID\W+USER\W+.*`)
	var header_line string

	for {
		header_line = response_lines[0]
		response_lines = response_lines[1:]

		if validTitle.MatchString(header_line) {
			break
		}
	}
	header_line = strings.Trim(header_line, " ")
	header_columns := strings.Fields(header_line)

	top_ret := make(map[string]interface{})
	top_data := make(map[string]interface{})
	top_ret["data"] = top_data

	top := make(map[string]interface{})
	for index, line := range response_lines {
		line = strings.Trim(line, " ")
		if line == "lepdendstring" {
			break
		}

		values := strings.Fields(line)

		column := make(map[string]interface{})
		for i, col := range header_columns {
			if col == "Name" || col == "CMD" {
				column[col] = values[i:]
			} else {
				column[col] = values[i]
			}
		}
		top[Itoa(index)] = column
	}

	top_data["top"] = top
	top_data["headerline"] = &header_line

	var androidTitle = regexp.MustCompile(`\W*PID\W+USER\W+PR\W+.*`)
	var linuxTitle = regexp.MustCompile(`\W*PID\W+USER\W+PRI\W+NI\W+VSZ\W+RSS\W+.*`)
	if androidTitle.MatchString(header_line) {
		top_data["os"] = "android"
	} else if linuxTitle.MatchString(header_line) {
		top_data["os"] = "linux"
	} else {
		log.Println("GetCmdTop command returned data from unrecognized system")
	}

	return top_ret
}
