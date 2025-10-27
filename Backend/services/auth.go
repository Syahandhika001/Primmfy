package services

import (
    "context"
    "errors"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/jackc/pgx/v5"
    "golang.org/x/crypto/bcrypt"
    "primmfy_db/models"
)

// ═══════════════════════════════════════════════════════════
// PASSWORD HASHING FUNCTIONS
// ═══════════════════════════════════════════════════════════

// HashPassword meng-hash password menggunakan bcrypt
func HashPassword(password string) (string, error) {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedBytes), nil
}

// CheckPasswordHash membandingkan password dengan hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// ═══════════════════════════════════════════════════════════
// JWT TOKEN FUNCTIONS
// ═══════════════════════════════════════════════════════════

// GenerateJWT membuat JWT token untuk user
func GenerateJWT(userID int, email string, role string) (string, error) {
    secretKey := os.Getenv("JWT_SECRET")
    if secretKey == "" {
        return "", errors.New("JWT_SECRET tidak ditemukan di environment variables")
    }

    // Claims berisi data yang akan disimpan di token
    claims := jwt.MapClaims{
        "user_id": userID,
        "email":   email,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Token valid 7 hari
    }

    // Buat token dengan claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign token dengan secret key
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateJWT memvalidasi JWT token dan return claims
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
    secretKey := os.Getenv("JWT_SECRET")
    if secretKey == "" {
        return nil, errors.New("JWT_SECRET tidak ditemukan")
    }

    // Parse token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validasi signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return []byte(secretKey), nil
    })

    if err != nil {
        return nil, err
    }

    // Extract claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}

// ═══════════════════════════════════════════════════════════
// AUTH SERVICE FUNCTIONS
// ═══════════════════════════════════════════════════════════

// Register membuat user baru
func Register(db *pgx.Conn, req models.RegisterRequest) (*models.User, error) {
    // 1. Hash password
    hashedPassword, err := HashPassword(req.Password)
    if err != nil {
        return nil, errors.New("gagal hash password: " + err.Error())
    }

    // 2. Cek apakah email sudah terdaftar
    var existingID int
    err = db.QueryRow(context.Background(),
        "SELECT id FROM users WHERE email = $1", req.Email).Scan(&existingID)

    if err == nil {
        // Email sudah ada
        return nil, errors.New("email sudah terdaftar")
    }

    if err != pgx.ErrNoRows {
        // Error lain selain "no rows"
        return nil, errors.New("gagal cek email: " + err.Error())
    }

    // 3. Insert user baru DENGAN profile_picture dan bio
    var user models.User
    err = db.QueryRow(context.Background(), `
        INSERT INTO users (email, password_hash, full_name, role, level, total_coins, experience_points)
        VALUES ($1, $2, $3, $4, 1, 0, 0)
        RETURNING id, email, full_name, role, level, total_coins, experience_points, 
                  profile_picture, bio, created_at, updated_at`,
        req.Email, hashedPassword, req.FullName, req.Role).Scan(
        &user.ID, &user.Email, &user.FullName, &user.Role,
        &user.Level, &user.TotalCoins, &user.ExperiencePoints,
        &user.ProfilePicture, &user.Bio, &user.CreatedAt, &user.UpdatedAt)

    if err != nil {
        return nil, errors.New("gagal membuat user: " + err.Error())
    }

    return &user, nil
}

// Login memvalidasi kredensial dan return JWT token
func Login(db *pgx.Conn, req models.LoginRequest) (*models.LoginResponse, error) {
    // 1. Ambil user berdasarkan email DENGAN profile_picture dan bio
    var user models.User
    var passwordHash string

    err := db.QueryRow(context.Background(),
        `SELECT id, email, password_hash, full_name, role, level, total_coins, experience_points, 
                profile_picture, bio, created_at, updated_at 
         FROM users WHERE email = $1`,
        req.Email).Scan(
        &user.ID, &user.Email, &passwordHash, &user.FullName, &user.Role,
        &user.Level, &user.TotalCoins, &user.ExperiencePoints,
        &user.ProfilePicture, &user.Bio, &user.CreatedAt, &user.UpdatedAt)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("email atau password salah")
        }
        return nil, errors.New("gagal login: " + err.Error())
    }

    // 2. Validasi password
    if !CheckPasswordHash(req.Password, passwordHash) {
        return nil, errors.New("email atau password salah")
    }

    // 3. Generate JWT token
    token, err := GenerateJWT(user.ID, user.Email, user.Role)
    if err != nil {
        return nil, errors.New("gagal generate token: " + err.Error())
    }

    // 4. Return response dengan token
    return &models.LoginResponse{
        Message: "Login berhasil!",
        Token:   token,
        User:    &user,
    }, nil
}

// GetUserByID mengambil user berdasarkan ID DENGAN profile_picture dan bio
func GetUserByID(db *pgx.Conn, userID int) (*models.User, error) {
    var user models.User
    err := db.QueryRow(context.Background(),
        `SELECT id, email, full_name, role, level, total_coins, experience_points, 
                profile_picture, bio, created_at, updated_at 
         FROM users WHERE id = $1`,
        userID).Scan(
        &user.ID, &user.Email, &user.FullName, &user.Role,
        &user.Level, &user.TotalCoins, &user.ExperiencePoints,
        &user.ProfilePicture, &user.Bio, &user.CreatedAt, &user.UpdatedAt)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("user tidak ditemukan")
        }
        return nil, errors.New("gagal mengambil user: " + err.Error())
    }

    return &user, nil
}