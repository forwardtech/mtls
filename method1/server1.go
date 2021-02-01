package main

import(
	"fmt"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func reqHandler(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint( w, "Hello From Server 1\n" )
}

func main() {

	http.HandleFunc("/", reqHandler)

        // Add the selfca certificate to the certificate pool
        // Adding selfca to the trusted CAs helps to verify the certificate 
        // presented by the client (i.e., to trust the client)
	cacert, err := ioutil.ReadFile("selfca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cacert)

        // tls config (adds the CA cert pool and 
        // the option to verify client cert)
	tlsConfig := &tls.Config{
    		ClientCAs: caCertPool,
    		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	//http server setting using the tls config above 
	httpserver := &http.Server{
    		Addr:      ":9443",
    		TLSConfig: tlsConfig,
	}

	// Listen 
	log.Fatal( httpserver.ListenAndServeTLS("server1.crt", "server1.key") )

}
