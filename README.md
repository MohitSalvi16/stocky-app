# 🏦 Stocky App — Golang Backend

---

## Overview

**Stocky** is a backend system built using **Golang (Gin + Logrus)** that simulates a platform where users can earn fractional shares of Indian stocks (e.g., *RELIANCE*, *TCS*, *INFY*) as rewards for completing milestones.

When a user is rewarded:
- The user receives **full stock units** (no deductions)
- Stocky pays for brokerage, taxes, and regulatory fees internally, recorded in a **ledger**
- Stock prices are periodically updated (mocked with random prices for this assignment)

---

## Tech Stack

| Component | Technology Used |
|------------|-----------------|
| **Language** | Go 1.22 |
| **Web Framework** | Gin |
| **Logging** | Logrus |
| **Database** | PostgreSQL (DB name: `assignment`) |
| **Environment Management** | godotenv |
| **Testing Tool** | Postman |

---

## Project Structure

```
stocky-assignment/
├── main.go
├── go.mod
├── .env
├── README.md
├── postman_collection.json
├── config/
│   └── config.go
├── db/
│   └── db.go
├── controllers/
│   ├── reward_controller.go
│   ├── portfolio_controller.go
│   └── price_service.go
├── routes/
│   └── routes.go
├── models/
│   ├── reward.go
│   └── ledger.go
```

---

## Database Schema

**Database Name:** `assignment`

### Table: `rewards`
| Column | Type | Description |
|---------|------|-------------|
| id | SERIAL PRIMARY KEY | Unique ID |
| user_id | INT | ID of user receiving reward |
| stock_symbol | VARCHAR(50) | Stock symbol (e.g., RELIANCE) |
| shares | NUMERIC(18,6) | Fractional shares allowed |
| timestamp | TIMESTAMP | Reward timestamp |

### Table: `ledger`
| Column | Type | Description |
|---------|------|-------------|
| id | SERIAL PRIMARY KEY | Ledger entry ID |
| user_id | INT | Related user |
| stock_symbol | VARCHAR(50) | Related stock |
| units | NUMERIC(18,6) | Number of shares |
| cash_outflow | NUMERIC(18,4) | INR cash spent by Stocky |
| fees | NUMERIC(18,4) | Brokerage, STT, GST, etc. |
| timestamp | TIMESTAMP | Record timestamp |

**Relationship:**
- `user_id` links `rewards` and `ledger`
- Aggregations and valuations derived using SQL `SUM()` + `GROUP BY`

---

## API Specifications

### **1. POST /reward**
Record a stock reward event for a user.

**Request:**
```json
{
  "user_id": 1,
  "stock": "RELIANCE",
  "shares": 2.5
}
```

**Response:**
```json
{
  "message": "Reward recorded successfully"
}
```

---

### **2. GET /today-stocks/{userId}**
Return all stock rewards for today for a user.

**Response:**
```json
[
  {
    "stock": "RELIANCE",
    "shares": 2.5,
    "timestamp": "2025-11-11T15:45:00Z"
  }
]
```

---

### **3. GET /stats/{userId}**
Returns total shares rewarded today (grouped by stock) and current INR portfolio value.

**Response:**
```json
{
  "total_today": [
    {"stock": "RELIANCE", "total_shares": 2.5, "inr_value": 6875.0}
  ],
  "portfolio_value": {"total_inr": 6875.0}
}
```

---

### **4. GET /historical-inr/{userId}**
Returns INR value of all stock rewards for past days (up to yesterday).

**Response:**
```json
{
  "historical_inr": {
    "2025-11-10": 9450.0,
    "2025-11-09": 8900.0
  }
}
```

---

### **5. GET /portfolio/{userId}**
Returns current holdings per stock symbol with INR value.

**Response:**
```json
{
  "portfolio": [
    {"stock": "RELIANCE", "shares": 2.5, "current_inr": 6850.0},
    {"stock": "TCS", "shares": 1.2, "current_inr": 4120.0}
  ],
  "total_inr": 10970.0
}
```

---

## Environment Configuration

`.env`
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=assignment
```

Loaded automatically using `godotenv`.

---

## Logging Example (Logrus)

Example log entries:
```
time="2025-11-11T17:00:05Z" level=info msg="Reward recorded successfully for user 1"
time="2025-11-11T17:00:07Z" level=info msg="Fetched portfolio for user 1"
```

---

## Edge Cases & Handling

| Case | Handling Strategy |
|------|--------------------|
| **Duplicate rewards** | Add unique constraint `(user_id, stock_symbol, timestamp)` |
| **Stock splits / mergers** | Isolated in `price_service.go` for future API integration |
| **Rounding errors** | NUMERIC(18,6) & NUMERIC(18,4) prevent precision loss |
| **Price API downtime** | Random price generator fallback |
| **Refunds / adjustments** | Can use negative shares |
| **Stale data** | Price refreshed hourly (mocked)

---

## Scaling & Architecture Notes

| Concern | Approach |
|----------|-----------|
| **Concurrency** | Gin supports concurrent requests natively |
| **Horizontal scaling** | Stateless API with PostgreSQL backend |
| **Maintainability** | Modular folder structure (controllers, routes, db, models) |
| **Logging** | Logrus ensures structured logs with timestamps |
| **Extensibility** | Mock price generator replaceable with live NSE/BSE API |

---

## Setup Instructions

### 1. Clone Repository
```bash
git clone https://github.com/<your-username>/stocky-app.git
cd stocky-app
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Run Application
```bash
go run main.go
```

**Expected Output:**
```
✅ Connected to PostgreSQL
INFO[0000] 🚀 Stocky server started on :8080
```

### 4. Test APIs via Postman
Import `postman_collection.json` → Run all five endpoints.

---

##  Database Creation Script

```sql
CREATE DATABASE assignment;

\c assignment;

CREATE TABLE rewards (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    stock_symbol VARCHAR(50) NOT NULL,
    shares NUMERIC(18,6) NOT NULL,
    timestamp TIMESTAMP DEFAULT NOW()
);

CREATE TABLE ledger (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    stock_symbol VARCHAR(50) NOT NULL,
    units NUMERIC(18,6),
    cash_outflow NUMERIC(18,4),
    fees NUMERIC(18,4),
    timestamp TIMESTAMP DEFAULT NOW()
);
```

**End of README**

