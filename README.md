# mTLS
mTLS implementation in Go

Provides sample implementation of mTLS using two methods.

METHOD I:

Here we create a custom CA certificate which both Server and Client can trust
Instructions:

o Generate the certificates using the commands in method1/setup.txt

o Run method1/server1.go and method1/client1.go to test.

METHOD II:

In this method you generate a certificate and add it both to the client and server's certificate pool before communicating. This is simpler but it’s not a typical implementation and it has its drawbacks. Method I is closer to what happens in practice.

o Generate the certificates using the commands in method2/setup.txt

o Run method2/server2.go and method2/client2.go to test.

