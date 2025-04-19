# 🚀 API Gateway

A modern API Gateway implementation in Go, providing authentication, authorization, and request routing capabilities.

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14-blue.svg)](https://www.postgresql.org)
[![JWT](https://img.shields.io/badge/JWT-Auth-orange.svg)](https://jwt.io)

</div>

## 📋 Table of Contents
- [Architecture](#-architecture)
- [Features](#-features)
- [Current Status](#-current-status)
- [Roadmap](#-roadmap)
- [Getting Started](#-getting-started)
- [API Documentation](#-api-documentation)
- [Development](#-development)

## 🏗 Architecture

```mermaid
graph TD
    A[Client] --> B[API Gateway]
    B --> C[Authentication Service]
    B --> D[Authorization Service]
    B --> E[Request Router]
    C --> F[PostgreSQL]
    D --> F
    E --> F
    
    subgraph "Core Services"
        C
        D
        E
    end
    
    subgraph "Data Layer"
        F
    end
```

### Component Overview

```mermaid
graph LR
    A[API Gateway] --> B[JWT Auth]
    A --> C[User Management]
    A --> D[Request Routing]
    B --> E[Token Generation]
    B --> F[Token Validation]
    C --> G[User CRUD]
    C --> H[Password Hashing]
    D --> I[Route Protection]
    D --> J[Request Forwarding]
```

## ✨ Features

- 🔐 JWT-based authentication
- 👥 User management with PostgreSQL
- 🛡 Protected routes with middleware
- 📊 Request logging and monitoring
- 🔄 Request routing and forwarding
- 🧪 Debug endpoints for testing

## 🚧 Current Status

### What's Working
- Basic user registration
- JWT token generation
- PostgreSQL integration
- Debug endpoints

### Known Issues
- Database connection issues with certain credentials
- JWT token validation errors
- Password hashing inconsistencies

## 🗺 Roadmap

### Week 1: Core Stability
- [ ] Fix database connection issues
- [ ] Resolve JWT token validation
- [ ] Standardize password hashing
- [ ] Add comprehensive error handling

### Week 2: Enhanced Security
- [ ] Implement rate limiting
- [ ] Add request validation
- [ ] Set up secure headers
- [ ] Add API key authentication

### Week 3: Monitoring & Logging
- [ ] Add request logging
- [ ] Implement metrics collection
- [ ] Set up health checks
- [ ] Add performance monitoring

### Week 4: Advanced Features
- [ ] Add request caching
- [ ] Implement circuit breakers
- [ ] Add request transformation
- [ ] Set up API documentation

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or later
- PostgreSQL 14 or later
- Make (optional)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/apigateway.git
cd apigateway
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Start the server:
```bash
go run main.go
```

## 📚 API Documentation

### Authentication Endpoints

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant DB
    
    Client->>API: POST /register
    API->>DB: Create User
    DB-->>API: User Created
    API-->>Client: Success
    
    Client->>API: POST /login
    API->>DB: Verify Credentials
    DB-->>API: User Verified
    API-->>Client: JWT Token
```

### Protected Endpoints

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant DB
    
    Client->>API: GET /api/profile
    API->>API: Validate JWT
    API->>DB: Fetch User Data
    DB-->>API: User Data
    API-->>Client: Profile Data
```

## 🛠 Development

### Project Structure
```
.
├── auth/           # Authentication logic
├── db/            # Database operations
├── models/        # Data models
├── main.go        # Entry point
└── README.md      # Documentation
```

### Running Tests
```bash
go test ./...
```

### Code Style
- Follow Go standard formatting
- Use meaningful variable names
- Add comments for complex logic
- Write tests for new features

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
Made with ❤️ by [Your Name]
</div> 