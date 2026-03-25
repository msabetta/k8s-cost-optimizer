#!/bin/bash

echo "🧪 Testing environment..."

# kubectl check
if command -v kubectl &> /dev/null; then
    echo "🔍 Checking Kubernetes cluster..."
    if kubectl cluster-info &> /dev/null; then
        echo "✅ Kubernetes cluster reachable"
    else
        echo "❌ Kubernetes cluster NOT reachable"
    fi
else
    echo "⚠️ kubectl not installed"
fi

# Prometheus check
echo "🔍 Checking Prometheus..."
if curl -s http://localhost:9090/-/healthy | grep -q "Prometheus"; then
    echo "✅ Prometheus is healthy"
else
    echo "❌ Prometheus not reachable"
fi

# Test query
echo "🔍 Testing Prometheus query..."
curl -s -G http://localhost:9090/api/v1/query \
  --data-urlencode 'query=up' | grep -q "result" && \
  echo "✅ Query works" || echo "❌ Query failed"

echo "✅ Test completed"
