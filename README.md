
#  Go Network Monitor (WebSocket-based)

A lightweight real-time server status monitor built with **Go** and **WebSocket**.

This project demonstrates how to use Go’s `net/http` and `gorilla/websocket` to implement a basic real-time dashboard that checks the online/offline status of multiple servers and streams the result to the frontend using WebSocket.

---

##  Features

- Real-time server status updates (online/offline)
- WebSocket-based communication between server and browser
- Minimal and clean architecture
- Easy to extend (authentication, persistence, UI enhancements)

---

##  Tech Stack

- **Backend:** Go (Golang), [gorilla/websocket](https://github.com/gorilla/websocket)
- **Frontend:** Plain HTML + JavaScript (WebSocket API)
- **Protocol:** WebSocket over TCP

---

##  How to Run

1. Clone the repository:

```bash
git clone https://github.com/RezaHaddad29/network-monitor.git
cd network-monitor
```

2. Run the server:

```bash
go run .
```

3. Open test.html in your browser to view the real-time status.

