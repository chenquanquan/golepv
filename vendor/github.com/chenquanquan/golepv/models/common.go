package models

import (
	"encoding/json"
	"log"
	"net"
	"strings"
)

//"github.com/astaxie/beego"

type Client struct {
	Addr string
}

var (
	bufferLen  int
	ClientList []Client
)

func init() {
	bufferLen = 1024 * 1024
}

func AddClient(addr string) {
	for _, c := range ClientList {
		if addr == c.Addr {
			goto out
		}
	}

	ClientList = append(ClientList, Client{Addr: addr})
out:
}

func ClientResponse(server, method string) (map[string]interface{}, error) {
	body := make(map[string]string)
	body["method"] = method
	jbody, _ := json.Marshal(body)
	sbody := string(jbody)

	//log.Println(sbody)
	// connect to this socket
	conn, err := net.Dial("tcp", server+":12307")
	if err != nil {
		return nil, err
	}

	_, err = conn.Write([]byte(sbody))
	if err != nil {
		return nil, err
	}

	message := make([]byte, bufferLen)
	len, err := conn.Read(message)
	if err != nil {
		return nil, err
	}
	conn.Close()

	ret := make(map[string]interface{})
	err = json.Unmarshal(message[:len], &ret)
	if err != nil {
		log.Println(err.Error())
	}

	return ret, nil
}

func ClientResponseResult(client, method string) string {
	response, err := ClientResponse(client, method)
	if err != nil || response == nil {
		return ""
	}

	if response["result"] == nil {
		log.Printf("%s:%s failed", client, method)
		return ""
	}

	result := response["result"].(string)

	if result == "" {
		return ""
	}

	return result
}

func ClientResponseString(client, method string) []string {
	response, err := ClientResponse(client, method)
	if err != nil || response == nil {
		return nil
	}

	if response["result"] == nil {
		log.Printf("%s:%s failed", client, method)
		return nil
	}

	result := response["result"].(string)
	response_lines := strings.Split(result, "\n")

	if len(response_lines) == 0 {
		return nil
	}

	return response_lines
}
