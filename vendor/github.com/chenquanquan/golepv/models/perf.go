package models

import (
	"strings"
)

func PerfCpuClock(client string) map[string]interface{} {
	lepd_command := "GetCmdPerfCpuclock"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	var perf_line []string
	for i, line := range response_lines {
		if strings.Contains(line, "# Overhead") {
			perf_line = response_lines[i+3:]
		}
	}

	perf_ret := make(map[string]interface{})
	var result_list []interface{}

	for _, line := range perf_line {
		if strings.Trim(line, " ") == "" {
			break
		}
		values := strings.Fields(line)

		if len(values) < 5 {
			continue
		}

		if !strings.Contains(values[0], "%") {
			continue
		}

		result := make(map[string]interface{})

		result["Overhead"] = values[0]
		result["Command"] = values[1]
		result["Shared Object"] = values[2]
		result["Symbol"] = values[3:]

		result_list = append(result_list, result)
	}

	perf_ret["data"] = result_list

	return perf_ret
}

// TODO: no work
func PerfFlame(client string) map[string]interface{} {
	//lepd_command := "GetCmdPerfFlame"
	//response_lines := ClientResponseString(client, lepd_command)
	//if response_lines == nil {
	//	return nil
	//}

	//flame := make(map[string]interface{})

	//flame["flame"] = fameBurner(response_lines)
	//flame["hierarchy"] = ""
	//flame["perf_script_output"] = response_lines

	flame := make(map[string]interface{})
	children1 := make(map[string]interface{})
	children2 := make(map[string]interface{})
	//children3 := make(map[string]interface{})
	flame["name"] = "root"
	flame["value"] = 7412
	flame["children"] = []map[string]interface{}{children1, children2}
	children1["name"] = "children1"
	children1["value"] = 89
	children2["name"] = "children2"
	children2["value"] = 89

	return flame
}
