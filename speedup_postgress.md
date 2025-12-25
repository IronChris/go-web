# ♟️ FIDE Player Search Engine: Optimization Log
**Target Hardware:** Raspberry Pi 2B (ARMv7)
**Date:** December 25, 2025

---

## 🚀 The Stack
- **Backend:** Go (Golang) with `pgx/v5`.
- **Frontend:** HTMX (Low-overhead SSR approach).
- **Database:** PostgreSQL 17 with `pg_trgm` extensions.
- **Server Address:** `192.168.178.120`.

## ⚡ The 24x Speed Optimization
The major win was moving from a brute-force sequential scan to a high-efficiency Trigram index.

| Phase | Strategy | Execution Time | Experience |
| :--- | :--- | :--- | :--- |
| **Initial** | Parallel Sequential Scan | **1,191ms** | Sluggish |
| **Final** | **GIN Trigram Index** | **49ms** | **Instant** |

### The SQL "Secret Sauce"
```sql
-- Enable trigram matching
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Create the high-performance GIN index
CREATE INDEX idx_players_name_trgm ON fide_players USING gin ("Name" gin_trgm_ops);
