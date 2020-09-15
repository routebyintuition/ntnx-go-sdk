package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/routebyintuition/ntnx-go-sdk/karbon"

	nutanix "github.com/routebyintuition/ntnx-go-sdk"
)

type KubeConf struct {
	KubeConfig string `json:"kube_config"`
}

func main() {
	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	con, err := nutanix.NewClient(httpClient, &nutanix.Config{Karbon: new(karbon.ServiceConfig)})
	if err != nil {
		panic(err)
	}

	cKubeRequest := &karbon.ClusterGetKubeConfigRequest{ClusterName: "single01-CL"}

	getRes, _, err := con.Karbon.Cluster.GetKubeConfig(cKubeRequest)
	if err != nil {
		fmt.Println("kuneconfig get error: ", err)
		return
	}

	fmt.Println(getRes.KubeConfig)

	err = ioutil.WriteFile("/tmp/kube.txt", []byte(getRes.KubeConfig), 0644)
	if err != nil {
		fmt.Println("error writing to file: ", err)
	}
}
