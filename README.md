# Go REST API with PostgreSQL

A simple Go REST API application with PostgreSQL database for user authentication (registration and login).

## Features

- User Registration
- User Login with password hashing (bcrypt)
- PostgreSQL database integration
- Environment variable configuration
- Database migrations
- JSON request/response handling
- Input validation

## Prerequisites

- Go 1.24.4 or higher
- PostgreSQL (or Docker for local development)
- Git

## Project Structure

```
├── auth/           # Password hashing utilities
├── cmd/            # Application entry points
│   ├── api/        # API server setup
│   ├── main.go     # Main application
│   └── migrate/    # Database migration tool
├── config/         # Environment configuration
├── db/             # Database connection
├── services/       # Business logic services
│   └── users/      # User service (routes & store)
├── types/          # Type definitions
└── utils/          # Utility functions
```

## API Endpoints

### POST /register
Register a new user.

**Request Body:**
```json
{
  "username": "your_username",
  "email": "your_email@example.com",
  "password": "your_password"
}
```

**Response (Success):**
```json
{
  "register": true
}
```

**Response (Error):**
```json
{
  "error": "user already exists"
}
```

### POST /login
Login with existing credentials.

**Request Body:**
```json
{
  "email": "your_email@example.com",
  "password": "your_password"
}
```

**Response (Success):**
```json
{
  "login": "success"
}
```

**Response (Error):**
```json
{
  "error": "wrong password"
}
```

## Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/AnshSinghSonkhia/go-rest-api-postgresql.git
cd go-rest-api-postgresql
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Set up PostgreSQL Database

#### Option A: Using Docker (Recommended for local development)
```bash
# Start PostgreSQL container
docker run --name postgres-go-api \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=go_rest_api_db \
  -p 5432:5432 \
  -d postgres:13

# Verify container is running
docker ps
```

#### Option B: Using local PostgreSQL installation
1. Install PostgreSQL on your system
2. Create a database named `go_rest_api_db`
3. Update the `.env` file with your credentials

### 4. Configure Environment Variables
Create a `.env` file in the root directory:
```bash
# PostgreSQL Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=go_rest_api_db
```

### 5. Run Database Migrations
```bash
# Apply migrations to create tables
go run cmd/migrate/main.go up

# To rollback migrations (if needed)
go run cmd/migrate/main.go down
```

### 6. Build and Run the Application
```bash
# Build the application
go build -o bin/restapi cmd/main.go

# Run the application
./bin/restapi
```

Alternative using npm scripts:
```bash
npm run build
npm run run
```

The API server will start on `http://localhost:8080`

## Testing the API

### Using curl

#### Register a new user:
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "testpass123"
  }'
```

#### Login with credentials:
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "testpass123"
  }'
```

### Using Postman

1. **Set up Postman environment:**
   - Base URL: `http://localhost:8080`

2. **Register User Request:**
   - Method: `POST`
   - URL: `{{base_url}}/register`
   - Headers: `Content-Type: application/json`
   - Body (raw JSON):
     ```json
     {
       "username": "testuser",
       "email": "test@example.com",
       "password": "testpass123"
     }
     ```

3. **Login User Request:**
   - Method: `POST`
   - URL: `{{base_url}}/login`
   - Headers: `Content-Type: application/json`
   - Body (raw JSON):
     ```json
     {
       "email": "test@example.com",
       "password": "testpass123"
     }
     ```

## Database Schema

### Users Table
```sql
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  createdAT TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)
```

## Technologies Used

- **Backend Framework:** Go with Gorilla Mux
- **Database:** PostgreSQL
- **Password Hashing:** bcrypt
- **Environment Management:** godotenv
- **Database Migrations:** golang-migrate
- **Validation:** go-playground/validator

## Development

### Available Make Commands
```bash
make build      # Build the application
make run        # Build and run the application
make migrate-up # Apply database migrations
make migrate-down # Rollback database migrations
```

### Available npm Scripts
```bash
npm run build        # Build the application
npm run run          # Build and run the application
npm run migrate-up   # Apply database migrations
npm run migrate-down # Rollback database migrations
```

## Troubleshooting

### Common Issues

1. **Database connection failed:**
   - Ensure PostgreSQL is running
   - Check your `.env` file configuration
   - Verify database exists and credentials are correct

2. **Migration errors:**
   - Ensure database is accessible
   - Check if migrations have already been applied
   - Verify migration files exist in `cmd/migrate/migrations/`

3. **API not responding:**
   - Check if port 8080 is available
   - Ensure the application started without errors
   - Verify PostgreSQL connection is established

### Using pgAdmin4

1. **Install pgAdmin4** from [pgAdmin website](https://www.pgadmin.org/download/)

2. **Connect to PostgreSQL:**
   - Host: `localhost`
   - Port: `5432`
   - Database: `go_rest_api_db`
   - Username: `postgres`
   - Password: `password`

3. **View Tables:**
   - Navigate to: Servers → PostgreSQL → Databases → go_rest_api_db → Schemas → public → Tables
   - You should see the `users` table with the registered user data

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

