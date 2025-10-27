package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
    "primmfy_db/models"
    "primmfy_db/services"
)

// AuthHandler mengelola endpoint authentication
type AuthHandler struct {
    DB *pgx.Conn
}

// NewAuthHandler membuat instance AuthHandler baru
func NewAuthHandler(db *pgx.Conn) *AuthHandler {
    return &AuthHandler{DB: db}
}

// Register handler untuk POST /api/register
// Purpose: Mendaftarkan user baru (student/teacher/admin)
func (h *AuthHandler) Register(c *gin.Context) {
    var req models.RegisterRequest

    // 1. Bind & validate JSON request
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data tidak valid: " + err.Error(),
        })
        return
    }

    // 2. Register user via service
    user, err := services.Register(h.DB, req)
    if err != nil {
        // Cek error type untuk response code yang sesuai
        if err.Error() == "email sudah terdaftar" {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    // 3. Return success response
    c.JSON(http.StatusCreated, gin.H{
        "message": "User berhasil didaftarkan!",
        "user":    user,
    })
}

// Login handler untuk POST /api/login
// Purpose: Login user dan return JWT token
func (h *AuthHandler) Login(c *gin.Context) {
    var req models.LoginRequest

    // 1. Bind & validate JSON request
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Data tidak valid: " + err.Error(),
        })
        return
    }

    // 2. Login via service (validate credentials & generate JWT)
    response, err := services.Login(h.DB, req)
    if err != nil {
        // Semua login error return 401 Unauthorized
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // 3. Return success response dengan token
    c.JSON(http.StatusOK, response)
}

// GetProfile handler untuk GET /api/profile (protected route)
// Purpose: Mendapatkan detail profile user yang login
func (h *AuthHandler) GetProfile(c *gin.Context) {
    // User ID sudah di-set oleh AuthMiddleware
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    // Get user details from database
    user, err := services.GetUserByID(h.DB, userID.(int))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}