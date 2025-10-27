# 🎓 STUDENT FLOW - Complete Documentation

> **Purpose:** Dokumentasi lengkap untuk Student Flow (Test 10-24) - learning journey dari enrollment sampai course completion dengan reward system.

**Last Updated:** October 25, 2025  
**Version:** 1.0  
**Status:** Student Flow Complete ✅

---

## 📋 Table of Contents

1. [Overview](#1-overview)
2. [Student Learning Journey](#2-student-learning-journey)
3. [API Endpoints](#3-api-endpoints)
4. [PRIMM Methodology Implementation](#4-primm-methodology-implementation)
5. [Submission & Grading System](#5-submission--grading-system)
6. [Progress Tracking & Rewards](#6-progress-tracking--rewards)
7. [Database Schema](#7-database-schema)
8. [Test Flow](#8-test-flow)
9. [Code Structure](#9-code-structure)
10. [Troubleshooting](#10-troubleshooting)

---

## 🎯 1. OVERVIEW

### Student Flow Purpose

Student Flow memungkinkan student untuk:

- ✅ Register & Login ke platform
- ✅ Browse available lessons
- ✅ Enroll ke lessons yang diminati
- ✅ View enrolled lessons dengan progress tracking
- ✅ Access courses dan PRIMM stages
- ✅ Submit answers untuk setiap stage
- ✅ Receive automatic grading & feedback
- ✅ Earn coins upon course completion
- ✅ Track learning progress

### Learning Philosophy: PRIMM Methodology

**PRIMM** = **P**redict → **R**un → **I**nvestigate → **M**odify → **M**ake

Scaffolded learning approach yang membantu students:

1. **Predict:** Develop code reading & prediction skills
2. **Run:** Observe actual code behavior
3. **Investigate:** Deep dive into concepts
4. **Modify:** Apply understanding through modification
5. **Make:** Create solutions from scratch

---

## 🎓 2. STUDENT LEARNING JOURNEY

### Complete Learning Flow

```
┌─────────────────────────────────────────────────────────┐
│  PHASE 1: REGISTRATION & ENROLLMENT                     │
│  • Register student account                              │
│  • Login & get JWT token                                 │
│  • Browse available lessons                              │
│  • Enroll to desired lesson                              │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  PHASE 2: VIEW LEARNING CONTENT                         │
│  • View enrolled lessons with progress                   │
│  • View courses in enrolled lesson                       │
│  • View all PRIMM stages in course                       │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  PHASE 3: COMPLETE PRIMM STAGES                         │
│  • Submit PREDICT (multiple choice)                      │
│  • Submit RUN (code execution)                           │
│  • Submit INVESTIGATE (reflection essay)                 │
│  • Submit MODIFY (code modification)                     │
│  • Submit MAKE (create from scratch)                     │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  PHASE 4: REWARDS & PROGRESS                            │
│  • Receive automatic coin reward                         │
│  • Course marked as completed                            │
│  • Progress updated to 100%                              │
│  • Ready for next course                                 │
└─────────────────────────────────────────────────────────┘
```

---

## 🔌 3. API ENDPOINTS

### 3.1 Authentication Endpoints

#### POST /api/register (Student)

Register new student account.

**Request:**

```json
{
  "email": "student@example.com",
  "password": "Student123!",
  "full_name": "Alice Student",
  "role": "student"
}
```

**Response (201 Created):**

```json
{
  "message": "User berhasil didaftarkan!",
  "user": {
    "id": 8,
    "email": "student@example.com",
    "full_name": "Alice Student",
    "role": "student",
    "level": 1,
    "total_coins": 0,
    "experience_points": 0,
    "created_at": "2025-10-25T19:51:55.556944+07:00"
  }
}
```

---

#### POST /api/login (Student)

Login student and get JWT token.

**Request:**

```json
{
  "email": "student@example.com",
  "password": "Student123!"
}
```

**Response (200 OK):**

```json
{
  "message": "Login berhasil!",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 8,
    "email": "student@example.com",
    "full_name": "Alice Student",
    "role": "student",
    "level": 1,
    "total_coins": 0,
    "experience_points": 0
  }
}
```

**Token Usage:**

```
Authorization: Bearer <TOKEN>
```

---

### 3.2 Lesson Browsing & Enrollment

#### GET /api/lessons

Browse all available lessons (public or authenticated).

**Headers:**

```
Authorization: Bearer <TOKEN> (optional)
```

**Response (200 OK):**

```json
{
  "count": 2,
  "lessons": [
    {
      "id": 2,
      "teacher_id": 7,
      "title": "Python Basics",
      "description": "Learn fundamental Python programming concepts from scratch",
      "category": "python",
      "difficulty": "beginner",
      "thumbnail_url": "https://example.com/python-basics.jpg",
      "is_active": true,
      "teacher_name": "John Teacher"
    }
  ]
}
```

---

#### POST /api/lessons/:id/enroll

Enroll student to a lesson.

**Headers:**

```
Authorization: Bearer <TOKEN>
```

**No Body Required**

**Response (201 Created):**

```json
{
  "message": "Berhasil enroll ke lesson!",
  "enrollment": {
    "id": 1,
    "user_id": 8,
    "lesson_id": 2,
    "enrolled_at": "2025-10-25T19:55:36.52239Z"
  }
}
```

**Business Rules:**

- ✅ Student can only enroll once per lesson
- ✅ Lesson must be active
- ✅ Duplicate enrollment returns error

---

### 3.3 View Learning Content

#### GET /api/my-lessons

Get all enrolled lessons with progress tracking.

**Headers:**

```
Authorization: Bearer <TOKEN>
```

**Response (200 OK):**

```json
{
  "count": 1,
  "lessons": [
    {
      "id": 2,
      "teacher_id": 7,
      "title": "Python Basics",
      "description": "Learn fundamental Python programming concepts from scratch",
      "category": "python",
      "difficulty": "beginner",
      "thumbnail_url": "https://example.com/python-basics.jpg",
      "is_active": true,
      "teacher_name": "John Teacher",
      "total_courses": 1,
      "completed_courses": 0,
      "progress_percent": 0,
      "is_enrolled": true
    }
  ]
}
```

**Progress Calculation:**

```
progress_percent = (completed_courses / total_courses) * 100
```

---

#### GET /api/lessons/:id/courses

Get all courses in a lesson.

**Headers:**

```
Authorization: Bearer <TOKEN>
```

**Response (200 OK):**

```json
{
  "courses": [
    {
      "id": 1,
      "lesson_id": 2,
      "title": "Variables and Data Types",
      "description": "Understanding Python variables, integers, strings, and basic data types",
      "order_index": 1,
      "coin_reward": 200,
      "is_active": true,
      "created_at": "2025-10-25T14:36:39.431253Z",
      "updated_at": "2025-10-25T14:36:39.431253Z"
    }
  ]
}
```

---

#### GET /api/courses/:id/stages

Get all PRIMM stages in a course.

**Headers:**

```
Authorization: Bearer <TOKEN>
```

**Response (200 OK):**

```json
{
  "stages": [
    {
      "id": 1,
      "course_id": 1,
      "stage_type": "predict",
      "title": "Predict: Variable Assignment Output",
      "description": "What will be the output of this code?",
      "order_index": 1,
      "code_snippet": "x = 5\ny = 10\nresult = x + y\nprint(result)",
      "predict_options": {
        "A": "5",
        "B": "10",
        "C": "15",
        "D": "510"
      },
      "correct_answer": "C"
    },
    {
      "id": 2,
      "course_id": 1,
      "stage_type": "run",
      "title": "Run: Execute Variable Code",
      "description": "Run the code and observe the output",
      "order_index": 2,
      "code_snippet": "name = \"Alice\"\nage = 25\nprint(f\"My name is {name} and I am {age} years old\")"
    },
    {
      "id": 3,
      "course_id": 1,
      "stage_type": "investigate",
      "title": "Investigate: Understanding Variables",
      "description": "Watch the video and reflect on how variables work in Python",
      "order_index": 3,
      "video_embed_url": "https://www.youtube.com/watch?v=example",
      "explanation_text": "Variables are containers for storing data values...",
      "reflection_prompt": "Write a short reflection about how variables help make programs more flexible..."
    },
    {
      "id": 4,
      "course_id": 1,
      "stage_type": "modify",
      "title": "Modify: Calculate Sum of Three Numbers",
      "description": "Modify the existing code to work with three variables instead of two",
      "order_index": 4,
      "code_snippet": "x = 5\ny = 10\nresult = x + y\nprint(result)",
      "task_description": "Add a third variable 'z' with value 15...",
      "modify_test_cases": [
        {
          "input": "",
          "expected_output": "30"
        }
      ]
    },
    {
      "id": 5,
      "course_id": 1,
      "stage_type": "make",
      "title": "Make: Create a Greeting Program",
      "description": "Write a complete program from scratch that greets users with their name and age",
      "order_index": 5,
      "task_description": "Create a program that:\n1. Uses input() to ask for the user's name\n2. Uses input() to ask for the user's age\n3. Prints: 'Hello [name], you are [age] years old!'",
      "make_test_cases": [
        {
          "input": "Alice\n25",
          "expected_output": "Hello Alice, you are 25 years old!"
        },
        {
          "input": "Bob\n30",
          "expected_output": "Hello Bob, you are 30 years old!"
        }
      ]
    }
  ]
}
```

---

### 3.4 Stage Submissions

#### POST /api/stages/:id/submit

Submit answer for a PRIMM stage.

**Headers:**

```
Authorization: Bearer <TOKEN>
Content-Type: application/json
```

**Request varies by stage type:**

##### PREDICT Stage Submission:

```json
{
  "submission_type": "predict",
  "selected_answer": "C"
}
```

##### RUN Stage Submission:

```json
{
  "submission_type": "run",
  "code_output": "15"
}
```

##### INVESTIGATE Stage Submission:

```json
{
  "submission_type": "investigate",
  "reflection_text": "After watching the video and reading the explanation, I now understand that variables are containers for storing data values in Python. Unlike other programming languages, Python doesn't require explicit type declaration - the type is automatically determined when you assign a value..."
}
```

##### MODIFY Stage Submission:

```json
{
  "submission_type": "modify",
  "modified_code": "x = 5\ny = 10\nz = 15\nresult = x + y + z\nprint(result)"
}
```

##### MAKE Stage Submission:

```json
{
  "submission_type": "make",
  "code": "name = input()\nage = input()\nprint(f'Hello {name}, you are {age} years old!')"
}
```

**Response (201 Created):**

```json
{
  "message": "Jawaban berhasil disubmit!",
  "is_correct": true,
  "score": 100,
  "submission": {
    "id": 1,
    "user_id": 8,
    "stage_id": 1,
    "submission_type": "predict",
    "submission_data": {
      "selected_answer": "C",
      "correct_answer": "C"
    },
    "is_correct": true,
    "score": 100,
    "submitted_at": "2025-10-25T21:01:02.356494Z"
  }
}
```

---

## 🎨 4. PRIMM METHODOLOGY IMPLEMENTATION

### Stage 1: PREDICT

**Learning Objective:** Develop code reading and prediction skills

**Format:** Multiple Choice Question

**Student Task:**

- Read provided code snippet
- Predict the output without running
- Select answer from options A, B, C, D

**Example:**

```python
Code: x = 5
      y = 10
      result = x + y
      print(result)

Options:
A) 5
B) 10
C) 15      ← Correct Answer
D) 510
```

**Grading:**

```go
isCorrect = (selectedAnswer == correctAnswer)
score = isCorrect ? 100 : 0
```

---

### Stage 2: RUN

**Learning Objective:** Observe actual code behavior

**Format:** Code Execution

**Student Task:**

- Run the provided code
- Observe the actual output
- Submit observed output

**Example:**

```python
Code: name = "Alice"
      age = 25
      print(f"My name is {name} and I am {age} years old")

Expected Output: My name is Alice and I am 25 years old
```

**Grading (Current):**

```go
// Auto-accept for observation stage
isCorrect = true
score = 100
```

**Future Enhancement:**

```go
// Validate actual output
isCorrect = (codeOutput == expectedOutput)
```

---

### Stage 3: INVESTIGATE

**Learning Objective:** Deep understanding through reflection

**Format:** Video + Essay Reflection

**Student Task:**

- Watch educational video
- Read explanation text
- Answer guiding questions (mentally)
- Write reflection essay

**Example:**

```
Video: "Understanding Python Variables"
URL: https://www.youtube.com/watch?v=example

Explanation:
"Variables are containers for storing data values. In Python,
you don't need to declare a variable type explicitly..."

Guiding Questions:
1. What is a variable in Python?
2. How do you assign a value to a variable?
3. What are the naming rules for variables?
4. Can you change a variable's value after assignment?

Reflection Prompt:
"Write a short reflection about how variables help make programs
more flexible and reusable. Include at least one example."
```

**Grading:**

```go
// Check minimum length (50 characters)
isCorrect = len(reflectionText) >= 50
score = isCorrect ? 100 : 0
```

**Future Enhancement:**

- AI-powered content analysis
- Check for key concepts
- Evaluate depth of understanding

---

### Stage 4: MODIFY

**Learning Objective:** Apply understanding through code modification

**Format:** Code Modification with Test Cases

**Student Task:**

- Read original code
- Understand modification requirements
- Modify code to meet requirements
- Submit modified code

**Example:**

```python
Original Code:
x = 5
y = 10
result = x + y
print(result)  # Output: 15

Task: Add a third variable 'z' with value 15 and calculate sum of all three

Modified Code:
x = 5
y = 10
z = 15
result = x + y + z
print(result)  # Output: 30

Test Cases:
Input: (none)
Expected Output: 30
```

**Grading (Current):**

```go
// Simple validation: code not empty
isCorrect = len(modifiedCode) > 0
score = isCorrect ? 100 : 0
```

**Future Enhancement:**

```go
// Run code in sandbox and validate output
output = runCodeInSandbox(modifiedCode)
isCorrect = validateTestCases(output, testCases)
score = calculateScore(testCases)
```

---

### Stage 5: MAKE

**Learning Objective:** Create solutions from scratch

**Format:** Code Creation with Multiple Test Cases

**Student Task:**

- Read challenge description
- Design solution
- Write complete code from scratch
- Submit for validation

**Example:**

```python
Challenge:
"Create a program that greets users with their name and age"

Requirements:
1. Use input() to ask for the user's name
2. Use input() to ask for the user's age
3. Print: "Hello [name], you are [age] years old!"

Student Solution:
name = input()
age = input()
print(f'Hello {name}, you are {age} years old!')

Test Cases:
Test 1: Input: "Alice\n25"  → Output: "Hello Alice, you are 25 years old!"
Test 2: Input: "Bob\n30"    → Output: "Hello Bob, you are 30 years old!"
Test 3: Input: "Charlie\n18" → Output: "Hello Charlie, you are 18 years old!"
```

**Grading (Current):**

```go
// Simple validation: code not empty
isCorrect = len(code) > 0
score = isCorrect ? 100 : 0
```

**Future Enhancement:**

```go
// Run code with test cases and validate all outputs
testResults = runMultipleTestCases(code, testCases)
isCorrect = allTestsPassed(testResults)
score = calculateScore(testResults)  // Partial credit possible
```

---

## ✅ 5. SUBMISSION & GRADING SYSTEM

### Submission Flow

```
Student Submit Answer
    ↓
Validate Submission Type
    ↓
Extract Stage Details from Database
    ↓
Grade Based on Stage Type
    ↓
Save to stage_submissions Table
    ↓
Check Course Completion (Background)
    ↓
Award Coins if All Stages Complete
```

---

### Grading Logic

**File: `services/primm_stage_service.go` → `SubmitStage` function**

```go
func SubmitStage(db *pgx.Conn, userID int, stageID int, req models.StageSubmissionRequest) (*models.StageSubmission, error) {
    // 1. Get stage details
    var stage models.PRIMMStage
    // ... fetch stage from database

    // 2. Validate submission type matches stage type
    if req.SubmissionType != stage.StageType {
        return nil, errors.New("tipe submission tidak sesuai dengan tipe stage")
    }

    // 3. Grade based on stage type
    var isCorrect bool
    var score int

    switch stage.StageType {
    case "predict":
        if stage.CorrectAnswer != nil {
            isCorrect = req.SelectedAnswer == *stage.CorrectAnswer
            if isCorrect {
                score = 100
            }
        }

    case "run":
        // Auto-accept for observation stage
        isCorrect = true
        score = 100

    case "investigate":
        // Check minimum length (50 characters)
        isCorrect = len(req.ReflectionText) >= 50
        if isCorrect {
            score = 100
        }

    case "modify":
        // Check code not empty
        isCorrect = len(req.ModifiedCode) > 0
        if isCorrect {
            score = 100
        }

    case "make":
        // Check code not empty
        isCorrect = len(req.Code) > 0
        if isCorrect {
            score = 100
        }
    }

    // 4. Save submission
    // ... insert into stage_submissions

    // 5. Check course completion (async)
    go checkCourseCompletion(db, userID, stage.CourseID)

    return &submission, nil
}
```

---

### Submission Data Structure

**Database Column:** `submission_data` (JSONB)

**Per Stage Type:**

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
  "reflection_text": "After watching the video..."
}
```

**MODIFY:**

```json
{
  "modified_code": "x = 5\ny = 10\nz = 15\nresult = x + y + z\nprint(result)"
}
```

**MAKE:**

```json
{
  "code": "name = input()\nage = input()\nprint(f'Hello {name}, you are {age} years old!')"
}
```

---

## 🏆 6. PROGRESS TRACKING & REWARDS

### Course Completion Logic

**Function:** `checkCourseCompletion` (runs asynchronously after each submission)

```go
func checkCourseCompletion(db *pgx.Conn, userID int, courseID int) {
    // 1. Count total stages and completed stages
    var totalStages, completedStages int
    db.QueryRow(context.Background(), `
        SELECT
            COUNT(DISTINCT ps.id) as total_stages,
            COUNT(DISTINCT ss.stage_id) FILTER (WHERE ss.is_correct = true) as completed_stages
        FROM primm_stages ps
        LEFT JOIN stage_submissions ss ON ps.id = ss.stage_id AND ss.user_id = $1
        WHERE ps.course_id = $2
    `, userID, courseID).Scan(&totalStages, &completedStages)

    // 2. Check if all stages completed
    if totalStages > 0 && totalStages == completedStages {
        // 3. Get coin reward
        var coinReward int
        db.QueryRow(context.Background(), `
            SELECT coin_reward FROM courses WHERE id = $1
        `, courseID).Scan(&coinReward)

        // 4. Insert course completion record
        db.Exec(context.Background(), `
            INSERT INTO user_course_completion (user_id, course_id, coins_awarded)
            VALUES ($1, $2, $3)
            ON CONFLICT (user_id, course_id) DO NOTHING
        `, userID, courseID, coinReward)

        // 5. Award coins to user
        db.Exec(context.Background(), `
            UPDATE users SET coins = coins + $1 WHERE id = $2
        `, coinReward, userID)
    }
}
```

---

### Reward System

**Completion Criteria:**

```
Course is completed when:
✅ All 5 PRIMM stages submitted
✅ All 5 submissions marked as correct (is_correct = true)
```

**Automatic Actions on Completion:**

1. **Insert Record to `user_course_completion`:**

   ```sql
   INSERT INTO user_course_completion (user_id, course_id, coins_awarded)
   VALUES (8, 1, 200)
   ON CONFLICT (user_id, course_id) DO NOTHING
   ```

2. **Award Coins to User:**
   ```sql
   UPDATE users
   SET coins = coins + 200
   WHERE id = 8
   ```

**Result:**

- ✅ Student's `coins` column updated: 0 → 200
- ✅ Course marked as completed
- ✅ Can view completion status in progress endpoints

---

### Progress Calculation

**Lesson Level Progress:**

```sql
SELECT
    COUNT(DISTINCT c.id) AS total_courses,
    COUNT(DISTINCT ucc.course_id) AS completed_courses
FROM courses c
LEFT JOIN user_course_completion ucc ON c.id = ucc.course_id AND ucc.user_id = ?
WHERE c.lesson_id = ?

progress_percent = (completed_courses / total_courses) * 100
```

**Course Level Progress:**

```sql
SELECT
    COUNT(DISTINCT ps.id) AS total_stages,
    COUNT(DISTINCT ss.stage_id) FILTER (WHERE ss.is_correct = true) AS completed_stages
FROM primm_stages ps
LEFT JOIN stage_submissions ss ON ps.id = ss.stage_id AND ss.user_id = ?
WHERE ps.course_id = ?

progress_percent = (completed_stages / total_stages) * 100
```

---

## 🗄️ 7. DATABASE SCHEMA

### 7.1 stage_submissions Table

```sql
CREATE TABLE stage_submissions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    stage_id INT NOT NULL REFERENCES primm_stages(id) ON DELETE CASCADE,
    submission_type VARCHAR(50) NOT NULL,
    submission_data JSONB NOT NULL,
    is_correct BOOLEAN DEFAULT false,
    score INT DEFAULT 0,
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, stage_id)
);

CREATE INDEX idx_stage_submissions_user ON stage_submissions(user_id);
CREATE INDEX idx_stage_submissions_stage ON stage_submissions(stage_id);
```

**Key Points:**

- `UNIQUE(user_id, stage_id)` → Student can only submit once per stage (can update)
- `submission_data` → JSONB for flexible data per stage type
- `is_correct` → Boolean for pass/fail
- `score` → Integer 0-100 for scoring

---

### 7.2 user_course_completion Table

```sql
CREATE TABLE user_course_completion (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id INT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    coins_awarded INT DEFAULT 0,
    UNIQUE(user_id, course_id)
);

CREATE INDEX idx_course_completion_user ON user_course_completion(user_id);
CREATE INDEX idx_course_completion_course ON user_course_completion(course_id);
```

**Key Points:**

- `UNIQUE(user_id, course_id)` → Course can only be completed once
- `coins_awarded` → Track coins given for this completion
- `completed_at` → Timestamp of completion

---

### 7.3 user_lessons Table

```sql
CREATE TABLE user_lessons (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_id INT NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, lesson_id)
);
```

**Key Points:**

- `UNIQUE(user_id, lesson_id)` → Student can only enroll once per lesson
- Tracks enrollment relationship

**TODO for Future:**

- Add `last_accessed_at` column for tracking activity

---

## 🧪 8. TEST FLOW

### Complete Test Sequence (Test 10-24)

```
┌─────────────────────────────────────────────────────────┐
│  TEST 10: REGISTER STUDENT                              │
│  POST /api/register                                     │
│  • Create student account (student_id: 8)               │
│  • Initial coins: 0                                     │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 11: LOGIN STUDENT                                 │
│  POST /api/login                                        │
│  • Get JWT token                                        │
│  • Token valid for 7 days                               │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 12: BROWSE LESSONS                                │
│  GET /api/lessons                                       │
│  • View all available lessons                           │
│  • See lesson details (title, description, difficulty)  │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 13: ENROLL TO LESSON                              │
│  POST /api/lessons/2/enroll                             │
│  • Enroll to "Python Basics" lesson                     │
│  • Get enrollment_id: 1                                 │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 14: VIEW ENROLLED LESSONS                         │
│  GET /api/my-lessons                                    │
│  • See enrolled lessons with progress                   │
│  • Progress: 0% (not started)                           │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 15: VIEW COURSES IN LESSON                        │
│  GET /api/lessons/2/courses                             │
│  • See course: "Variables and Data Types"               │
│  • Coin reward: 200                                     │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 16: VIEW ALL PRIMM STAGES                         │
│  GET /api/courses/1/stages                              │
│  • View 5 stages: PREDICT, RUN, INVESTIGATE, MODIFY, MAKE │
│  • See details for each stage                           │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 17: SUBMIT PREDICT STAGE                          │
│  POST /api/stages/1/submit                              │
│  • Submit answer: "C"                                   │
│  • Result: Correct! Score: 100                          │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 18: SUBMIT RUN STAGE                              │
│  POST /api/stages/2/submit                              │
│  • Submit output: "15"                                  │
│  • Result: Correct! Score: 100                          │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 19: SUBMIT INVESTIGATE STAGE                      │
│  POST /api/stages/3/submit                              │
│  • Submit reflection essay (571 chars)                  │
│  • Result: Correct! Score: 100                          │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 20: SUBMIT MODIFY STAGE                           │
│  POST /api/stages/4/submit                              │
│  • Submit modified code (added variable z)              │
│  • Result: Correct! Score: 100                          │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 21: SUBMIT MAKE STAGE                             │
│  POST /api/stages/5/submit                              │
│  • Submit greeting program from scratch                 │
│  • Result: Correct! Score: 100                          │
│  ✅ ALL 5 STAGES COMPLETED!                             │
│  ✅ Course marked as complete                           │
│  ✅ 200 coins automatically awarded                     │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 22: VERIFY COIN REWARD                            │
│  GET /api/profile                                       │
│  • Check coins: 0 → 200 ✅                              │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 23: VERIFY LESSON PROGRESS                        │
│  GET /api/my-lessons                                    │
│  • Progress: 0% → 100% ✅                               │
│  • Completed courses: 0 → 1 ✅                          │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  TEST 24: VERIFY COURSE COMPLETION                      │
│  GET /api/lessons/2/courses                             │
│  • Course marked as completed ✅                        │
│  • Ready for next course                                │
└─────────────────────────────────────────────────────────┘
```

---

### Data Created After Test 10-24

```
Student (id: 8) "Alice Student"
├── Coins: 200 🪙
├── Enrolled Lessons: 1
│   └── Lesson (id: 2) "Python Basics"
│       └── Course (id: 1) "Variables and Data Types" [COMPLETED ✅]
│            ├── Stage 1 [PREDICT]     - Submitted ✅ (Score: 100)
│            ├── Stage 2 [RUN]         - Submitted ✅ (Score: 100)
│            ├── Stage 3 [INVESTIGATE] - Submitted ✅ (Score: 100)
│            ├── Stage 4 [MODIFY]      - Submitted ✅ (Score: 100)
│            └── Stage 5 [MAKE]        - Submitted ✅ (Score: 100)
└── Progress: 100% (5/5 stages completed)
```

---

## 📂 9. CODE STRUCTURE

### 9.1 Handlers (HTTP Controllers)

#### handlers/primm_stage_handler.go

**Method: `SubmitStage`**

```go
func (h *PRIMMStageHandler) SubmitStage(c *gin.Context) {
    // 1. Parse stage ID from URL
    stageID, err := strconv.Atoi(c.Param("id"))

    // 2. Get user ID from JWT context
    userID, exists := c.Get("user_id")

    // 3. Parse request body
    var req models.StageSubmissionRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 4. Call service to process submission
    submission, err := services.SubmitStage(h.DB, userID.(int), stageID, req)

    // 5. Return response
    c.JSON(http.StatusCreated, gin.H{
        "message":    "Jawaban berhasil disubmit!",
        "is_correct": submission.IsCorrect,
        "score":      submission.Score,
        "submission": submission,
    })
}
```

---

#### handlers/lesson_handler.go

**Method: `GetMyEnrolledLessons`**

```go
func (h *LessonHandler) GetMyEnrolledLessons(c *gin.Context) {
    // 1. Get user ID from JWT
    userID, exists := c.Get("user_id")

    // 2. Get enrolled lessons via service
    lessons, err := services.GetMyEnrolledLessons(h.DB, userID.(int))

    // 3. Return response with progress
    c.JSON(http.StatusOK, gin.H{
        "count":   len(lessons),
        "lessons": lessons,
    })
}
```

**Method: `EnrollLesson`**

```go
func (h *LessonHandler) EnrollLesson(c *gin.Context) {
    // 1. Parse lesson ID from URL
    lessonID, err := strconv.Atoi(c.Param("id"))

    // 2. Get user ID from JWT
    userID, exists := c.Get("user_id")

    // 3. Enroll via service
    enrollment, err := services.EnrollLesson(h.DB, userID.(int), lessonID)

    // 4. Return response
    c.JSON(http.StatusCreated, gin.H{
        "message":    "Berhasil enroll ke lesson!",
        "enrollment": enrollment,
    })
}
```

---

### 9.2 Services (Business Logic)

#### services/primm_stage_service.go

**Key Functions:**

**`SubmitStage` - Process student submission**

```go
func SubmitStage(db *pgx.Conn, userID int, stageID int, req models.StageSubmissionRequest) (*models.StageSubmission, error) {
    // 1. Get stage details from database
    var stage models.PRIMMStage
    // ... query stage details

    // 2. Validate submission type
    if req.SubmissionType != stage.StageType {
        return nil, errors.New("tipe submission tidak sesuai")
    }

    // 3. Grade based on stage type
    var isCorrect bool
    var score int

    switch stage.StageType {
    case "predict":
        isCorrect = req.SelectedAnswer == *stage.CorrectAnswer
    case "run":
        isCorrect = true
    case "investigate":
        isCorrect = len(req.ReflectionText) >= 50
    case "modify":
        isCorrect = len(req.ModifiedCode) > 0
    case "make":
        isCorrect = len(req.Code) > 0
    }

    if isCorrect {
        score = 100
    }

    // 4. Save submission to database
    // ... insert into stage_submissions

    // 5. Check course completion (async)
    go checkCourseCompletion(db, userID, stage.CourseID)

    return &submission, nil
}
```

**`checkCourseCompletion` - Auto-award coins**

```go
func checkCourseCompletion(db *pgx.Conn, userID int, courseID int) {
    // 1. Count stages
    var totalStages, completedStages int
    // ... query stage counts

    // 2. Check if all completed
    if totalStages > 0 && totalStages == completedStages {
        // 3. Get coin reward
        var coinReward int
        // ... query coin_reward from courses

        // 4. Mark course complete
        db.Exec(`
            INSERT INTO user_course_completion (user_id, course_id, coins_awarded)
            VALUES ($1, $2, $3)
            ON CONFLICT (user_id, course_id) DO NOTHING
        `, userID, courseID, coinReward)

        // 5. Award coins
        db.Exec(`
            UPDATE users SET coins = coins + $1 WHERE id = $2
        `, coinReward, userID)
    }
}
```

**`GetStagesByCourse` - Get all stages in course**

```go
func GetStagesByCourse(db *pgx.Conn, courseID int) ([]models.PRIMMStage, error) {
    rows, err := db.Query(context.Background(), `
        SELECT
            id, course_id, stage_type, title, description, order_index,
            code_snippet, predict_options, correct_answer,
            run_code_template, video_embed_url, explanation_text,
            reflection_prompt, modify_challenge, modify_code_template,
            modify_expected_output, modify_test_cases,
            make_challenge, make_expected_output, make_test_cases,
            created_at, updated_at
        FROM primm_stages
        WHERE course_id = $1
        ORDER BY order_index ASC
    `, courseID)

    // Parse and return stages
    // ...
}
```

---

#### services/lesson_service.go

**`EnrollLesson` - Enroll student to lesson**

```go
func EnrollLesson(db *pgx.Conn, userID int, lessonID int) (*UserLessonEnrollment, error) {
    // 1. Check if lesson exists and active
    var isActive bool
    err := db.QueryRow(context.Background(), `
        SELECT is_active FROM lessons WHERE id = $1
    `, lessonID).Scan(&isActive)

    if !isActive {
        return nil, errors.New("lesson tidak aktif")
    }

    // 2. Check if already enrolled
    var existingID int
    err = db.QueryRow(context.Background(), `
        SELECT id FROM user_lessons WHERE user_id = $1 AND lesson_id = $2
    `, userID, lessonID).Scan(&existingID)

    if err == nil {
        return nil, errors.New("anda sudah terdaftar di lesson ini")
    }

    // 3. Insert enrollment
    var enrollment UserLessonEnrollment
    err = db.QueryRow(context.Background(), `
        INSERT INTO user_lessons (user_id, lesson_id)
        VALUES ($1, $2)
        RETURNING id, user_id, lesson_id, enrolled_at
    `, userID, lessonID).Scan(
        &enrollment.ID,
        &enrollment.UserID,
        &enrollment.LessonID,
        &enrollment.EnrolledAt,
    )

    return &enrollment, nil
}
```

**`GetMyEnrolledLessons` - Get enrolled lessons with progress**

```go
func GetMyEnrolledLessons(db *pgx.Conn, userID int) ([]models.LessonWithProgress, error) {
    rows, err := db.Query(context.Background(), `
        SELECT
            l.id, l.teacher_id, l.title, l.description,
            l.category, l.difficulty, l.thumbnail_url,
            l.is_active, l.created_at, l.updated_at,
            u.full_name as teacher_name,
            ul.enrolled_at,
            COALESCE(COUNT(DISTINCT c.id), 0) as total_courses
        FROM user_lessons ul
        JOIN lessons l ON ul.lesson_id = l.id
        JOIN users u ON l.teacher_id = u.id
        LEFT JOIN courses c ON l.id = c.lesson_id AND c.is_active = true
        WHERE ul.user_id = $1
        GROUP BY
            l.id, l.teacher_id, l.title, l.description,
            l.category, l.difficulty, l.thumbnail_url,
            l.is_active, l.created_at, l.updated_at,
            u.full_name, ul.enrolled_at
        ORDER BY ul.enrolled_at DESC
    `, userID)

    // Parse results with progress calculation
    // ...
}
```

---

### 9.3 Models (Data Structures)

#### models/primm_stage.go

**Submission Models:**

```go
// StageSubmissionRequest untuk request submit stage
type StageSubmissionRequest struct {
    SubmissionType string `json:"submission_type" binding:"required,oneof=predict run investigate modify make"`
    SelectedAnswer string `json:"selected_answer,omitempty"` // For PREDICT
    CodeOutput     string `json:"code_output,omitempty"`     // For RUN
    ReflectionText string `json:"reflection_text,omitempty"` // For INVESTIGATE
    ModifiedCode   string `json:"modified_code,omitempty"`   // For MODIFY
    Code           string `json:"code,omitempty"`            // For MAKE
}

// StageSubmission untuk response submission
type StageSubmission struct {
    ID             int                    `json:"id"`
    UserID         int                    `json:"user_id"`
    StageID        int                    `json:"stage_id"`
    SubmissionType string                 `json:"submission_type"`
    SubmissionData map[string]interface{} `json:"submission_data"`
    IsCorrect      bool                   `json:"is_correct"`
    Score          int                    `json:"score"`
    SubmittedAt    time.Time              `json:"submitted_at"`
}
```

---

#### models/lesson.go

**Progress Models:**

```go
// LessonWithProgress untuk response dengan progress tracking
type LessonWithProgress struct {
    ID               int       `json:"id"`
    TeacherID        int       `json:"teacher_id"`
    Title            string    `json:"title"`
    Description      string    `json:"description"`
    Category         string    `json:"category"`
    Difficulty       string    `json:"difficulty"`
    ThumbnailURL     string    `json:"thumbnail_url"`
    IsActive         bool      `json:"is_active"`
    CreatedAt        time.Time `json:"created_at"`
    UpdatedAt        time.Time `json:"updated_at"`
    TeacherName      string    `json:"teacher_name"`
    TotalCourses     int       `json:"total_courses"`
    CompletedCourses int       `json:"completed_courses"`
    ProgressPercent  int       `json:"progress_percent"`
    IsEnrolled       bool      `json:"is_enrolled"`
}
```

---

## 🔧 10. TROUBLESHOOTING

### Common Issues & Solutions

#### Issue 1: Table Not Found

```
ERROR: relation "stage_submissions" does not exist
```

**Solution:**

```sql
-- Create stage_submissions table
CREATE TABLE stage_submissions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    stage_id INT NOT NULL REFERENCES primm_stages(id) ON DELETE CASCADE,
    submission_type VARCHAR(50) NOT NULL,
    submission_data JSONB NOT NULL,
    is_correct BOOLEAN DEFAULT false,
    score INT DEFAULT 0,
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, stage_id)
);
```

---

#### Issue 2: Column Not Found

```
ERROR: column "last_accessed_at" does not exist
```

**Solution:**
Remove reference to non-existent column or add it:

```sql
-- Option 1: Add column (future enhancement)
ALTER TABLE user_lessons ADD COLUMN last_accessed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Option 2: Remove from query (current approach)
-- Just don't SELECT or use last_accessed_at
```

---

#### Issue 3: Type Mismatch - Pointer vs Value

```
ERROR: invalid operation: req.SelectedAnswer == stage.CorrectAnswer (mismatched types string and *string)
```

**Solution:**

```go
// Wrong:
isCorrect = req.SelectedAnswer == stage.CorrectAnswer

// Correct:
if stage.CorrectAnswer != nil {
    isCorrect = req.SelectedAnswer == *stage.CorrectAnswer
}
```

---

#### Issue 4: Coins Not Awarded

**Symptoms:** Course completed but student coins still 0

**Debug Steps:**

1. **Check if course completion was triggered:**

```sql
SELECT * FROM user_course_completion WHERE user_id = 8 AND course_id = 1;
```

2. **Check stage completion count:**

```sql
SELECT
    COUNT(DISTINCT ps.id) as total_stages,
    COUNT(DISTINCT ss.stage_id) FILTER (WHERE ss.is_correct = true) as completed_stages
FROM primm_stages ps
LEFT JOIN stage_submissions ss ON ps.id = ss.stage_id AND ss.user_id = 8
WHERE ps.course_id = 1;
```

3. **Check user coins:**

```sql
SELECT id, full_name, coins FROM users WHERE id = 8;
```

**Possible Causes:**

- ❌ Not all stages marked as `is_correct = true`
- ❌ `checkCourseCompletion` function not running
- ❌ Database transaction failed silently

---

#### Issue 5: Progress Not Updating

**Symptoms:** Submitted all stages but progress still 0%

**Solution:**
Check query in `GetMyEnrolledLessons`:

```go
// Make sure query joins user_course_completion properly
LEFT JOIN user_course_completion ucc ON c.id = ucc.course_id AND ucc.user_id = $1
```

---

#### Issue 6: Duplicate Submission Error

**Symptoms:**

```
ERROR: duplicate key value violates unique constraint "stage_submissions_user_id_stage_id_key"
```

**Solution:**
This is expected behavior! Use `ON CONFLICT` to update:

```sql
INSERT INTO stage_submissions (user_id, stage_id, ...)
VALUES ($1, $2, ...)
ON CONFLICT (user_id, stage_id)
DO UPDATE SET
    submission_data = EXCLUDED.submission_data,
    is_correct = EXCLUDED.is_correct,
    score = EXCLUDED.score,
    submitted_at = CURRENT_TIMESTAMP
```

---

## 📝 SUMMARY

### What Student Flow Does:

1. ✅ **Registration & Authentication** - Student can create account and login
2. ✅ **Content Discovery** - Browse and view available lessons
3. ✅ **Enrollment** - Enroll to lessons of interest
4. ✅ **Learning Journey** - Complete 5 PRIMM stages sequentially
5. ✅ **Auto-Grading** - Instant feedback on submissions
6. ✅ **Progress Tracking** - Visual progress indicators
7. ✅ **Rewards** - Automatic coin distribution on course completion
8. ✅ **Achievement System** - Track completed courses and earned coins

### Key Design Decisions:

1. **PRIMM Methodology** - Evidence-based teaching approach
2. **Auto-grading** - Instant feedback without teacher intervention
3. **JSONB storage** - Flexible submission data per stage type
4. **Async processing** - Course completion check doesn't block submission
5. **Unique constraints** - Prevent duplicate enrollments/submissions
6. **Progress calculation** - Real-time progress based on stage completion
7. **Coin rewards** - Gamification to motivate learning

### Future Enhancements:

**Priority 1: Code Execution**

- ✅ Integrate Python sandbox (Piston API / Judge0)
- ✅ Auto-validate RUN stage output against expected
- ✅ Run test cases for MODIFY and MAKE stages
- ✅ Partial credit for partially correct solutions

**Priority 2: Enhanced Grading**

- ✅ AI-powered reflection analysis for INVESTIGATE
- ✅ Code quality metrics for MODIFY/MAKE
- ✅ Multiple attempt tracking with score decay
- ✅ Detailed feedback on incorrect answers

**Priority 3: Student Experience**

- ✅ Add hints system (reveal incrementally)
- ✅ Discussion forum per course/stage
- ✅ Peer code review features
- ✅ Certificate generation on completion

**Priority 4: Progress Enhancement**

- ✅ Add `last_accessed_at` tracking
- ✅ Time spent per stage analytics
- ✅ Leaderboard system
- ✅ Achievement badges

---

## 📚 Additional Resources

- **Teacher Flow Documentation:** `docs/TEACHER_FLOW_DOCUMENTATION.md`
- **Full System Documentation:** `docs/FULL_SYSTEM_DOCUMENTATION.md`
- **API Reference:** (Coming soon)
- **PRIMM Methodology:** https://primmportal.com/

---

**Last Updated:** October 25, 2025  
**Version:** 1.0  
**Status:** Student Flow Complete ✅
