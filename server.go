package main
import (
	"log"
	"net"
	"fmt"
	"net/rpc"
)
// rpc server host and port
const (
	HOST = "localhost"
	PORT = 9000
)
type RPCServer int

type DataPayload struct {
	Data string
}

// rpc endpoint
func(r *RPCServer)GetData(payload *DataPayload, reply *string) error{
	log.Printf("%+v", *payload)
	*reply = "hi client here is your data"
	return nil
}

// magic happens here
func run(){
	log.Printf("Starting up RPC server.")

	listen, err:= net.Listen("tcp", fmt.Sprintf("%v:%d", HOST, PORT))
	if err != err{
		log.Println(err)
	}
	defer listen.Close()

	// accept connection
	for {
		rpcConnect, err := listen.Accept()
		if err !=nil{
			log.Println(err)
		}
		// register rpc server
		err = rpc.Register(new(RPCServer))
		if err !=nil{
			log.Println(err)
		}
		go rpc.ServeConn(rpcConnect)
	}
}


func main(){
	run()
}