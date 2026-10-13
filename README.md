# Notification Agent

Notification Agent is a Go-based microservice for sending SMS and email notifications through Twilio and SendGrid. It also exposes health checks and gRPC endpoints for notification workflows.

## Features

- Producer and consumer notification flow
- Azure Service Bus integration
- PostgreSQL support
- SMS notification support using Twilio
- Email notification support using SendGrid
- gRPC and HTTP services
- Health check endpoint
- Worker-based processing for sending notifications

---

## Prerequisites

- Go 1.22+
- Azure Service Bus connection string
- Azure PostgreSQL database (if enabled)
- Twilio account
- SendGrid account
- Git

---

## Project Structure

```text
notification-agent/
├── cmd/
│   ├── grpcclient/
│   └── grpcserver/
├── internal/
│   ├── config/
│   ├── dto/
│   ├── grpc/
│   ├── handler/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   └── service/
├── migrations/
├── proto/
├── .env.example
├── .gitignore
├── dockerfile
├── go.mod
├── main.go
├── producer.go
├── README.md
└── client_temp.go.txt
```

---

## Environment Variables

Create a local `.env` file from the template:

```bash
copy .env.example .env
```

Example:

```env
PORT=8085
API_TOKEN=test111git
GRPC_TOKEN=your_grpc_token

DB_HOST=your_db_host
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=notificationdb

SERVICEBUS_CONNECTION=Endpoint=sb://your-service-bus;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=your_key
QUEUE_NAME=notification-events

SENDGRID_API_KEY=your_sendgrid_api_key
TWILIO_ACCOUNT_SID=your_twilio_account_sid
TWILIO_AUTH_TOKEN=your_twilio_auth_token
TWILIO_PHONE_NUMBER=+17372212163
SMS_TO=+919691129372
SMS_FROM=+17372212163

FROM_EMAIL=your_verified_sender@example.com
TO_EMAIL=recipient@example.com
```

Important:
- `.env` should not be committed
- Twilio trial accounts require the destination number to be verified
- SendGrid `FROM_EMAIL` must match a verified sender identity

---

## Installation

Clone repository:

```bash
git clone https://github.com/santipriya114/notifyagent.git
cd notifyagent
```

Install dependencies:

```bash
go mod tidy
```

---

## Run Application

```bash
go run .
```

Expected startup logs:

```text
[INFO] Worker Pool Started
[INFO] HTTP Server started on :8085
gRPC Server started on :50051
```

---

## Health Check

```bash
curl http://localhost:8085/health
```

Response:

```json
{"status":"UP"}
```

---

## Producer API

Send a notification request:

```bash
curl -X POST "http://localhost:8085/send?message=Hello%20santi"
```

---

## Consumer

The consumer reads messages from Azure Service Bus and routes them through the notification worker.

Example log output:

```text
[WORKER 1] Processing: Hello santi
```

---

## Database Migration

If migrations are needed:

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
https://github.com/santipriya114/notifyagent.git
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
  - notifyagentpgserverdev

- Azure Service Bus
  - notifyagentsbdev

### Verify Deployment

```bash
az container show \
  --resource-group rg-notification-agent-poc \
  --name notification-agent \
  --query "{State:instanceView.state,IP:ipAddress.ip}"
```

---

## Common Issues

### Twilio

- Error `572002` means the destination number is not verified for a Twilio trial account.
- Use E.164 format such as `+919691129372`.

### SendGrid

- `FROM_EMAIL` must be a verified sender identity in SendGrid.
- If using a dynamic template, ensure the template ID exists in your SendGrid account.

---

## Notes

This project was created for notification processing and is intended to be configured with real provider credentials in a local `.env` file before running in development or deployment environments.

