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

	// pccon := &pc.ServiceConfig{URL: nutanix.String("https://10.1.1.10:9440/api/nutanix/v3/")}
	// con, err := nutanix.NewClient(httpClient, &nutanix.Config{PrismCentral: pccon})
	con, err := nutanix.NewClient(httpClient, &nutanix.Config{PrismCentral: new(pc.ServiceConfig)})
	if err != nil {
		fmt.Println("error on NewClient: ", err)
	}

	cGetRequest := new(pc.ClusterGetRequest)
	cGetRequest.UUID = "0005ad73-a7b6-cb49-2e13-ac1f6b35bfb8"

	fmt.Println("performing new get request")

	getRes, _, err := con.PC.Cluster.Get(cGetRequest)
	if err != nil {
		fmt.Println("cluster list error: ", err)
		return
	}

	// fmt.Println("cluster list response: ", listResp)
	out, _ := json.MarshalIndent(getRes, "", "  ")
	fmt.Println("cluster list result: ", string(out))
}
