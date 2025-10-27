package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "primmfy_db/services"
)

// AuthMiddleware memvalidasi JWT token dari request header
// Header format: Authorization: Bearer <token>
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. Ambil token dari Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authorization header tidak ditemukan",
            })
            c.Abort()
            return
        }

        // 2. Extract token dari "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Format Authorization header salah. Gunakan: Bearer <token>",
            })
            c.Abort()
            return
        }

        token := parts[1]

        // 3. Validasi token
        claims, err := services.ValidateJWT(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Token tidak valid atau sudah expired: " + err.Error(),
            })
            c.Abort()
            return
        }

        // 4. Simpan user info ke context (bisa diakses di handler)
        // MapClaims diakses seperti map, bukan struct
        userID := int(claims["user_id"].(float64)) // JWT number selalu float64
        userRole := claims["role"].(string)
        
        c.Set("user_id", userID)
        c.Set("user_role", userRole)

        // 5. Lanjutkan ke handler berikutnya
        c.Next()
    }
}

// RequireRole middleware untuk membatasi akses berdasarkan role
// Contoh: RequireRole("admin", "teacher")
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Ambil role dari context (di-set oleh AuthMiddleware)
        userRole, exists := c.Get("user_role")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "User role tidak ditemukan",
            })
            c.Abort()
            return
        }

        // Cek apakah role user diizinkan
        roleStr := userRole.(string)
        allowed := false
        for _, role := range allowedRoles {
            if roleStr == role {
                allowed = true
                break
            }
        }

        if !allowed {
            c.JSON(http.StatusForbidden, gin.H{
                "error": "Anda tidak memiliki akses ke resource ini",
            })
            c.Abort()
            return
        }

        c.Next()
    }
}