# FleetGuard

> A production-oriented real-time fleet monitoring and command platform built with Go, React, PostgreSQL, Redis, and WebSockets.

## 🚀 Overview

FleetGuard is a real-time fleet management platform designed to monitor and communicate with a large number of autonomous vehicles.

The project simulates autonomous vehicles so that the entire system can be developed and tested without requiring physical hardware.

The goal is not to build a simple CRUD application.

FleetGuard focuses on solving real-world backend engineering problems such as:

* Real-time communication
* Concurrent connections
* Vehicle health monitoring
* Failure detection
* Authentication and authorization
* Rate limiting
* Caching
* Idempotency
* Reliable command processing
* Graceful shutdown
* Observability
* Load testing
* Fault handling

---

## 🎯 Problem

Imagine a company operating hundreds or thousands of autonomous vehicles.

Each vehicle continuously sends information such as:

* Location
* Speed
* Battery level
* Temperature
* Connection status
* Sensor status
* Mission status

Operators need to monitor these vehicles in real time and, depending on their permissions, send commands to them.

The system must continue behaving correctly when:

* A vehicle disconnects
* Network connectivity becomes unstable
* Duplicate messages are received
* The database becomes unavailable
* Redis becomes unavailable
* A command receives no response
* Thousands of vehicles connect simultaneously
* The backend is restarted

FleetGuard is built to explore and solve these problems.

---

## 🏗️ High-Level Architecture

```text
                         ┌────────────────────┐
                         │   React Dashboard  │
                         └─────────┬──────────┘
                                   │
                              HTTP / WebSocket
                                   │
                         ┌─────────▼──────────┐
                         │     Go API         │
                         │                    │
                         │ Authentication     │
                         │ Authorization      │
                         │ Rate Limiting      │
                         │ Request Handling   │
                         └──────┬───────┬─────┘
                                │       │
                         ┌──────▼──┐ ┌──▼───────┐
                         │  Redis  │ │PostgreSQL│
                         └─────────┘ └──────────┘
                                │
                         ┌──────▼──────────┐
                         │ Go Workers / Hub │
                         └──────┬──────────┘
                                │
                ┌───────────────┼────────────────┐
                │               │                │
          ┌─────▼─────┐   ┌─────▼─────┐   ┌─────▼─────┐
          │ Vehicle   │   │ Vehicle   │   │ Vehicle   │
          │ Simulator │   │ Simulator │   │ Simulator │
          │   001     │   │   002     │   │   003     │
          └───────────┘   └───────────┘   └───────────┘
```

The vehicle simulators behave like real autonomous vehicles and allow us to test the backend under realistic conditions.

---

## 🧩 Core Features

### Vehicle Simulation

Simulate hundreds or thousands of vehicles generating realistic telemetry.

Example:

```json
{
  "vehicle_id": "vehicle-001",
  "latitude": 19.0760,
  "longitude": 72.8777,
  "speed": 12.4,
  "battery": 82.5,
  "temperature": 41.2
}
```

### Real-Time Monitoring

Operators can see:

* Online/offline vehicles
* Current location
* Battery
* Speed
* Temperature
* Active alerts
* Mission status

### Heartbeat Monitoring

Vehicles periodically send heartbeat messages.

If a vehicle stops communicating for a configured period, FleetGuard detects the failure and marks it offline.

### Command Processing

Authorized operators can send commands such as:

```text
START MISSION
STOP MISSION
PAUSE
RETURN TO BASE
```

Commands are processed reliably and tracked through their lifecycle.

### Authentication & Authorization

Different users have different permissions.

Example roles:

```text
ADMIN
OPERATOR
VIEWER
```

Authorization is enforced by the backend rather than relying on frontend UI restrictions.

### Rate Limiting

Protect APIs from excessive traffic using Redis-based rate limiting.

### Caching

Frequently accessed vehicle information can be cached using Redis to reduce database load.

### Idempotency

The system prevents duplicate processing of messages and commands where required.

### Observability

The system will eventually provide:

* Structured logs
* Request metrics
* Error metrics
* Request latency
* Active connections
* Worker queue information
* Vehicle health information

---

## 🛠️ Technology Stack

### Backend

* Go
* REST API
* WebSockets
* Goroutines
* Channels
* Context
* PostgreSQL
* Redis

### Frontend

* React
* TypeScript

### Infrastructure

* Docker
* Docker Compose
* GitHub Actions

### Observability

* Prometheus
* Grafana
* Structured logging

### Testing

* Go unit tests
* Integration tests
* API tests
* Load testing

---

## 📁 Project Structure

The project will evolve as features are added.

Initial structure:

```text
fleetguard/
│
├── cmd/
│   ├── server/
│   │   └── main.go
│   │
│   └── simulator/
│       └── main.go
│
├── internal/
│   └── vehicle/
│
├── go.mod
│
├── README.md
│
└── .gitignore
```

The structure will grow as the system becomes more complex.

---

## 🔄 Development Roadmap

### Phase 1 — Foundation

* [ ] Project setup
* [ ] Go HTTP server
* [ ] Health endpoint
* [ ] Vehicle heartbeat API
* [ ] Request validation
* [ ] Error handling
* [ ] Vehicle simulator

### Phase 2 — Persistence

* [ ] PostgreSQL integration
* [ ] Database migrations
* [ ] Repository layer
* [ ] Vehicle persistence
* [ ] Telemetry storage
* [ ] Indexing
* [ ] Connection pooling

### Phase 3 — Redis

* [ ] Redis integration
* [ ] Vehicle status caching
* [ ] Cache expiration
* [ ] Rate limiting
* [ ] Distributed coordination

### Phase 4 — Real-Time Communication

* [ ] WebSocket server
* [ ] Real-time vehicle updates
* [ ] Connection management
* [ ] Backpressure handling
* [ ] Reconnection handling

### Phase 5 — Security

* [ ] Authentication
* [ ] JWT
* [ ] Role-based authorization
* [ ] Resource-level authorization
* [ ] API validation
* [ ] Security logging

### Phase 6 — Command System

* [ ] Command API
* [ ] Command queue
* [ ] Worker pool
* [ ] Command acknowledgement
* [ ] Timeout handling
* [ ] Retry policy
* [ ] Idempotency

### Phase 7 — Reliability

* [ ] Context cancellation
* [ ] Graceful shutdown
* [ ] Timeouts
* [ ] Retry strategies
* [ ] Failure detection
* [ ] Dead-letter handling

### Phase 8 — Observability

* [ ] Structured logging
* [ ] Metrics
* [ ] Prometheus
* [ ] Grafana
* [ ] Health checks
* [ ] Readiness checks

### Phase 9 — Testing & Performance

* [ ] Unit tests
* [ ] Integration tests
* [ ] Race detection
* [ ] Load testing
* [ ] Concurrent vehicle simulation
* [ ] Performance profiling

### Phase 10 — Deployment

* [ ] Docker
* [ ] Docker Compose
* [ ] CI pipeline
* [ ] Automated tests
* [ ] Production configuration
* [ ] Deployment documentation

---

## 🧠 Engineering Problems We Intend to Solve

FleetGuard is specifically designed around production engineering challenges.

### Concurrency

How can Go efficiently handle thousands of simultaneous vehicle connections?

### Backpressure

What happens when vehicles produce telemetry faster than the system can process it?

### Reliability

What happens when a dependency becomes unavailable?

### Idempotency

How do we safely handle duplicate messages?

### Distributed State

How do multiple backend instances share vehicle state?

### Security

How do we prevent an authenticated user from accessing or controlling a vehicle they aren't authorized to access?

### Graceful Shutdown

How can the server restart without abruptly terminating important work?

### Observability

How do we know what is happening inside the system when something goes wrong?

### Performance

How does the system behave when the number of vehicles increases from 10 to 1,000 or 10,000?

---

## 🧪 Testing Philosophy

The system will not only be tested under normal conditions.

We will deliberately introduce failures.

Examples:

```text
Vehicle disconnects
        ↓
Does the system detect it?
```

```text
Redis goes down
        ↓
Does the API fail safely?
```

```text
Duplicate command
        ↓
Is it processed twice?
```

```text
10,000 vehicles connect
        ↓
Does the system remain stable?
```

```text
Server receives SIGTERM
        ↓
Does it shut down gracefully?
```

---

## 📊 Performance Goals

The exact targets will be established through benchmarking rather than assumed upfront.

We will measure:

* Requests per second
* API latency
* WebSocket connections
* Memory usage
* CPU usage
* Database query latency
* Redis latency
* Message processing rate
* Queue depth
* Error rate

Results will be documented as the project develops.

---

## 🎓 What This Project Demonstrates

By completing FleetGuard, the project should demonstrate practical knowledge of:

```text
Go
├── HTTP
├── REST
├── WebSockets
├── Goroutines
├── Channels
├── Mutexes
├── Context
├── Error handling
└── Graceful shutdown

Backend
├── API design
├── PostgreSQL
├── Redis
├── Authentication
├── Authorization
├── Caching
├── Rate limiting
├── Idempotency
└── Distributed systems

Production
├── Docker
├── Testing
├── Logging
├── Metrics
├── Monitoring
├── Load testing
└── CI/CD
```

---

## 📌 Project Status

**Status:** 🚧 Under active development

Current phase:

> **Phase 1 — Foundation**

The project is being built incrementally, with each feature introduced to solve a specific engineering problem.

---

## 📄 License

License will be added as the project develops.
