package main

import (
	"../gen-go/hellothrift"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"context"
	"os"
)

func main() {
	var transport thrift.TTransport
	args := os.Args
	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, _ = thrift.NewTSocket(args[1])
	//transportFactory := thrift.NewTBufferedTransportFactory(8192)
	//transport, _ = transportFactory.GetTransport(transport)
	//transportFactory := thrift.NewTJSONProtocolFactory()
	//defer transport.Close()
	transport.Open()
	inputProtocol:= thrift.NewTJSONProtocolFactory().GetProtocol(transport)
	outputProtocol:= thrift.NewTJSONProtocolFactory().GetProtocol(transport)
	tStandardClient := thrift.NewTStandardClient(inputProtocol,outputProtocol)
	client := hellothrift.NewHelloThriftClient(tStandardClient)
	str, _ := client.SayHello(context.Background(), "Go-Client")
	fmt.Println(str)
}