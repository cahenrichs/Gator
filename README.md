## 🧰 Requirements

Before running this program, you must have the following installed:

### 1. Go
The project is written in Go, so you'll need Go installed (version 1.20+ recommended).  
Download Go here: https://go.dev/dl/

### 2. PostgreSQL
The application uses PostgreSQL as its database.  
Install Postgres from: https://www.postgresql.org/download/

Ensure PostgreSQL is running and you know your database URL (e.g., `postgres://user:pass@localhost:5432/dbname`).

---

## 📦 Installing the gator CLI

You can install the CLI using `go install`:
 
```bash
go install github.com/cahenrichs/gator@latest