
# Microservices Web Crawler

This project implements a **microservices-based web crawler** with JWT-based authentication and multiple services responsible for crawling, data storage, and API gateway functionality. Each service is containerized using Docker and can be scaled independently.

## Services Overview

### 1. **Authentication Service**:
- Handles user login and generates JWT tokens for secure authentication.

### 2. **Crawl Management Service**:
- Manages the crawling process, handles URL fetching, parsing, and concurrency.

### 3. **Data Storage Service**:
- Stores crawled data in a database or persistent storage.

### 4. **API Gateway**:
- Routes requests to appropriate microservices and handles JWT token validation.

## Features

- JWT Authentication with token-based access control.
- URL Filtering to avoid revisiting already crawled links.
- Concurrency for faster crawling using goroutines and channels.
- Data storage in a local or cloud database.
- Microservices architecture for scalability and flexibility.
- Dockerized for easy deployment and scaling.

## Prerequisites

- Docker
- Docker Compose
- Go (for development and building individual services)

## Getting Started

Follow the steps below to set up and run the project locally.

### 1. Clone the Repository

Clone the repository to your local machine.

```bash
git clone https://github.com/iampawans/microservices-crawler.git
cd microservices-crawler
```

### 2. Set Up Environment Variables (Optional)

If necessary, configure environment variables or create a `config.yaml` file in the `/config` folder to manage secret keys or database configurations.

Example `config.yaml`:
```yaml
jwtSecret: "your-secret-key"
```

### 3. Build and Run the Services with Docker Compose

Make sure Docker and Docker Compose are installed, then build and start the services:

```bash
docker-compose build
docker-compose up
```

This will start all the microservices: `auth-service`, `crawl-service`, `data-service`, and `api-gateway`.

### 4. Accessing the Services

Once the services are running, you can interact with them using the following endpoints:

- **Authentication Service (Login)**:  
  `POST http://localhost:8000/login`  
  This endpoint accepts a JSON body with `username` and `password`. On successful login, it returns a JWT token.

  Example request:
  ```json
  {
    "username": "admin",
    "password": "password"
  }
  ```

  Example response:
  ```json
  {
    "token": "your-jwt-token"
  }
  ```

- **API Gateway (Start Crawl)**:  
  `GET http://localhost:8082/start-crawl`  
  This endpoint requires a valid JWT token in the `Authorization` header as a `Bearer token`. It triggers the crawl process.

- **Crawl Management Service (Start Crawl)**:  
  `GET http://localhost:8080/start-crawl`  
  This service starts crawling the URLs in the list. It will be triggered via the API Gateway.

- **Data Storage Service (Store Data)**:  
  `POST http://localhost:8081/store`  
  This service stores crawled data in the database or other persistent storage.

  Example request:
  ```json
  {
    "url": "https://example.com",
    "timestamp": "2024-10-10T10:00:00Z"
  }
  ```

### 5. Stopping the Services

To stop the services, run the following command:

```bash
docker-compose down
```

This will stop and remove the containers.

## Service Communication

The services communicate via HTTP requests, and the API Gateway ensures that requests are routed to the correct service. The **JWT token** issued by the **Authentication Service** must be included in the header of each request to secure the system.

### JWT Authentication

JWT tokens are used for authentication and authorization. When the user logs in via the Authentication Service, a JWT token is returned. This token should be included in the `Authorization` header in subsequent requests to protected services.

Example:
```bash
curl -H "Authorization: Bearer your-jwt-token" http://localhost:8082/start-crawl
```

### Service Dependencies

- **API Gateway** depends on the **Authentication Service**, **Crawl Management Service**, and **Data Storage Service**.
- **Crawl Management Service** communicates with the **Data Storage Service** to store crawled data.

## Folder Structure

```
/microservices-crawler
│
├── /auth-service
│   ├── Dockerfile
│   ├── main.go                # Authentication service code
│   ├── go.mod                 # Go module for auth service
│   └── go.sum                 # Go dependencies for auth service
│
├── /crawl-service
│   ├── Dockerfile
│   ├── main.go                # Crawl management service code
│   ├── go.mod                 # Go module for crawl service
│   └── go.sum                 # Go dependencies for crawl service
│
├── /data-service
│   ├── Dockerfile
│   ├── main.go                # Data storage service code
│   ├── go.mod                 # Go module for data service
│   └── go.sum                 # Go dependencies for data service
│
├── /api-gateway
│   ├── Dockerfile
│   ├── main.go                # API Gateway (routes & JWT validation)
│   ├── go.mod                 # Go module for API Gateway
│   └── go.sum                 # Go dependencies for API Gateway
│
├── /config
│   └── config.yaml            # Configuration file for environment variables
│
├── docker-compose.yml         # Docker Compose configuration for all services
└── README.md                  # Project description and setup instructions
```

## Scalability

Each microservice is isolated and can be scaled independently based on traffic. Docker Compose makes it easy to manage the entire system locally, and each service can be deployed in a containerized environment such as Kubernetes for production.

## Conclusion

This project demonstrates the power of microservices by creating a modular and scalable web crawler with JWT authentication. Each service can be developed, tested, and deployed independently, providing flexibility and scalability for large-scale applications.
