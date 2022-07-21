package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	downloadDir := flag.String("download-dir", "", "Directory to unpack certificates into.")
	region := flag.String("region", "", "Cloud resource region. (e.g. 'us-east-1')")
	secretName := flag.String("secret-name", "", "Cloud resource name.")
	flag.Parse()

	private_key, certificate_crt, ca_bundle_crt, err := TLSCerts(*secretName, *region)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s//%s", *downloadDir, "private.key"), []byte(private_key), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s//%s", *downloadDir, "ca_bundle.crt"), []byte(ca_bundle_crt), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s//%s", *downloadDir, "certificate.crt"), []byte(certificate_crt), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

}
