package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	Desc       string   `xml:"desc,attr"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
	TestDesc   string   `xml:",innerxml"`
}

func main() {

	file, err := os.Open("D:/git_projects/beegoApp/src/gray-revel/resources/sampleXml")
	if err != nil {
		fmt.Printf("%s", err.Error())
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("%s", err.Error())
		panic(err)
	}
	fmt.Println(v.Svs[0].ServerName)
	fmt.Println(v.Svs[1].ServerName)

}