package models

import (
	"log"
	"strings"
	"time"
)

func getIostatResult(client string) []string {
	lepd_command := "GetCmdIostat"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	for i, line := range response_lines {
		if strings.Contains(line, "Device:") {
			return response_lines[i:]
		}
	}

	return nil
}

func IoStatus(client string) map[string]interface{} {
	start_time := time.Now()
	response_lines := getIostatResult(client)
	if response_lines == nil {
		return nil
	}
	end_time := time.Now()

	//header_line := response_lines[0]
	value_line := response_lines[1:]

	stat_ret := make(map[string]interface{})
	status := make(map[string]interface{})
	disks := make(map[string]interface{})

	stat_ret["data"] = status
	stat_ret["raw_ret"] = response_lines
	status["disks"] = disks
	status["lepdDuration"] = end_time.UnixNano() - start_time.UnixNano()
	status["diskCount"] = 0

	for _, line := range value_line {
		if strings.Trim(line, " ") == "" {
			break
		}
		values := strings.Fields(line)

		device := make(map[string]interface{})
		name := values[0]

		disks[name] = device
		device["rkbs"] = values[5]
		device["wkbs"] = values[6]
		device["radio"] = values[13]

		count := status["diskCount"].(int)
		count++
		status["diskCount"] = count
	}

	end_time_2 := time.Now()

	status["lepvParsingDuration"] = end_time_2.UnixNano() - start_time.UnixNano()

	return stat_ret
}

// TODO: no work
func IoTop(client string) map[string]interface{} {
	lepd_command := "GetCmdIotop"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	for i, line := range response_lines {
		log.Println(i)
		log.Println(line)
	}

	return nil
}
