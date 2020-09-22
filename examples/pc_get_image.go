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

	GetRequest := new(pc.ImageGetRequest)
	GetRequest.UUID = "1a2d3b2b-a253-4c0f-9283-eb2b36e553ae"

	getRes, listResp, err := con.PC.Image.Get(GetRequest)
	if err != nil {
		fmt.Println("cluster list error: ", err)
		return
	}

	fmt.Println("image list response: ", listResp)
	out, _ := json.MarshalIndent(getRes, "", "  ")
	fmt.Println("image list result: ", string(out))
}
