package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/routebyintuition/ntnx-go-sdk/karbon"

	nutanix "github.com/routebyintuition/ntnx-go-sdk"
)

func main() {
	fmt.Println("testing karbon list registries...")

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	// pccon := &karbon.ServiceConfig{URL: nutanix.String("https://10.1.1.10:9440/karbon/")}
	con, err := nutanix.NewClient(httpClient, &nutanix.Config{Karbon: new(karbon.ServiceConfig)})
	if err != nil {
		fmt.Println("error on NewClient: ", err)
		os.Exit(1)
	}

	cListRequest := new(karbon.ClusterListRegistriesRequest)
	cListRequest.ClusterName = "K8-PHX-POC159-F"

	fmt.Println("performing new get request")

	getRes, _, err := con.Karbon.Cluster.ListRegistries(cListRequest)
	if err != nil {
		fmt.Println("cluster list error: ", err)
		return
	}

	fmt.Println("cluster list response: ", getRes)
	out, _ := json.MarshalIndent(getRes, "", "  ")
	fmt.Println("cluster list result: ", string(out))
}
