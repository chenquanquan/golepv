package models

import (
	//"log"
	"strings"
)

func _createJson(line []string, value int, children_list []interface{}) []interface{} {
	if len(line) == 0 {
		return nil
	}

	//flag := 0
	//for _, children := range children_list {
	//	item, err := children.(map[string]interface{})
	//	log.Println(err)
	//	if item["name"] == line[0] {
	//		flag++
	//		v := item["value"].(int)
	//		v += value
	//		item["value"] = v
	//		item["children"] = _createJson(line[1:], value, item["children"].([]interface{}))
	//		break;
	//	}
	//}

	//if flag == 0 {
	//	dict := make(map[string]interface{})
	//	dict["name"] = line[0]
	//	dict["value"] = value
	//	dict["children"] = nil
	//	children_list = append(children_list, dict)
	//	children_list[len(children_list) - 1] = _createJson(line[1:], value, dict["children"])
	//}

	dict := make(map[string]interface{})
	dict["name"] = line[0]
	dict["value"] = value
	dict["children"] = _createJson(line[1:], value, nil)
	children_list = append(children_list, dict)

	return children_list

}

func flameBurner(input []string) interface{} {
	separate := "@"
	var list_for_flame []string
	dict_for_flame := make(map[string]interface{})
	var children_list []interface{}
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
			dict_for_flame[flame]=count
		} else {
			dict_for_flame[flame] = 1
		}
	}

	for flame, dict := range dict_for_flame {
		li := strings.Split(flame, separate)
		li = append(li, li[0])
		li = li[1:]
		li_len := len(li)
		new_li := make([]string, li_len)
		for i, v := range li {
			new_li[li_len - 1 - i] = v
		}
		children_list = _createJson(new_li, dict.(int), children_list)
	}

	//log.Println("list:")
	//log.Println(list_for_flame)
	//log.Println("dict")
	//log.Println(dict_for_flame)
	//for i, j := range dict_for_flame {
	//	log.Println(i)
	//	log.Println(j)
	//}


	result["children"] = children_list
	result["value"] = len(list_for_flame)
	result["name"] = "root"
	result["children"] = list_for_flame

	return result
}
