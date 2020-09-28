package main

import (
	"encoding/json"
	"fmt"

	nutanix "github.com/routebyintuition/ntnx-go-sdk"
	"github.com/routebyintuition/ntnx-go-sdk/pc"
)

func main() {
	fmt.Println("testing...")

	// pccon := &pc.ServiceConfig{URL: nutanix.String("https://10.1.1.10:9440/api/nutanix/v3/", User: nutanix.String("admin"), Pass: nutanix.String("password")}
	nconf := &nutanix.Config{
		PrismCentral:       new(pc.ServiceConfig),
		InsecureSkipVerify: true,
	}

	con, err := nutanix.NewClient(nil, nconf)
	if err != nil {
		fmt.Println("error on NewClient: ", err)
		return
	}

	cListRequest := new(pc.ClusterListRequest)
	fmt.Println("performing new list request")
	listRes, _, err := con.PC.Cluster.List(cListRequest)
	if err != nil {
		fmt.Println("cluster list error: ", err)
		return
	}

	// fmt.Println("cluster list response: ", listResp)
	out, _ := json.MarshalIndent(listRes, "", "  ")
	fmt.Println("cluster list result: ", string(out))
}
