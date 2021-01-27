
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

        // Send the certificate so the server can verify you
        clcert, err := tls.LoadX509KeyPair("cert.crt", "cert.key")
        if err != nil {
                log.Fatal(err)
        }

        // Add the certificate to the certificate pool
        // So you can verify the certificate presented by the server 
 // (i.e., to trust the server)
        cacert, err := ioutil.ReadFile("cert.crt")
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

