package main

import (
	"fmt"
	"github.com/roccobarbi/autorecon/pkg/network"
	"os"
	"strings"
)

type crt struct {
	IssuerCaId     int    `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	Id             int64  `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore      string `json:"not_before"`
	NotAfter       string `json:"not_after"`
	SerialNumber   string `json:"serial_number"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments. Usage:")
		fmt.Println("get_crt_sh <domain_with_crt.sh_syntax>")
		os.Exit(0)
	}
	baseUrl := "https://crt.sh/"
	//query := map[string]string{"q": os.Args[1], "output": "json"}
	//req := network.JsonGetRequest{BaseUrl: baseUrl, Query: query}
	//var req network.GetRequest = &network.JsonGetRequest{BaseUrl: baseUrl}
	var req network.GetRequest[crt] = &network.JsonGetRequest[crt]{}
	req.SetBaseUrl(baseUrl)
	req.SetQueryKeyValue("q", os.Args[1])
	req.SetQueryKeyValue("output", "json")
	fmt.Println("Requesting...")
	var crtArray []crt
	crtArray = req.Request()
	domain := make(map[string]int)
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Printf("crt.sh entries: %d\n", len(crtArray))
	for _, element := range crtArray {
		for _, nameValue := range strings.Split(element.NameValue, "\n") {
			domain[nameValue] = 1
		}
	}
	fmt.Printf("%d unique domains found:\n", len(domain))
	fmt.Println("-------------------------------------------------------------------------")
	for key := range domain {
		fmt.Println(key)
	}
}
