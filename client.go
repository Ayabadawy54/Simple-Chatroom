package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🗨️  Welcome to the Chatroom! Type 'exit' to quit.\n")

	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("👋 Exiting chatroom...")
			break
		}

		var reply []string
		err = client.Call("ChatServer.SendMessage", &struct{ Message string }{text}, &reply)
		if err != nil {
			fmt.Println("⚠️ Error sending message:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, msg := range reply {
			fmt.Println(msg)
		}
		fmt.Println("--------------------\n")
	}
}
