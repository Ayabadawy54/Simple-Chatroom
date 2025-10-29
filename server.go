package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// ChatServer struct holds the messages
type ChatServer struct {
	messages []string
	mu       sync.Mutex
}

// Args represents the message argument
type Args struct {
	Message string
}

// SendMessage stores the message and returns the chat history
func (cs *ChatServer) SendMessage(args *Args, reply *[]string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.messages = append(cs.messages, args.Message)
	*reply = cs.messages
	return nil
}

// GetHistory returns the entire chat history
func (cs *ChatServer) GetHistory(dummy *struct{}, reply *[]string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	*reply = cs.messages
	return nil
}

func main() {
	chatServer := new(ChatServer)

	err := rpc.Register(chatServer)
	if err != nil {
		log.Fatal("Error registering ChatServer:", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listener error:", err)
	}
	defer listener.Close()

	fmt.Println("🚀 Chat server is running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
