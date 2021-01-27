package main

import(
        "fmt"
        "net/http"
       	"crypto/tls"
        "crypto/x509"
	"io/ioutil"
        "log"
)


func main() {

	// create certificate with client crt and key so the server can verify this client
	clcert, err := tls.LoadX509KeyPair("client1.crt", "client1.key")
	if err != nil {
		log.Fatal(err)
	}

        // Add the selfca certificate to the certificate pool
	// Adding selfca to the trusted CAs helps to verify the certificate 
	// presented by the server (ie. to trust the server)
        // Read cert
        cacert, err := ioutil.ReadFile("selfca.crt")
        if err != nil {
                log.Fatal(err)
        }
        caCertPool := x509.NewCertPool()
        caCertPool.AppendCertsFromPEM(cacert)

	//create https client using the trust CA pool and client certificate
	httpsclient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
				Certificates: []tls.Certificate{clcert},
			},
		},
	}

	resp, err := httpsclient.Get("https://localhost:9443/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}

