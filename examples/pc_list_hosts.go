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
	fmt.Println("testing prism central list all hosts...")

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	con, err := nutanix.NewClient(httpClient, &nutanix.Config{PrismCentral: new(pc.ServiceConfig)})
	if err != nil {
		fmt.Println("error on NewClient: ", err)
	}

	// GetRequest := &pc.HostListRequest{}
	GetRequest := new(pc.HostListRequest)
	fmt.Println("get request: ", GetRequest)

	getRes, listResp, err := con.PC.Host.List(GetRequest)
	if err != nil {
		fmt.Println("cluster list error: ", err)
		return
	}

	fmt.Println("cluster list response: ", listResp)
	out, _ := json.MarshalIndent(getRes, "", "  ")
	fmt.Println("cluster list result: ", string(out))
}
