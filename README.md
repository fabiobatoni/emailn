## EMAILN - Golang Project

[![Golang](https://img.shields.io/badge/-Golang-00ADD8?logo=go&logoColor=white&style=for-the-badge)](https://go.dev/)
[![GORM](https://img.shields.io/badge/-GORM-720026?logo=go&logoColor=white&style=for-the-badge)](https://gorm.io/)
[![Chi](https://img.shields.io/badge/-Go%20Chi-00BFFF?logo=go&logoColor=white&style=for-the-badge)](https://github.com/go-chi/chi)
[![JWT](https://img.shields.io/badge/-JWT-000000?logo=jsonwebtokens&logoColor=white&style=for-the-badge)](https://jwt.io/)
[![Keycloak](https://img.shields.io/badge/-Keycloak-0078D7?logo=keycloak&logoColor=white&style=for-the-badge)](https://www.keycloak.org/)
[![Dotenv](https://img.shields.io/badge/-Go%20Dotenv-4CAF50?logo=dotenv&logoColor=white&style=for-the-badge)](https://github.com/joho/godotenv)

## 1. Introduction

EMAILN is a project built with Golang that integrates several technologies for authentication, database management, and email handling. This project was created as a campaign generator, allowing users to create marketing campaigns and send emails to clients with relevant information about those campaigns.

- **GORM** - ORM for Golang
- **Go Chi** - Lightweight router for HTTP handling
- **Testify** - Testing utilities
- **Air** - Live reloading for development
- **Keycloak** - Identity and access management
- **Go-OIDC** - OpenID Connect authentication
- **JWT-Go** - JSON Web Token authentication
- **Go-Dotenv** - Load environment variables
- **Go-Mail** - Email sending functionality

## 2. Getting Started

To set up a local copy of this repository:

- For SSH (recommended for secure, key-based authentication), use:

  ```bash
  git clone git@github.com:fabiobatoni/emailn.git
  ```

- For HTTPS (simpler setup, ideal for quick trials), use:

  ```bash
  git clone https://github.com/fabiobatoni/emailn.git
  ```

### 2.1 Installation

Follow these steps to set up the EMAILN environment on your local machine:

1. **Clone the Repository and Change Directory**

```bash
cd emailn
```

2. **Install Project Dependencies**

```bash
go mod tidy
```

### 2.2 Local Development

**Generate Database Tables**

To generate the necessary tables, run:

```bash
go run main.go
```

**Run the Application**

For local development with live reload, use:

```bash
air
```

Once the server is running, the API will be accessible at:

- [http://localhost:3000](http://localhost:3000)

## 3. Testing

To run tests with Testify, use:

```bash
go test ./...
```

---


