__author__ = 'Ran Meng <1329597253@qq.com>'


__perf_script_lines = []

def burn(perf_script_lines):
    separate='@'
    # for p_s_l in perf_script_lines:
    #     print(p_s_l)
    # print("=================")

    __perf_script_lines = perf_script_lines

    list_for_flame = []
    dict_for_flame = {}
    json_for_flame = {"name": "root", "value": 0, "children": []}

    # get list_for_flame
    str = ''
    try:
        for line in __perf_script_lines:
            # print(line)

            if not line.strip():
                continue

            if not line.isspace() and line:
                if not line[0].isspace():
                    list_for_flame.append(str)
                    line = line.strip()
                    str = line.split()[0]
                elif line.split()[1] == '[unknown]':
                    str = str + separate + line.split()[2]
                else:
                    str = str + separate + line.split()[1]
        list_for_flame[0] = str

        # get dict_for_flame
        for line in list_for_flame:
            if line in dict_for_flame.keys():
                dict_for_flame[line] += 1
            else:
                dict_for_flame[line] = 1

        # get json_for_flame
        json_for_flame["value"] = len(list_for_flame)
        for line in dict_for_flame.items():
            li = line[0].split(separate)
            li.append(li[0])
            li.pop(0)
            li.reverse()
            __create_json(li, line[1], json_for_flame["children"])

        return json_for_flame
    except Exception as err:
        print(err, "-------  frame burn")
        return {}


def __create_json(line, value, children_list):
    if not line:
        return

    flags = 0
    for i in range(len(children_list)):
        if line[0] == children_list[i]["name"]:
            flags += 1
            children_list[i]["value"] += value
            __create_json(line[1:], value, children_list[i]["children"])
            break
    if flags == 0:
        dict = {"name": line[0], "value": value, "children": []}
        children_list.append(dict)
        __create_json(line[1:], value, children_list[-1]["children"])

