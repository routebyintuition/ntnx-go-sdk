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
	fmt.Println("testing prism central get one cluster...")

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	con, err := nutanix.NewClient(httpClient, &nutanix.Config{PrismCentral: new(pc.ServiceConfig)})
	if err != nil {
		fmt.Println("error on NewClient: ", err)
	}

	hGetRequest := new(pc.HostGetRequest)
	hGetRequest.UUID = "878c5684-dbb5-4306-a94f-0849244ea4fb"

	getRes, listResp, err := con.PC.Host.Get(hGetRequest)
	if err != nil {
		fmt.Println("cluster list error: ", err)
		return
	}

	fmt.Println("cluster list response: ", listResp)
	out, _ := json.MarshalIndent(getRes, "", "  ")
	fmt.Println("cluster list result: ", string(out))
}
