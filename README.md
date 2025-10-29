# 🗨️ Simple Chatroom (Assignment 04)

A basic **client-server chatroom** built in **Go** using **Remote Procedure Call (RPC)**.  
This project demonstrates how distributed systems communicate through remote procedure calls, where multiple clients can send and retrieve chat messages via a central coordinating server.

---

🎥 Demo

👉 Watch Demo Video: https://drive.google.com/file/d/1py8NU3LiEEz-GJJfGyLVaC_QF5XLUkMy/view?usp=sharing

----

## 📋 Description

- The **server** stores all messages sent by clients in memory.
- The **client** connects to the server via RPC and:
  - Sends messages remotely to the server.
  - Fetches the complete chat history from the server after each message.
- The client runs indefinitely until you type `exit` or press `Ctrl + C`.
- The system handles connection errors gracefully if the server goes offline.

---

## ⚙️ Technologies Used

- **Language:** Go (Golang)
- **Packages:** `net`, `net/rpc`, `fmt`, `sync`, `bufio`, `os`
- **Concepts:** Distributed Systems, RPC, Client-Server Communication

---

## 🚀 How to Run

### 1. Initialize the Go Module  
If not already done:  
```bash
go mod init Ass4
go mod tidy
