package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serversslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	var s Serversslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"192.168.20.132"}]}`
	json.Unmarshal([]byte(str), &s)
	//fmt.Printf("%#v", s)
	fmt.Println(s.Servers[0].ServerName)
	fmt.Println(s.Servers[0].ServerIP)

}