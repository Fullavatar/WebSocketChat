# Real-Time Chat Server

A real-time chat server developed in Go, with a client in Godot. This project serves as a learning exercise for Go, Godot, and MySQL.

## 📌 Features

- WebSocket communication
- User authentication (coming soon)
- Private messages and chat rooms (coming soon)
- Message storage in a database (coming soon)
- Client hosting via GitHub Releases (coming soon)

## 🛠 Installation

### 1️⃣ Clone the repository
```bash
git clone https://github.com/Fullavatar/WebSocketChat.git
cd WebSocketChat
```

### 2️⃣ Run the server

Compile and run the Go server:
```bash
go build -o Builds/Server.exe./cmd/main.go
./Builds/Server.exe
```

### 3️⃣ Configure the client (Godot)
The client's Godot Project is available in the `Client/` folder. You need Godot 4.4 to open and export it.
The default export folder is `Builds/` -> `Client.exe`


## 🔜 Upcoming Improvements

- [ ] User authentication
- [ ] Room management
- [ ] Database write optimization with batching
- [ ] Notifications and avatars
- [ ] Transition to HTTPS/WSS
- [ ] Code signing to avoid security warnings

## 📜 License

This project is licensed under MIT. Contributions are welcome! 🚀
