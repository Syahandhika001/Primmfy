# 📚 TEACHER FLOW - Complete Documentation

> **Purpose:** Dokumentasi lengkap untuk Teacher Flow (Test 1-9) - untuk maintenance, learning reference, dan AI context memory.

**Last Updated:** October 25, 2025  
**Version:** 1.0  
**Status:** Teacher Flow Complete ✅

---

## 📋 Table of Contents

1. [Overview](#1-overview)
2. [Architecture](#2-architecture)
3. [Database Schema](#3-database-schema)
4. [API Endpoints](#4-api-endpoints)
5. [Test Flow](#5-test-flow)
6. [Code Structure](#6-code-structure)
7. [Troubleshooting](#7-troubleshooting)
8. [Summary](#summary)

---

## 🎯 1. OVERVIEW

### Teacher Flow Purpose

Teacher Flow memungkinkan teacher untuk:

- ✅ Register & Login
- ✅ Membuat Lesson (big topic)
- ✅ Membuat Course di dalam Lesson (sub-topic)
- ✅ Membuat 5 PRIMM Stages di dalam Course:
  - **PREDICT** - Multiple choice prediction
  - **RUN** - Execute & observe code
  - **INVESTIGATE** - Watch video & reflect
  - **MODIFY** - Modify existing code
  - **MAKE** - Create code from scratch

### PRIMM Methodology

**PRIMM** = Predict → Run → Investigate → Modify → Make

- Pedagogical approach untuk teaching programming
- Scaffolded learning dari observation → creation
- Evidence-based teaching method

---

## 🏗️ 2. ARCHITECTURE

### Tech Stack

```
Backend:
├── Language: Go 1.21+
├── Framework: Gin (HTTP router)
├── Database: PostgreSQL 15+
├── DB Driver: pgx/v5 (native PostgreSQL driver)
├── Auth: JWT (HS256)
└── Validation: go-playground/validator/v10
```

### Folder Structure

```
Backend/
├── main.go                     # Entry point & routing
├── database.go                 # PostgreSQL connection
├── .env                        # Environment variables
├── go.mod / go.sum             # Go dependencies
├── handlers/                   # HTTP handlers (controllers)
│   ├── auth_handler.go         # Register & Login
│   ├── lesson_handler.go       # Lesson CRUD
│   ├── course_handler.go       # Course CRUD
│   ├── primm_stage_handler.go  # PRIMM Stage CRUD
│   └── progress_handler.go     # Progress tracking
├── services/                   # Business logic
│   ├── auth.go                 # Auth logic (JWT, bcrypt)
│   ├── lesson_service.go       # Lesson operations
│   ├── course_service.go       # Course operations
│   ├── primm_stage_service.go  # Stage operations
│   └── progress_service.go     # Progress calculations
├── models/                     # Data structures
│   ├── user.go                 # User models
│   ├── lesson.go               # Lesson models
│   ├── course.go               # Course models
│   ├── primm_stage.go          # Stage models
│   └── progress.go             # Progress models
├── middleware/                 # HTTP middleware
│   └── middleware.go           # JWT auth & role validation
└── database/                   # SQL scripts
    ├── schema_v2.sql           # Complete database schema
    └── drop_and_recreate.sql   # Reset database
```

### Request Flow

```
Client Request
    ↓
Gin Router (main.go)
    ↓
Middleware (JWT validation, role check)
    ↓
Handler (parse request, validate)
    ↓
Service (business logic, database operations)
    ↓
Database (PostgreSQL)
    ↓
Response to Client
```

---

## 🗄️ 3. DATABASE SCHEMA

### 3.1 Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('teacher', 'student')),
    level INTEGER DEFAULT 1,
    total_coins INTEGER DEFAULT 0,
    experience_points INTEGER DEFAULT 0,
    profile_picture TEXT,
    bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Key Points:**

- `email` adalah unique identifier (no username)
- `password_hash` menggunakan bcrypt
- `role` untuk authorization (teacher/student)
- `level`, `total_coins`, `experience_points` untuk gamification

---

### 3.2 Lessons Table

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
```

**Key Points:**

- Lesson = Big Topic (e.g., "Python Basics")
- `teacher_id` untuk ownership (teacher hanya bisa edit lesson miliknya)
- `category` terbatas pada programming languages yang supported
- `is_active` untuk soft delete (tidak menghapus data)

---

### 3.3 Courses Table

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
```

**Key Points:**

- Course = Sub-topic dalam Lesson (e.g., "Variables and Data Types")
- `order_index` untuk urutan course (1, 2, 3, ...)
- `coin_reward` adalah coins yang didapat student setelah complete course
- `UNIQUE(lesson_id, order_index)` mencegah duplicate order dalam 1 lesson

---

### 3.4 PRIMM Stages Table

```sql
CREATE TABLE primm_stages (
    id SERIAL PRIMARY KEY,
    course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    stage_type VARCHAR(20) NOT NULL CHECK (stage_type IN ('predict', 'run', 'investigate', 'modify', 'make')),
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    order_index INTEGER NOT NULL CHECK (order_index >= 1 AND order_index <= 5),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- Common fields
    code_snippet TEXT,
    task_description TEXT,

    -- PREDICT stage specific
    predict_options JSONB,
    correct_answer VARCHAR(10),

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
```

**Key Points:**

- **Single Table Design** untuk semua 5 stage types (polymorphic approach)
- `stage_type` menentukan field mana yang dipakai
- `order_index` fixed 1-5 untuk PRIMM sequence
- **JSONB fields** untuk flexible data (options, questions, test cases)
- `UNIQUE(course_id, order_index)` mencegah duplicate stage dalam 1 course

---

### 3.5 Progress Tracking Tables

#### User Lessons (Enrollment)

```sql
CREATE TABLE user_lessons (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_accessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, lesson_id)
);
```

#### Stage Submissions

```sql
CREATE TABLE user_stage_completions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    stage_id INTEGER NOT NULL REFERENCES primm_stages(id) ON DELETE CASCADE,

    -- PREDICT submissions
    predict_selected_answer VARCHAR(10),
    predict_is_correct BOOLEAN,

    -- RUN submissions
    run_submitted_code TEXT,
    run_output TEXT,

    -- INVESTIGATE submissions
    investigate_reflection TEXT,
    investigate_completed BOOLEAN,

    -- MODIFY submissions
    modify_submitted_code TEXT,
    modify_output TEXT,
    modify_is_correct BOOLEAN,
    modify_attempts INTEGER DEFAULT 0,

    -- MAKE submissions
    make_submitted_code TEXT,
    make_output TEXT,
    make_is_correct BOOLEAN,
    make_attempts INTEGER DEFAULT 0,

    is_completed BOOLEAN DEFAULT FALSE,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(user_id, stage_id)
);
```

#### Course Completion

```sql
CREATE TABLE user_course_completions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    is_completed BOOLEAN DEFAULT FALSE,
    completed_at TIMESTAMP,
    coins_earned INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, course_id)
);
```

---

## 🔌 4. API ENDPOINTS

### 4.1 Authentication Endpoints

#### POST /api/register

Register new user (teacher or student)

**Request:**

```json
{
  "email": "teacher@example.com",
  "password": "Teacher123!",
  "full_name": "John Teacher",
  "role": "teacher"
}
```

**Response (201 Created):**

```json
{
  "message": "User berhasil didaftarkan!",
  "user": {
    "id": 7,
    "email": "teacher@example.com",
    "full_name": "John Teacher",
    "role": "teacher",
    "level": 1,
    "total_coins": 0,
    "experience_points": 0,
    "created_at": "2025-10-25T14:25:38Z"
  }
}
```

**Validation Rules:**

- ✅ Email must be valid format
- ✅ Email must be unique
- ✅ Password minimum 6 characters
- ✅ Full name required
- ✅ Role must be 'teacher' or 'student'

---

#### POST /api/login

Login user and get JWT token

**Request:**

```json
{
  "email": "teacher@example.com",
  "password": "Teacher123!"
}
```

**Response (200 OK):**

```json
{
  "message": "Login berhasil!",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 7,
    "email": "teacher@example.com",
    "full_name": "John Teacher",
    "role": "teacher",
    "level": 1,
    "total_coins": 0,
    "experience_points": 0
  }
}
```

**Token Payload:**

```json
{
  "user_id": 7,
  "email": "teacher@example.com",
  "role": "teacher",
  "exp": 1729876800
}
```

**Error Responses:**

- `400 Bad Request` - Invalid email/password format
- `401 Unauthorized` - Wrong password
- `404 Not Found` - Email not found

---

### 4.2 Lesson Endpoints (Teacher Only)

#### POST /api/lessons

Create new lesson

**Headers:**

```
Authorization: Bearer <TOKEN>
Content-Type: application/json
```

**Request:**

```json
{
  "title": "Python Basics",
  "description": "Learn fundamental Python programming concepts from scratch",
  "category": "python",
  "difficulty": "beginner",
  "thumbnail_url": "https://example.com/python-basics.jpg"
}
```

**Response (201 Created):**

```json
{
  "message": "Lesson berhasil dibuat!",
  "lesson": {
    "id": 2,
    "teacher_id": 7,
    "title": "Python Basics",
    "description": "Learn fundamental Python programming concepts from scratch",
    "category": "python",
    "difficulty": "beginner",
    "thumbnail_url": "https://example.com/python-basics.jpg",
    "is_active": true,
    "created_at": "2025-10-25T14:35:40Z",
    "updated_at": "2025-10-25T14:35:40Z"
  }
}
```

**Validation Rules:**

- ✅ Title: 3-200 characters
- ✅ Description: minimum 10 characters
- ✅ Category: must be one of ['python', 'javascript', 'html', 'c']
- ✅ Difficulty: must be one of ['beginner', 'intermediate', 'advanced']
- ✅ Thumbnail URL: valid URL format (optional)

**Authorization:**

- ❌ 401 if no token
- ❌ 403 if role is not 'teacher'

---

### 4.3 Course Endpoints (Teacher Only)

#### POST /api/courses

Create new course in a lesson

**Headers:**

```
Authorization: Bearer <TOKEN>
Content-Type: application/json
```

**Request:**

```json
{
  "lesson_id": 2,
  "title": "Variables and Data Types",
  "description": "Understanding Python variables, integers, strings, and basic data types",
  "order_index": 1,
  "coin_reward": 200
}
```

**Response (201 Created):**

```json
{
  "message": "Course berhasil dibuat!",
  "course": {
    "id": 1,
    "lesson_id": 2,
    "title": "Variables and Data Types",
    "description": "Understanding Python variables, integers, strings, and basic data types",
    "order_index": 1,
    "coin_reward": 200,
    "is_active": true,
    "created_at": "2025-10-25T14:37:20Z",
    "updated_at": "2025-10-25T14:37:20Z"
  }
}
```

**Validation Rules:**

- ✅ Lesson ID: must exist
- ✅ Teacher must own the lesson
- ✅ Title: 3-200 characters
- ✅ Description: minimum 10 characters
- ✅ Order Index: minimum 1
- ✅ Coin Reward: 10-1000

**Authorization:**

- ❌ 401 if no token
- ❌ 403 if not the lesson owner

---

### 4.4 PRIMM Stage Endpoints (Teacher Only)

#### POST /api/stages/predict

Create PREDICT stage (multiple choice prediction)

**Headers:**

```
Authorization: Bearer <TOKEN>
Content-Type: application/json
```

**Request:**

```json
{
  "course_id": 1,
  "title": "Predict: Variable Assignment Output",
  "description": "What will be the output of this code?",
  "code_snippet": "x = 5\ny = 10\nresult = x + y\nprint(result)",
  "predict_options": {
    "A": "5",
    "B": "10",
    "C": "15",
    "D": "510"
  },
  "correct_answer": "C"
}
```

**Response (201 Created):**

```json
{
  "message": "PREDICT stage berhasil dibuat!",
  "stage": {
    "id": 1,
    "course_id": 1,
    "stage_type": "predict",
    "order_index": 1,
    "title": "Predict: Variable Assignment Output",
    "description": "What will be the output of this code?",
    "code_snippet": "x = 5\ny = 10\nresult = x + y\nprint(result)",
    "predict_options": {
      "A": "5",
      "B": "10",
      "C": "15",
      "D": "510"
    },
    "correct_answer": "C",
    "is_active": true,
    "created_at": "2025-10-25T14:38:15Z",
    "updated_at": "2025-10-25T14:38:15Z"
  }
}
```

**Validation Rules:**

- ✅ Course ID: must exist
- ✅ Teacher must own the course's lesson
- ✅ Title & Description: required
- ✅ Code Snippet: required
- ✅ Predict Options: must have at least 2 options
- ✅ Correct Answer: must match one of the option keys

---

#### POST /api/stages/run

Create RUN stage (execute & observe code)

**Request:**

```json
{
  "course_id": 1,
  "title": "Run: Execute Variable Code",
  "description": "Run the code and observe the output",
  "code_snippet": "name = \"Alice\"\nage = 25\nprint(f\"My name is {name} and I am {age} years old\")",
  "run_code_template": "# Optional: editable template for students"
}
```

**Response (201 Created):**

```json
{
  "message": "RUN stage berhasil dibuat!",
  "stage": {
    "id": 2,
    "course_id": 1,
    "stage_type": "run",
    "order_index": 2,
    "title": "Run: Execute Variable Code",
    "description": "Run the code and observe the output",
    "code_snippet": "name = \"Alice\"...",
    "run_code_template": null,
    "is_active": true,
    "created_at": "2025-10-25T14:39:49Z"
  }
}
```

**Validation Rules:**

- ✅ Code Snippet: required
- ✅ Run Code Template: optional (for editable version)

---

#### POST /api/stages/investigate

Create INVESTIGATE stage (watch video & reflect)

**Request:**

```json
{
  "course_id": 1,
  "title": "Investigate: Understanding Variables",
  "description": "Watch the video and reflect on how variables work in Python",
  "video_url": "https://www.youtube.com/watch?v=example",
  "explanation_text": "Variables are containers for storing data values...",
  "guiding_questions": ["What is a variable in Python?", "How do you assign a value to a variable?", "What are the naming rules for variables?", "Can you change a variable's value after assignment?"],
  "reflection_prompt": "Write a short reflection about how variables help make programs more flexible..."
}
```

**Response (201 Created):**

```json
{
  "message": "INVESTIGATE stage berhasil dibuat!",
  "stage": {
    "id": 3,
    "course_id": 1,
    "stage_type": "investigate",
    "order_index": 3,
    "title": "Investigate: Understanding Variables",
    "description": "Watch the video and reflect on how variables work in Python",
    "video_embed_url": "https://www.youtube.com/watch?v=example",
    "explanation_text": "Variables are containers...",
    "guiding_questions": [...],
    "reflection_prompt": "Write a short reflection...",
    "is_active": true,
    "created_at": "2025-10-25T14:45:56Z"
  }
}
```

**Validation Rules:**

- ✅ Video URL: required, must be valid URL
- ✅ Explanation Text: optional
- ✅ Guiding Questions: required, minimum 1 question
- ✅ Reflection Prompt: optional

---

#### POST /api/stages/modify

Create MODIFY stage (modify existing code)

**Request:**

```json
{
  "course_id": 1,
  "title": "Modify: Calculate Sum of Three Numbers",
  "description": "Modify the existing code to work with three variables instead of two",
  "code_snippet": "x = 5\ny = 10\nresult = x + y\nprint(result)",
  "task_description": "Add a third variable 'z' with value 15. Modify the code to calculate the sum of x, y, and z.",
  "modify_challenge": "Can you modify the code to sum three numbers?",
  "modify_code_template": "x = 5\ny = 10\n# TODO: Add variable z here\nresult = x + y\nprint(result)",
  "modify_expected_output": "30",
  "modify_test_cases": [
    {
      "input": "",
      "expected_output": "30"
    }
  ]
}
```

**Response (201 Created):**

```json
{
  "message": "MODIFY stage berhasil dibuat!",
  "stage": {
    "id": 4,
    "course_id": 1,
    "stage_type": "modify",
    "order_index": 4,
    "title": "Modify: Calculate Sum of Three Numbers",
    "description": "Modify the existing code to work with three variables instead of two",
    "code_snippet": "x = 5\ny = 10\nresult = x + y\nprint(result)",
    "task_description": "Add a third variable 'z'...",
    "modify_test_cases": [...],
    "is_active": true,
    "created_at": "2025-10-25T17:09:49Z"
  }
}
```

**Validation Rules:**

- ✅ Code Snippet: required
- ✅ Task Description: required
- ✅ Modify Challenge: optional
- ✅ Modify Code Template: optional
- ✅ Modify Expected Output: optional
- ✅ Modify Test Cases: required, minimum 1 test case

**Test Case Format:**

```json
{
  "input": "", // Input untuk program (optional)
  "expected_output": "30" // Expected output dari program
}
```

---

#### POST /api/stages/make

Create MAKE stage (create code from scratch)

**Request:**

```json
{
  "course_id": 1,
  "title": "Make: Create a Greeting Program",
  "description": "Write a complete program from scratch that greets users with their name and age",
  "task_description": "Create a program that:\n1. Uses input() to ask for the user's name\n2. Uses input() to ask for the user's age\n3. Prints: 'Hello [name], you are [age] years old!'\n\nExample:\nInput: Alice\nInput: 25\nOutput: Hello Alice, you are 25 years old!",
  "make_challenge": "Build a program that asks for user input and creates a personalized greeting",
  "make_expected_output": "Hello Alice, you are 25 years old!",
  "make_test_cases": [
    {
      "input": "Alice\n25",
      "expected_output": "Hello Alice, you are 25 years old!"
    },
    {
      "input": "Bob\n30",
      "expected_output": "Hello Bob, you are 30 years old!"
    },
    {
      "input": "Charlie\n18",
      "expected_output": "Hello Charlie, you are 18 years old!"
    }
  ]
}
```

**Response (201 Created):**

```json
{
  "message": "MAKE stage berhasil dibuat!",
  "stage": {
    "id": 5,
    "course_id": 1,
    "stage_type": "make",
    "order_index": 5,
    "title": "Make: Create a Greeting Program",
    "description": "Write a complete program from scratch that greets users with their name and age",
    "task_description": "Create a program that:\n1. Uses input()...",
    "make_test_cases": [...],
    "is_active": true,
    "created_at": "2025-10-25T19:00:07Z"
  }
}
```

**Validation Rules:**

- ✅ Task Description: required
- ✅ Make Challenge: optional
- ✅ Make Expected Output: optional
- ✅ Make Test Cases: required, minimum 1 test case

**Test Case Format:**

```json
{
  "input": "Alice\n25", // Multi-line input (separated by \n)
  "expected_output": "Hello Alice, you are 25 years old!" // Expected output
}
```

---

## 🧪 5. TEST FLOW

### Complete Test Sequence (Test 1-9)

```
┌─────────────────────────────────────────────────────────┐
│  TEST 1: REGISTER TEACHER                               │
│  POST /api/register                                     │
│  • Create teacher account                               │
│  • Get user_id: 7                                       │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 2: LOGIN TEACHER                                  │
│  POST /api/login                                        │
│  • Authenticate teacher                                 │
│  • Get JWT token (valid 7 days)                         │
│  • Token contains: user_id, email, role                 │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 3: CREATE LESSON                                  │
│  POST /api/lessons                                      │
│  • Create "Python Basics" lesson                        │
│  • Get lesson_id: 2                                     │
│  • Teacher owns this lesson                             │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 4: CREATE COURSE                                  │
│  POST /api/courses                                      │
│  • Create "Variables and Data Types" course             │
│  • Course belongs to lesson_id: 2                       │
│  • Get course_id: 1                                     │
│  • Coin reward: 200                                     │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 5: CREATE PREDICT STAGE                           │
│  POST /api/stages/predict                               │
│  • Stage 1 of 5 (order_index: 1)                        │
│  • Multiple choice prediction                           │
│  • Get stage_id: 1                                      │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 6: CREATE RUN STAGE                               │
│  POST /api/stages/run                                   │
│  • Stage 2 of 5 (order_index: 2)                        │
│  • Execute & observe code                               │
│  • Get stage_id: 2                                      │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 7: CREATE INVESTIGATE STAGE                       │
│  POST /api/stages/investigate                           │
│  • Stage 3 of 5 (order_index: 3)                        │
│  • Watch video & reflect                                │
│  • Get stage_id: 3                                      │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 8: CREATE MODIFY STAGE                            │
│  POST /api/stages/modify                                │
│  • Stage 4 of 5 (order_index: 4)                        │
│  • Modify existing code                                 │
│  • Get stage_id: 4                                      │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 9: CREATE MAKE STAGE                              │
│  POST /api/stages/make                                  │
│  • Stage 5 of 5 (order_index: 5)                        │
│  • Create code from scratch                             │
│  • Get stage_id: 5                                      │
│  ✅ COURSE COMPLETE (all 5 stages created)              │
└─────────────────────────────────────────────────────────┘
```

---

### Data Created After Test 1-9

```
User (id: 7) "John Teacher"
  └── Lesson (id: 2) "Python Basics"
       └── Course (id: 1) "Variables and Data Types" [200 coins]
            ├── Stage 1 [PREDICT]   - "Variable Assignment Output"
            ├── Stage 2 [RUN]       - "Execute Variable Code"
            ├── Stage 3 [INVESTIGATE] - "Understanding Variables"
            ├── Stage 4 [MODIFY]    - "Calculate Sum of Three Numbers"
            └── Stage 5 [MAKE]      - "Create a Greeting Program"
```

---

## 📂 6. CODE STRUCTURE

### 6.1 Models (Data Structures)

#### models/user.go

```go
// User model
type User struct {
    ID               int       `json:"id"`
    Email            string    `json:"email"`
    PasswordHash     string    `json:"-"` // Hidden from JSON
    FullName         string    `json:"full_name"`
    Role             string    `json:"role"`
    Level            int       `json:"level"`
    TotalCoins       int       `json:"total_coins"`
    ExperiencePoints int       `json:"experience_points"`
    ProfilePicture   *string   `json:"profile_picture"`
    Bio              *string   `json:"bio"`
    CreatedAt        time.Time `json:"created_at"`
    UpdatedAt        time.Time `json:"updated_at"`
}

// Request models
type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
    FullName string `json:"full_name" binding:"required"`
    Role     string `json:"role" binding:"required,oneof=teacher student"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// Response models
type LoginResponse struct {
    Message string `json:"message"`
    Token   string `json:"token"`
    User    *User  `json:"user"`
}
```

---

### 6.2 Services (Business Logic)

#### services/auth.go

```go
// Register creates new user
func Register(db *pgx.Conn, req models.RegisterRequest) (*models.User, error) {
    // 1. Hash password with bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

    // 2. Check if email already exists
    var existingID int
    err = db.QueryRow(context.Background(),
        "SELECT id FROM users WHERE email = $1", req.Email).Scan(&existingID)

    if err == nil {
        return nil, errors.New("email sudah terdaftar")
    }

    // 3. Insert user to database
    var user models.User
    err = db.QueryRow(context.Background(), `
        INSERT INTO users (email, password_hash, full_name, role, level, total_coins, experience_points)
        VALUES ($1, $2, $3, $4, 1, 0, 0)
        RETURNING id, email, full_name, role, level, total_coins, experience_points,
                  profile_picture, bio, created_at, updated_at
    `, req.Email, string(hashedPassword), req.FullName, req.Role).Scan(...)

    return &user, nil
}

// Login authenticates user and returns JWT token
func Login(db *pgx.Conn, req models.LoginRequest) (*models.LoginResponse, error) {
    // 1. Get user from database
    var user models.User
    var passwordHash string
    err := db.QueryRow(context.Background(), `
        SELECT id, email, password_hash, full_name, role, level, total_coins, experience_points,
               profile_picture, bio, created_at, updated_at
        FROM users WHERE email = $1
    `, req.Email).Scan(...)

    if err != nil {
        return nil, errors.New("email atau password salah")
    }

    // 2. Verify password
    if !CheckPasswordHash(req.Password, passwordHash) {
        return nil, errors.New("email atau password salah")
    }

    // 3. Generate JWT token
    token, err := GenerateJWT(user.ID, user.Email, user.Role)
    if err != nil {
        return nil, err
    }

    return &models.LoginResponse{
        Message: "Login berhasil!",
        Token:   token,
        User:    &user,
    }, nil
}

// GenerateJWT creates JWT token
func GenerateJWT(userID int, email string, role string) (string, error) {
    secretKey := os.Getenv("JWT_SECRET")

    claims := jwt.MapClaims{
        "user_id": userID,
        "email":   email,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secretKey))

    return tokenString, err
}
```

---

### 6.3 Handlers (HTTP Controllers)

#### handlers/auth_handler.go

```go
type AuthHandler struct {
    DB *pgx.Conn
}

// Register handler for POST /api/register
func (h *AuthHandler) Register(c *gin.Context) {
    var req models.RegisterRequest

    // 1. Bind & validate JSON
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    // 2. Call service
    user, err := services.Register(h.DB, req)
    if err != nil {
        if err.Error() == "email sudah terdaftar" {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    // 3. Return response
    c.JSON(http.StatusCreated, gin.H{
        "message": "User berhasil didaftarkan!",
        "user":    user,
    })
}

// Login handler for POST /api/login
func (h *AuthHandler) Login(c *gin.Context) {
    var req models.LoginRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    response, err := services.Login(h.DB, req)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}
```

---

### 6.4 Middleware (Authentication & Authorization)

#### middleware/middleware.go

```go
// AuthMiddleware validates JWT token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. Get token from Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        // 2. Parse token (format: Bearer <token>)
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
            c.Abort()
            return
        }

        tokenString := parts[1]

        // 3. Validate token
        claims, err := services.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
            c.Abort()
            return
        }

        // 4. Extract claims and set to context
        userID := int(claims["user_id"].(float64))
        userRole := claims["role"].(string)

        c.Set("user_id", userID)
        c.Set("user_role", userRole)

        c.Next()
    }
}

// RequireRole validates user role
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("user_role")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found"})
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

        c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
        c.Abort()
    }
}
```

---

## 🔧 7. TROUBLESHOOTING

### Common Errors & Solutions

#### Error 1: Unused Imports

```
error: "context" imported and not used
```

**Solution:**

- Hapus import yang tidak dipakai
- Go sangat strict terhadap unused imports

---

#### Error 2: Field Not Found

```
error: req.MakeHints undefined (type models.CreateMakeStageRequest has no field or method MakeHints)
```

**Solution:**

- Tambahkan field ke struct model
- Atau hapus referensi ke field tersebut dari service

---

#### Error 3: Database Column Not Found

```
error: ERROR: column "task_description" of relation "primm_stages" does not exist (SQLSTATE 42703)
```

**Solution:**

```sql
ALTER TABLE primm_stages
ADD COLUMN IF NOT EXISTS task_description TEXT;
```

---

#### Error 4: Validation Failed

```
error: Key: 'CreateLessonRequest.Category' Error:Field validation for 'Category' failed on the 'oneof' tag
```

**Solution:**

- Check validation tag di model
- Gunakan value yang sesuai dengan `oneof` list
- Example: `category: "python"` (bukan "programming")

---

#### Error 5: Authorization Failed

```
error: anda tidak memiliki akses untuk membuat stage di course ini
```

**Solution:**

- Teacher hanya bisa create stage di course milik lesson yang dia buat
- Check `teacher_id` ownership chain: Stage → Course → Lesson → Teacher

---

#### Error 6: Unique Constraint Violation

```
error: ERROR: duplicate key value violates unique constraint "primm_stages_course_id_order_index_key"
```

**Solution:**

- Order index sudah ada untuk course tersebut
- System auto-increment order_index, tidak perlu manual
- Check apakah ada bug di auto-increment logic

---

### Debugging Tips

#### 1. Check Token Validity

```bash
# Decode JWT token (online: jwt.io)
echo "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." | base64 -d
```

#### 2. Check Database State

```sql
-- Check user
SELECT * FROM users WHERE email = 'teacher@example.com';

-- Check lesson ownership
SELECT l.id, l.title, l.teacher_id, u.full_name
FROM lessons l
JOIN users u ON l.teacher_id = u.id;

-- Check course with stages
SELECT
    c.id AS course_id,
    c.title AS course_title,
    ps.id AS stage_id,
    ps.stage_type,
    ps.title AS stage_title,
    ps.order_index
FROM courses c
LEFT JOIN primm_stages ps ON c.id = ps.course_id
WHERE c.id = 1
ORDER BY ps.order_index;
```

#### 3. Test with cURL

```bash
# Register
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"email":"teacher@example.com","password":"Teacher123!","full_name":"John Teacher","role":"teacher"}'

# Login
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"teacher@example.com","password":"Teacher123!"}'

# Create Lesson (with token)
curl -X POST http://localhost:8080/api/lessons \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{"title":"Python Basics","description":"Learn Python","category":"python","difficulty":"beginner"}'
```

---

## 📝 SUMMARY

### What Teacher Flow Does:

1. ✅ **Authentication** - Register & Login dengan JWT
2. ✅ **Content Creation** - Buat Lesson, Course, dan 5 PRIMM Stages
3. ✅ **Authorization** - Teacher hanya bisa edit content miliknya
4. ✅ **Validation** - Input validation di semua endpoints
5. ✅ **Database Integrity** - Foreign key constraints, unique constraints

### Key Design Decisions:

1. **Email-based login** (no username)
2. **JWT authentication** (7 days expiry)
3. **Role-based authorization** (teacher/student)
4. **Single table for PRIMM stages** (polymorphic design)
5. **Auto-increment order_index** (no manual ordering)
6. **Ownership validation** (teacher can only edit own content)

### Next Steps:

- ✅ Teacher Flow Complete (Test 1-9)
- 🔜 Student Flow (Test 10-24)
- 🔜 Progress Tracking
- 🔜 Gamification (coins, XP, levels)
- 🔜 Code Execution Engine

---

## 📚 Additional Resources

- **Go Gin Documentation:** https://gin-gonic.com/docs/
- **pgx Documentation:** https://pkg.go.dev/github.com/jackc/pgx/v5
- **JWT Best Practices:** https://jwt.io/introduction
- **PRIMM Methodology:** https://primmportal.com/

---

**Last Updated:** October 25, 2025  
**Version:** 1.0  
**Status:** Teacher Flow Complete ✅
