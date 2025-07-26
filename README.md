# 🛡️ grpc_demo_auth

A demo gRPC-based authentication service written in Go, implementing user **registration**, **login**, and **JWT token management** (access + refresh tokens).

---

## ✨ Features

- ✅ User registration (`Register`)
- ✅ User login (`Login`)  
- 🔐 JWT-based authentication
- 🔄 Access & refresh token generation
- 🧠 Clear project structure with separation of logic into layers: service, repository, domain

---

## 🔑 JWT Implementation

JWT token generation, parsing, and validation logic is located in:


The service issues both **access tokens** (short-lived) and **refresh tokens** (longer-lived) to support secure and scalable authentication.

---

## 🚀 Getting Started

```bash
# Run the server
go run app/cmd/server/main.go

# Run the client
go run app/cmd/client/main.go
