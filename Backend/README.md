# ⚙️ PRIMMFY Backend

Go backend API for PRIMMFY learning platform.

## 🚀 Quick Start

See [main README](../README.md) for full setup instructions.

### Development

```bash
go mod download
cp .env.example .env
# Edit .env
go run main.go
```

Server runs on http://localhost:8080

## 📁 Structure

```
Backend/
├── main.go
├── config/          # Configuration
├── controllers/     # HTTP handlers
├── models/          # Database models
├── routes/          # API routes
├── middleware/      # Middleware (auth, cors)
├── services/        # Business logic
└── utils/           # Helper functions
```

## 🔐 Authentication

Uses JWT tokens:
- Access token expires: 24 hours
- Refresh token expires: 7 days

## 🗄️ Database

PostgreSQL with GORM ORM.

### Tables
- `users` - User accounts (teacher/student)
- `lessons` - Learning content
- `progress` - Student progress tracking

## 📡 API Endpoints

### Auth
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login
- `POST /api/auth/logout` - Logout
- `GET /api/auth/me` - Get current user

### Lessons (Teacher)
- `GET /api/lessons` - List all lessons
- `POST /api/lessons` - Create lesson
- `PUT /api/lessons/:id` - Update lesson
- `DELETE /api/lessons/:id` - Delete lesson

### Student
- `GET /api/student/lessons` - Available lessons
- `POST /api/student/progress` - Update progress

For more details, see [main README](../README.md).