package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	nutanix "github.com/routebyintuition/ntnx-go-sdk"
	"github.com/routebyintuition/ntnx-go-sdk/pc"
)

func main() {
	fmt.Println("testing prism central list subnets...")

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	pcConf := &pc.ServiceConfig{}

	con, err := nutanix.NewClient(httpClient, &nutanix.Config{PrismCentral: pcConf})
	if err != nil {
		fmt.Println("error on NewClient: ", err)
	}

	GetRequest := new(pc.SubnetListRequest)
	fmt.Println("get request: ", GetRequest)

	getRes, listResp, err := con.PC.Subnet.List(GetRequest)
	if err != nil {
		fmt.Println("subnet list err response: ", listResp)
		fmt.Println("subnet list error: ", err)
		return
	}

	fmt.Println("subnet list response: ", listResp)
	out, _ := json.MarshalIndent(getRes, "", "  ")
	fmt.Println("subnet list result: ", string(out))
}
