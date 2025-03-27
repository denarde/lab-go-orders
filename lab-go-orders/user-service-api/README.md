# User Service API

This is a simple user management service with user registration and data persistence using a PostgreSQL database. The API implements secure password authentication using bcrypt hashing and follows good security practices.

## Features

- **User Registration** (`POST /register`)
  - Receives name, email, and password.
  - Validates the complexity of the password (minimum of 8 characters, including at least one uppercase letter, one lowercase letter, and one number).
  - Hashes the password with bcrypt before persisting it in the database.
  - Returns the user's data (without the password) after registration.

## Technologies

- **Go** (Golang)
- **PostgreSQL** for data persistence
- **GORM** (Go ORM) for database interaction
- **bcrypt** for password hashing
- **Chi** for routing

## Setup

### Prerequisites

- **Go** (version 1.18+)
- **PostgreSQL** (or use the provided Docker Compose setup)

### Environment Variables

Make sure to set the following environment variables for the database connection:

- `DB_HOST`: Host of the PostgreSQL database (default: `localhost`)
- `DB_PORT`: Port for PostgreSQL (default: `5432`)
- `DB_USER`: PostgreSQL username (default: `postgres`)
- `DB_PASSWORD`: PostgreSQL password (default: `password`)
- `DB_NAME`: PostgreSQL database name (default: `user_service_db`)

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/user-service-api.git
   cd user-service-api
