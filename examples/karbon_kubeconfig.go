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

	/*
		karbonConfig := &karbon.ServiceConfig{
			User: nutanix.String("username"),
			Pass: nutanix.String("password"),
			URL:  nutanix.String("https://0.0.0.0:9440/karbon"),
		}
		peConfig := &pe.ServiceConfig{
			User: nutanix.String("username"),
			Pass: nutanix.String("password"),
			URL:  nutanix.String("https://0.0.0.0:9440/api/nutanix/v2/"),
		}
		pcConfig := &pc.ServiceConfig{
			User: nutanix.String("username"),
			Pass: nutanix.String("password"),
			URL:  nutanix.String("https://0.0.0.0:9440/api/nutanix/v3/"),
		}

		ntnxConfig := &nutanix.Config{
			Karbon:       karbonConfig,
			PrismElement: peConfig,
			PrismCentral: pcConfig,
		}

		ntnx, err := nutanix.NewClient(httpClient, ntnxConfig)
		if err != nil {
			panic(err)
		}

		ntnxEnvConfig := &nutanix.Config{
			Karbon:       &karbon.ServiceConfig{},
			PrismCentral: &pc.ServiceConfig{},
			PrismElement: &pe.ServiceConfig{},
		}
		ntnx, err = nutanix.NewClient(httpClient, ntnxEnvConfig)
		if err != nil {
			panic(err)
		}
	*/
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
