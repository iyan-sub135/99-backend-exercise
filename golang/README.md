# Go Microservices Implementation

## Tech Stack
- Go 1.25+
- Gin (HTTP framework)
- GORM (ORM)
- SQLite (Database)
- MVC Architecture

## Architecture

Simple MVC pattern dengan 3 microservices:

- **listing-service**: Manage property listings
- **user-service**: Manage users
- **public-api**: API gateway for aggregate data

*For a more advanced API gateway implementation, check out my project: [go-getway](https://github.com/i-sub135/go-getway)*

## Run Services

```bash
cd golang

# Install dependencies first
make deps

# Run each service in separate terminal
make run-listing    # Port 6000
make run-user      # Port 7001
make run-public    # Port 8000
```

## Test

```bash
# Create user
curl -X POST http://localhost:8000/public-api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe"}'

# Create listing  
curl -X POST http://localhost:8000/public-api/listings \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "listing_type": "rent", "price": 5000}'

# Get listings with user data
curl http://localhost:8000/public-api/listings
```
