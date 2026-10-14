# Devora Backend

<p align="center">
  <img src="https://img.shields.io/badge/Devora-Backend-6366F1?style=for-the-badge&logo=go&logoColor=white" alt="Devora Backend" />
</p>

<p align="center">
  <strong>Build together. Learn together. Grow together.</strong>
  <br />
  A backend foundation for a collaborative platform built for students and developers.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-Backend-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Gin-HTTP_Framework-008ECF?style=flat-square" alt="Gin" />
  <img src="https://img.shields.io/badge/MySQL-8.4-4479A1?style=flat-square&logo=mysql&logoColor=white" alt="MySQL" />
  <img src="https://img.shields.io/badge/Docker-Containerized-2496ED?style=flat-square&logo=docker&logoColor=white" alt="Docker" />
  <img src="https://img.shields.io/badge/Status-In_Development-orange?style=flat-square" alt="In Development" />
</p>

---

## <img src="https://img.shields.io/badge/-Overview-6366F1?style=flat-square" alt="Overview" />

Devora is a platform concept designed to help students collaborate, manage projects, share knowledge, and use AI-powered tools throughout their learning journey.

The goal is to create a unified environment where students can turn ideas into real projects, build teams, and learn by working together.

This repository contains the backend API and supporting infrastructure for the Devora platform.

> **Project status:** Active development. Features and architecture will evolve as the project progresses.

## <img src="https://img.shields.io/badge/-Core_Goals-6366F1?style=flat-square" alt="Core Goals" />

- **User Management** — Accounts, profiles, and authentication.
- **Project Collaboration** — Project creation, team formation, and task management.
- **Community** — Groups, posts, and knowledge sharing.
- **AI Assistance** — Support for learning, project planning, and development.
- **Security** — Secure authentication, authorization, and input validation.
- **Scalability** — A maintainable foundation that can grow with the platform.

## <img src="https://img.shields.io/badge/-Technology_Stack-6366F1?style=flat-square" alt="Technology Stack" />

| Technology | Role |
|---|---|
| [Go](https://go.dev/) | Backend development |
| [Gin](https://gin-gonic.com/) | HTTP routing and API development |
| [MySQL](https://www.mysql.com/) | Relational data storage |
| [Docker](https://www.docker.com/) | Local database environment |
| [Git](https://git-scm.com/) | Version control |

## <img src="https://img.shields.io/badge/-Architecture-6366F1?style=flat-square" alt="Architecture" />

The backend is being organized into focused components to separate HTTP handling, business logic, data access, and infrastructure.

```text
Devora_Backend/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   ├── controllers/
│   ├── services/
│   ├── repositories/
│   └── models/
├── platform/
│   └── database/
│       └── mysql.go
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

*The structure is a development guide; directories and components may change as implementation progresses.*

## <img src="https://img.shields.io/badge/-Getting_Started-6366F1?style=flat-square" alt="Getting Started" />

### Prerequisites

- Go
- Docker
- Git

### 1. Clone the repository

```bash
git clone https://github.com/3mrkhalid/Devora_Backend.git
cd Devora_Backend
```

### 2. Configure the database

Start a local MySQL 8.4 container if you do not already have one:

```bash
docker run -d \
  --name devora-mysql \
  -e MYSQL_ROOT_PASSWORD=change_me \
  -e MYSQL_DATABASE=devora \
  -p 3306:3306 \
  mysql:8.4
```

Replace `change_me` with a secure local development password.

### 3. Configure environment variables

```bash
cp .env.example .env
```

Update `.env` to match your database credentials and the variable names expected by the application.

Example:

```env
PORT=:8080
DATABASE_URL=devora_user:your_password@tcp(127.0.0.1:3306)/devora?parseTime=true
```

Ensure the database user exists and has the required permissions.

### 4. Install dependencies

```bash
go mod download
```

### 5. Run the application

```bash
go run ./cmd/api
```

> Setup instructions may change as development continues.

## <img src="https://img.shields.io/badge/-Roadmap-6366F1?style=flat-square" alt="Roadmap" />

- [ ] Environment configuration
- [ ] MySQL connection and database initialization
- [ ] User registration and authentication
- [ ] User profiles
- [ ] Project and task management
- [ ] Teams and group membership
- [ ] Posts and collaboration
- [ ] Authorization and security hardening
- [ ] Automated testing
- [ ] API documentation
- [ ] AI-powered student assistance

## <img src="https://img.shields.io/badge/-Security-6366F1?style=flat-square" alt="Security" />

Security is an important design consideration for Devora.

Planned security practices include:

- Password hashing.
- Request validation.
- Authentication and authorization.
- Parameterized database queries.
- Secure secret management.
- Rate limiting and abuse protection.

These practices will be implemented as the corresponding features are developed.

## <img src="https://img.shields.io/badge/-Testing-6366F1?style=flat-square" alt="Testing" />

Automated tests will be introduced as the application grows, covering business logic, HTTP handlers, and database interactions.

## <img src="https://img.shields.io/badge/-Contributing-6366F1?style=flat-square" alt="Contributing" />

Devora is under active development. Feedback, suggestions, and contributions are welcome.

Please discuss significant changes before implementing them to help maintain a consistent architecture.

## <img src="https://img.shields.io/badge/-License-6366F1?style=flat-square" alt="License" />

No license has been selected yet. Licensing terms will be added when a license is chosen.

---

<p align="center">
  <strong>DEVORA</strong>
  <br />
  <sub>Learn together. Build together.</sub>
</p>
    