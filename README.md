# nutanix-pc-go-sdk

Go Client SDK for various Nutanix APIs. Currently, only Karbon is fully supported but additional APIs are being added.

This client API SDK for Go is based upon the Google, Github Go client SDK found here: https://github.com/google/go-github

## Current APIs and status:

### Karbon - fully supported.

Documentation available [KARBON](karbon/README.md)

### Prism Central (v3 API) - partial support

Documentation available [PrismCentral](pc/README.md)

### Prism Element (v2 API) - limited support

Documentation available [PrismElement](pe/README.md)

### Calm - limited support

Documentation available [Calm](calm/README.md)

### Era - no support currently

### Examples

A basic example is included below but additional examples are available in the [examples folder](examples/)

The following example will both print out the KubeConfig for the cluster name identified but also write it to a file for use.

```go
func main() {
	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	con, err := nutanix.NewClient(httpClient, &nutanix.Config{Karbon: new(karbon.ServiceConfig)})
	if err != nil {
		panic(err)
	}

	cKubeRequest := &karbon.ClusterGetKubeConfigRequest{ClusterName: "cluster-name-01"}

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
```

### Configuration

It is NOT RECOMMENDED to hardcode connection and authentication information directly in the code. While this can be done, environmental variables are recommended. The environmental variable list used for the API are below:

#### Environmental Variables

Setting environmental variables is optional but recommended over hard-coded values.

###### Karbon API

- NUTANIX_KARBON_USER = Nutanix Karbon username
- NUTANIX_KARBON_PASS = Nutanix Karbon password
- NUTANIX_KARBON_URL = Nutanix Karbon URL (e.g. `export NUTANIX_KARBON_URL="https://10.0.0.10:9440/karbon/"`)

###### Prism Central API

- NUTANIX_PC_USER = Nutanix Prism Central username
- NUTANIX_PC_PASS = Nutanix Prism Central password
- NUTANIX_PC_URL = Nutanix Prism Central URL (e.g. `export NUTANIX_PC_URL="https://10.0.0.10:9440/api/nutanix/v3/"`)

###### Prism Element API

- NUTANIX_PE_USER = Nutanix Prism Element username
- NUTANIX_PE_PASS = Nutanix Prism Element password
- NUTANIX_PE_URL = Nutanix Prism Element URL (e.g. `export NUTANIX_PC_URL="https://10.0.0.11:9440/api/nutanix/v2/"`)

###### Calm API

#### Manually setting credentials and configurations

Not recommended! You may end up checking your credentials into a source repository. Use this for testing only!

#### Initialization

We initialize our service clients by providing struct configs for each service we wish to use. Only services that have a struct config will be initialized. In the examples below we will work through some examples.

###### Initialize just the Karbon API

###### Initialize both Karbon and Prism Central v3 API
