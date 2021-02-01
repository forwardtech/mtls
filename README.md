# mTLS
mTLS implementation in Go
-------------------------

Provides sample implementation of mTLS using two methods.

## What is mTLS
While the familiar TLS protocol is used to authenticate a Server, mTLS or Mutual TLS is used to authenticate both the Server and the Client before exchanging information. As in TLS it does provide secure delivery of the data by encrypting it during transmission. The primary use case for mTLS is to authenticate the client as an additional assurance - to make sure the server doesn’t serve unidentified clients. 

### METHOD I:

#Here we create a custom CA certificate which both Server and Client can trust
##Instructions

```Shell
# Generate a CA Certificate for signing
openssl req -new -x509 -nodes -days 365 -keyout selfca.key -out selfca.crt -subj '/CN=selfCA'
```

#### Server-side steps

```Shell
# Step1: Generate Private Key for the Server (Key Pair generation is clearer)
openssl genrsa -out server1.key 2048

# Step 2: Generate CSR  
openssl req -new -key server1.key -subj '/CN=localhost' -out server1.csr

#Step 3: Self signing (instead of sending the CSR it to an external CA). 
#This should generate a certificate (.crt) file which is the signed certificate for the server.
openssl x509 -req -in server1.csr -CA selfca.crt -CAkey selfca.key -CAcreateserial -out server1.crt -days 365
```

#### Client-side steps 
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
#### Test it
```Shell
go run server1.go
go run client1.go
```

### METHOD II:

In this method you generate a certificate and add it both to the client and server's certificate pool before communicating. This is simpler but it’s not a typical implementation and it has its drawbacks. Method I is closer to what happens in practice.

```Shell
#Generate a certificate and a private key (single step process)
openssl req -x509 -sha256 -nodes -newkey rsa:2048 -keyout cert.key -out cert.crt -subj '/CN=localhost' -days 365
```

#### Test it
```Shell
go run server2.go
go run client2.go
```

