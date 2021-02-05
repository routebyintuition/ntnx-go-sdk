package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	nutanix "github.com/routebyintuition/ntnx-go-sdk"
	"github.com/routebyintuition/ntnx-go-sdk/pe"
)

func main() {
	fmt.Println("testing prism element list disks...")

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	peConf := &pe.ServiceConfig{}

	con, err := nutanix.NewClient(httpClient, &nutanix.Config{PrismElement: peConf})
	if err != nil {
		fmt.Println("error on NewClient: ", err)
	}

	GetRequest := new(pe.DiskListRequest)
	GetRequest.Count = 1
	fmt.Println("get request: ", GetRequest)

	// test
	getRes, listResp, err := con.PE.Disk.List(GetRequest)
	if err != nil {
		fmt.Println("disk list err response: ", listResp)
		fmt.Println("disk list error: ", err)
		return
	}

	fmt.Println("disk list response: ", listResp)
	out, _ := json.MarshalIndent(getRes, "", "  ")
	fmt.Println("disk list result: ", string(out))
}
