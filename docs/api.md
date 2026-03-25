```markdown
# API Documentation

## Base URL

```
[http://localhost:8080](http://localhost:8080)

````

---

## 📌 Endpoints

### 1. Health Check

**GET** `/health`

Verifica lo stato del servizio.

#### Response

```json
{
  "status": "ok"
}
````

---

### 2. Get Costs

**GET** `/costs`

Restituisce i costi calcolati per container.

#### Response

```json
[
  {
    "container": "nginx",
    "cpu_cost": 0.02,
    "memory_cost": 0.01,
    "total_cost": 0.03
  }
]
```

---

### 3. Get Recommendations

**GET** `/recommendations`

Restituisce suggerimenti per ottimizzazione costi.

#### Response

```json
[
  {
    "action": "Reduce CPU requests",
    "estimated_savings": 12.5
  }
]
```

---

### 4. Summary

**GET** `/summary`

Riepilogo generale dei costi.

#### Response

```json
{
  "total_cost": 120.5,
  "total_savings": 30.2
}
```

---

### 5. Forecast

**GET** `/forecast`

Previsioni dei costi.

#### Response

```json
{
  "hourly_cost": 5.2,
  "daily_cost": 124.8,
  "hourly_costs": [5.0, 5.1, 5.2]
}
```

---

## ⚠️ Error Handling

Tutte le API restituiscono:

```json
{
  "error": "description"
}
```

Codici HTTP:

* 200 → OK
* 400 → Bad Request
* 500 → Internal Server Error

---

## 🔐 Future Improvements

* Autenticazione JWT
* Rate limiting
* API versioning (/v1/)

````

---

# 2️⃣ `docs/architecture.md`

```markdown


