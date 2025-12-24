# CinemaAPI 🍿📽️ 

> **CinemaAPI** is a high-performance RESTful service built on a **microservice architecture**, designed to streamline cinema content management and ticketing operations within a distributed ecosystem.  
>  
> Leveraging the power of **HTTP/2** for ultra-fast inter-service communication and **Redis caching** for performance optimization, the system ensures reliability, scalability, and exceptional speed.

---

## 🔮 Features

- [x] **Service-to-Service Communication via HTTP/2** — Enables low-latency, multiplexed communication between microservices for optimal performance and scalability.  
- [x] **JWT Authentication with RS256 Algorithm** — Implements secure, asymmetric token-based authentication to safeguard data integrity and user sessions.  
- [x] **Advanced Role & Permission Control** — Provides fine-grained access management across different user roles, ensuring robust authorization.  
- [x] **Email Verification & OTP Authentication** — Enhances account security with verified email workflows and one-time password (OTP) mechanisms.  
- [x] **Redis Cache Integration** — Accelerates data retrieval and minimizes database load with a high-speed caching layer.  

---

## ⚙️ Requirements  

Before running this project, make sure you have the following installed on your system:  

| Dependency | Minimum Version | Description | Link |
|-------------|----------------|-------------|------|
| **Golang** | `1.24.6` or higher | Main programming language for building microservices. | [golang.org](https://go.dev/dl/) |
| **sqlc** | Latest | Generates type-safe Go code from SQL queries. | [sqlc.dev](https://docs.sqlc.dev/en/latest/) |
| **Redis** | Latest | In-memory data store for caching and session management. | [redis.io](https://redis.io/download) |
| **PostgreSQL** | Latest | Relational database for data storage and management. | [postgresql.org](https://www.postgresql.org/download/) |
| **Docker** | Latest | Containerization platform to run the services seamlessly. | [docker.com](https://www.docker.com/get-started) |

---

## 📥 Instalation
### 1. Clone Repository
```bash
git clone https://github.com/Nuvantim/CinemaAPI.git && \
cd CinemaAPI
```
### 2. Install Modules for Each Service
Run the following commands to install dependencies for each service:

**🎬 Cinema Service**
```bash
cd cinema-service && go mod tidy
```
**🎟️ Booking Service**
```bash
cd booking-service && go mod tidy
```
**🌐 API Gateway**
```bash
cd cinema-service && go mod tidy
```

### 3. Build Binary  

Run the following command to build all services in the workspace:  
```bash
chmod +x build.sh &&\
./build.sh
```

## 🐳 Deployment
CinemaAPI is fully containerized using Docker Compose, allowing one-command deployment across all services.
### 1. Build & Start Containers
```bash
docker-compose up -d --build
```
### 2. Verify Running Containers
```bash
docker ps
```
