# K8s Cost Optimizer

**K8s Cost Optimizer** è una soluzione completa per monitorare, analizzare e ottimizzare i costi dei cluster Kubernetes.
Fornisce metriche in tempo reale, raccomandazioni di ottimizzazione, forecast dei costi e visualizzazioni tramite API e frontend.

---

## 🏗 Architettura

```
Kubernetes Cluster
├── k8s-cost-optimizer (Go API)
│   ├── internal/collector     # fetch metriche Prometheus
│   ├── internal/cost          # calcolo costi e rightsizing
│   ├── internal/forecasting   # forecast e anomaly detection
│   ├── api/handlers           # API REST
│   └── pkg/                   # utilità, config, logger
├── Prometheus                 # raccolta metriche
├── Grafana (opzionale)        # dashboard
└── Frontend (Next.js/React)   # visualizzazione dashboard

```

## 🏗 Architettura del Sistema (ASCII Diagram)

Diagramma semplificato che mostra come i componenti interagiscono:

                     +-------------------------+ <br>
                     |   Kubernetes Cluster    | <br>
                     |  (Pods, Nodes, cAdvisor)| <br>
                     +-----------+-------------+ <br>
                                 | <br>
                                 v <br>
                      +----------v-----------+ <br>
                      |     Prometheus       | <br>
                      | (metrics store + API)| <br>
                      +----------+-----------+ <br>
                                 | <br>
                                 | <br>
                                 v <br>
      +----------------+    +--------------------+ <br>
      |  API Server    |    | Frontend (React)   | <br>
      |(Go / REST API) |<-->| (Dashboard UI)     | <br>
      +------+---------+    +--------------------+ <br>
             | <br>
             v <br>
   +--------------------------+ <br>
   | Cost & Forecast Engine   | <br>
   | (rightsizing, anomalies) | <br>
   +--------------------------+ <br>



*(Il diagramma ASCII aiuta a visualizzare componenti e flussi di dati in un README semplice senza immagini esterne.)* :contentReference[oaicite:0]{index=0}

---

## 🔄 Dettaglio del Flusso dei Dati

+------------++----------------+ +-----------------------------+ <br>
| Prometheus ||     Cost Algorithms    || Anomaly/Forecasting  | <br>
| (metrics)  |---> | (CPU/Memory cost) | ----> |     Engine    | <br>
+------------++----------------+ +-----------------------------+ <br>
                  ||  <br>
                  ||  <br>
                  VV  <br>
+-----------------++-------------------+ <br>
| Collector       || API Response JSON | <br>
| (parse & fetch) || (cost / forecast) | <br>
+-----------------++-------------------+ <br>

---

## 🚀 Come Funziona il Processo

1) **Raccolta metriche**
   Prometheus raccoglie metriche dai nodi e container (es. CPU, memory).

2) **Collector**
   Il collector Go interroga Prometheus usando query PromQL.

3) **Calcolo dei costi**
   Il modulo cost utilizza i prezzi del file `pricing.yaml` per calcolare costi per container.

4) **Forecast & Anomaly**
   L’engine di forecasting stima costi futuri e rileva anomalie.

5) **API REST**
   Le API (/costs, /forecast, /summary, /recommendations) espongono i risultati.

6) **Frontend**
   Interfaccia React/Next.js visualizza dashboard e insight.

---

## 🧠 Cos’è un Diagramma di Flusso

Un diagramma di flusso rappresenta una sequenza di passaggi o flussi all’interno di un processo, collegando elementi tramite frecce per descriverne la dipendenza sequenziale. :contentReference[oaicite:1]{index=1}

---

## 🏷 Componenti chiave
+------------------------+ <br>
| API Endpoints          | <br>
| - /health              | <br>
| - /costs               | <br>
| - /summary             | <br>
| - /forecast            | <br>
| - /recommendations     | <br>
+------------------------+ <br>
           | <br>
           v <br>
+------------------------+ <br>
| Internal Modules       | <br>
| - collector            | <br>
| - cost                 | <br>
| - forecasting          | <br>
| - optimizer            | <br>
+------------------------+ <br>


---

## ⚡ Caratteristiche principali

- Raccolta metriche CPU/Memory dai container
- Calcolo costi basato su prezzi configurabili
- Analisi rightsizing e rilevamento idle/overprovision
- Forecast dei costi e rilevamento anomalie
- API REST: `/health`, `/costs`, `/recommendations`, `/summary`, `/forecast`
- Frontend Next.js con visualizzazioni: costi, risparmi, raccomandazioni
- Supporto deploy via Helm e YAML plain

---

## 📦 Prerequisiti

- Go 1.21+
- Docker & Docker Compose
- Kubernetes cluster (opzionale per Helm)
- Node.js 18+ e npm/yarn (per frontend)

---

## 🚀 Esecuzione locale con Docker Compose

```bash
# Build e avvia tutti i servizi (API + Prometheus + Grafana)
docker compose up --build

# Test endpoint API
curl http://localhost:8080/health
curl http://localhost:8080/costs
````

---

## 🛠 Deploy su Kubernetes

### 1. Con Helm

```bash
# Lint chart
helm lint deployments/helm/k8s-cost-optimizer

# Deploy in namespace finops
helm install k8s-cost-optimizer deployments/helm/k8s-cost-optimizer --namespace finops --create-namespace

# Controllo pods
kubectl get pods -n finops
```

### 2. Con YAML plain

```bash
kubectl apply -f deployments/plain-yaml/configmap.yaml
kubectl apply -f deployments/plain-yaml/rbac.yaml
kubectl apply -f deployments/plain-yaml/deployment.yaml
kubectl apply -f deployments/plain-yaml/service.yaml
kubectl apply -f deployments/plain-yaml/ingress.yaml
```

---

## 📁 Struttura progetto

```
k8s-cost-optimizer/
├── api/                 # API REST
├── internal/
│   ├── collector/
│   ├── cost/
│   └── forecasting/
├── pkg/                 # utilità (config, logger, utils)
├── configs/             # config.yaml, pricing.yaml
├── deployments/
│   ├── helm/            # Helm chart completo
│   └── plain-yaml/      # manifest YAML
├── web/                 # frontend Next.js
└── main.go
```

---

## 🔧 Test

```bash
# Test unitari
go test ./tests/unit/...

# Test integrazione
go test ./tests/integration/...

# Test completo
go test ./...
```

---

## 🌐 Frontend

* React / Next.js
* Componenti principali:
  * `CostCard` → visualizza costi e risparmi
  * `SavingsChart` → grafico forecast dei costi
  * `RecommendationList` → lista raccomandazioni
* Accessibile via `http://localhost:3000` (Docker Compose)

---

## ⚙️ Configurazioni

* `configs/config.yaml` → server, collector, forecast
* `configs/pricing.yaml` → prezzi CPU / Memory
* Config map Kubernetes automatica con Helm o YAML

---

## 🔗 Link utili

* [Prometheus](https://prometheus.io/)
* [Grafana](https://grafana.com/)
* [Next.js](https://nextjs.org/)

---

## 📄 Licenza

MIT License
