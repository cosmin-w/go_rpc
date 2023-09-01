package main

import (
	"log"
	"fmt"
	"net/rpc"
)

// use same payload as server receiver
type DataPayload struct {
	Data string
}
// rpc server host and port
const (
	HOST = "localhost"
	PORT = 9000
)


func callToRPC(){
	// init a client 
	client, err := rpc.Dial("tcp",fmt.Sprintf("%v:%d", HOST, PORT))
	if err != nil {
		log.Println(err)
	}
	// 
	d := DataPayload {
		Data:"client data",
	}

	var result string
	err = client.Call("RPCServer.GetData", d, &result)
	if err != nil {
		log.Printf("error client call:%+v", err)
	}

	log.Println(result)
}

func main(){
	callToRPC()
}