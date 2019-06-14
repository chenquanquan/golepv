package models

import (
	"regexp"
	. "strconv"
	"strings"
)

func MemStatus(client string) map[string]interface{} {
	lepd_command := "GetProcMeminfo"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	if len(response_lines) == 0 {
		return nil
	}

	stat_ret := make(map[string]interface{})
	stat_component := make(map[string]interface{})
	stat_ret["data"] = stat_component
	stat_result := make(map[string]string)

	for _, line := range response_lines {
		line = strings.Trim(line, " ")
		if line == "lepdendstring" {
			break
		}
		linePairs := strings.Split(line, ":")
		lineKey := strings.Trim(linePairs[0], " ")
		lineValue := strings.ReplaceAll(linePairs[1], "kB", " ")
		lineValue = strings.Trim(lineValue, " ")
		stat_result[lineKey] = lineValue
	}

	var val, total, used int

	stat_component["name"] = "memory"
	val, _ = Atoi(stat_result["MemTotal"])
	total = val
	used = val
	stat_component["total"] = val / 1024
	val, _ = Atoi(stat_result["MemFree"])
	used -= val
	stat_component["free"] = val / 1024
	val, _ = Atoi(stat_result["Buffers"])
	used -= val
	stat_component["buffers"] = val / 1024
	val, _ = Atoi(stat_result["Cached"])
	used -= val
	stat_component["cached"] = val / 1024

	stat_component["used"] = used / 1024
	stat_component["ratio"] = (used * 100 / total)
	stat_component["unit"] = "MB"

	return stat_ret
}

func MemProcrank(client string) map[string]interface{} {
	lepd_command := "GetCmdProcrank"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	procrank_ret := make(map[string]interface{})
	procrank_data := make(map[string]interface{})
	procrank_proc := make(map[string]interface{})
	procrank_sum := make(map[string]interface{})
	procrank_ret["data"] = procrank_data
	procrank_data["procranks"] = procrank_proc
	procrank_data["sum"] = procrank_sum

	response_lines = response_lines[1:] // discard the first line

	var validTail = regexp.MustCompile(`\W+-+\W+-+\W-+.*`)

	for index, line := range response_lines {
		if validTail.MatchString(line) {
			break
		}

		values := strings.Fields(line)
		procrank_index := make(map[string]interface{})

		procrank_index["pid"] = values[0]
		procrank_index["vss"] = values[1]
		procrank_index["rss"] = values[2]
		procrank_index["pss"] = values[3]
		procrank_index["uss"] = values[4]
		procrank_index["cmdline"] = values[5]

		procrank_proc[Itoa(index)] = procrank_index
	}

	// now parse from end, which contains summary info
	last_line := response_lines[len(response_lines)-1]
	last_line = strings.ReplaceAll(last_line, "RAM:", "")
	last_values := strings.Split(last_line, ", ")
	for _, value := range last_values {
		kv := strings.Fields(value)
		if len(kv) < 2 {
			return nil
		}
		name := strings.Trim(kv[1], " ")
		value := strings.Trim(kv[0], " ")
		procrank_sum[name+"Unit"] = string(value[len(value)-1])
		procrank_sum[name] = value[:len(value)-1]
	}

	xss_sum_line := response_lines[len(response_lines)-3]
	xss_total := strings.Fields(xss_sum_line)

	uss := xss_total[0]
	procrank_sum["ussTotalUnit"] = string(uss[len(uss)-1])
	procrank_sum["ussTotal"] = uss[:len(uss)-1]

	pss := xss_total[1]
	procrank_sum["pssTotalUnit"] = string(pss[len(pss)-1])
	procrank_sum["pssTotal"] = pss[:len(pss)-1]

	return procrank_ret
}
