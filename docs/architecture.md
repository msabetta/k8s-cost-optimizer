# Architecture

## 🏗 Overview

K8s Cost Optimizer è composto da più moduli:

- Collector → recupera metriche da Prometheus
- Cost Engine → calcola costi
- Optimizer → suggerisce miglioramenti
- Forecasting → predice costi futuri
- API → espone dati
- Frontend → visualizza dashboard

---

## 📊 High-Level Architecture

+----------------------+
| Kubernetes Cluster   |
| (Pods / Nodes)       |
+----------+-----------+
						|
						v
+----------------------+
| Prometheus           |
| (metrics collection) |
+----------+-----------+
						|
						v
+--------------------------+
| Collector (Go)           |
| - Fetch metrics          |
| - Parse data             |
+----------+---------------+
						|
						v
+--------------------------+
| Cost Engine              |
| - CPU cost               |
| - Memory cost            |
+----------+---------------+
						|
						v
+--------------------------+
| Forecasting              |
| - Time series            |
| - Anomaly detection      |
+----------+---------------+
						|
						v
+--------------------------+
| API Layer               |
| (REST endpoints)        |
+----------+---------------+
						|
						v
+--------------------------+
| Frontend (Next.js)      |
+--------------------------+


---

## 🔄 Data Flow

```

Prometheus → Collector → Cost Engine → Forecast → API → UI

```

---

## 🧩 Component Breakdown

### Collector
- Query Prometheus (PromQL)
- Parse metrics
- Normalize data

### Cost Engine
- CPU pricing
- Memory pricing
- Aggregation per container

### Optimizer
- Idle detection
- Overprovision detection

### Forecasting
- Trend analysis
- Cost prediction
- Anomaly detection

---

## ⚙️ Scalabilità

- Stateless API → scalabile orizzontalmente
- Prometheus → central metrics store
- Possibile integrazione con:
  - Kafka (streaming)
  - Redis (caching)

---

## 🔮 Evoluzioni future

- Multi-cluster support
- Multi-cloud pricing
- AI-based optimization
