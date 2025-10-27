# 🎨 PRIMMFY FRONTEND DEVELOPMENT PLAN

> **Purpose:** Complete development roadmap for PRIMMFY frontend using Next.js 14 + TypeScript + TailwindCSS

**Created:** October 26, 2025  
**Status:** Planning Phase 📋  
**Target:** Production-Ready Frontend

---

## 📋 Table of Contents

1. [Tech Stack & Architecture](#1-tech-stack--architecture)
2. [Project Structure](#2-project-structure)
3. [Development Phases](#3-development-phases)
4. [Feature Specifications](#4-feature-specifications)
5. [UI/UX Design Guidelines](#5-uiux-design-guidelines)
6. [API Integration](#6-api-integration)
7. [Testing Strategy](#7-testing-strategy)
8. [Deployment Plan](#8-deployment-plan)

---

## 🏗️ 1. TECH STACK & ARCHITECTURE

### Core Technologies

```
┌─────────────────────────────────────────┐
│  FRONTEND STACK                         │
├─────────────────────────────────────────┤
│  Framework:    Next.js 14 (App Router)  │
│  Language:     TypeScript 5.0+          │
│  Styling:      TailwindCSS 3.4+         │
│  UI Library:   shadcn/ui + Radix UI     │
│  Icons:        Lucide React             │
│  Code Editor:  Monaco Editor            │
└─────────────────────────────────────────┘

┌─────────────────────────────────────────┐
│  STATE MANAGEMENT                       │
├─────────────────────────────────────────┤
│  Auth State:   Context API              │
│  Server State: TanStack Query (React)   │
│  Form State:   React Hook Form          │
│  Local State:  React useState/useRef    │
└─────────────────────────────────────────┘

┌─────────────────────────────────────────┐
│  DATA FETCHING & VALIDATION             │
├─────────────────────────────────────────┤
│  HTTP Client:  Axios                    │
│  Validation:   Zod                      │
│  Cookies:      js-cookie                │
└─────────────────────────────────────────┘
```

### Architecture Pattern

```
┌──────────────────────────────────────────────────────────┐
│                    USER INTERFACE                        │
│              (Next.js Pages & Components)                │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                  REACT COMPONENTS                        │
│   • Presentational Components (UI)                       │
│   • Container Components (Logic)                         │
│   • Layout Components (Structure)                        │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                  CUSTOM HOOKS                            │
│   • useAuth()      - Authentication state                │
│   • useLessons()   - Lesson data fetching                │
│   • useCourses()   - Course data fetching                │
│   • useStages()    - Stage data fetching                 │
│   • useSubmit()    - Submission handling                 │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                  API CLIENT LAYER                        │
│   • axios.ts           - HTTP client config              │
│   • auth.ts            - Auth API calls                  │
│   • lessons.ts         - Lesson API calls                │
│   • courses.ts         - Course API calls                │
│   • stages.ts          - Stage API calls                 │
│   • submissions.ts     - Submission API calls            │
└──────────────────────────────────────────────────────────┘
                           ↓
┌──────────────────────────────────────────────────────────┐
│                  BACKEND API                             │
│              Go Backend (localhost:8080)                 │
└──────────────────────────────────────────────────────────┘
```

---

## 📁 2. PROJECT STRUCTURE

```
Frontend/
├── app/                           # Next.js 14 App Router
│   ├── (auth)/                    # Authentication routes
│   │   ├── login/
│   │   │   └── page.tsx           # Login page
│   │   ├── register/
│   │   │   └── page.tsx           # Register page
│   │   └── layout.tsx             # Auth layout (no navbar)
│   │
│   ├── (teacher)/                 # Teacher-only routes
│   │   ├── dashboard/
│   │   │   └── page.tsx           # Teacher dashboard
│   │   ├── lessons/
│   │   │   ├── page.tsx           # My lessons list
│   │   │   ├── create/
│   │   │   │   └── page.tsx       # Create new lesson
│   │   │   └── [id]/
│   │   │       ├── page.tsx       # Lesson detail
│   │   │       └── edit/
│   │   │           └── page.tsx   # Edit lesson
│   │   ├── courses/
│   │   │   ├── create/
│   │   │   │   └── page.tsx       # Create course
│   │   │   └── [id]/
│   │   │       ├── edit/
│   │   │       │   └── page.tsx   # Edit course
│   │   │       └── stages/
│   │   │           ├── page.tsx   # View all stages
│   │   │           └── create/
│   │   │               └── page.tsx # Create stages (5 types)
│   │   └── layout.tsx             # Teacher layout (with navbar)
│   │
│   ├── (student)/                 # Student-only routes
│   │   ├── dashboard/
│   │   │   └── page.tsx           # Student dashboard
│   │   ├── browse/
│   │   │   └── page.tsx           # Browse all lessons
│   │   ├── my-lessons/
│   │   │   └── page.tsx           # Enrolled lessons
│   │   ├── lessons/
│   │   │   └── [id]/
│   │   │       ├── page.tsx       # Lesson overview
│   │   │       └── courses/
│   │   │           └── [courseId]/
│   │   │               ├── page.tsx       # Course learning view
│   │   │               └── stages/
│   │   │                   └── [stageId]/
│   │   │                       └── page.tsx # Stage submission
│   │   └── layout.tsx             # Student layout (with navbar)
│   │
│   ├── layout.tsx                 # Root layout
│   ├── page.tsx                   # Landing page
│   ├── globals.css                # Global Tailwind styles
│   └── providers.tsx              # React Query & Auth providers
│
├── components/
│   ├── ui/                        # shadcn/ui base components
│   │   ├── button.tsx
│   │   ├── card.tsx
│   │   ├── dialog.tsx
│   │   ├── input.tsx
│   │   ├── label.tsx
│   │   ├── select.tsx
│   │   ├── tabs.tsx
│   │   ├── toast.tsx
│   │   ├── dropdown-menu.tsx
│   │   ├── avatar.tsx
│   │   ├── badge.tsx
│   │   ├── progress.tsx
│   │   └── separator.tsx
│   │
│   ├── auth/                      # Authentication components
│   │   ├── LoginForm.tsx          # Login form with validation
│   │   ├── RegisterForm.tsx       # Register form (teacher/student)
│   │   ├── ProtectedRoute.tsx     # Route guard component
│   │   └── RoleGuard.tsx          # Role-based access control
│   │
│   ├── teacher/                   # Teacher-specific components
│   │   ├── LessonCard.tsx         # Lesson display card
│   │   ├── LessonForm.tsx         # Create/edit lesson form
│   │   ├── CourseCard.tsx         # Course display card
│   │   ├── CourseForm.tsx         # Create/edit course form
│   │   ├── StageFormPredict.tsx   # PREDICT stage form
│   │   ├── StageFormRun.tsx       # RUN stage form
│   │   ├── StageFormInvestigate.tsx # INVESTIGATE stage form
│   │   ├── StageFormModify.tsx    # MODIFY stage form
│   │   ├── StageFormMake.tsx      # MAKE stage form
│   │   ├── DashboardStats.tsx     # Teacher statistics widget
│   │   └── StudentProgressTable.tsx # Student progress view
│   │
│   ├── student/                   # Student-specific components
│   │   ├── LessonBrowser.tsx      # Browse lessons with filters
│   │   ├── LessonCard.tsx         # Lesson card (student view)
│   │   ├── CourseProgress.tsx     # Course progress tracker
│   │   ├── StageViewer.tsx        # View stage content
│   │   ├── PredictSubmission.tsx  # PREDICT answer submission
│   │   ├── RunSubmission.tsx      # RUN code execution
│   │   ├── InvestigateSubmission.tsx # INVESTIGATE reflection
│   │   ├── ModifySubmission.tsx   # MODIFY code editor
│   │   ├── MakeSubmission.tsx     # MAKE code editor
│   │   ├── ProgressDashboard.tsx  # Student progress overview
│   │   └── CoinDisplay.tsx        # Coin balance widget
│   │
│   ├── shared/                    # Shared components
│   │   ├── Navbar.tsx             # Navigation bar
│   │   ├── Footer.tsx             # Footer
│   │   ├── Sidebar.tsx            # Sidebar navigation
│   │   ├── LoadingSpinner.tsx     # Loading state
│   │   ├── ErrorMessage.tsx       # Error display
│   │   ├── EmptyState.tsx         # Empty state illustration
│   │   ├── ConfirmDialog.tsx      # Confirmation modal
│   │   ├── CodeEditor.tsx         # Monaco code editor wrapper
│   │   ├── VideoPlayer.tsx        # Video embed wrapper
│   │   └── ProgressBar.tsx        # Progress bar component
│   │
│   └── layouts/                   # Layout components
│       ├── AuthLayout.tsx         # Auth pages layout
│       ├── TeacherLayout.tsx      # Teacher pages layout
│       └── StudentLayout.tsx      # Student pages layout
│
├── lib/
│   ├── api/                       # API client functions
│   │   ├── axios.ts               # Axios instance config
│   │   ├── auth.ts                # Auth API (login, register)
│   │   ├── lessons.ts             # Lesson CRUD API
│   │   ├── courses.ts             # Course CRUD API
│   │   ├── stages.ts              # Stage CRUD API
│   │   ├── submissions.ts         # Submission API
│   │   └── progress.ts            # Progress tracking API
│   │
│   ├── contexts/                  # React Context providers
│   │   ├── AuthContext.tsx        # Auth state management
│   │   └── ToastContext.tsx       # Toast notifications
│   │
│   ├── hooks/                     # Custom React hooks
│   │   ├── useAuth.ts             # Auth hook
│   │   ├── useLessons.ts          # Lessons data hook
│   │   ├── useCourses.ts          # Courses data hook
│   │   ├── useStages.ts           # Stages data hook
│   │   ├── useSubmissions.ts      # Submissions hook
│   │   ├── useProgress.ts         # Progress tracking hook
│   │   └── useDebounce.ts         # Debounce utility hook
│   │
│   ├── types/                     # TypeScript type definitions
│   │   ├── auth.ts                # Auth types (User, LoginRequest, etc.)
│   │   ├── lesson.ts              # Lesson types
│   │   ├── course.ts              # Course types
│   │   ├── stage.ts               # Stage types (all 5 types)
│   │   ├── submission.ts          # Submission types
│   │   ├── progress.ts            # Progress types
│   │   └── api.ts                 # API response types
│   │
│   ├── utils/                     # Utility functions
│   │   ├── cn.ts                  # Tailwind class merger
│   │   ├── validation.ts          # Zod schemas
│   │   ├── formatters.ts          # Date, number formatters
│   │   ├── storage.ts             # LocalStorage/Cookies helpers
│   │   └── constants.ts           # App constants
│   │
│   └── schemas/                   # Zod validation schemas
│       ├── auth.schema.ts         # Auth validation
│       ├── lesson.schema.ts       # Lesson validation
│       ├── course.schema.ts       # Course validation
│       └── stage.schema.ts        # Stage validation
│
├── public/
│   ├── images/
│   │   ├── logo.svg
│   │   ├── hero-illustration.svg
│   │   └── empty-states/
│   ├── icons/
│   │   ├── coin.svg
│   │   └── badge/
│   └── fonts/
│
├── .env.local                     # Environment variables
├── .env.example                   # Example env file
├── next.config.js                 # Next.js configuration
├── tailwind.config.ts             # Tailwind configuration
├── tsconfig.json                  # TypeScript configuration
├── package.json                   # Dependencies
└── README.md                      # Frontend documentation
```

---

## 🎯 3. DEVELOPMENT PHASES

### Phase 1: Foundation & Authentication (Week 1) ✅

**Goal:** Setup project, implement authentication, protected routes

#### Tasks:

- [x] Initialize Next.js project with TypeScript
- [x] Install all dependencies
- [x] Configure TailwindCSS
- [ ] Setup shadcn/ui components
- [ ] Create base TypeScript types
- [ ] Implement Axios client with interceptors
- [ ] Build Auth Context
- [ ] Create Login page
- [ ] Create Register page (with role selection)
- [ ] Implement JWT token management
- [ ] Build ProtectedRoute component
- [ ] Build RoleGuard component
- [ ] Test authentication flow

#### Deliverables:

```
✅ Working login system
✅ Working registration (teacher/student)
✅ JWT stored in cookies
✅ Protected routes working
✅ Role-based access control
✅ Error handling for auth
```

---

### Phase 2: Teacher Dashboard - Lesson Management (Week 2)

**Goal:** Teacher can create, edit, delete lessons

#### Tasks:

- [ ] Create Teacher Dashboard layout
- [ ] Build Teacher Navbar with stats
- [ ] Create "My Lessons" page
- [ ] Build LessonCard component
- [ ] Create LessonForm component
- [ ] Implement Create Lesson page
- [ ] Implement Edit Lesson page
- [ ] Implement Delete Lesson (with confirmation)
- [ ] Add lesson filters (category, difficulty)
- [ ] Add lesson search functionality
- [ ] Integrate with backend API
- [ ] Test all CRUD operations

#### API Endpoints Used:

```
GET    /api/lessons              # Get teacher's lessons
POST   /api/lessons              # Create lesson
PUT    /api/lessons/:id          # Update lesson
DELETE /api/lessons/:id          # Delete lesson
```

#### Deliverables:

```
✅ Teacher can view all their lessons
✅ Teacher can create new lesson
✅ Teacher can edit lesson
✅ Teacher can delete lesson
✅ Proper validation & error handling
✅ Loading states
```

---

### Phase 3: Teacher Dashboard - Course Management (Week 3)

**Goal:** Teacher can create courses within lessons

#### Tasks:

- [ ] Create "Lesson Detail" page (view courses)
- [ ] Build CourseCard component
- [ ] Create CourseForm component
- [ ] Implement Create Course page
- [ ] Implement Edit Course page
- [ ] Implement Delete Course (with confirmation)
- [ ] Add course ordering (drag & drop - optional)
- [ ] Show coin rewards clearly
- [ ] Integrate with backend API
- [ ] Test course CRUD operations

#### API Endpoints Used:

```
GET    /api/lessons/:id/courses  # Get courses in lesson
POST   /api/courses              # Create course
PUT    /api/courses/:id          # Update course
DELETE /api/courses/:id          # Delete course
```

#### Deliverables:

```
✅ Teacher can view courses in lesson
✅ Teacher can create course
✅ Teacher can edit course
✅ Teacher can delete course
✅ Course ordering works
✅ Validation & error handling
```

---

### Phase 4: Teacher Dashboard - PRIMM Stage Creation (Week 4-5)

**Goal:** Teacher can create all 5 PRIMM stage types

#### Tasks:

**PREDICT Stage:**

- [ ] Build StageFormPredict component
- [ ] Multiple choice options input
- [ ] Set correct answer selector
- [ ] Preview functionality

**RUN Stage:**

- [ ] Build StageFormRun component
- [ ] Code snippet input (Monaco Editor)
- [ ] Expected output field
- [ ] Syntax highlighting

**INVESTIGATE Stage:**

- [ ] Build StageFormInvestigate component
- [ ] Video URL input (with preview)
- [ ] Guiding questions (dynamic list)
- [ ] Reflection prompt input

**MODIFY Stage:**

- [ ] Build StageFormModify component
- [ ] Code template editor
- [ ] Challenge description
- [ ] Test cases input (JSON)
- [ ] Expected output

**MAKE Stage:**

- [ ] Build StageFormMake component
- [ ] Challenge description
- [ ] Hints input
- [ ] Test cases input (JSON)
- [ ] Expected output

**Integration:**

- [ ] Create unified "Create Stage" page (tabs for 5 types)
- [ ] Stage preview component
- [ ] Integrate all 5 stage APIs
- [ ] Test each stage creation
- [ ] View/edit/delete stages

#### API Endpoints Used:

```
GET    /api/courses/:id/stages         # Get stages in course
POST   /api/stages/predict             # Create PREDICT
POST   /api/stages/run                 # Create RUN
POST   /api/stages/investigate         # Create INVESTIGATE
POST   /api/stages/modify              # Create MODIFY
POST   /api/stages/make                # Create MAKE
GET    /api/stages/:id                 # Get stage details
DELETE /api/stages/:id                 # Delete stage
```

#### Deliverables:

```
✅ All 5 stage types can be created
✅ Monaco editor integrated for code
✅ Video preview for INVESTIGATE
✅ Test cases input for MODIFY/MAKE
✅ Stage preview before save
✅ Edit/delete stages
✅ Proper validation
```

---

### Phase 5: Student Interface - Browse & Enroll (Week 6)

**Goal:** Student can browse lessons and enroll

#### Tasks:

- [ ] Create Student Dashboard layout
- [ ] Build Student Navbar with coins display
- [ ] Create "Browse Lessons" page
- [ ] Build LessonCard (student view)
- [ ] Add category filters
- [ ] Add difficulty filters
- [ ] Add search functionality
- [ ] Implement "Enroll" button
- [ ] Create "My Lessons" page
- [ ] Show enrollment status
- [ ] Integrate with backend API
- [ ] Test enrollment flow

#### API Endpoints Used:

```
GET    /api/lessons                    # Browse all lessons
POST   /api/lessons/:id/enroll         # Enroll to lesson
GET    /api/my-lessons                 # Get enrolled lessons
```

#### Deliverables:

```
✅ Student can browse all lessons
✅ Filters work (category, difficulty)
✅ Search functionality
✅ Student can enroll to lesson
✅ "My Lessons" shows enrolled lessons
✅ Progress indicators
```

---

### Phase 6: Student Interface - Learning Flow (Week 7-8)

**Goal:** Student can view content and submit all 5 PRIMM stages

#### Tasks:

**Lesson & Course View:**

- [ ] Create Lesson Detail page (student)
- [ ] Show courses in lesson
- [ ] Course progress indicators
- [ ] Navigate to course learning page

**Course Learning Interface:**

- [ ] Create Course Learning page
- [ ] Show all 5 stages (navigation)
- [ ] Stage progress tracker (1/5, 2/5, etc.)
- [ ] Lock/unlock stages (sequential)

**PREDICT Submission:**

- [ ] Build PredictSubmission component
- [ ] Show question & options
- [ ] Radio button selection
- [ ] Submit answer
- [ ] Show correct/incorrect feedback

**RUN Submission:**

- [ ] Build RunSubmission component
- [ ] Show code snippet
- [ ] "Run Code" button (simulated)
- [ ] Output display area
- [ ] Submit output

**INVESTIGATE Submission:**

- [ ] Build InvestigateSubmission component
- [ ] Embed video player
- [ ] Show guiding questions
- [ ] Text area for reflection (min 50 chars)
- [ ] Submit reflection

**MODIFY Submission:**

- [ ] Build ModifySubmission component
- [ ] Show challenge description
- [ ] Monaco code editor (editable)
- [ ] "Test Code" button (optional)
- [ ] Submit modified code

**MAKE Submission:**

- [ ] Build MakeSubmission component
- [ ] Show challenge
- [ ] Show hints (reveal button)
- [ ] Monaco code editor (blank)
- [ ] Submit created code

**Progress & Rewards:**

- [ ] Show submission feedback (score)
- [ ] Update progress after submission
- [ ] Award coins on course completion
- [ ] Show coin notification/animation
- [ ] Update coin balance in navbar

#### API Endpoints Used:

```
GET    /api/lessons/:id/courses        # Get courses
GET    /api/courses/:id/stages         # Get all stages
POST   /api/stages/:id/submit          # Submit stage
GET    /api/profile                    # Get updated coins
```

#### Deliverables:

```
✅ Student can view all 5 stages
✅ All 5 submission forms work
✅ Auto-grading feedback displayed
✅ Progress updates after submission
✅ Coins awarded on completion
✅ Coin balance updates in real-time
✅ Smooth learning flow
```

---

### Phase 7: UI/UX Polish & Optimization (Week 9)

**Goal:** Improve user experience, animations, responsiveness

#### Tasks:

- [ ] Add loading skeletons
- [ ] Add smooth transitions
- [ ] Implement toast notifications
- [ ] Add empty states with illustrations
- [ ] Responsive design for mobile/tablet
- [ ] Add keyboard shortcuts
- [ ] Improve error messages
- [ ] Add confirmation dialogs
- [ ] Loading states for all actions
- [ ] Optimize images
- [ ] Code splitting
- [ ] Accessibility improvements (ARIA)

#### Deliverables:

```
✅ Smooth animations
✅ Loading states everywhere
✅ Beautiful empty states
✅ Mobile-responsive
✅ Accessible (WCAG AA)
✅ Fast page loads
```

---

### Phase 8: Testing & Bug Fixes (Week 10)

**Goal:** Ensure everything works perfectly

#### Tasks:

- [ ] Manual testing (all flows)
- [ ] Cross-browser testing
- [ ] Mobile device testing
- [ ] Fix identified bugs
- [ ] Performance optimization
- [ ] SEO optimization
- [ ] Security audit (XSS, CSRF)
- [ ] Write unit tests (Jest)
- [ ] Write integration tests (Playwright)
- [ ] Load testing

#### Deliverables:

```
✅ All bugs fixed
✅ Tests passing
✅ Performance optimized
✅ Security checked
✅ Ready for production
```

---

## 🎨 4. FEATURE SPECIFICATIONS

### 4.1 Authentication System

#### Login Page (`/login`)

**UI Elements:**

```
┌─────────────────────────────────────────┐
│           PRIMMFY LOGO                  │
│                                         │
│         Welcome Back! 👋                │
│                                         │
│  [Email Input]                          │
│  [Password Input]                       │
│                                         │
│  [Remember Me] [Forgot Password?]       │
│                                         │
│  [ Login Button ]                       │
│                                         │
│  Don't have account? [Register]         │
└─────────────────────────────────────────┘
```

**Validation:**

- Email: Valid email format
- Password: Min 6 characters

**Success Flow:**

1. User enters credentials
2. API call to POST /api/login
3. Receive JWT token
4. Store token in httpOnly cookie
5. Redirect to dashboard (based on role)

**Error Handling:**

- Invalid credentials → "Invalid email or password"
- Network error → "Connection failed, try again"
- Server error → "Something went wrong"

---

#### Register Page (`/register`)

**UI Elements:**

```
┌─────────────────────────────────────────┐
│           PRIMMFY LOGO                  │
│                                         │
│         Create Account 🚀               │
│                                         │
│  I am a:  [Teacher] [Student]           │
│                                         │
│  [Full Name Input]                      │
│  [Email Input]                          │
│  [Password Input]                       │
│  [Confirm Password Input]               │
│                                         │
│  [ ] I agree to Terms & Conditions      │
│                                         │
│  [ Register Button ]                    │
│                                         │
│  Already have account? [Login]          │
└─────────────────────────────────────────┘
```

**Validation:**

- Role: Required (teacher/student)
- Full Name: Min 3 characters
- Email: Valid + unique
- Password: Min 6 chars
- Confirm Password: Must match

---

### 4.2 Teacher Dashboard

#### Dashboard Home (`/teacher/dashboard`)

**Widgets:**

```
┌─────────────────────────────────────────────────────────┐
│  👨‍🏫 Teacher Dashboard                                   │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐             │
│  │ Lessons  │  │ Courses  │  │ Students │             │
│  │    5     │  │    12    │  │    45    │             │
│  └──────────┘  └──────────┘  └──────────┘             │
│                                                         │
│  Recent Activity                                        │
│  ├─ Alice completed "Variables" course                 │
│  ├─ Bob enrolled in "Python Basics"                    │
│  └─ 3 new enrollments today                            │
│                                                         │
│  Quick Actions                                          │
│  [ Create Lesson ] [ Create Course ]                   │
└─────────────────────────────────────────────────────────┘
```

---

#### My Lessons (`/teacher/lessons`)

**Features:**

- Grid/List view toggle
- Filter by category (Python, JavaScript, HTML, C)
- Filter by difficulty
- Search by title
- Sort by date/title

**Lesson Card:**

```
┌─────────────────────────────────────┐
│  🐍 Python Basics                  │
│  Category: Python • Beginner        │
│                                     │
│  Courses: 3                         │
│  Students: 15                       │
│                                     │
│  [View] [Edit] [Delete]             │
└─────────────────────────────────────┘
```

---

#### Create Lesson (`/teacher/lessons/create`)

**Form Fields:**

```
Title:          [_________________________]
Description:    [_________________________]
                [_________________________]
Category:       [Python ▼]
Difficulty:     [Beginner ▼]
Thumbnail URL:  [_________________________]

[ Cancel ]  [ Create Lesson ]
```

---

#### Create Stage (`/teacher/courses/:id/stages/create`)

**Tab Navigation:**

```
[ PREDICT ] [ RUN ] [ INVESTIGATE ] [ MODIFY ] [ MAKE ]
```

**PREDICT Form:**

```
Title:          [_________________________]
Description:    [_________________________]
Code Snippet:   [Code Editor with syntax highlighting]

Options:
  A) [_________________________]
  B) [_________________________]
  C) [_________________________]
  D) [_________________________]

Correct Answer: [C ▼]

[ Cancel ]  [ Create Stage ]
```

---

### 4.3 Student Dashboard

#### Dashboard Home (`/student/dashboard`)

**Widgets:**

```
┌─────────────────────────────────────────────────────────┐
│  🎓 Student Dashboard               Coins: 500 🪙       │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  Continue Learning                                      │
│  ┌─────────────────────────────────────────────────┐   │
│  │ 🐍 Python Basics                                │   │
│  │ Variables and Data Types                        │   │
│  │ Progress: [████████░░] 80%                      │   │
│  │ [ Continue → ]                                  │   │
│  └─────────────────────────────────────────────────┘   │
│                                                         │
│  My Progress                                            │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐             │
│  │ Enrolled │  │Completed │  │  Coins   │             │
│  │    3     │  │    5     │  │   500    │             │
│  └──────────┘  └──────────┘  └──────────┘             │
│                                                         │
│  Recommended Lessons                                    │
│  [ Browse All Lessons ]                                 │
└─────────────────────────────────────────────────────────┘
```

---

#### Browse Lessons (`/student/browse`)

**Features:**

- Grid view with lesson cards
- Category filter (All, Python, JavaScript, HTML, C)
- Difficulty filter (All, Beginner, Intermediate, Advanced)
- Search bar
- "Enrolled" badge on enrolled lessons

**Lesson Card:**

```
┌─────────────────────────────────────┐
│  [Thumbnail Image]                  │
│  🐍 Python Basics                  │
│  Master Python fundamentals         │
│                                     │
│  Category: Python • Beginner        │
│  Courses: 3 • Students: 45          │
│                                     │
│  [ Enroll Now ]                     │
└─────────────────────────────────────┘
```

---

#### Course Learning Page (`/student/lessons/:id/courses/:courseId`)

**Layout:**

```
┌─────────────────────────────────────────────────────────┐
│  ← Back to Lesson                 Progress: 3/5 (60%)   │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  Stage Navigation                                       │
│  [✓ 1. PREDICT] [✓ 2. RUN] [→ 3. INVESTIGATE] [🔒 4] [🔒 5] │
│                                                         │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  STAGE 3: INVESTIGATE                                   │
│  Understanding Variables                                │
│                                                         │
│  [Video Player]                                         │
│                                                         │
│  Guiding Questions:                                     │
│  • What is a variable?                                  │
│  • Why do we need variables?                            │
│                                                         │
│  Reflection (min 50 characters):                        │
│  [___________________________________________]          │
│  [___________________________________________]          │
│  [___________________________________________]          │
│                                                         │
│  [ Submit Reflection ]                                  │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

**Stage Icons:**

- ✓ = Completed (green)
- → = Current (blue)
- 🔒 = Locked (gray)

---

## 🎨 5. UI/UX DESIGN GUIDELINES

### Color Scheme

```css
/* Primary Colors */
--primary: #C2E2FA (Light Blue)       /* Main brand color - calm & learning */
--primary-dark: #9BCEF5              /* Hover states */
--primary-light: #E5F3FD             /* Backgrounds */

/* Secondary Colors */
--secondary: #B7A3E3 (Lavender)      /* Accents & interactive elements */
--secondary-dark: #9B81D9            /* Hover states */
--secondary-light: #E0D7F4           /* Backgrounds */

/* Accent Colors */
--accent-pink: #FF8F8F (Coral Pink)  /* Highlights & important actions */
--accent-cream: #FFF1CB (Cream)      /* Soft backgrounds & cards */

/* Semantic Colors */
--success: #10B981 (Green)           /* Success states */
--warning: #F59E0B (Orange)          /* Warnings */
--error: #EF4444 (Red)               /* Errors */

/* Neutral Colors */
--white: #FFFFFF
--gray-50: #F9FAFB
--gray-100: #F3F4F6
--gray-200: #E5E7EB
--gray-300: #D1D5DB
--gray-500: #6B7280
--gray-700: #374151
--gray-900: #111827
--black: #000000

/* PRIMM Stage Colors */
--stage-predict: #C2E2FA (Light Blue)     /* PREDICT - prediction/thinking */
--stage-run: #B7A3E3 (Lavender)           /* RUN - execution/action */
--stage-investigate: #FF8F8F (Coral Pink) /* INVESTIGATE - exploration */
--stage-modify: #FFF1CB (Cream)           /* MODIFY - modification */
--stage-make: #10B981 (Green)             /* MAKE - creation/completion */

/* UI Element Colors */
--coin-gold: #FCD34D              /* Coin rewards */
--button-primary: #C2E2FA         /* Primary buttons */
--button-secondary: #B7A3E3       /* Secondary buttons */
--button-danger: #FF8F8F          /* Delete/dangerous actions */
--card-bg: #FFFFFF                /* Card backgrounds */
--card-border: #E5E7EB            /* Card borders */
```

---

### Color Usage Guidelines

**Primary (Light Blue - #C2E2FA):**
- Main navigation background
- Primary buttons
- Links
- Active states
- Focus indicators
- Progress bars

**Secondary (Lavender - #B7A3E3):**
- Secondary buttons
- Tags/badges
- Hover states
- Interactive elements
- Icons

**Accent Pink (#FF8F8F):**
- Call-to-action buttons
- Important notifications
- Delete/warning actions
- Highlights
- Active course indicators

**Accent Cream (#FFF1CB):**
- Card backgrounds
- Section backgrounds
- Hover states for cards
- Soft highlights
- Empty states

**Example Usage:**

```css
/* Buttons */
.btn-primary {
  background: #C2E2FA;
  color: #374151;
  border: 2px solid #9BCEF5;
}

.btn-primary:hover {
  background: #9BCEF5;
  transform: translateY(-2px);
}

.btn-secondary {
  background: #B7A3E3;
  color: #FFFFFF;
}

.btn-danger {
  background: #FF8F8F;
  color: #FFFFFF;
}

/* Cards */
.card {
  background: #FFFFFF;
  border: 1px solid #E5E7EB;
}

.card:hover {
  background: #FFF1CB; /* Cream accent on hover */
  border-color: #C2E2FA;
}

/* PRIMM Stage Indicators */
.stage-predict {
  background: #C2E2FA;
  color: #374151;
}

.stage-run {
  background: #B7A3E3;
  color: #FFFFFF;
}

.stage-investigate {
  background: #FF8F8F;
  color: #FFFFFF;
}

.stage-modify {
  background: #FFF1CB;
  color: #374151;
}

.stage-make {
  background: #10B981;
  color: #FFFFFF;
}

/* Progress Bar */
.progress-bar {
  background: #E5E7EB;
}

.progress-fill {
  background: linear-gradient(90deg, #C2E2FA 0%, #B7A3E3 100%);
}

/* Navigation */
.navbar {
  background: linear-gradient(135deg, #C2E2FA 0%, #E5F3FD 100%);
  border-bottom: 2px solid #9BCEF5;
}

/* Sidebar Active Item */
.sidebar-item.active {
  background: #FFF1CB;
  border-left: 4px solid #B7A3E3;
}
```

### Typography

```css
/* Font Family */
font-family: 'Inter', sans-serif;

/* Headings */
h1: 2.5rem (40px), font-weight: 700
h2: 2rem (32px), font-weight: 700
h3: 1.5rem (24px), font-weight: 600
h4: 1.25rem (20px), font-weight: 600

/* Body Text */
body: 1rem (16px), font-weight: 400
small: 0.875rem (14px), font-weight: 400
```

---

### Spacing System

```
4px   (space-1)
8px   (space-2)
12px  (space-3)
16px  (space-4)
24px  (space-6)
32px  (space-8)
48px  (space-12)
64px  (space-16)
```

---

### Component Design

**Buttons:**

```
Primary:   Blue background, white text, rounded-lg
Secondary: Gray background, gray-900 text
Success:   Green background, white text
Danger:    Red background, white text

Sizes:
- Small:   px-3 py-1.5 text-sm
- Medium:  px-4 py-2 text-base
- Large:   px-6 py-3 text-lg
```

**Cards:**

```
- White background
- Border: 1px solid gray-200
- Rounded: rounded-lg
- Shadow: shadow-sm
- Padding: p-6
- Hover: shadow-md transition
```

**Inputs:**

```
- Border: 1px solid gray-300
- Rounded: rounded-md
- Padding: px-3 py-2
- Focus: ring-2 ring-blue-500
- Error: border-red-500 ring-red-500
```

---

### Animations

```css
/* Transitions */
transition: all 0.2s ease-in-out;

/* Hover Effects */
.card:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
}

/* Loading Spinner */
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* Fade In */
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

/* Slide In */
@keyframes slideIn {
  from {
    transform: translateY(10px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}
```

---

### Responsive Breakpoints

```css
/* Mobile First Approach */
sm:  640px   /* Tablet */
md:  768px   /* Small laptop */
lg:  1024px  /* Desktop */
xl:  1280px  /* Large desktop */
2xl: 1536px  /* Extra large */
```

---

## 🔌 6. API INTEGRATION

### Axios Configuration

**File:** `lib/api/axios.ts`

```typescript
import axios from "axios";
import Cookies from "js-cookie";

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api",
  headers: {
    "Content-Type": "application/json",
  },
});

// Request interceptor (add JWT token)
api.interceptors.request.use(
  (config) => {
    const token = Cookies.get("auth_token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor (handle errors)
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Unauthorized - redirect to login
      Cookies.remove("auth_token");
      window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);

export default api;
```

---

### API Functions Structure

**Example:** `lib/api/lessons.ts`

```typescript
import api from "./axios";
import { Lesson, CreateLessonRequest } from "@/lib/types/lesson";

export const lessonsAPI = {
  // Get all lessons
  getAll: async () => {
    const { data } = await api.get<Lesson[]>("/lessons");
    return data;
  },

  // Get lesson by ID
  getById: async (id: number) => {
    const { data } = await api.get<Lesson>(`/lessons/${id}`);
    return data;
  },

  // Create lesson (teacher only)
  create: async (lesson: CreateLessonRequest) => {
    const { data } = await api.post<Lesson>("/lessons", lesson);
    return data;
  },

  // Update lesson (teacher only)
  update: async (id: number, lesson: Partial<CreateLessonRequest>) => {
    const { data } = await api.put<Lesson>(`/lessons/${id}`, lesson);
    return data;
  },

  // Delete lesson (teacher only)
  delete: async (id: number) => {
    await api.delete(`/lessons/${id}`);
  },

  // Enroll to lesson (student only)
  enroll: async (id: number) => {
    const { data } = await api.post(`/lessons/${id}/enroll`);
    return data;
  },

  // Get enrolled lessons (student only)
  getEnrolled: async () => {
    const { data } = await api.get("/my-lessons");
    return data;
  },
};
```

---

### React Query Integration

**Custom Hook:** `lib/hooks/useLessons.ts`

```typescript
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { lessonsAPI } from "@/lib/api/lessons";
import { useToast } from "@/components/ui/use-toast";

export const useLessons = () => {
  const queryClient = useQueryClient();
  const { toast } = useToast();

  // Get all lessons
  const { data: lessons, isLoading } = useQuery({
    queryKey: ["lessons"],
    queryFn: lessonsAPI.getAll,
  });

  // Create lesson
  const createMutation = useMutation({
    mutationFn: lessonsAPI.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["lessons"] });
      toast({ title: "Lesson created successfully!" });
    },
    onError: (error: any) => {
      toast({
        title: "Error",
        description: error.response?.data?.error || "Failed to create lesson",
        variant: "destructive",
      });
    },
  });

  // Delete lesson
  const deleteMutation = useMutation({
    mutationFn: lessonsAPI.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["lessons"] });
      toast({ title: "Lesson deleted successfully!" });
    },
  });

  return {
    lessons,
    isLoading,
    createLesson: createMutation.mutate,
    deleteLesson: deleteMutation.mutate,
    isCreating: createMutation.isPending,
    isDeleting: deleteMutation.isPending,
  };
};
```

---

### Error Handling Strategy

```typescript
// Error types
interface APIError {
  message: string;
  code?: string;
  field?: string;
}

// Error handler utility
export const handleAPIError = (error: any): string => {
  if (error.response) {
    // Server responded with error
    const { status, data } = error.response;

    if (status === 400) return data.error || "Invalid request";
    if (status === 401) return "Unauthorized. Please login.";
    if (status === 403) return "Access denied";
    if (status === 404) return "Resource not found";
    if (status === 500) return "Server error. Please try again.";

    return data.error || "Something went wrong";
  }

  if (error.request) {
    // Network error
    return "Network error. Check your connection.";
  }

  return "An unexpected error occurred";
};
```

---

## ✅ 7. TESTING STRATEGY

### Unit Testing (Jest + React Testing Library)

**Test Files Structure:**

```
components/
├── auth/
│   ├── LoginForm.tsx
│   └── LoginForm.test.tsx
├── teacher/
│   ├── LessonCard.tsx
│   └── LessonCard.test.tsx
```

**Example Test:**

```typescript
// LoginForm.test.tsx
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { LoginForm } from "./LoginForm";

describe("LoginForm", () => {
  it("should render email and password inputs", () => {
    render(<LoginForm />);

    expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
  });

  it("should show validation error for invalid email", async () => {
    render(<LoginForm />);

    const emailInput = screen.getByLabelText(/email/i);
    fireEvent.change(emailInput, { target: { value: "invalid" } });
    fireEvent.blur(emailInput);

    await waitFor(() => {
      expect(screen.getByText(/invalid email/i)).toBeInTheDocument();
    });
  });

  it("should submit form with valid data", async () => {
    const onSubmit = jest.fn();
    render(<LoginForm onSubmit={onSubmit} />);

    fireEvent.change(screen.getByLabelText(/email/i), {
      target: { value: "test@example.com" },
    });
    fireEvent.change(screen.getByLabelText(/password/i), {
      target: { value: "password123" },
    });

    fireEvent.click(screen.getByRole("button", { name: /login/i }));

    await waitFor(() => {
      expect(onSubmit).toHaveBeenCalledWith({
        email: "test@example.com",
        password: "password123",
      });
    });
  });
});
```

---

### Integration Testing (Playwright)

**Test Scenarios:**

```typescript
// tests/auth.spec.ts
import { test, expect } from "@playwright/test";

test.describe("Authentication Flow", () => {
  test("should login as teacher", async ({ page }) => {
    await page.goto("/login");

    await page.fill('input[name="email"]', "teacher@example.com");
    await page.fill('input[name="password"]', "password123");
    await page.click('button[type="submit"]');

    await expect(page).toHaveURL("/teacher/dashboard");
    await expect(page.locator("h1")).toContainText("Teacher Dashboard");
  });

  test("should prevent access to teacher routes for students", async ({ page }) => {
    // Login as student
    await page.goto("/login");
    await page.fill('input[name="email"]', "student@example.com");
    await page.fill('input[name="password"]', "password123");
    await page.click('button[type="submit"]');

    // Try to access teacher route
    await page.goto("/teacher/lessons/create");

    // Should redirect to unauthorized or student dashboard
    await expect(page).not.toHaveURL("/teacher/lessons/create");
  });
});

// tests/lesson-flow.spec.ts
test.describe("Teacher - Lesson Management", () => {
  test("should create new lesson", async ({ page }) => {
    // Login as teacher first
    await page.goto("/login");
    await page.fill('input[name="email"]', "teacher@example.com");
    await page.fill('input[name="password"]', "password123");
    await page.click('button[type="submit"]');

    // Navigate to create lesson
    await page.goto("/teacher/lessons/create");

    // Fill form
    await page.fill('input[name="title"]', "Test Lesson");
    await page.fill('textarea[name="description"]', "Test description");
    await page.selectOption('select[name="category"]', "python");
    await page.selectOption('select[name="difficulty"]', "beginner");

    // Submit
    await page.click('button[type="submit"]');

    // Should redirect to lessons list
    await expect(page).toHaveURL("/teacher/lessons");
    await expect(page.locator("text=Test Lesson")).toBeVisible();
  });
});
```

---

### Manual Testing Checklist

**Authentication:**

- [ ] Login with valid credentials (teacher)
- [ ] Login with valid credentials (student)
- [ ] Login with invalid credentials (show error)
- [ ] Register as teacher
- [ ] Register as student
- [ ] Logout
- [ ] Protected routes redirect to login
- [ ] JWT token persists on refresh

**Teacher Flow:**

- [ ] Create lesson
- [ ] Edit lesson
- [ ] Delete lesson (with confirmation)
- [ ] Create course within lesson
- [ ] Edit course
- [ ] Delete course
- [ ] Create all 5 PRIMM stages
- [ ] Edit stages
- [ ] Delete stages
- [ ] View lesson/course/stage list

**Student Flow:**

- [ ] Browse all lessons
- [ ] Filter lessons (category, difficulty)
- [ ] Search lessons
- [ ] Enroll to lesson
- [ ] View enrolled lessons with progress
- [ ] View courses in lesson
- [ ] View all 5 stages in course
- [ ] Submit PREDICT stage
- [ ] Submit RUN stage
- [ ] Submit INVESTIGATE stage
- [ ] Submit MODIFY stage
- [ ] Submit MAKE stage
- [ ] See submission feedback
- [ ] See progress update
- [ ] See coins awarded
- [ ] Coin balance updates

**UI/UX:**

- [ ] All pages responsive (mobile/tablet/desktop)
- [ ] Loading states visible
- [ ] Error messages clear
- [ ] Success messages visible
- [ ] Animations smooth
- [ ] Forms validate properly
- [ ] Buttons have hover states
- [ ] Empty states show illustrations

---

## 🚀 8. DEPLOYMENT PLAN

### Environment Setup

**Development:**

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
NODE_ENV=development
```

**Production:**

```env
NEXT_PUBLIC_API_URL=https://api.primmfy.com/api
NODE_ENV=production
NEXT_PUBLIC_SITE_URL=https://primmfy.com
```

---

### Build Process

```bash
# Install dependencies
npm install

# Build for production
npm run build

# Start production server
npm start
```

---

### Deployment Options

#### Option 1: Vercel (Recommended for Next.js)

```bash
# Install Vercel CLI
npm i -g vercel

# Deploy
vercel

# Production deployment
vercel --prod
```

**Vercel Configuration:**

- Automatic deployments from Git
- Preview deployments for PRs
- Environment variables in dashboard
- Edge functions support
- Automatic HTTPS

---

#### Option 2: Docker + VPS

**Dockerfile:**

```dockerfile
FROM node:20-alpine AS base

# Install dependencies
FROM base AS deps
WORKDIR /app
COPY package*.json ./
RUN npm ci

# Build application
FROM base AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build

# Production image
FROM base AS runner
WORKDIR /app
ENV NODE_ENV production

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000

ENV PORT 3000

CMD ["node", "server.js"]
```

**docker-compose.yml:**

```yaml
version: "3.8"

services:
  frontend:
    build: .
    ports:
      - "3000:3000"
    environment:
      - NEXT_PUBLIC_API_URL=http://backend:8080/api
    depends_on:
      - backend
    restart: unless-stopped

  backend:
    build: ./Backend
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
    depends_on:
      - postgres
    restart: unless-stopped

  postgres:
    image: postgres:15
    environment:
      - POSTGRES_DB=primmfy_db
      - POSTGRES_PASSWORD=yourpassword
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

volumes:
  postgres_data:
```

---

### Performance Optimization

**Next.js Config:**

```javascript
// next.config.js
module.exports = {
  // Enable compiler optimizations
  compiler: {
    removeConsole: process.env.NODE_ENV === "production",
  },

  // Image optimization
  images: {
    domains: ["your-cdn.com"],
    formats: ["image/avif", "image/webp"],
  },

  // Compression
  compress: true,

  // Headers for caching
  async headers() {
    return [
      {
        source: "/static/:path*",
        headers: [
          {
            key: "Cache-Control",
            value: "public, max-age=31536000, immutable",
          },
        ],
      },
    ];
  },
};
```

**Bundle Analysis:**

```bash
# Analyze bundle size
npm run build
npx @next/bundle-analyzer
```

---

### Monitoring & Analytics

**Setup:**

```typescript
// lib/analytics.ts
export const analytics = {
  pageView: (url: string) => {
    if (window.gtag) {
      window.gtag("config", "GA_MEASUREMENT_ID", {
        page_path: url,
      });
    }
  },

  event: (action: string, params?: any) => {
    if (window.gtag) {
      window.gtag("event", action, params);
    }
  },
};

// Usage in pages
useEffect(() => {
  analytics.pageView(window.location.pathname);
}, []);
```

---

## 📊 PROGRESS TRACKING

### Phase Checklist

```
Phase 1: Foundation & Authentication
├── [x] Project initialization
├── [x] Dependencies installed
├── [x] TailwindCSS configured
├── [ ] shadcn/ui setup
├── [ ] TypeScript types
├── [ ] Axios client
├── [ ] Auth Context
├── [ ] Login page
├── [ ] Register page
├── [ ] Protected routes
└── [ ] Testing

Phase 2: Teacher - Lesson Management
├── [ ] Dashboard layout
├── [ ] My Lessons page
├── [ ] Create Lesson
├── [ ] Edit Lesson
├── [ ] Delete Lesson
├── [ ] Filters & Search
└── [ ] Testing

Phase 3: Teacher - Course Management
├── [ ] Lesson Detail page
├── [ ] Create Course
├── [ ] Edit Course
├── [ ] Delete Course
└── [ ] Testing

Phase 4: Teacher - PRIMM Stages
├── [ ] PREDICT form
├── [ ] RUN form
├── [ ] INVESTIGATE form
├── [ ] MODIFY form
├── [ ] MAKE form
├── [ ] Monaco Editor
├── [ ] Stage preview
└── [ ] Testing

Phase 5: Student - Browse & Enroll
├── [ ] Student dashboard
├── [ ] Browse lessons
├── [ ] Filters & Search
├── [ ] Enroll functionality
├── [ ] My Lessons page
└── [ ] Testing

Phase 6: Student - Learning Flow
├── [ ] Lesson detail
├── [ ] Course learning page
├── [ ] PREDICT submission
├── [ ] RUN submission
├── [ ] INVESTIGATE submission
├── [ ] MODIFY submission
├── [ ] MAKE submission
├── [ ] Progress tracking
├── [ ] Coin rewards
└── [ ] Testing

Phase 7: UI/UX Polish
├── [ ] Loading states
├── [ ] Animations
├── [ ] Toast notifications
├── [ ] Empty states
├── [ ] Responsive design
├── [ ] Accessibility
└── [ ] Performance

Phase 8: Testing & Deployment
├── [ ] Unit tests
├── [ ] Integration tests
├── [ ] Manual testing
├── [ ] Bug fixes
├── [ ] Build optimization
└── [ ] Deployment
```

---

## 📝 NOTES & CONSIDERATIONS

### Important Decisions Made

1. **Next.js 14 App Router** - Modern approach with Server Components
2. **TailwindCSS + shadcn/ui** - Consistent, customizable UI
3. **TypeScript** - Type safety, better DX
4. **React Query** - Efficient server state management
5. **Monaco Editor** - Professional code editing experience
6. **JWT in httpOnly cookies** - Secure token storage

### Potential Challenges

1. **Monaco Editor** - Large bundle size (lazy load)
2. **Real-time code execution** - May need sandbox API later
3. **Video embedding** - YouTube/Vimeo iframe security
4. **File uploads** - Thumbnail images (future)
5. **Mobile code editing** - Monaco may need mobile alternative

### Future Enhancements

- [ ] Real-time collaboration
- [ ] Chat/discussion per course
- [ ] Video recording for teachers
- [ ] Peer code review
- [ ] Advanced analytics dashboard
- [ ] Gamification (badges, leaderboard)
- [ ] Content import/export
- [ ] API documentation (Swagger)
- [ ] Mobile app (React Native)
- [ ] AI code suggestions

---

## 📚 RESOURCES

### Documentation

- **Next.js:** https://nextjs.org/docs
- **React:** https://react.dev
- **TailwindCSS:** https://tailwindcss.com/docs
- **shadcn/ui:** https://ui.shadcn.com
- **React Query:** https://tanstack.com/query/latest
- **Monaco Editor:** https://microsoft.github.io/monaco-editor/

### Design Inspiration

- **Dribbble:** https://dribbble.com/tags/education
- **Behance:** https://www.behance.net/search/projects?search=learning+platform
- **Coursera:** https://www.coursera.org
- **Udemy:** https://www.udemy.com
- **Khan Academy:** https://www.khanacademy.org

### Tools

- **Figma:** UI/UX design mockups
- **Excalidraw:** Quick wireframes
- **VS Code:** Development
- **Chrome DevTools:** Debugging
- **React DevTools:** Component inspection

---

## ✅ SUCCESS CRITERIA

**Frontend is considered complete when:**

1. ✅ All 24 backend test flows work through UI
2. ✅ Authentication flow secure & smooth
3. ✅ Teacher can create lessons/courses/stages easily
4. ✅ Student can learn & submit all 5 PRIMM stages
5. ✅ Progress tracking accurate
6. ✅ Coin rewards work automatically
7. ✅ UI responsive on all devices
8. ✅ Loading/error states handled properly
9. ✅ Performance optimized (<3s page load)
10. ✅ No critical bugs
11. ✅ Accessible (WCAG AA)
12. ✅ Tests passing (>80% coverage)

---

**Created:** October 26, 2025  
**Last Updated:** October 26, 2025  
**Status:** Planning Complete ✅  
**Next Step:** Phase 1 - Authentication Implementation 🚀

---

**END OF FRONTEND DEVELOPMENT PLAN**
