package models

import (
	"regexp"
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

func IoTop(client string) map[string]interface{} {
	lepd_command := "GetCmdIotop"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	reg := regexp.MustCompile("\\s+")
	//var totalline, actualline string
	//var headline string
	var values_lines []string
	for i, line := range response_lines {
		if strings.Contains(line, "Total DISK READ") {
			/* head line like this:
			 * Total DISK READ :       0.00 B/s | Total DISK WRITE :       5.42 M/s
			 * Actual DISK READ:       0.00 B/s | Actual DISK WRITE:       0.00 B/s
			 *   TID  PRIO  USER     DISK READ  DISK WRITE  SWAPIN      IO    COMMAND
			 */
			//totalline = response_lines[0]
			//actualline = response_lines[1]
			//headline = response_lines[2]
			values_lines = response_lines[i+3:]
		}
	}

	//headline = strings.Trim(headline, " ")
	//headline = reg.ReplaceAllString(headline, " ")
	//title := strings.Split(headline, " ")
	//title_len := len(title)

	index := make(map[int]interface{})
	for i, line := range values_lines {
		item := make(map[string]interface{})

		line = strings.Trim(line, " ")
		line = reg.ReplaceAllString(line, " ")
		values := strings.Split(line, " ")

		if len(values) < 11 {
			continue
		}

		item["TID"] = values[0]
		item["PRIO"] = values[1]
		item["USER"] = values[2]
		item["READ"] = strings.Join(values[3:5], " ")
		item["WRITE"] = strings.Join(values[5:7], " ")
		item["SWAPIN"] = strings.Join(values[7:9], " ")
		item["IO"] = strings.Join(values[9:11], " ")
		item["COMMAND"] = strings.Join(values[11:], " ")

		index[i] = item
	}

	result := make(map[string]interface{})
	result["data"] = index

	return result
}

func JnetTop(client string) map[string]interface{} {
	lepd_command := "GetCmdJnettop"
	response_lines := ClientResponseString(client, lepd_command)
	if response_lines == nil {
		return nil
	}

	index := make(map[int]interface{})
	for i, line := range response_lines {
		item := make(map[string]interface{})

		line = strings.Trim(line, " ")
		values := strings.Split(line, ",")

		if len(values) < 20 {
			continue
		}

		/* jnettop --display text -t 5 --format
		   $src$,$srcname$,$srcport$,$srcbytes$,$srcpackets$,
		   $srcbps$,$srcpps$,$dst$,$dstname$,$dstport$,
		   $dstbytes$,$dstpackets$,$dstbps$,$dstpps$,$proto$,
		   $totalbytes$,$totalpackets$,$totalbps$,$totalpps$,$filterdata$
		*/
		item["Src"] = values[0]
		item["Src name"] = values[1]
		item["Src port"] = values[2]
		item["Src bytes"] = values[3]
		item["Src packets"] = values[4]
		item["Src bps"] = values[5]
		item["Src pps"] = values[6]
		item["Dst"] = values[7]
		item["Dst name"] = values[8]
		item["Dst port"] = values[9]
		item["Dst bytes"] = values[10]
		item["Dst packets"] = values[11]
		item["Dst bps"] = values[12]
		item["Dst pps"] = values[13]
		item["Proto"] = values[14]
		item["Total bytes"] = values[15]
		item["Total packets"] = values[16]
		item["Total bps"] = values[17]
		item["Total pps"] = values[18]
		item["Filter data"] = values[19]

		index[i] = item
	}

	result := make(map[string]interface{})
	result["data"] = index

	return result
}
