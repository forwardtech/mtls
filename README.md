# mTLS
mTLS implementation in Go
-------------------------

Provides sample implementation of mTLS using two methods.

## What is mTLS
While the familiar TLS protocol is used to authenticate a Server, mTLS or Mutual TLS is used to authenticate both the Server and the Client before exchanging information. As in TLS it does provide secure delivery of the data by encrypting it during transmission. The primary use case for mTLS is to authenticate the client as an additional assurance - to make sure the server doesn’t serve unidentified clients. 

To be clear, client authentication is NOT user authentication which is an application-level authentication performed by the application (think User ID and Password). In mTLS the client application itself is authenticated. For example, a mobile app you want to trust before exchanging information or say two cloud workloads want to trust each other before communicating. You can look at it as an additional factor in the authentication process. In a 2-factor authentication terminology, you can consider this as “something you have”. mTLS brings us one step closer to the concept of Zero Trust Networking[1].

Let’s implement one to understand this further. We will use Go for our example here with two different methods to implement mTLS.

Before we start let us cover some basics quickly:

## Certificate Authority (CA)

CA is a 3rd party organization that validates and provides a digital certificate to an entity. Digital certificate primarily provides authentication and helps in encrypted communication between a server and a client.

## Self-signed CA or Root CA

Instead of using a 3rd party CA you could create a custom CA to use within your organization to sign certificates. You will get your own CA which you can use to sign certificates within your organization. Like the regular CA certificate, this is pre-installed on the clients (and servers in case of mTLS) to trust certificates signed by this CA. This is shown in Method I below.

## Self-signed Certificate

Instead of using a 3rd party CA or custom CA you can just create a certificate, sign it with your private key and preinstall this on the applications that need to trust this certificate. This is implemented in Method II below.

This option is obviously simpler, but you must install this on all the applications each time a new certificate is generated, whereas in the first method once you have the CA certificate installed across the organization any certificate signed but this CA will be accepted across the organization - without having to re-install the new certificate. The self-signed certificate method is used for demo systems or at best used within a small team or organization.

mTLS brings us one step closer to Zero Trust Networking

## Certificate generation process

The entity which needs a certificate would start the process by generating a private key[2] and then creating a certificate signing request (CSR) using the private key and the organization details. This CSR is sent to the CA which validates the organization and generates a certificate and sends its back to the organization.

The organization then configures its servers to provide this certificate on request.

## Certificate verification process

When a client (say a browser) initiates a connection to a server, it receives a certificate. The client then looks for the CA which signed the certificate and checks if it is one of the CAs it trusts and uses the CA’s public key to decrypt the certificate, which proves it was in fact signed by the trusted CA.

With this background let’s get into the implementation example.

### METHOD I:

Here we create a custom CA certificate which both Server and Client can trust

#### Instructions

(Note: cd to the method1 directory before doing the following)

```Shell
# Generate a CA Certificate for signing
openssl req -new -x509 -nodes -days 365 -keyout selfca.key -out selfca.crt -subj '/CN=selfCA'
```

##### Server-side steps

```Shell
# Step1: Generate Private Key for the Server (Key Pair generation is clearer)
openssl genrsa -out server1.key 2048

# Step 2: Generate CSR  
openssl req -new -key server1.key -subj '/CN=localhost' -out server1.csr

#Step 3: Self signing (instead of sending the CSR it to an external CA). 
#This should generate a certificate (.crt) file which is the signed certificate for the server.
openssl x509 -req -in server1.csr -CA selfca.crt -CAkey selfca.key -CAcreateserial -out server1.crt -days 365
```

##### Client-side steps 
(The certification generation process is the same as the server)


```Shell
#Step 1: Generate Private Key for the Client
openssl genrsa -out client1.key 2048

#Step 2: Generate the CSR 
openssl req  -new -key client1.key -subj '/CN=localhost' -out client1.csr

#Step 3: Self signing (instead of sending the CSR it to an external CA). 
#This should generate a certificate (.crt) file which is the signed certificate for the client.
openssl x509 -req -in client1.csr -CA selfca.crt -CAkey selfca.key -CAcreateserial -out client1.crt -days 365
```
##### Test it
```Shell
go run server1.go
go run client1.go
```

### METHOD II:

#### Instructions
(Note: cd to the method2 directory before doing the following)

In this method you generate a certificate and add it both to the client and server's certificate pool before communicating. This is simpler but it’s not a typical implementation and it has its drawbacks. Method I is closer to what happens in practice.

```Shell
#Generate a certificate and a private key (single step process)
openssl req -x509 -sha256 -nodes -newkey rsa:2048 -keyout cert.key -out cert.crt -subj '/CN=localhost' -days 365
```

##### Test it
```Shell
go run server2.go
go run client2.go
```

