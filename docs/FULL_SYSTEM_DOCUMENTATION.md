# 🎓 PRIMMFY FULL SYSTEM DOCUMENTATION

> **Purpose:** Dokumentasi lengkap sistem PRIMMFY - Complete Learning Platform dengan PRIMM Methodology

**Last Updated:** October 25, 2025  
**Version:** 1.0  
**Status:** Production Ready ✅

---

## 📋 Table of Contents

1. [System Overview](#1-system-overview)
2. [Architecture](#2-architecture)
3. [Complete Database Schema](#3-complete-database-schema)
4. [All API Endpoints](#4-all-api-endpoints)
5. [Authentication & Authorization](#5-authentication--authorization)
6. [Complete Test Suite](#6-complete-test-suite)
7. [Deployment Guide](#7-deployment-guide)
8. [Maintenance Guide](#8-maintenance-guide)

---

## 🎯 1. SYSTEM OVERVIEW

### What is PRIMMFY?

PRIMMFY adalah **Educational Learning Platform** yang mengimplementasikan **PRIMM Methodology** untuk mengajarkan programming melalui 5-stage scaffolded learning approach.

### System Purpose

**For Teachers:**

- ✅ Create structured programming lessons
- ✅ Organize content into courses and stages
- ✅ Implement PRIMM methodology easily
- ✅ Track student progress (future)
- ✅ Manage learning content

**For Students:**

- ✅ Learn programming systematically
- ✅ Follow PRIMM learning path
- ✅ Get instant feedback
- ✅ Earn rewards (coins)
- ✅ Track progress visually

### Key Features

```
┌─────────────────────────────────────────────────────────┐
│  PRIMMFY PLATFORM FEATURES                              │
├─────────────────────────────────────────────────────────┤
│  ✅ Role-Based Access (Teacher/Student)                 │
│  ✅ JWT Authentication & Authorization                  │
│  ✅ Content Management System                           │
│  ✅ 5-Stage PRIMM Learning Flow                         │
│  ✅ Auto-Grading System                                 │
│  ✅ Progress Tracking                                   │
│  ✅ Gamification (Coins & Rewards)                      │
│  ✅ PostgreSQL Database                                 │
│  ✅ RESTful API Design                                  │
│  ✅ Production-Ready Architecture                       │
└─────────────────────────────────────────────────────────┘
```

---

## 🏗️ 2. ARCHITECTURE

### Tech Stack

```
┌─────────────────────────────────────────┐
│  BACKEND                                │
├─────────────────────────────────────────┤
│  Language:     Go 1.21+                 │
│  Framework:    Gin (HTTP Router)        │
│  Database:     PostgreSQL 15+           │
│  DB Driver:    pgx/v5                   │
│  Auth:         JWT (HS256)              │
│  Validation:   go-playground/validator  │
└─────────────────────────────────────────┘

┌─────────────────────────────────────────┐
│  FRONTEND (Future)                      │
├─────────────────────────────────────────┤
│  Framework:    Next.js / React          │
│  Language:     TypeScript               │
│  Styling:      TailwindCSS              │
│  State:        Context API / Redux      │
└─────────────────────────────────────────┘
```

### System Architecture

```
┌──────────────────────────────────────────────────────────┐
│                     CLIENT (Browser)                     │
│              HTTP Requests with JWT Token                │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                   GIN HTTP ROUTER                        │
│                  (main.go - Routing)                     │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                   MIDDLEWARE LAYER                       │
│         • JWT Validation (RequireAuth)                   │
│         • Role Authorization (RequireRole)               │
│         • CORS Handling                                  │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                   HANDLERS LAYER                         │
│         (HTTP Controllers - Parse & Validate)            │
│  • auth_handler.go      - Register/Login                 │
│  • lesson_handler.go    - Lesson CRUD                    │
│  • course_handler.go    - Course CRUD                    │
│  • primm_stage_handler.go - Stage CRUD & Submissions     │
│  • progress_handler.go  - Progress Tracking              │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                   SERVICES LAYER                         │
│            (Business Logic & Database Ops)               │
│  • auth.go              - Authentication & JWT           │
│  • lesson_service.go    - Lesson operations              │
│  • course_service.go    - Course operations              │
│  • primm_stage_service.go - Stage ops & Grading          │
│  • progress_service.go  - Progress calculations          │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                   MODELS LAYER                           │
│              (Data Structures & Validation)              │
│  • user.go              - User models                    │
│  • lesson.go            - Lesson models                  │
│  • course.go            - Course models                  │
│  • primm_stage.go       - Stage models                   │
│  • progress.go          - Progress models                │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                   DATABASE LAYER                         │
│                 PostgreSQL Database                      │
│  • users                - User accounts                  │
│  • lessons              - Course topics                  │
│  • courses              - Sub-topics                     │
│  • primm_stages         - 5 learning stages              │
│  • user_lessons         - Enrollments                    │
│  • stage_submissions    - Student answers                │
│  • user_course_completion - Completion tracking          │
└──────────────────────────────────────────────────────────┘
```

### Project Structure

```
d:\Project\PRIMMFY\
├── Backend/
│   ├── main.go                    # Entry point & routing
│   ├── database.go                # PostgreSQL connection
│   ├── .env                       # Environment variables
│   ├── go.mod                     # Go dependencies
│   ├── go.sum                     # Dependency checksums
│   │
│   ├── handlers/                  # HTTP Controllers
│   │   ├── auth_handler.go        # Register & Login
│   │   ├── lesson_handler.go      # Lesson CRUD
│   │   ├── course_handler.go      # Course CRUD
│   │   ├── primm_stage_handler.go # Stage CRUD & Submissions
│   │   └── progress_handler.go    # Progress tracking
│   │
│   ├── services/                  # Business Logic
│   │   ├── auth.go                # Auth & JWT
│   │   ├── lesson_service.go      # Lesson operations
│   │   ├── course_service.go      # Course operations
│   │   ├── primm_stage_service.go # Stage ops & grading
│   │   └── progress_service.go    # Progress calculations
│   │
│   ├── models/                    # Data Structures
│   │   ├── user.go                # User models
│   │   ├── lesson.go              # Lesson models
│   │   ├── course.go              # Course models
│   │   ├── primm_stage.go         # Stage models
│   │   └── progress.go            # Progress models
│   │
│   ├── middleware/                # HTTP Middleware
│   │   └── middleware.go          # JWT auth & role validation
│   │
│   └── Database/                  # SQL Scripts
│       ├── schema_v2.sql          # Complete database schema
│       └── drop_and_recreate.sql  # Reset database
│
├── Frontend/                      # Next.js Frontend (future)
│   ├── src/
│   ├── public/
│   └── package.json
│
├── docs/                          # Documentation
│   ├── TEACHER_FLOW_DOCUMENTATION.md
│   ├── STUDENT_FLOW_DOCUMENTATION.md
│   └── FULL_SYSTEM_DOCUMENTATION.md
│
└── README.md                      # Project overview
```

---

## 🗄️ 3. COMPLETE DATABASE SCHEMA

### 3.1 users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('teacher', 'student')),
    level INTEGER DEFAULT 1,
    coins INTEGER DEFAULT 0,
    experience_points INTEGER DEFAULT 0,
    profile_picture TEXT,
    bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
```

**Purpose:** Store user accounts (teachers & students)

**Key Fields:**

- `email` - Unique identifier (no username)
- `password_hash` - bcrypt hashed password
- `role` - 'teacher' or 'student' for authorization
- `coins` - Gamification currency
- `level`, `experience_points` - Future gamification

---

### 3.2 lessons Table

```sql
CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    teacher_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50) NOT NULL CHECK (category IN ('python', 'javascript', 'html', 'c')),
    difficulty VARCHAR(20) NOT NULL CHECK (difficulty IN ('beginner', 'intermediate', 'advanced')),
    thumbnail_url TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_lessons_teacher ON lessons(teacher_id);
CREATE INDEX idx_lessons_category ON lessons(category);
CREATE INDEX idx_lessons_active ON lessons(is_active);
```

**Purpose:** Store lessons (big topics) created by teachers

**Key Fields:**

- `teacher_id` - Ownership (teacher only edits own lessons)
- `category` - Programming language
- `difficulty` - Beginner/Intermediate/Advanced
- `is_active` - Soft delete (hide without deleting)

---

### 3.3 courses Table

```sql
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    order_index INTEGER NOT NULL,
    coin_reward INTEGER NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(lesson_id, order_index)
);

CREATE INDEX idx_courses_lesson ON courses(lesson_id);
```

**Purpose:** Store courses (sub-topics) within lessons

**Key Fields:**

- `order_index` - Sequential ordering within lesson
- `coin_reward` - Coins awarded on completion
- `UNIQUE(lesson_id, order_index)` - Prevent duplicate ordering

---

### 3.4 primm_stages Table

```sql
CREATE TABLE primm_stages (
    id SERIAL PRIMARY KEY,
    course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    stage_type VARCHAR(20) NOT NULL CHECK (stage_type IN ('predict', 'run', 'investigate', 'modify', 'make')),
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    order_index INTEGER NOT NULL CHECK (order_index >= 1 AND order_index <= 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- Common fields
    code_snippet TEXT,
    task_description TEXT,

    -- PREDICT stage specific
    predict_options JSONB,
    correct_answer TEXT,

    -- RUN stage specific
    run_code_template TEXT,

    -- INVESTIGATE stage specific
    video_embed_url TEXT,
    explanation_text TEXT,
    guiding_questions JSONB,
    reflection_prompt TEXT,

    -- MODIFY stage specific
    modify_challenge TEXT,
    modify_code_template TEXT,
    modify_expected_output TEXT,
    modify_test_cases JSONB,

    -- MAKE stage specific
    make_challenge TEXT,
    make_hints TEXT,
    make_expected_output TEXT,
    make_test_cases JSONB,

    UNIQUE(course_id, order_index)
);

CREATE INDEX idx_stages_course ON primm_stages(course_id);
CREATE INDEX idx_stages_type ON primm_stages(stage_type);
```

**Purpose:** Store 5 PRIMM learning stages per course

**Design Decision:** Single table (polymorphic) vs separate tables

- ✅ Chosen: Single table with nullable fields
- ✅ Easier querying and ordering
- ✅ Flexible JSONB for complex data

---

### 3.5 user_lessons Table

```sql
CREATE TABLE user_lessons (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, lesson_id)
);

CREATE INDEX idx_user_lessons_user ON user_lessons(user_id);
CREATE INDEX idx_user_lessons_lesson ON user_lessons(lesson_id);
```

**Purpose:** Track student enrollments to lessons

**Key Fields:**

- `UNIQUE(user_id, lesson_id)` - Student can only enroll once

**TODO:**

- Add `last_accessed_at` column for activity tracking

---

### 3.6 stage_submissions Table

```sql
CREATE TABLE stage_submissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    stage_id INTEGER NOT NULL REFERENCES primm_stages(id) ON DELETE CASCADE,
    submission_type VARCHAR(50) NOT NULL,
    submission_data JSONB NOT NULL,
    is_correct BOOLEAN DEFAULT false,
    score INTEGER DEFAULT 0,
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, stage_id)
);

CREATE INDEX idx_stage_submissions_user ON stage_submissions(user_id);
CREATE INDEX idx_stage_submissions_stage ON stage_submissions(stage_id);
```

**Purpose:** Store student submissions for each stage

**Key Fields:**

- `submission_data` - JSONB for flexible data per stage type
- `is_correct` - Pass/fail status
- `score` - 0-100 scoring
- `UNIQUE(user_id, stage_id)` - One submission per stage (can update)

**Submission Data Examples:**

**PREDICT:**

```json
{
  "selected_answer": "C",
  "correct_answer": "C"
}
```

**RUN:**

```json
{
  "code_output": "15"
}
```

**INVESTIGATE:**

```json
{
  "reflection_text": "Variables are containers..."
}
```

**MODIFY:**

```json
{
  "modified_code": "x = 5\ny = 10\nz = 15..."
}
```

**MAKE:**

```json
{
  "code": "name = input()..."
}
```

---

### 3.7 user_course_completion Table

```sql
CREATE TABLE user_course_completion (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    coins_awarded INTEGER DEFAULT 0,
    UNIQUE(user_id, course_id)
);

CREATE INDEX idx_course_completion_user ON user_course_completion(user_id);
CREATE INDEX idx_course_completion_course ON user_course_completion(course_id);
```

**Purpose:** Track course completions and coin rewards

**Key Fields:**

- `coins_awarded` - Amount of coins given for this completion
- `UNIQUE(user_id, course_id)` - Course can only be completed once

---

### Database Relationships

```
users (1) ──────┬─── (N) lessons
                │
                └─── (N) user_lessons ─── (N) lessons

lessons (1) ──── (N) courses

courses (1) ──── (N) primm_stages

primm_stages (1) ──── (N) stage_submissions ─── (N) users

courses (1) ──── (N) user_course_completion ─── (N) users
```

---

## 🔌 4. ALL API ENDPOINTS

### 4.1 Public Endpoints (No Authentication)

#### POST /api/register

Register new user (teacher or student)

#### POST /api/login

Login and get JWT token

#### GET /api/lessons

Browse all active lessons (optional auth)

---

### 4.2 Teacher-Only Endpoints

**Require:**

- ✅ Authorization header with valid JWT
- ✅ Role = "teacher"

#### Lesson Management

- `POST /api/lessons` - Create lesson
- `PUT /api/lessons/:id` - Update lesson
- `DELETE /api/lessons/:id` - Delete lesson (soft delete)

#### Course Management

- `POST /api/courses` - Create course
- `PUT /api/courses/:id` - Update course
- `DELETE /api/courses/:id` - Delete course

#### Stage Management (PRIMM)

- `POST /api/stages/predict` - Create PREDICT stage
- `POST /api/stages/run` - Create RUN stage
- `POST /api/stages/investigate` - Create INVESTIGATE stage
- `POST /api/stages/modify` - Create MODIFY stage
- `POST /api/stages/make` - Create MAKE stage
- `DELETE /api/stages/:id` - Delete stage

---

### 4.3 Student-Only Endpoints

**Require:**

- ✅ Authorization header with valid JWT
- ✅ Role = "student"

#### Enrollment

- `POST /api/lessons/:id/enroll` - Enroll to lesson
- `GET /api/my-lessons` - View enrolled lessons with progress

#### View Content

- `GET /api/lessons/:id/courses` - View courses in lesson
- `GET /api/courses/:id/stages` - View stages in course

#### Submissions

- `POST /api/stages/:id/submit` - Submit answer for stage

---

### 4.4 Shared Authenticated Endpoints

**Require:**

- ✅ Authorization header with valid JWT
- ✅ Any role (teacher or student)

- `GET /api/profile` - Get current user profile
- `GET /api/stages/:id` - Get stage details

---

### Request/Response Examples

See detailed examples in:

- **Teacher Flow:** `docs/TEACHER_FLOW_DOCUMENTATION.md`
- **Student Flow:** `docs/STUDENT_FLOW_DOCUMENTATION.md`

---

## 🔐 5. AUTHENTICATION & AUTHORIZATION

### JWT Token Structure

**Generation:**

```go
func GenerateJWT(userID int, email string, role string) (string, error) {
    secretKey := os.Getenv("JWT_SECRET")

    claims := jwt.MapClaims{
        "user_id": userID,
        "email":   email,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secretKey))
}
```

**Token Payload:**

```json
{
  "user_id": 8,
  "email": "student@example.com",
  "role": "student",
  "exp": 1762001600
}
```

**Usage in Requests:**

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

### Middleware: RequireAuth

**File:** `middleware/middleware.go`

```go
func RequireAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. Get token from header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Authorization required"})
            c.Abort()
            return
        }

        // 2. Parse "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(401, gin.H{"error": "Invalid format"})
            c.Abort()
            return
        }

        // 3. Validate token
        claims, err := services.ValidateJWT(parts[1])
        if err != nil {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // 4. Set user info to context
        c.Set("user_id", int(claims["user_id"].(float64)))
        c.Set("user_role", claims["role"].(string))

        c.Next()
    }
}
```

---

### Middleware: RequireRole

```go
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("user_role")
        if !exists {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        roleStr := role.(string)
        for _, allowedRole := range allowedRoles {
            if roleStr == allowedRole {
                c.Next()
                return
            }
        }

        c.JSON(403, gin.H{"error": "Access denied"})
        c.Abort()
    }
}
```

---

### Routing with Middleware

**File:** `main.go`

```go
func main() {
    router := gin.Default()

    // Public routes
    router.POST("/api/register", authHandler.Register)
    router.POST("/api/login", authHandler.Login)
    router.GET("/api/lessons", lessonHandler.GetAllLessons)

    // Protected routes (require auth)
    protected := router.Group("/api")
    protected.Use(middleware.RequireAuth())
    {
        // Teacher-only routes
        teacher := protected.Group("")
        teacher.Use(middleware.RequireRole("teacher"))
        {
            teacher.POST("/lessons", lessonHandler.CreateLesson)
            teacher.POST("/courses", courseHandler.CreateCourse)
            teacher.POST("/stages/predict", stageHandler.CreatePredictStage)
            // ... more teacher routes
        }

        // Student-only routes
        student := protected.Group("")
        student.Use(middleware.RequireRole("student"))
        {
            student.POST("/lessons/:id/enroll", lessonHandler.EnrollLesson)
            student.GET("/my-lessons", lessonHandler.GetMyEnrolledLessons)
            student.POST("/stages/:id/submit", stageHandler.SubmitStage)
            // ... more student routes
        }

        // Shared authenticated routes
        protected.GET("/profile", authHandler.GetProfile)
        protected.GET("/stages/:id", stageHandler.GetStageByID)
    }

    router.Run(":8080")
}
```

---

## 🧪 6. COMPLETE TEST SUITE

### Test Summary

```
┌─────────────────────────────────────────────────────────┐
│  COMPLETE TEST SUITE (24 Tests)                         │
├─────────────────────────────────────────────────────────┤
│  Teacher Flow (Test 1-9)     ✅ All Passed              │
│  Student Flow (Test 10-24)   ✅ All Passed              │
│  Total Success Rate:         100%                       │
└─────────────────────────────────────────────────────────┘
```

### Teacher Flow Tests (1-9)

| Test | Endpoint                     | Purpose                  | Status |
| ---- | ---------------------------- | ------------------------ | ------ |
| 1    | POST /api/register           | Register teacher         | ✅     |
| 2    | POST /api/login              | Login teacher            | ✅     |
| 3    | POST /api/lessons            | Create lesson            | ✅     |
| 4    | POST /api/courses            | Create course            | ✅     |
| 5    | POST /api/stages/predict     | Create PREDICT stage     | ✅     |
| 6    | POST /api/stages/run         | Create RUN stage         | ✅     |
| 7    | POST /api/stages/investigate | Create INVESTIGATE stage | ✅     |
| 8    | POST /api/stages/modify      | Create MODIFY stage      | ✅     |
| 9    | POST /api/stages/make        | Create MAKE stage        | ✅     |

---

### Student Flow Tests (10-24)

| Test | Endpoint                   | Purpose               | Status |
| ---- | -------------------------- | --------------------- | ------ |
| 10   | POST /api/register         | Register student      | ✅     |
| 11   | POST /api/login            | Login student         | ✅     |
| 12   | GET /api/lessons           | Browse lessons        | ✅     |
| 13   | POST /api/lessons/2/enroll | Enroll to lesson      | ✅     |
| 14   | GET /api/my-lessons        | View enrolled lessons | ✅     |
| 15   | GET /api/lessons/2/courses | View courses          | ✅     |
| 16   | GET /api/courses/1/stages  | View all stages       | ✅     |
| 17   | POST /api/stages/1/submit  | Submit PREDICT        | ✅     |
| 18   | POST /api/stages/2/submit  | Submit RUN            | ✅     |
| 19   | POST /api/stages/3/submit  | Submit INVESTIGATE    | ✅     |
| 20   | POST /api/stages/4/submit  | Submit MODIFY         | ✅     |
| 21   | POST /api/stages/5/submit  | Submit MAKE           | ✅     |
| 22   | GET /api/profile           | Verify coins awarded  | ✅     |
| 23   | GET /api/my-lessons        | Verify progress 100%  | ✅     |
| 24   | GET /api/lessons/2/courses | Verify completion     | ✅     |

---

### Complete Test Data After All Tests

```
TEACHER ACCOUNT:
├── ID: 7
├── Email: teacher@example.com
├── Name: John Teacher
├── Role: teacher
└── Created Content:
    ├── 1 Lesson: "Python Basics"
    ├── 1 Course: "Variables and Data Types"
    └── 5 Stages: PREDICT, RUN, INVESTIGATE, MODIFY, MAKE

STUDENT ACCOUNT:
├── ID: 8
├── Email: student@example.com
├── Name: Alice Student
├── Role: student
├── Coins: 200 🪙
└── Learning Progress:
    ├── Enrolled: 1 lesson
    ├── Completed: 1 course (100%)
    └── Submissions: 5/5 stages (all correct)

DATABASE STATE:
├── users: 2 records (1 teacher, 1 student)
├── lessons: 2 records (1 active)
├── courses: 1 record
├── primm_stages: 5 records
├── user_lessons: 1 record
├── stage_submissions: 5 records
└── user_course_completion: 1 record
```

---

## 🚀 7. DEPLOYMENT GUIDE

### Prerequisites

```
✅ Go 1.21+ installed
✅ PostgreSQL 15+ installed
✅ Git installed
```

### Step 1: Clone Repository

```bash
git clone <repository-url>
cd PRIMMFY
```

### Step 2: Setup Database

```bash
# Create database
createdb primmfy_db

# Or via psql
psql -U postgres
CREATE DATABASE primmfy_db;
\q

# Run schema
psql -U postgres -d primmfy_db -f Backend/Database/schema_v2.sql
```

### Step 3: Configure Environment

Create `Backend/.env`:

```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=primmfy_db

# JWT Secret (generate strong random string)
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Server
PORT=8080
```

**Generate JWT Secret:**

```bash
# Linux/Mac
openssl rand -base64 32

# Windows PowerShell
[Convert]::ToBase64String((1..32 | ForEach-Object { Get-Random -Maximum 256 }))
```

### Step 4: Install Dependencies

```bash
cd Backend
go mod download
```

### Step 5: Build & Run

```bash
# Development
go run .

# Production build
go build -o primmfy-server
./primmfy-server

# Windows
go build -o primmfy-server.exe
./primmfy-server.exe
```

### Step 6: Verify Server

```bash
# Check health
curl http://localhost:8080/api/lessons

# Should return lesson list or empty array
```

---

## 🔧 8. MAINTENANCE GUIDE

### Daily Tasks

**Monitor Server Health:**

```bash
# Check server logs
tail -f server.log

# Check database connections
psql -U postgres -d primmfy_db -c "SELECT count(*) FROM pg_stat_activity;"
```

---

### Weekly Tasks

**Database Backup:**

```bash
# Full backup
pg_dump -U postgres primmfy_db > backup_$(date +%Y%m%d).sql

# Restore if needed
psql -U postgres -d primmfy_db < backup_20251025.sql
```

---

### Common Maintenance Queries

**Check User Statistics:**

```sql
SELECT
    role,
    COUNT(*) as total_users,
    SUM(coins) as total_coins
FROM users
GROUP BY role;
```

**Check Popular Lessons:**

```sql
SELECT
    l.id,
    l.title,
    COUNT(ul.user_id) as enrolled_students
FROM lessons l
LEFT JOIN user_lessons ul ON l.id = ul.lesson_id
GROUP BY l.id, l.title
ORDER BY enrolled_students DESC;
```

**Check Course Completion Rate:**

```sql
SELECT
    c.id,
    c.title,
    COUNT(DISTINCT ul.user_id) as enrolled,
    COUNT(DISTINCT ucc.user_id) as completed,
    ROUND(COUNT(DISTINCT ucc.user_id)::numeric / NULLIF(COUNT(DISTINCT ul.user_id), 0) * 100, 2) as completion_rate
FROM courses c
JOIN lessons l ON c.lesson_id = l.id
LEFT JOIN user_lessons ul ON l.id = ul.lesson_id
LEFT JOIN user_course_completion ucc ON c.id = ucc.course_id AND ucc.user_id = ul.user_id
GROUP BY c.id, c.title;
```

---

### Adding New Features

**1. Add New Stage Type:**

Update `primm_stages` table:

```sql
ALTER TABLE primm_stages
ADD COLUMN new_stage_field TEXT;
```

Update validation:

```go
// models/primm_stage.go
CHECK (stage_type IN ('predict', 'run', 'investigate', 'modify', 'make', 'newtype'))
```

Add handler:

```go
// handlers/primm_stage_handler.go
func (h *PRIMMStageHandler) CreateNewTypeStage(c *gin.Context) {
    // Implementation
}
```

---

**2. Add Progress Analytics:**

Create new table:

```sql
CREATE TABLE user_analytics (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    time_spent_minutes INT DEFAULT 0,
    stages_completed INT DEFAULT 0,
    avg_score DECIMAL(5,2) DEFAULT 0,
    last_activity TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Track in submission:

```go
// After successful submission
go updateUserAnalytics(db, userID, stageID, timeSpent)
```

---

**3. Add Teacher Dashboard:**

Create new endpoints:

```go
// GET /api/teacher/dashboard
func (h *TeacherHandler) GetDashboard(c *gin.Context) {
    teacherID := c.Get("user_id")

    stats := services.GetTeacherStats(db, teacherID.(int))

    c.JSON(200, stats)
}
```

Query student progress:

```sql
SELECT
    u.full_name,
    COUNT(DISTINCT ucc.course_id) as courses_completed,
    SUM(ucc.coins_awarded) as total_coins_earned
FROM users u
JOIN user_lessons ul ON u.id = ul.user_id
JOIN lessons l ON ul.lesson_id = l.id
LEFT JOIN courses c ON l.id = c.lesson_id
LEFT JOIN user_course_completion ucc ON c.id = ucc.course_id AND ucc.user_id = u.id
WHERE l.teacher_id = $1
GROUP BY u.id, u.full_name;
```

---

### Troubleshooting Common Issues

**Issue: Server won't start**

```bash
# Check if port 8080 is in use
netstat -ano | findstr :8080  # Windows
lsof -i :8080                  # Linux/Mac

# Kill process if needed
taskkill /PID <PID> /F         # Windows
kill -9 <PID>                  # Linux/Mac
```

**Issue: Database connection failed**

```go
// Check connection string in database.go
connStr := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
)
```

**Issue: JWT validation fails**

```
Error: "Invalid token"

Solutions:
1. Check JWT_SECRET in .env matches generation secret
2. Check token hasn't expired (7 days validity)
3. Verify Authorization header format: "Bearer <token>"
```

---

## 📊 SYSTEM METRICS

### Current Statistics (After Test 24)

```
USERS:
├── Total: 2
├── Teachers: 1
└── Students: 1

CONTENT:
├── Lessons: 2
├── Courses: 1
├── PRIMM Stages: 5
└── Active: All

ENGAGEMENT:
├── Enrollments: 1
├── Submissions: 5
├── Completions: 1
└── Coins Distributed: 200

DATABASE:
├── Total Tables: 7
├── Total Indexes: 15
├── Total Records: ~15
└── Size: <1MB
```

---

## 🎯 FUTURE ROADMAP

### Phase 2: Enhanced Features

**Code Execution Engine:**

- [ ] Integrate Python sandbox (Piston API / Judge0)
- [ ] Auto-validate RUN stage outputs
- [ ] Run MODIFY & MAKE test cases
- [ ] Support multiple programming languages

**Advanced Grading:**

- [ ] AI-powered reflection analysis
- [ ] Code quality metrics
- [ ] Partial credit system
- [ ] Detailed feedback generation

**Progress Enhancement:**

- [ ] Add `last_accessed_at` tracking
- [ ] Time spent per stage analytics
- [ ] Leaderboard system
- [ ] Achievement badges

**Teacher Tools:**

- [ ] Teacher dashboard with analytics
- [ ] Student progress reports
- [ ] Bulk content import/export
- [ ] Content templates

**Student Experience:**

- [ ] Hints system (reveal incrementally)
- [ ] Discussion forum per course
- [ ] Peer code review
- [ ] Certificate generation

---

## 📚 ADDITIONAL RESOURCES

### Documentation Files

- **Teacher Flow:** `docs/TEACHER_FLOW_DOCUMENTATION.md`
- **Student Flow:** `docs/STUDENT_FLOW_DOCUMENTATION.md`
- **This File:** `docs/FULL_SYSTEM_DOCUMENTATION.md`

### External Resources

- **Go Gin Framework:** https://gin-gonic.com/docs/
- **PostgreSQL Docs:** https://www.postgresql.org/docs/
- **pgx Driver:** https://pkg.go.dev/github.com/jackc/pgx/v5
- **JWT.io:** https://jwt.io/
- **PRIMM Methodology:** https://primmportal.com/

---

## 📝 CHANGELOG

### Version 1.0 (October 25, 2025)

**Initial Release:**

- ✅ Complete Teacher Flow (Test 1-9)
- ✅ Complete Student Flow (Test 10-24)
- ✅ JWT Authentication & Authorization
- ✅ 5-Stage PRIMM Implementation
- ✅ Auto-Grading System
- ✅ Progress Tracking
- ✅ Coin Reward System
- ✅ PostgreSQL Database
- ✅ RESTful API Design
- ✅ Complete Documentation

---

## 🏆 ACKNOWLEDGMENTS

**Built with:**

- Go Programming Language
- Gin Web Framework
- PostgreSQL Database
- JWT for Authentication
- PRIMM Methodology

**Special Thanks:**

- PRIMM Portal for pedagogy framework
- Go Community for excellent tools
- OpenAI for development assistance

---

**Last Updated:** October 25, 2025  
**Version:** 1.0  
**Status:** Production Ready ✅

---

## 📞 SUPPORT

For questions, issues, or contributions:

- **Documentation:** Read this file and related docs
- **Database Issues:** Check `Database/schema_v2.sql`
- **API Reference:** See Teacher/Student flow docs
- **Code Issues:** Check inline comments in code

---

**END OF DOCUMENTATION**

This system is now fully functional and ready for:

- ✅ Development continuation
- ✅ Feature enhancement
- ✅ Production deployment
- ✅ Educational use

**Happy Teaching & Learning with PRIMMFY! 🎓🚀**
