package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/routebyintuition/ntnx-go-sdk/karbon"

	nutanix "github.com/routebyintuition/ntnx-go-sdk"
)

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
		fmt.Println("kubeconfig get error: ", err)
		return
	}

	fmt.Println(getRes.KubeConfig)

	err = ioutil.WriteFile("/tmp/kube.txt", []byte(getRes.KubeConfig), 0644)
	if err != nil {
		fmt.Println("error writing to file: ", err)
		os.Exit(1)
	}
}
