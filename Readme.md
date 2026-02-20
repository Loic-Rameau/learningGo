
## 🗓 Week 1 — Language & Runtime Mastery

### 🎯 Goal:

Stop writing “PHP in Go”.

### Focus Areas:

* Pointers & value semantics
* Zero values
* Interfaces (implicit implementation)
* Error wrapping
* Package structure
* Escape analysis
* Memory basics

---

### 📌 Exercises

#### 1️⃣ Build a CLI Task Manager (Production Style)

Not toy-level. Include:

* Clean package boundaries
* File persistence (JSON)
* Manual dependency injection
* Unit tests
* No globals
* Context usage where meaningful

You are training discipline.

---

#### 2️⃣ Deep Concurrency Drill

Build:

* Worker pool
* Bounded concurrency
* Context cancellation
* Panic recovery inside workers
* Timeout handling
* Benchmark throughput

Then answer:

* When does memory escape to heap?
* How do goroutine leaks happen?
* What happens if channel is never read?
* When to use buffered vs unbuffered channels?

If you can answer those confidently, you’re progressing correctly.

---

## 🗓 Week 2 — HTTP & Systems Thinking

### 🎯 Goal:

Understand how Go services actually run.

No frameworks.

Use `net/http`.

---

### 📌 Build a Small REST API

* CRUD endpoints
* In-memory store with mutex
* Middleware chain
* Structured logging
* Graceful shutdown
* Request context timeouts

Then refactor:

* Replace memory store with PostgreSQL
* Use `database/sql`
* Add migrations
* Handle transaction rollback properly

---

### You Must Understand:

* Connection pooling
* Blocking behavior
* Context propagation
* Error classification

At this point, you should feel uncomfortable. That’s good.

---

## 🗓 Week 3 — Production Concerns

### 🎯 Goal:

Think like infrastructure engineers at companies like Cloudflare or Uber.

---

### Add to Your API:

* Redis caching
* Token bucket rate limiter
* Health checks (liveness vs readiness)
* Dockerfile (multi-stage)
* docker-compose
* Environment config
* Structured logs (zap/zerolog)
* Makefile
* Proper folder structure

Now you’re building deployable software.

---

## 🗓 Week 4 — Advanced Go

### 🎯 Goal:

Concurrency + reliability mastery.

---

### Build:

#### 1️⃣ Worker Queue System

* API receives job
* Stores in DB
* Worker pool processes
* Retry with exponential backoff
* Dead letter logic
* Metrics endpoint

#### 2️⃣ Implement a Circuit Breaker

* Failure threshold
* Half-open state
* Timeout tracking
