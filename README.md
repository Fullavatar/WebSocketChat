# Real-Time Chat Server

A real-time chat server developed in Go, with a client in Godot. This project serves as a learning exercise for Go, Godot, and MySQL.

## 📌 Features

- WebSocket communication
- Hosting for both client and server
- User authentication (coming soon)
- Private messages and chat rooms (coming soon)
- Message storage in a database (coming soon)
- Client hosting via GitHub Releases (coming soon)

## 🛠 Installation

### 1️⃣ Clone the repository
```bash
git clone https://github.com/your-repo/chat-server.git
cd chat-server
```

### 2️⃣ Run the server

Compile and run the Go server:
```bash
go build -o chat-server ./cmd/main.go
./chat-server
```

### 3️⃣ Configure the client (Godot)
The client is available in the `client/` folder. It must be configured to connect to your WebSocket server address.

## 🌐 Deployment

To make your server accessible externally:
1. Use a reverse proxy (e.g., Nginx) or Cloudflare to mask your IP.
2. Host the client on GitHub Releases or a web server.

## 🔜 Upcoming Improvements

- [ ] User authentication
- [ ] Room management
- [ ] Database write optimization with batching
- [ ] Notifications and avatars
- [ ] Transition to HTTPS/WSS
- [ ] Code signing to avoid security warnings

## 📜 License

This project is licensed under MIT. Contributions are welcome! 🚀
