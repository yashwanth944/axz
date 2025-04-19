#!/bin/bash
set -e

# Create necessary directories
mkdir -p api/cmd/api
mkdir -p api/internal/{config,handlers,middleware,models,store}
mkdir -p api/pkg/client

mkdir -p gateway/cmd/gateway
mkdir -p gateway/internal/{config,proxy,metrics}
mkdir -p gateway/pkg/client

mkdir -p web/{src,public}
mkdir -p deploy/{docker,k8s/{control-plane,data-plane},terraform}
mkdir -p docs
mkdir -p scripts

# Make scripts executable
chmod +x scripts/*.sh

# Initialize Go modules
cd api && go mod init github.com/axz/api && go mod tidy
cd ../gateway && go mod init github.com/axz/gateway && go mod tidy

# Initialize Next.js project
cd ../web && npx create-next-app@latest . \
  --typescript \
  --tailwind \
  --eslint \
  --app \
  --src-dir \
  --import-alias "@/*"

echo "Development environment setup completed" 