package models

import (
	//"bytes"
	"io/ioutil"
	"encoding/json"
	"log"
	"net"
	"strings"
	//"compress/zlib"
	//"errors"
)

//"github.com/astaxie/beego"

type Client struct {
	Addr string
}

var (
	ClientList []Client
	endString string
)

func init() {
	endString = "lepdendstring"
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
	defer conn.Close()

	_, err = conn.Write([]byte(sbody))
	if err != nil {
		return nil, err
	}

	message, err := ioutil.ReadAll(conn)//buf.Bytes()
	if err != nil {
		return nil, err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal(message, &ret)
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

	p := strings.Index(result, endString)
	if p <= 0 {
		return ""
	}
	result = result[0: p - 1]

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
	p := strings.Index(result, endString)
	if p <= 0 {
		return nil
	}

	result = result[0: p - 1]
	response_lines := strings.Split(result, "\n")

	if len(response_lines) == 0 {
		return nil
	}

	return response_lines
}
