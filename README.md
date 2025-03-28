# Monorepo for User Service, Order Service, and Notification Service

This repository contains a monorepo with multiple microservices for a complete system architecture. The services included in this project are:

- **User Service** (User registration, login, profile management)
- **Order Service** (Order management, event publishing)
- **Notification Service** (Event-driven notifications)

The architecture leverages key technologies like **Go**, **gRPC**, **REST APIs**, **GraphQL**, **PostgreSQL**, **Redis**, and **Kafka** for building scalable and efficient microservices.

## Services Overview

### 1. User Service (Usuários)

- **Description**: Provides APIs for user registration, login, and profile management.
- **Key Features**:
  - REST API for user registration, login, and profile management.
  - JWT-based authentication and authorization for secure access.
  - User data stored in PostgreSQL.
  - Cache user session and data in Redis for fast access.
  - Exposes gRPC endpoints for other services to query user data.
  
- **Technologies**:
  - Go
  - PostgreSQL
  - Redis
  - gRPC
  - JWT for authentication

### 2. Order Service (Pedidos)

- **Description**: Manages customer orders with a REST API for order creation and GraphQL API for querying orders.
- **Key Features**:
  - REST API to create and manage orders.
  - GraphQL API for querying orders.
  - Order data stored in PostgreSQL.
  - Cache recent orders in Redis for quick access.
  - Publishes events to Kafka when a new order is created.

- **Technologies**:
  - Go
  - PostgreSQL
  - Redis
  - GraphQL
  - Kafka for event publishing

### 3. Notification Service (Notificações)

- **Description**: Listens to events from Kafka and sends notifications based on order creation.
- **Key Features**:
  - Consumes events from Kafka related to order creation.
  - Uses gRPC to query User Service for user details.
  - Sends notifications via email, SMS, or other channels.

- **Technologies**:
  - Go
  - Kafka for event consumption
  - gRPC to communicate with User Service
  - Email/SMS services for notifications