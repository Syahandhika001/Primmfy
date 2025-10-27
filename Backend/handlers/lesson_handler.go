package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
    "primmfy_db/models"
    "primmfy_db/services"
)

// LessonHandler mengelola endpoint lesson (Big Topic)
type LessonHandler struct {
    DB *pgx.Conn
}

// NewLessonHandler membuat instance LessonHandler baru
func NewLessonHandler(db *pgx.Conn) *LessonHandler {
    return &LessonHandler{DB: db}
}

// ═══════════════════════════════════════════════════════════
// LESSON CRUD ENDPOINTS (Teacher Only)
// ═══════════════════════════════════════════════════════════

// CreateLesson handler untuk POST /api/lessons (teacher only)
// Purpose: Teacher membuat lesson baru (big topic seperti "Python Basics")
func (h *LessonHandler) CreateLesson(c *gin.Context) {
    // 1. Get teacher ID dari JWT middleware
    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    // 2. Bind & validate request body
    var req models.CreateLessonRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    // 3. Create lesson via service
    lesson, err := services.CreateLesson(h.DB, teacherID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Lesson berhasil dibuat!",
        "lesson":  lesson,
    })
}

// GetAllLessons handler untuk GET /api/lessons (public)
// Purpose: Menampilkan semua lesson yang aktif
func (h *LessonHandler) GetAllLessons(c *gin.Context) {
    lessons, err := services.GetAllLessons(h.DB)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "lessons": lessons,
        "count":   len(lessons),
    })
}

// GetLessonByID handler untuk GET /api/lessons/:id
// Purpose: Menampilkan detail lesson beserta courses-nya
func (h *LessonHandler) GetLessonByID(c *gin.Context) {
    lessonID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Lesson ID tidak valid"})
        return
    }

    // Get lesson with all courses
    lesson, courses, err := services.GetLessonWithCourses(h.DB, lessonID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "lesson":  lesson,
        "courses": courses,
        "total_courses": len(courses),
    })
}

// GetMyLessons handler untuk GET /api/lessons/my (teacher only)
// Purpose: Teacher melihat semua lesson yang dia buat
func (h *LessonHandler) GetMyLessons(c *gin.Context) {
    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    lessons, err := services.GetLessonsByTeacher(h.DB, teacherID.(int))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "lessons": lessons,
        "count":   len(lessons),
    })
}

// UpdateLesson handler untuk PUT /api/lessons/:id (teacher only)
func (h *LessonHandler) UpdateLesson(c *gin.Context) {
    lessonID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Lesson ID tidak valid"})
        return
    }

    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    var req models.UpdateLessonRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    lesson, err := services.UpdateLesson(h.DB, lessonID, teacherID.(int), req)
    if err != nil {
        if err.Error() == "lesson tidak ditemukan" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else if err.Error() == "anda tidak memiliki akses untuk mengupdate lesson ini" {
            c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Lesson berhasil diupdate!",
        "lesson":  lesson,
    })
}

// DeleteLesson handler untuk DELETE /api/lessons/:id (teacher only)
func (h *LessonHandler) DeleteLesson(c *gin.Context) {
    lessonID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Lesson ID tidak valid"})
        return
    }

    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    err = services.DeleteLesson(h.DB, lessonID, teacherID.(int))
    if err != nil {
        if err.Error() == "lesson tidak ditemukan" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else if err.Error() == "anda tidak memiliki akses untuk menghapus lesson ini" {
            c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Lesson berhasil dihapus!"})
}

// ═══════════════════════════════════════════════════════════
// LESSON ENROLLMENT ENDPOINTS (Student)
// ═══════════════════════════════════════════════════════════

// EnrollLesson handler untuk POST /api/lessons/:id/enroll (student only)
// Purpose: Siswa mendaftar ke lesson untuk akses courses
func (h *LessonHandler) EnrollLesson(c *gin.Context) {
    lessonID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Lesson ID tidak valid"})
        return
    }

    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    enrollment, err := services.EnrollLesson(h.DB, userID.(int), lessonID)
    if err != nil {
        if err.Error() == "lesson tidak ditemukan" || err.Error() == "lesson tidak aktif" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else if err.Error() == "anda sudah terdaftar di lesson ini" {
            c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message":    "Berhasil enroll ke lesson!",
        "enrollment": enrollment,
    })
}

// GetMyEnrolledLessons handler untuk GET /api/my-lessons (student only)
// Purpose: Siswa melihat semua lesson yang sudah di-enroll dengan progress
func (h *LessonHandler) GetMyEnrolledLessons(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    lessons, err := services.GetMyEnrolledLessons(h.DB, userID.(int))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "lessons": lessons,
        "count":   len(lessons),
    })
}

// GetCoursesInLesson handler untuk GET /api/lessons/:id/courses
func (h *LessonHandler) GetCoursesInLesson(c *gin.Context) {
    // 1. Parse lesson ID dari URL
    lessonID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Lesson ID tidak valid"})
        return
    }

    // 2. Check jika user authenticated (optional: untuk progress)
    userID, hasAuth := c.Get("user_id")

    var result interface{}

    if hasAuth {
        // User authenticated - return courses dengan progress
        courses, err := services.GetCoursesByLessonWithProgress(h.DB, lessonID, userID.(int))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        result = gin.H{"courses": courses}
    } else {
        // User not authenticated - return courses tanpa progress
        courses, err := services.GetCoursesByLesson(h.DB, lessonID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        result = gin.H{"courses": courses}
    }

    c.JSON(http.StatusOK, result)
}