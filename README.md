# Notification Agent

Notification Agent is a Go-based microservice that processes notifications through Azure Service Bus and sends notifications through SMS and Email channels.

## Features

- Producer API to send notifications
- Consumer to receive notifications from Azure Service Bus
- Azure PostgreSQL integration
- gRPC APIs
- SMS notifications using Twilio
- Email notification support
- Health Check API
- Azure deployment support

---

## Prerequisites

- Go 1.24+
- Azure Subscription
- Azure Service Bus
- Azure PostgreSQL
- Twilio Account
- Git

---

## Project Structure

```
notification-agent/
├── cmd/
├── internal/
├── proto/
├── migrations/
├── main.go
├── producer.go
├── consumer.go
├── go.mod
├── Dockerfile
└── README.md
```

---

## Environment Variables

Create a `.env` file:

```env
PORT=8085

DB_HOST=<postgres-host>
DB_PORT=5432
DB_USER=<db-user>
DB_PASSWORD=<db-password>
DB_NAME=notificationdb

SERVICE_BUS_CONNECTION_STRING=<service-bus-connection-string>
QUEUE_NAME=notifications

TWILIO_ACCOUNT_SID=<twilio-sid>
TWILIO_AUTH_TOKEN=<twilio-token>
TWILIO_PHONE_NUMBER=<twilio-number>

GRPC_PORT=50051
```

---

## Installation

Clone repository:

```bash
git clone https://github.com/monikaab-ops/notification-agent.git
cd notification-agent
```

Install dependencies:

```bash
go mod tidy
```

---

## Run Application

Start application:

```bash
go run .
```

Expected output:

```text
Connected to Azure PostgreSQL
HTTP Server started on :8085
gRPC Server started on :50051
```

---

## Producer API

Send notification:

```bash
curl -X POST http://localhost:8085/send
```

Expected:

```text
Message Sent
```

---

## Consumer

Consumer receives messages from Azure Service Bus.

Example log:

```text
[CONSUMER] Message Received: Hello Monika
```

---

## Health Check

```bash
curl http://localhost:8085/health
```

Response:

```json
{
  "status": "healthy"
}
```

---

## Database Migration

Run migrations:

```bash
go run migrations/migrate.go
```

---

## Build

```bash
go build -o notification-agent
```

---

## Docker Build

```bash
docker build -t notification-agent:v1 .
```

---

## GitHub Repository

```text
https://github.com/monikaab-ops/notification-agent
```

---

## Azure Resources

- Azure Service Bus
- Azure PostgreSQL
- Azure Container Registry (ACR)
- Azure Container Instance (ACI)

---

## Azure Deployment

### Azure Resources

- Azure Container Registry (ACR)
  - notificationagentacr2026

- Azure Container Instance (ACI)
  - notification-agent

- Azure PostgreSQL
  - psqlmonika

- Azure Service Bus
  - sbmonika

### Verify Deployment

```bash
az container show \
  --resource-group rg-notification-agent-poc \
  --name notification-agent \
  --query "{State:instanceView.state,IP:ipAddress.ip}"