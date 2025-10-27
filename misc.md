# 🎓 Platform E-Learning PRIMM - Panduan Lengkap untuk Pemula

## 📚 Daftar Isi

1. [Pengenalan Proyek](#pengenalan-proyek)
2. [Apa itu PRIMM?](#apa-itu-primm)
3. [Arsitektur Aplikasi](#arsitektur-aplikasi)
4. [Teknologi yang Digunakan](#teknologi-yang-digunakan)
5. [Penjelasan Database](#penjelasan-database)
6. [Setup Awal](#setup-awal)
7. [Fase Pengembangan](#fase-pengembangan)
8. [Struktur Folder](#struktur-folder)
9. [Tutorial Pemula](#tutorial-pemula)

---

## 🎯 Pengenalan Proyek

Platform e-learning ini adalah aplikasi web untuk belajar pemrograman dengan pendekatan **PRIMM** (Predict, Run, Investigate, Modify, Make). Aplikasi ini memiliki **sistem gamifikasi** di mana siswa mendapatkan koin untuk membeli item dan menghias avatar mereka.

### Fitur Utama:

- 🔐 **Autentikasi User**: Registrasi dan login yang aman
- 📖 **Sistem Pembelajaran**: Kursus, lessons, dan tahapan PRIMM
- 🎮 **Gamifikasi**: Koin, level, experience points
- 🛒 **Toko Virtual**: Beli item dengan koin
- 👤 **Kustomisasi Avatar**: Personalisasi karakter
- 🏆 **Achievement System**: Pencapaian dan badge

---

## 🤔 Apa itu PRIMM?

PRIMM adalah metodologi pembelajaran pemrograman yang terdiri dari 5 tahap:

1. **Predict (Prediksi)**: Siswa memprediksi apa yang akan dilakukan kode
2. **Run (Jalankan)**: Siswa menjalankan kode dan melihat hasilnya
3. **Investigate (Investigasi)**: Siswa menganalisis bagaimana kode bekerja
4. **Modify (Modifikasi)**: Siswa memodifikasi kode yang ada
5. **Make (Buat)**: Siswa membuat kode sendiri dari awal

Setiap tahap yang diselesaikan memberikan koin dan experience points!

---

## 🏗️ Arsitektur Aplikasi

Aplikasi ini menggunakan arsitektur **Client-Server**:

```
┌─────────────────┐         HTTP/HTTPS          ┌─────────────────┐
│                 │  ◄─────────────────────────► │                 │
│  Frontend       │         REST API             │   Backend       │
│  (Next.js)      │    (JSON Request/Response)   │   (Golang)      │
│                 │                              │                 │
└─────────────────┘                              └────────┬────────┘
                                                          │
                                                          │ SQL
                                                          │
                                                   ┌──────▼─────────┐
                                                   │   PostgreSQL   │
                                                   │   Database     │
                                                   └────────────────┘
```

### Penjelasan Komponen:

#### 1. **Frontend (Next.js + Tailwind CSS)**

- **Apa itu?**: Bagian yang dilihat dan diinteraksi oleh user (tampilan website)
- **Teknologi**:
  - **Next.js**: Framework React untuk membuat aplikasi web modern
  - **Tailwind CSS**: Framework CSS untuk styling yang cepat dan responsive
- **Tugas**: Menampilkan data, menangani input user, mengirim request ke backend

#### 2. **Backend (Golang + Gin)**

- **Apa itu?**: Bagian yang menangani logika bisnis dan berkomunikasi dengan database
- **Teknologi**:
  - **Go (Golang)**: Bahasa pemrograman yang cepat dan efisien
  - **Gin**: Framework web untuk Go (seperti Express.js di Node.js)
- **Tugas**:
  - Memproses request dari frontend
  - Validasi data
  - Berkomunikasi dengan database
  - Mengirim response ke frontend

#### 3. **Database (PostgreSQL)**

- **Apa itu?**: Tempat penyimpanan data secara permanen
- **Teknologi**: PostgreSQL (database relational yang powerful)
- **Tugas**: Menyimpan data users, courses, progress, dll

---

## 💻 Teknologi yang Digunakan

### Backend Stack:

| Teknologi       | Kegunaan                   | Analogi                            |
| --------------- | -------------------------- | ---------------------------------- |
| **Go (Golang)** | Bahasa pemrograman backend | "Otak" yang memproses logika       |
| **Gin**         | Web framework              | "Jalanan" yang mengarahkan traffic |
| **pgx**         | PostgreSQL driver          | "Jembatan" ke database             |
| **JWT**         | Autentikasi                | "KTP digital" untuk user           |
| **bcrypt**      | Hashing password           | "Brankas" untuk password           |

### Frontend Stack:

| Teknologi        | Kegunaan               | Analogi                    |
| ---------------- | ---------------------- | -------------------------- |
| **Next.js**      | React framework        | "Kerangka" aplikasi web    |
| **TypeScript**   | JavaScript dengan type | JavaScript yang lebih aman |
| **Tailwind CSS** | Styling framework      | "Cat dan dekorasi" website |
| **Axios**        | HTTP client            | "Kurir" yang kirim data    |

### Database:

| Teknologi      | Kegunaan            |
| -------------- | ------------------- |
| **PostgreSQL** | Relational database |

---

## 🗄️ Penjelasan Database

### Mengapa Butuh Database?

Database adalah tempat kita menyimpan semua data aplikasi. Tanpa database, setiap kali server restart, semua data akan hilang!

### Struktur Database Kita

#### 1. **Tabel `users`** - Menyimpan Data Pengguna

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,           -- ID unik otomatis
    username VARCHAR(50),             -- Nama user untuk login
    email VARCHAR(255),               -- Email user
    password_hash VARCHAR(255),       -- Password (di-encrypt!)
    full_name VARCHAR(100),           -- Nama lengkap
    total_coins INTEGER DEFAULT 0,    -- Jumlah koin yang dimiliki
    level INTEGER DEFAULT 1,          -- Level user
    experience_points INTEGER,        -- XP untuk naik level
    created_at TIMESTAMP              -- Kapan akun dibuat
);
```

**Analogi**: Seperti kartu identitas mahasiswa yang menyimpan info pribadi dan progress belajar.

#### 2. **Tabel `courses`** - Menyimpan Kursus

```sql
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),               -- Judul kursus (misal: "Python Basics")
    description TEXT,                 -- Penjelasan kursus
    difficulty_level VARCHAR(20),     -- beginner/intermediate/advanced
    programming_language VARCHAR(50)  -- Bahasa pemrograman
);
```

**Analogi**: Seperti mata kuliah di universitas (Database, Web Programming, dll).

#### 3. **Tabel `lessons`** - Menyimpan Pelajaran

```sql
CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    course_id INTEGER,                -- Pelajaran ini bagian dari kursus apa?
    title VARCHAR(255),               -- Judul pelajaran
    order_number INTEGER,             -- Urutan (Lesson 1, 2, 3...)
    code_example TEXT,                -- Contoh kode
    FOREIGN KEY (course_id) REFERENCES courses(id)
);
```

**Analogi**: Seperti bab-bab dalam sebuah buku pelajaran.

**Konsep Penting - Foreign Key:**

- `course_id` adalah **Foreign Key** yang menghubungkan lesson ke course
- Artinya: Setiap lesson **harus** terikat ke satu course
- Relasi: **One-to-Many** (1 course bisa punya banyak lessons)

#### 4. **Tabel `primm_stages`** - Menyimpan Tahapan PRIMM

```sql
CREATE TABLE primm_stages (
    id SERIAL PRIMARY KEY,
    lesson_id INTEGER,                -- Tahap ini bagian dari lesson apa?
    stage_type VARCHAR(20),           -- predict/run/investigate/modify/make
    instruction TEXT,                 -- Instruksi untuk siswa
    coin_reward INTEGER,              -- Berapa koin yang didapat?
    FOREIGN KEY (lesson_id) REFERENCES lessons(id)
);
```

**Analogi**: Seperti soal-soal latihan dalam sebuah bab.

#### 5. **Tabel `user_progress`** - Tracking Progress User

```sql
CREATE TABLE user_progress (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,                  -- Siapa yang mengerjakan?
    stage_id INTEGER,                 -- Tahap apa yang dikerjakan?
    is_completed BOOLEAN,             -- Sudah selesai?
    user_code TEXT,                   -- Kode yang ditulis user
    coins_earned INTEGER,             -- Koin yang didapat
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (stage_id) REFERENCES primm_stages(id)
);
```

**Analogi**: Seperti rapor yang mencatat nilai dan progres belajar siswa.

**Konsep Penting - Many-to-Many:**

- 1 user bisa mengerjakan banyak stages
- 1 stage bisa dikerjakan banyak users
- Tabel `user_progress` adalah **junction table** yang menghubungkan keduanya

#### 6. **Tabel `avatar_components`** - Komponen Avatar

```sql
CREATE TABLE avatar_components (
    id SERIAL PRIMARY KEY,
    category VARCHAR(50),             -- 'head', 'body', 'accessory'
    name VARCHAR(100),                -- Nama item (misal: "Cool Hat")
    image_url VARCHAR(255),           -- URL gambar
    rarity VARCHAR(20)                -- common/rare/epic/legendary
);
```

**Analogi**: Seperti katalog item di game (topi, baju, aksesoris).

#### 7. **Tabel `shop_items`** - Item di Toko

```sql
CREATE TABLE shop_items (
    id SERIAL PRIMARY KEY,
    avatar_component_id INTEGER,      -- Item apa yang dijual?
    price INTEGER,                    -- Harga dalam koin
    is_available BOOLEAN,             -- Masih dijual atau tidak?
    FOREIGN KEY (avatar_component_id) REFERENCES avatar_components(id)
);
```

#### 8. **Tabel `user_inventory`** - Item yang Dimiliki User

```sql
CREATE TABLE user_inventory (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,                  -- Siapa pemiliknya?
    avatar_component_id INTEGER,      -- Item apa yang dimiliki?
    purchased_at TIMESTAMP,           -- Kapan dibeli?
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (avatar_component_id) REFERENCES avatar_components(id)
);
```

**Analogi**: Seperti lemari yang menyimpan koleksi item yang sudah dibeli.

### Diagram Relasi Database

```
users ──┬─── user_progress ─── primm_stages ─── lessons ─── courses
        │
        └─── user_inventory ─── avatar_components
                                        │
                                        └─── shop_items
```

---

## 🚀 Setup Awal

### Prasyarat

Pastikan sudah terinstall:

- **Go** (v1.21+): [Download](https://golang.org/dl/)
- **Node.js** (v18+): [Download](https://nodejs.org/)
- **PostgreSQL** (v14+): [Download](https://www.postgresql.org/download/)
- **Git**: [Download](https://git-scm.com/)

### Step 1: Setup Database

1. **Buka PostgreSQL** (pgAdmin atau psql)
2. **Buat database baru:**

```sql
CREATE DATABASE primmfy;
```

3. **Jalankan script SQL** untuk membuat tabel (gunakan script yang sudah saya berikan di atas)

### Step 2: Clone & Setup Backend

```bash
# Masuk ke folder project
cd d:\Project\PRIMMFY

# Install dependencies Go
go mod tidy

# Buat file .env
# (Copy isi dari contoh di bawah)
```

**File `.env`:**

```env
DATABASE_URL=postgres://postgres:dhika001@localhost:5432/primmfy?sslmode=disable
JWT_SECRET=rahasia-super-aman-ganti-ini
PORT=8080
```

**Penjelasan Environment Variables:**

- `DATABASE_URL`: Alamat dan kredensial database
- `JWT_SECRET`: Kunci rahasia untuk enkripsi token (HARUS DIGANTI!)
- `PORT`: Port di mana server akan berjalan

### Step 3: Setup Frontend

```bash
# Buat folder frontend
cd d:\Project\PRIMMFY
npx create-next-app@latest Frontend --typescript --tailwind --app

# Masuk ke folder frontend
cd Frontend

# Install dependencies tambahan
npm install axios
npm install @types/node
```

---

## 📋 Fase Pengembangan

### **FASE 1: Foundation & Authentication** ⭐ (Mulai dari sini!)

#### Apa yang Akan Dipelajari?

- Cara membuat REST API sederhana
- Konsep autentikasi dan keamanan
- Hash password menggunakan bcrypt
- JSON Web Token (JWT)
- HTTP methods (GET, POST)
- Database CRUD operations

#### Struktur File yang Akan Dibuat:

```
PRIMMFY/
├── main.go                  # Entry point aplikasi
├── database.go              # Koneksi database (SUDAH ADA)
├── config/
│   └── config.go           # Konfigurasi environment
├── models/
│   └── user.go             # Struct User dan request models
├── services/
│   └── auth_service.go     # Logika autentikasi
├── handlers/
│   └── auth_handler.go     # Handler untuk endpoint API
└── middleware/
    └── auth_middleware.go  # Middleware untuk proteksi route
```

#### Penjelasan Konsep:

##### 1. **MVC Pattern (Model-View-Controller)**

Kita menggunakan pola desain yang memisahkan aplikasi menjadi:

- **Models**: Struktur data (struct di Go)
- **Controllers/Handlers**: Menangani HTTP request
- **Services**: Logika bisnis
- **Views**: Frontend (Next.js)

##### 2. **REST API**

API adalah cara backend dan frontend berkomunikasi.

**Contoh Endpoint:**

```
POST /api/auth/register  → Daftar akun baru
POST /api/auth/login     → Login ke akun
GET  /api/auth/profile   → Lihat profil (butuh token)
```

**Request & Response:**

```
[Frontend]                [Backend]
   │                         │
   ├─ POST /api/auth/login ─►│
   │  Body: {                │
   │    "email": "...",      │
   │    "password": "..."    │
   │  }                      │
   │                         │
   │◄─ Response: 200 OK ─────┤
      Body: {                │
        "token": "...",      │
        "user": {...}        │
      }                      │
```

##### 3. **Password Hashing**

❌ **JANGAN PERNAH** simpan password dalam bentuk plain text!
✅ Gunakan **bcrypt** untuk hash password.

**Cara Kerja:**

```
Password User: "rahasia123"
      ↓ (bcrypt hash)
Hash: "$2a$14$xyz...abc" → Disimpan di database

Saat Login:
Password Input: "rahasia123"
      ↓ (bcrypt compare)
Cocok dengan hash? → Login sukses!
```

##### 4. **JWT (JSON Web Token)**

JWT adalah "tiket" yang membuktikan user sudah login.

**Cara Kerja:**

```
1. User login → Server cek email & password
2. Jika benar → Server buat JWT token
3. Frontend simpan token (localStorage/cookie)
4. Setiap request → Frontend kirim token di header
5. Backend cek token → Tahu siapa user-nya
```

**Contoh Token:**

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2OTk5OTk5OTl9.xyz...
```

#### Tutorial Langkah-demi-Langkah:

##### **Langkah 1.1: Buat Models**

**File: `models/user.go`**

```go
package models

import "time"

// User adalah struct yang merepresentasikan tabel users
// Setiap field punya "tag" json untuk serialize ke JSON
type User struct {
    ID              int       `json:"id"`
    Username        string    `json:"username"`
    Email           string    `json:"email"`
    PasswordHash    string    `json:"-"` // "-" artinya tidak di-export ke JSON
    FullName        string    `json:"full_name"`
    TotalCoins      int       `json:"total_coins"`
    Level           int       `json:"level"`
    ExperiencePoints int      `json:"experience_points"`
    CreatedAt       time.Time `json:"created_at"`
}

// RegisterRequest adalah data yang dikirim saat register
type RegisterRequest struct {
    Username string `json:"username" binding:"required,min=3"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
    FullName string `json:"full_name" binding:"required"`
}

// LoginRequest adalah data yang dikirim saat login
type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// AuthResponse adalah response setelah login/register berhasil
type AuthResponse struct {
    Token string `json:"token"`
    User  User   `json:"user"`
}
```

**Penjelasan:**

- `struct`: Cara Go mendefinisikan tipe data kompleks (seperti class)
- `json:"..."`: Tag untuk mapping field ke JSON
- `binding:"required"`: Validasi otomatis oleh Gin

##### **Langkah 1.2: Buat Service untuk Hash & JWT**

**File: `services/auth_service.go`**

```go
package services

import (
    "context"
    "errors"
    "time"
    "os"

    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    "primmfy-backend/models"
    "primmfy-backend/config"
)

// Ambil secret dari environment variable
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Claims adalah data yang disimpan di dalam JWT
type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}

// HashPassword mengubah password plain text menjadi hash
func HashPassword(password string) (string, error) {
    // Cost 14 = tingkat kesulitan (semakin tinggi semakin lambat tapi aman)
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPasswordHash membandingkan password dengan hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil // true jika cocok, false jika tidak
}

// GenerateJWT membuat token JWT untuk user
func GenerateJWT(userID int) (string, error) {
    // Token berlaku 24 jam
    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    // Buat token dengan algoritma HS256
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign token dengan secret key
    return token.SignedString(jwtSecret)
}

// RegisterUser membuat akun user baru
func RegisterUser(req models.RegisterRequest) (*models.User, error) {
    // 1. Cek apakah email/username sudah ada
    var exists bool
    err := config.DB.QueryRow(context.Background(),
        "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 OR username = $2)",
        req.Email, req.Username).Scan(&exists)

    if err != nil {
        return nil, err
    }

    if exists {
        return nil, errors.New("email atau username sudah digunakan")
    }

    // 2. Hash password
    hashedPassword, err := HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    // 3. Insert user baru ke database
    var user models.User
    err = config.DB.QueryRow(context.Background(), `
        INSERT INTO users (username, email, password_hash, full_name)
        VALUES ($1, $2, $3, $4)
        RETURNING id, username, email, full_name, total_coins, level,
                  experience_points, created_at`,
        req.Username, req.Email, hashedPassword, req.FullName).Scan(
        &user.ID, &user.Username, &user.Email, &user.FullName,
        &user.TotalCoins, &user.Level, &user.ExperiencePoints, &user.CreatedAt)

    if err != nil {
        return nil, err
    }

    return &user, nil
}

// LoginUser memverifikasi kredensial dan return user
func LoginUser(req models.LoginRequest) (*models.User, error) {
    var user models.User
    var passwordHash string

    // 1. Cari user berdasarkan email
    err := config.DB.QueryRow(context.Background(), `
        SELECT id, username, email, password_hash, full_name, total_coins,
               level, experience_points, created_at
        FROM users WHERE email = $1`, req.Email).Scan(
        &user.ID, &user.Username, &user.Email, &passwordHash, &user.FullName,
        &user.TotalCoins, &user.Level, &user.ExperiencePoints, &user.CreatedAt)

    if err != nil {
        return nil, errors.New("email atau password salah")
    }

    // 2. Cek apakah password cocok
    if !CheckPasswordHash(req.Password, passwordHash) {
        return nil, errors.New("email atau password salah")
    }

    return &user, nil
}
```

**Konsep Penting:**

- **Context**: Di Go, context digunakan untuk kontrol lifecycle request
- **$1, $2**: Placeholder untuk prepared statements (mencegah SQL injection)
- **RETURNING**: Keyword PostgreSQL untuk langsung mendapat data yang di-insert

##### **Langkah 1.3: Buat Handlers**

**File: `handlers/auth_handler.go`**

```go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "primmfy-backend/models"
    "primmfy-backend/services"
)

// Register handler untuk endpoint POST /api/auth/register
func Register(c *gin.Context) {
    var req models.RegisterRequest

    // 1. Bind JSON dari request body ke struct
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data tidak valid",
            "details": err.Error(),
        })
        return
    }

    // 2. Panggil service untuk register user
    user, err := services.RegisterUser(req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // 3. Generate JWT token
    token, err := services.GenerateJWT(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Gagal membuat token",
        })
        return
    }

    // 4. Return response sukses
    c.JSON(http.StatusCreated, models.AuthResponse{
        Token: token,
        User:  *user,
    })
}

// Login handler untuk endpoint POST /api/auth/login
func Login(c *gin.Context) {
    var req models.LoginRequest

    // 1. Bind JSON request
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data tidak valid",
        })
        return
    }

    // 2. Verifikasi kredensial
    user, err := services.LoginUser(req)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": err.Error(),
        })
        return
    }

    // 3. Generate token
    token, err := services.GenerateJWT(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Gagal membuat token",
        })
        return
    }

    // 4. Return response
    c.JSON(http.StatusOK, models.AuthResponse{
        Token: token,
        User:  *user,
    })
}
```

**HTTP Status Codes:**

- `200 OK`: Request berhasil
- `201 Created`: Resource baru dibuat (register berhasil)
- `400 Bad Request`: Data yang dikirim tidak valid
- `401 Unauthorized`: Autentikasi gagal
- `500 Internal Server Error`: Error di server

##### **Langkah 1.4: Buat Middleware untuk Proteksi Route**

**File: `middleware/auth_middleware.go`**

```go
package middleware

import (
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "primmfy-backend/services"
)

// AuthMiddleware memverifikasi JWT token di setiap request
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. Ambil token dari header Authorization
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Token tidak ditemukan",
            })
            c.Abort() // Stop request di sini
            return
        }

        // 2. Token biasanya dalam format: "Bearer <token>"
        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

        // 3. Parse dan validasi token
        token, err := jwt.ParseWithClaims(tokenString, &services.Claims{},
            func(token *jwt.Token) (interface{}, error) {
                return []byte(os.Getenv("JWT_SECRET")), nil
            })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Token tidak valid",
            })
            c.Abort()
            return
        }

        // 4. Ambil user ID dari token
        if claims, ok := token.Claims.(*services.Claims); ok {
            // Simpan user ID di context untuk diakses handler
            c.Set("userID", claims.UserID)
            c.Next() // Lanjutkan ke handler berikutnya
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Token tidak valid",
            })
            c.Abort()
        }
    }
}
```

**Cara Kerja Middleware:**

```
Request → Middleware (cek token) → Handler (process)
                ↓ (jika token invalid)
              401 Unauthorized
```

##### **Langkah 1.5: Update main.go**

**File: `main.go`**

```go
package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "primmfy-backend/handlers"
    "primmfy-backend/middleware"
)

func main() {
    // 1. Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found")
    }

    // 2. Initialize database
    InitDB()
    defer func() {
        if DB != nil {
            DB.Close(context.Background())
        }
    }()

    // 3. Setup Gin router
    r := gin.Default()

    // 4. CORS middleware (untuk frontend bisa akses API)
    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })

    // 5. Setup routes
    api := r.Group("/api")
    {
        // Auth routes (tidak perlu token)
        auth := api.Group("/auth")
        {
            auth.POST("/register", handlers.Register)
            auth.POST("/login", handlers.Login)
        }

        // Protected routes (perlu token)
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            protected.GET("/profile", handlers.GetProfile)
            // Tambahkan endpoint lain yang perlu autentikasi di sini
        }
    }

    // 6. Health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "OK",
            "message": "Server is running",
        })
    })

    // 7. Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("🚀 Server starting on port %s\n", port)
    log.Printf("📚 API Documentation: http://localhost:%s/health\n", port)
    r.Run(":" + port)
}
```

#### Testing Fase 1:

##### **Menggunakan cURL (Command Line):**

```bash
# 1. Test health check
curl http://localhost:8080/health

# 2. Register user baru
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "password": "password123",
    "full_name": "John Doe"
  }'

# 3. Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'

# Response akan berisi token:
# {
#   "token": "eyJhbGc...",
#   "user": { ... }
# }

# 4. Akses protected route (ganti TOKEN dengan token dari response login)
curl http://localhost:8080/api/profile \
  -H "Authorization: Bearer TOKEN"
```

##### **Menggunakan Postman atau Thunder Client:**

1. **Install Thunder Client** extension di VS Code
2. **Buat Request Baru:**
   - Method: POST
   - URL: `http://localhost:8080/api/auth/register`
   - Body: JSON
   ```json
   {
     "username": "testuser",
     "email": "test@example.com",
     "password": "test123",
     "full_name": "Test User"
   }
   ```
3. **Klik Send** dan lihat response

---

### **FASE 2: Core Learning System**

#### Apa yang Akan Dipelajari?

- CRUD operations (Create, Read, Update, Delete)
- Relasi database (Foreign Keys)
- Query kompleks dengan JOIN
- Pagination
- File upload untuk gambar

#### Fitur yang Akan Dibuat:

- ✅ Management kursus (create, read, update, delete)
- ✅ Management lessons
- ✅ PRIMM stages implementation
- ✅ Progress tracking

---

### **FASE 3: Frontend Development**

#### Apa yang Akan Dipelajari?

- React components
- State management
- API integration dengan Axios
- Form handling
- Routing dengan Next.js App Router

#### Halaman yang Akan Dibuat:

- 🎨 Landing page
- 🔐 Login/Register page
- 📚 Course catalog
- 📖 Lesson viewer
- ✍️ Code editor untuk PRIMM stages

---

### **FASE 4: Gamification System**

#### Apa yang Akan Dipelajari?

- Transaction logic
- Business rules implementation
- Real-time updates

#### Fitur:

- 🪙 Coin system
- 📈 Level & XP system
- 🏆 Achievement system
- 🛒 Shop system

---

### **FASE 5: Avatar System**

#### Apa yang Akan Dipelajari?

- Image handling
- Canvas API (optional)
- Inventory management

---

### **FASE 6: Deployment**

#### Apa yang Akan Dipelajari?

- Docker
- CI/CD
- Cloud deployment (Railway, Vercel, dll)

---

## 📁 Struktur Folder Final

```
PRIMMFY/
├── Backend/
│   ├── main.go
│   ├── database.go
│   ├── .env
│   ├── go.mod
│   ├── go.sum
│   ├── config/
│   │   └── config.go
│   ├── models/
│   │   ├── user.go
│   │   ├── course.go
│   │   ├── lesson.go
│   │   └── avatar.go
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── course_service.go
│   │   └── gamification_service.go
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── course_handler.go
│   │   └── avatar_handler.go
│   ├── middleware/
│   │   ├── auth_middleware.go
│   │   └── cors_middleware.go
│   └── utils/
│       ├── response.go
│       └── validator.go
│
├── Frontend/
│   ├── app/
│   │   ├── layout.tsx
│   │   ├── page.tsx
│   │   ├── login/
│   │   ├── register/
│   │   ├── dashboard/
│   │   ├── courses/
│   │   └── profile/
│   ├── components/
│   │   ├── Header.tsx
│   │   ├── CourseCard.tsx
│   │   ├── CodeEditor.tsx
│   │   └── AvatarBuilder.tsx
│   ├── lib/
│   │   ├── api.ts
│   │   └── auth.ts
│   ├── public/
│   │   └── images/
│   └── package.json
│
├── database/
│   ├── schema.sql
│   └── seed.sql
│
└── README.md (file ini!)
```

---

## 🎓 Tutorial Pemula

### Konsep-Konsep Penting

#### 1. **API (Application Programming Interface)**

API adalah "jembatan" yang menghubungkan frontend dan backend.

**Analogi**: Seperti pelayan di restoran

- Kamu (Frontend) → Pesan makanan
- Pelayan (API) → Bawa pesanan ke dapur
- Dapur (Backend) → Masak makanan
- Pelayan → Bawa makanan kembali ke kamu

#### 2. **HTTP Methods**

- **GET**: Ambil data (seperti membaca buku)
- **POST**: Kirim data baru (seperti menulis buku baru)
- **PUT**: Update data (seperti edit buku)
- **DELETE**: Hapus data (seperti buang buku)

#### 3. **JSON (JavaScript Object Notation)**

Format data yang digunakan untuk komunikasi API.

```json
{
  "name": "John",
  "age": 25,
  "hobbies": ["coding", "gaming"]
}
```

#### 4. **Authentication vs Authorization**

- **Authentication**: "Siapa kamu?" (Login)
- **Authorization**: "Apa yang boleh kamu lakukan?" (Permission)

#### 5. **Environment Variables**

Data sensitif yang tidak boleh di-commit ke Git (password, secret keys).

**File `.env`:**

```env
DATABASE_URL=postgres://...
JWT_SECRET=rahasia123
```

**Di Git, tambahkan `.env` ke `.gitignore`!**

---

## 🔧 Tips untuk Pemula

### 1. **Jangan Takut Error!**

Error adalah teman belajar. Setiap error mengajarkan sesuatu.

### 2. **Baca Error Message**

Go memberikan error message yang jelas. Baca pelan-pelan.

### 3. **Print/Log untuk Debug**

```go
fmt.Println("Debug: user ID =", userID)
log.Printf("Error: %v", err)
```

### 4. **Test Satu Fitur Sekaligus**

Jangan langsung bikin semua. Test tiap endpoint sebelum lanjut.

### 5. **Gunakan Git**

Commit setelah setiap fitur selesai:

```bash
git add .
git commit -m "Fase 1: Implement authentication"
```

---

## 📚 Resource Belajar Tambahan

### Go (Golang):

- [Tour of Go](https://tour.golang.org/) - Tutorial interaktif
- [Go by Example](https://gobyexample.com/) - Contoh kode praktis
- [Effective Go](https://golang.org/doc/effective_go) - Best practices

### Next.js & React:

- [Next.js Docs](https://nextjs.org/docs)
- [React Docs](https://react.dev/)
- [Tailwind CSS Docs](https://tailwindcss.com/docs)

### PostgreSQL:

- [PostgreSQL Tutorial](https://www.postgresqltutorial.com/)
- [SQL Zoo](https://sqlzoo.net/) - Interactive SQL learning

### API Testing:

- [Thunder Client](https://www.thunderclient.com/) - VS Code extension
- [Postman](https://www.postman.com/)

---

## 🆘 Troubleshooting

### Error: "Cannot connect to database"

**Solusi:**

1. Pastikan PostgreSQL sudah running
2. Cek username, password, dan nama database di connection string
3. Test koneksi dengan pgAdmin atau psql

### Error: "Port already in use"

**Solusi:**

```bash
# Windows (PowerShell)
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# Atau ganti port di .env
PORT=8081
```

### Error: "Module not found"

**Solusi:**

```bash
go mod tidy
go mod download
```

---

## 🎯 Next Steps

Setelah menyelesaikan Fase 1, kamu akan:

- ✅ Memahami dasar REST API
- ✅ Bisa membuat sistem autentikasi
- ✅ Mengerti cara kerja JWT
- ✅ Siap untuk Fase 2!

**Selamat belajar! 🚀**

Jika ada pertanyaan atau stuck di tahap tertentu, jangan ragu untuk bertanya!

---

## 📝 Changelog

- **v0.1.0** (Oct 2025): Initial project setup & Fase 1
- **v0.2.0** (Coming soon): Fase 2 - Core Learning System

---

## 👥 Contributors

- **Dhika** - Lead Developer & Creator

---

## 📄 License

This project is for educational purposes.

---

**Happy Coding! 💻✨**
