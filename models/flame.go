package models

import (
	"strings"
)

func _createJson(line []string, value int, list []map[string]interface{}) []map[string]interface{} {
	if len(line) == 0 {
		return nil
	}

	name := line[0]
	flag := 0

	for i := range list {
		dict := list[i]
		if dict["name"] == name {
			flag++
			count := dict["value"].(int)
			count += value
			dict["value"] = count
			dict["children"] = _createJson(line[1:], value, dict["children"].([]map[string]interface{}))
			break
		}

	}

	if flag == 0 {
		dict := make(map[string]interface{})
		dict["value"] = value
		dict["name"] = name
		dict["children"] = _createJson(line[1:], value, nil)

		list = append(list, dict)
	}

	return list

}

func flameBurner(input []string) interface{} {
	separate := "@"
	var list_for_flame []string
	dict_for_flame := make(map[string]interface{})
	children_list := []map[string]interface{}{}
	result := make(map[string]interface{})

	str := ""
	for _, line := range input {
		if strings.Trim(line, " ") == "" {
			continue
		}

		values := strings.Fields(line)

		if len(values) < 2 {
			continue
		}

		if strings.Index(line, "\t") != 0 {
			list_for_flame = append(list_for_flame, str)
			str = values[0]
		} else if strings.Contains(values[1], "[unknown]") {
			str = str + separate + values[2]
		} else {
			str = str + separate + values[1]
		}

	}
	list_for_flame[0] = str

	for _, flame := range list_for_flame {
		if dict_for_flame[flame] != nil {
			count := dict_for_flame[flame].(int)
			count++
			dict_for_flame[flame] = count
		} else {
			dict_for_flame[flame] = 1
		}
	}

	for flame, dict := range dict_for_flame {
		li := strings.Split(flame, separate)
		li = append(li[1:], li[0])
		li_len := len(li)
		new_li := make([]string, li_len)

		for i, v := range li {
			new_li[li_len-1-i] = v
		}

		children_list = _createJson(new_li, dict.(int), children_list)
	}

	result["children"] = children_list
	result["value"] = len(list_for_flame)
	result["name"] = "root"

	return result
}
