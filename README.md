# IDRM Go Common

[中文文档](README_ZH.md)

## Overview

`idrm-go-common` is a comprehensive Go common library for the **Intelligent Data Resource Management (IDRM)** system developed by kweaver-ai. It provides shared components, API definitions, utilities, and service integrations for building data governance and management applications.

## Features

- **Audit Logging**: Comprehensive audit trail system with distributed logging support
- **REST Clients**: Pre-built HTTP clients for microservice integration
- **Middleware**: HTTP middleware for authentication, authorization, and logging
- **Workflow Engine**: Event-driven workflow processing with message queue support
- **Session Management**: Redis-based session storage and lifecycle management
- **Database Callbacks**: Change data capture and data lineage tracking
- **Error Handling**: Centralized error code management across services

## Requirements

- Go 1.24 or higher

## Installation

```bash
go get github.com/kweaver-ai/idrm-go-common
```

## Project Structure

```
idrm-go-common/
├── api/                  # Protocol Buffer and API definitions
│   ├── auth-service/     # Authentication & authorization APIs
│   ├── audit/            # Audit event definitions
│   ├── configuration-center/  # Configuration management APIs
│   ├── data-view/        # Data view APIs
│   └── task_center/      # Task and workflow APIs
├── audit/                # Audit logging implementation
├── callback/             # gRPC callback clients
├── d_session/            # Distributed session management
├── database_callback/    # Database change tracking
├── errorcode/            # Error code definitions
├── interception/         # Token and user interception
├── middleware/           # HTTP middleware
├── reconcile/            # Reconciliation utilities
├── rest/                 # REST client implementations
├── trace/                # OpenTelemetry tracing
├── util/                 # Common utilities
└── workflow/             # Workflow engine
```

## Usage

### Audit Logging

```go
import "github.com/kweaver-ai/idrm-go-common/audit"

// Create audit context
ctx := audit.NewContext().
    WithOperator("user-id").
    WithOperation("create").
    WithResource("data-source")

// Log audit event
audit.Log(ctx, "Data source created")
```

### REST Client

```go
import "github.com/kweaver-ai/idrm-go-common/rest/auth-service/v1"

// Create auth service client
client := authservice.NewAuthClient(&authservice.Config{
    BaseURL: "http://localhost:8080",
})

// Check authorization
resp, err := client.CheckPermission(ctx, &authservice.CheckPermissionRequest{
    Subject: "user:123",
    Action:  "read",
    Object:  "resource:456",
})
```

### Middleware

```go
import "github.com/kweaver-ai/idrm-go-common/middleware/v1"

// Add authentication middleware
router.Use(middleware.TokenInterception())

// Add audit logging middleware
router.Use(middleware.AuditLogger())

// Add access control middleware
router.Use(middleware.AccessControl())
```

## Module Documentation

### API Definitions (`api/`)

Protocol Buffer definitions and Go types for service integration:

- **auth-service**: Authentication policies and authorization requests
- **audit**: Event types, operators, and resources
- **configuration-center**: User, role, and permission management
- **data-view**: Data view and sub-view definitions
- **task_center**: Work order and task management

### REST Clients (`rest/`)

HTTP client implementations for service integration:

- **auth-service**: Authorization and policy enforcement
- **anyrobot**: Data model and unified query clients
- **configuration-center**: Configuration and user management
- **task_center**: Task and workflow operations
- **data_catalog**: Data catalog integration

### Audit (`audit/`)

Comprehensive audit logging system:

- Agent detection (web, mobile)
- Event logging with source tracking
- Kafka integration for distributed logging
- Structured audit events

### Middleware (`middleware/`)

HTTP middleware for request processing:

- Token interception and validation
- User context management
- Permission checking
- Audit logging
- Error handling

### Workflow (`workflow/`)

Event-driven workflow processing:

- Message queue abstraction (Kafka, NSQ)
- Workflow definition and execution
- Event consumption and handling

## Dependencies

Key dependencies include:

- `github.com/gin-gonic/gin` - HTTP framework
- `github.com/google/uuid` - UUID generation
- `github.com/redis/go-redis/v9` - Redis client
- `github.com/IBM/sarama` - Kafka client
- `github.com/nsqio/go-nsq` - NSQ client
- `go.opentelemetry.io/otel` - OpenTelemetry tracing
- `go.uber.org/zap` - Structured logging
- `gorm.io/gorm` - ORM library
- `google.golang.org/grpc` - gRPC framework

## Development

### Running Tests

```bash
go test ./...
```

### Running Tests with Coverage

```bash
go test -cover ./...
```

### Generating Protocol Buffers

```bash
# Generate Go code from .proto files
make proto
```

## License

Copyright © kweaver-ai

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
