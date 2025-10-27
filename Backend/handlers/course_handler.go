package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
    "primmfy_db/models"
    "primmfy_db/services"
)

// CourseHandler mengelola endpoint course (Sub-topic dalam Lesson)
type CourseHandler struct {
    DB *pgx.Conn
}

// NewCourseHandler membuat instance CourseHandler baru
func NewCourseHandler(db *pgx.Conn) *CourseHandler {
    return &CourseHandler{DB: db}
}

// ═══════════════════════════════════════════════════════════
// COURSE CRUD ENDPOINTS (Teacher Only)
// ═══════════════════════════════════════════════════════════

// CreateCourse handler untuk POST /api/courses (teacher only)
// Purpose: Teacher membuat course baru dalam lesson
func (h *CourseHandler) CreateCourse(c *gin.Context) {
    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    var req models.CreateCourseRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    course, err := services.CreateCourse(h.DB, teacherID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Course berhasil dibuat!",
        "course":  course,
    })
}

// GetCourseByID handler untuk GET /api/courses/:id
// Purpose: Menampilkan detail course beserta semua PRIMM stages
func (h *CourseHandler) GetCourseByID(c *gin.Context) {
    courseID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Course ID tidak valid"})
        return
    }

    // Get course with all PRIMM stages
    course, stages, err := services.GetCourseWithStages(h.DB, courseID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "course":       course,
        "stages":       stages,
        "total_stages": len(stages),
    })
}

// UpdateCourse handler untuk PUT /api/courses/:id (teacher only)
func (h *CourseHandler) UpdateCourse(c *gin.Context) {
    courseID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Course ID tidak valid"})
        return
    }

    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    var req models.UpdateCourseRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    course, err := services.UpdateCourse(h.DB, courseID, teacherID.(int), req)
    if err != nil {
        if err.Error() == "course tidak ditemukan" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else if err.Error() == "anda tidak memiliki akses untuk mengupdate course ini" {
            c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Course berhasil diupdate!",
        "course":  course,
    })
}

// DeleteCourse handler untuk DELETE /api/courses/:id (teacher only)
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
    courseID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Course ID tidak valid"})
        return
    }

    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    err = services.DeleteCourse(h.DB, courseID, teacherID.(int))
    if err != nil {
        if err.Error() == "course tidak ditemukan" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else if err.Error() == "anda tidak memiliki akses untuk menghapus course ini" {
            c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Course berhasil dihapus!"})
}

// GetStagesInCourse handler untuk GET /api/courses/:id/stages
func (h *CourseHandler) GetStagesInCourse(c *gin.Context) {
    // 1. Parse course ID
    courseID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Course ID tidak valid"})
        return
    }

    // 2. Get stages via service
    stages, err := services.GetStagesByCourse(h.DB, courseID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 3. Return response
    c.JSON(http.StatusOK, gin.H{
        "stages": stages,
    })
}