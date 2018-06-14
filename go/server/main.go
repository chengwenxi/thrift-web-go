package main

import (
	"context"
	"../gen-go/hellothrift"
	"git.apache.org/thrift.git/lib/go/thrift"

	"net/http"
	"github.com/rs/cors"
)

const (
	NET_WORK_ADDR = "localhost:9092"
)

type HelloServiceTmpl struct {
}

func (this *HelloServiceTmpl) SayHello(ctx context.Context, str string) (s string, err error) {
	print("Go-Server:" + str)
	return str, nil
}

type MyHandler struct {
}

func (self MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var bodySlc = make([]byte, 1024)

	bodyLen, readErr := req.Body.Read(bodySlc)
	print(bodyLen)
	print(readErr)
	print("input %s", string(bodySlc))
	out := thriftRequest(bodySlc)
	println(string(out))
	w.WriteHeader(200)
	w.Write(out)
}

func thriftRequest(input []byte) []byte {
	//inbuffer := thrift.TMemoryBuffer{b,len(input)}
	//transport := thrift.NewTTransportFactory().GetTransport
	var inprotocol *thrift.TJSONProtocol
	var outprotocol *thrift.TJSONProtocol
	println("34444")
	var inbuffer thrift.TTransport
	inbuffer = thrift.NewTMemoryBuffer()
	inbuffer.Write(input)
	if inbuffer != nil {
		defer inbuffer.Close()
	}

	var outbuffer thrift.TTransport
	outbuffer = thrift.NewTMemoryBuffer()
	if outbuffer != nil {
		defer outbuffer.Close()
	}

	inprotocol = thrift.NewTJSONProtocol(inbuffer)
	outprotocol = thrift.NewTJSONProtocol(outbuffer)

	handler1 := &HelloServiceTmpl{}
	processor := hellothrift.NewHelloThriftProcessor(handler1)
	processor.Process(nil, inprotocol, outprotocol)

	out := make([]byte, outbuffer.RemainingBytes())
	outbuffer.Read(out)
	return out

}

func main() {

	handler1 := cors.Default().Handler(MyHandler{})
	http.ListenAndServe(":9092", handler1)

	//transportFactory := thrift.NewTTransportFactory()
	//protocolFactory := thrift.NewTJSONProtocolFactory()
	//transportFactory = thrift.NewTBufferedTransportFactory(8192)
	//transport, _ := thrift.NewTServerSocket(NET_WORK_ADDR)
	////transport1, _ := thrift.NewTHttpClient(NET_WORK_ADDR)
	//handler := &HelloServiceTmpl{}
	//processor := hellothrift.NewHelloThriftProcessor(handler)
	//server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	//server.AcceptLoop()
	//server.Serve()
}
