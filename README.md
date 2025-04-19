# Axe Platform

Simple and Reliable API Gateway

## Overview
Axe is a modern API Gateway platform designed for cloud-native environments. It provides dynamic routing, service discovery, and traffic management capabilities through a simple yet powerful interface.

## MVP Components

### Control Plane
- Service and Route management API
- PostgreSQL for configuration storage
- Consul integration for service discovery
- Vault integration for secrets management
- RabbitMQ for configuration distribution

### Data Plane
- Dynamic HTTP routing
- Health checks
- Basic metrics (requests, latency, errors)
- Configuration hot-reload

### Management UI
- Service management interface
- Route configuration
- Basic monitoring dashboard

## Quick Start

### Prerequisites
- Docker and Docker Compose
- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Consul
- Vault
- RabbitMQ

### Development Setup
1. Clone the repository
```bash
git clone https://github.com/yourusername/axz.git
cd axz
```

2. Start infrastructure services
```bash
docker-compose -f deploy/docker/docker-compose.yaml up -d
```

3. Initialize the database
```bash
./scripts/init-db.sh
```

4. Start the Control Plane
```bash
cd api && go run cmd/api/main.go
```

5. Start the Data Plane
```bash
cd gateway && go run cmd/gateway/main.go
```

6. Start the Management UI
```bash
cd web && npm install && npm run dev
```

## Architecture
```ascii
┌─────────────┐     ┌─────────────┐
│  Management │     │   Control   │
│     UI      │────▶│    Plane    │
└─────────────┘     └─────────────┘
                           │
                           ▼
                    ┌─────────────┐
                    │  Data Plane │
                    │  (Gateway)  │
                    └─────────────┘
                           │
                           ▼
                    ┌─────────────┐
                    │  Services   │
                    └─────────────┘
```

## License
MIT
