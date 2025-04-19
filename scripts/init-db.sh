#!/bin/bash
set -e

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
until PGPASSWORD=axz123 psql -h localhost -U axz -d axz -p 5433 -c '\q' 2>/dev/null; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done

echo "PostgreSQL is up - executing database initialization"

# Run the initialization SQL
PGPASSWORD=axz123 psql -h localhost -U axz -d axz -p 5433 -f deploy/docker/init.sql

echo "Database initialization completed" 