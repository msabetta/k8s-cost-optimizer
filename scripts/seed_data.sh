#!/bin/bash

echo "🌱 Seeding fake metrics..."

cat <<EOF > /tmp/fake_metrics.json
[
  {
    "name": "api-service",
    "namespace": "default",
    "pod": "api-pod",
    "cpu_usage": 0.1,
    "cpu_request": 0.5,
    "memory_usage_mb": 120,
    "memory_request_mb": 512
  },
  {
    "name": "worker",
    "namespace": "default",
    "pod": "worker-pod",
    "cpu_usage": 0.05,
    "cpu_request": 0.6,
    "memory_usage_mb": 80,
    "memory_request_mb": 512
  },
  {
    "name": "idle-service",
    "namespace": "default",
    "pod": "idle-pod",
    "cpu_usage": 0.01,
    "cpu_request": 0.3,
    "memory_usage_mb": 20,
    "memory_request_mb": 256
  }
]
EOF

echo "✅ Fake metrics written to /tmp/fake_metrics.json"
echo "👉 Ora puoi usarle nel collector come fallback"
