#!/bin/bash

echo "🚀 Starting K8s Cost Optimizer..."

# check Go
if ! command -v go &> /dev/null
then
    echo "❌ Go not installed"
    exit 1
fi

# check Prometheus
echo "🔍 Checking Prometheus..."
if curl -s http://localhost:9090/-/healthy | grep -q "Prometheus"; then
    echo "✅ Prometheus is running"
else
    echo "⚠️ Prometheus not detected at http://localhost:9090"
fi

# start API
echo "🔥 Starting API..."
go run cmd/api/main.go
