package models

import "time"

// User adalah model untuk tabel users
type User struct {
    ID               int        `json:"id"`
    Email            string     `json:"email"`
    PasswordHash     string     `json:"-"` // Tidak di-return ke client
    FullName         string     `json:"full_name"`
    Role             string     `json:"role"`
    Level            int        `json:"level"`
    TotalCoins       int        `json:"total_coins"`
    ExperiencePoints int        `json:"experience_points"`
    ProfilePicture   *string    `json:"profile_picture,omitempty"`
    Bio              *string    `json:"bio,omitempty"`
    CreatedAt        time.Time  `json:"created_at"`
    UpdatedAt        time.Time  `json:"updated_at"`
}

// RegisterRequest adalah struktur untuk request register
type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
    FullName string `json:"full_name" binding:"required"`
    Role     string `json:"role" binding:"required,oneof=student teacher admin"`
}

// LoginRequest adalah struktur untuk request login
type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// LoginResponse adalah struktur untuk response login
type LoginResponse struct {
    Message string `json:"message"`
    Token   string `json:"token"`
    User    *User  `json:"user"`
}