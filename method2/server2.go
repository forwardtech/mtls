package main

import (
        "fmt"
        "net/http"
        "crypto/tls"
        "crypto/x509"
        "io/ioutil"
        "log"
)

func reqHandler(w http.ResponseWriter, req *http.Request) {
                fmt.Fprint( w, "Hello From Server 2\n" )
}

func main() {

        http.HandleFunc("/", reqHandler)


        // Add the certificate to the certificate pool
        // So you can verify the certificate presented by the client 
 // (i.e., to trust the client)

        cacert, err := ioutil.ReadFile("cert.crt")
        if err != nil {
                log.Fatal(err)
        }
        caCertPool := x509.NewCertPool()
        caCertPool.AppendCertsFromPEM(cacert)

        // tls config
        
        tlsConfig := &tls.Config{
                ClientCAs: caCertPool,
                ClientAuth: tls.RequireAndVerifyClientCert,
        }

        //http server setting using the tls config above
        httpserver := &http.Server{
                Addr:      ":9443",
                TLSConfig: tlsConfig,
        }

        // Send cert on request. So, the client can verify you
        log.Fatal( httpserver.ListenAndServeTLS("cert.crt", "cert.key") )
}

