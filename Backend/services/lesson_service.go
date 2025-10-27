package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"primmfy_db/models"

	"github.com/jackc/pgx/v5"
)

// ═══════════════════════════════════════════════════════════
// TYPE DEFINITIONS
// ═══════════════════════════════════════════════════════════

// EnrolledLessonResponse untuk response enrolled lesson dengan progress
type EnrolledLessonResponse struct {
    ID             int            `json:"id"`
    TeacherID      int            `json:"teacher_id"`
    Title          string         `json:"title"`
    Description    string         `json:"description"`
    Category       string         `json:"category"`
    Difficulty     string         `json:"difficulty"`
    ThumbnailURL   string         `json:"thumbnail_url"`
    IsActive       bool           `json:"is_active"`
    CreatedAt      time.Time      `json:"created_at"`
    UpdatedAt      time.Time      `json:"updated_at"`
    EnrolledAt     time.Time      `json:"enrolled_at"`
    LastAccessedAt time.Time      `json:"last_accessed_at"`
    TeacherName    string         `json:"teacher_name"`
    Progress       LessonProgress `json:"progress"`
}

// LessonProgress untuk tracking progress lesson
type LessonProgress struct {
    TotalCourses      int     `json:"total_courses"`
    CompletedCourses  int     `json:"completed_courses"`
    CompletionPercent float64 `json:"completion_percentage"`
}

// UserLessonEnrollment untuk response enrollment
type UserLessonEnrollment struct {
    ID             int       `json:"id"`
    UserID         int       `json:"user_id"`
    LessonID       int       `json:"lesson_id"`
    EnrolledAt     time.Time `json:"enrolled_at"`
    LastAccessedAt time.Time `json:"last_accessed_at"`
}

// ═══════════════════════════════════════════════════════════
// LESSON CRUD OPERATIONS (Teacher Only)
// ═══════════════════════════════════════════════════════════

// CreateLesson membuat lesson baru (teacher only)
func CreateLesson(db *pgx.Conn, teacherID int, req models.CreateLessonRequest) (*models.Lesson, error) {
    var lesson models.Lesson

    err := db.QueryRow(context.Background(), `
        INSERT INTO lessons (teacher_id, title, description, category, difficulty, thumbnail_url)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, teacher_id, title, description, category, difficulty, 
                  thumbnail_url, is_active, created_at, updated_at`,
        teacherID, req.Title, req.Description, req.Category, req.Difficulty, req.ThumbnailURL).Scan(
        &lesson.ID, &lesson.TeacherID, &lesson.Title, &lesson.Description,
        &lesson.Category, &lesson.Difficulty, &lesson.ThumbnailURL,
        &lesson.IsActive, &lesson.CreatedAt, &lesson.UpdatedAt)

    if err != nil {
        return nil, errors.New("gagal membuat lesson: " + err.Error())
    }

    return &lesson, nil
}

// GetAllLessons mengambil semua lesson yang aktif
func GetAllLessons(db *pgx.Conn) ([]models.LessonWithTeacher, error) {
    rows, err := db.Query(context.Background(), `
        SELECT l.id, l.teacher_id, l.title, l.description, l.category, 
               l.difficulty, l.thumbnail_url, l.is_active, l.created_at, l.updated_at,
               u.full_name as teacher_name
        FROM lessons l
        JOIN users u ON l.teacher_id = u.id
        WHERE l.is_active = true
        ORDER BY l.created_at DESC`)

    if err != nil {
        return nil, errors.New("gagal mengambil daftar lesson: " + err.Error())
    }
    defer rows.Close()

    var lessons []models.LessonWithTeacher
    for rows.Next() {
        var lesson models.LessonWithTeacher
        err := rows.Scan(
            &lesson.ID, &lesson.TeacherID, &lesson.Title, &lesson.Description,
            &lesson.Category, &lesson.Difficulty, &lesson.ThumbnailURL,
            &lesson.IsActive, &lesson.CreatedAt, &lesson.UpdatedAt,
            &lesson.TeacherName)

        if err != nil {
            return nil, errors.New("gagal scan lesson: " + err.Error())
        }
        lessons = append(lessons, lesson)
    }

    return lessons, nil
}

// GetLessonByID mengambil detail lesson berdasarkan ID
func GetLessonByID(db *pgx.Conn, lessonID int) (*models.LessonWithTeacher, error) {
    var lesson models.LessonWithTeacher

    err := db.QueryRow(context.Background(), `
        SELECT l.id, l.teacher_id, l.title, l.description, l.category, 
               l.difficulty, l.thumbnail_url, l.is_active, l.created_at, l.updated_at,
               u.full_name as teacher_name
        FROM lessons l
        JOIN users u ON l.teacher_id = u.id
        WHERE l.id = $1`, lessonID).Scan(
        &lesson.ID, &lesson.TeacherID, &lesson.Title, &lesson.Description,
        &lesson.Category, &lesson.Difficulty, &lesson.ThumbnailURL,
        &lesson.IsActive, &lesson.CreatedAt, &lesson.UpdatedAt,
        &lesson.TeacherName)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("lesson tidak ditemukan")
        }
        return nil, errors.New("gagal mengambil lesson: " + err.Error())
    }

    return &lesson, nil
}

// GetLessonsByTeacher mengambil semua lesson milik teacher tertentu
func GetLessonsByTeacher(db *pgx.Conn, teacherID int) ([]models.Lesson, error) {
    rows, err := db.Query(context.Background(), `
        SELECT id, teacher_id, title, description, category, difficulty, 
               thumbnail_url, is_active, created_at, updated_at
        FROM lessons
        WHERE teacher_id = $1
        ORDER BY created_at DESC`, teacherID)

    if err != nil {
        return nil, errors.New("gagal mengambil lesson: " + err.Error())
    }
    defer rows.Close()

    var lessons []models.Lesson
    for rows.Next() {
        var lesson models.Lesson
        err := rows.Scan(
            &lesson.ID, &lesson.TeacherID, &lesson.Title, &lesson.Description,
            &lesson.Category, &lesson.Difficulty, &lesson.ThumbnailURL,
            &lesson.IsActive, &lesson.CreatedAt, &lesson.UpdatedAt)

        if err != nil {
            return nil, errors.New("gagal scan lesson: " + err.Error())
        }
        lessons = append(lessons, lesson)
    }

    return lessons, nil
}

// UpdateLesson mengupdate lesson (teacher only, hanya lesson miliknya)
func UpdateLesson(db *pgx.Conn, lessonID int, teacherID int, req models.UpdateLessonRequest) (*models.Lesson, error) {
    // Cek apakah lesson milik teacher ini
    var ownerID int
    err := db.QueryRow(context.Background(),
        "SELECT teacher_id FROM lessons WHERE id = $1", lessonID).Scan(&ownerID)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("lesson tidak ditemukan")
        }
        return nil, errors.New("gagal cek ownership: " + err.Error())
    }

    if ownerID != teacherID {
        return nil, errors.New("anda tidak memiliki akses untuk mengupdate lesson ini")
    }

    // Build dynamic update query
    setClauses := []string{"updated_at = NOW()"}
    args := []interface{}{}
    argPos := 1

    if req.Title != "" {
        setClauses = append(setClauses, fmt.Sprintf("title = $%d", argPos))
        args = append(args, req.Title)
        argPos++
    }
    if req.Description != "" {
        setClauses = append(setClauses, fmt.Sprintf("description = $%d", argPos))
        args = append(args, req.Description)
        argPos++
    }
    if req.Category != "" {
        setClauses = append(setClauses, fmt.Sprintf("category = $%d", argPos))
        args = append(args, req.Category)
        argPos++
    }
    if req.Difficulty != "" {
        setClauses = append(setClauses, fmt.Sprintf("difficulty = $%d", argPos))
        args = append(args, req.Difficulty)
        argPos++
    }
    if req.ThumbnailURL != "" {
        setClauses = append(setClauses, fmt.Sprintf("thumbnail_url = $%d", argPos))
        args = append(args, req.ThumbnailURL)
        argPos++
    }
    if req.IsActive != nil {
        setClauses = append(setClauses, fmt.Sprintf("is_active = $%d", argPos))
        args = append(args, *req.IsActive)
        argPos++
    }

    // Add WHERE clause
    args = append(args, lessonID)

    // Build final query
    query := fmt.Sprintf(`
        UPDATE lessons 
        SET %s 
        WHERE id = $%d
        RETURNING id, teacher_id, title, description, category, difficulty, 
                  thumbnail_url, is_active, created_at, updated_at`,
        strings.Join(setClauses, ", "),
        argPos)

    var lesson models.Lesson
    err = db.QueryRow(context.Background(), query, args...).Scan(
        &lesson.ID, &lesson.TeacherID, &lesson.Title, &lesson.Description,
        &lesson.Category, &lesson.Difficulty, &lesson.ThumbnailURL,
        &lesson.IsActive, &lesson.CreatedAt, &lesson.UpdatedAt)

    if err != nil {
        return nil, errors.New("gagal update lesson: " + err.Error())
    }

    return &lesson, nil
}

// DeleteLesson menghapus lesson (soft delete: set is_active = false)
func DeleteLesson(db *pgx.Conn, lessonID int, teacherID int) error {
    // Cek ownership
    var ownerID int
    err := db.QueryRow(context.Background(),
        "SELECT teacher_id FROM lessons WHERE id = $1", lessonID).Scan(&ownerID)

    if err != nil {
        if err == pgx.ErrNoRows {
            return errors.New("lesson tidak ditemukan")
        }
        return errors.New("gagal cek ownership: " + err.Error())
    }

    if ownerID != teacherID {
        return errors.New("anda tidak memiliki akses untuk menghapus lesson ini")
    }

    // Soft delete
    _, err = db.Exec(context.Background(),
        "UPDATE lessons SET is_active = false, updated_at = NOW() WHERE id = $1", lessonID)

    if err != nil {
        return errors.New("gagal menghapus lesson: " + err.Error())
    }

    return nil
}

// ═══════════════════════════════════════════════════════════
// LESSON ENROLLMENT & PROGRESS (Student)
// ═══════════════════════════════════════════════════════════

// GetUserEnrolledLessons mendapatkan semua lessons yang dienroll user beserta progress
func GetUserEnrolledLessons(db *pgx.Conn, userID int) ([]EnrolledLessonResponse, error) {
    rows, err := db.Query(context.Background(), `
        SELECT 
            l.id, 
            l.teacher_id, 
            l.title, 
            l.description, 
            l.category, 
            l.difficulty, 
            l.thumbnail_url, 
            l.is_active, 
            l.created_at, 
            l.updated_at,
            ul.enrolled_at, 
            ul.last_accessed_at,
            u.full_name AS teacher_name,
            COUNT(DISTINCT c.id) AS total_courses,
            COUNT(DISTINCT ucc.course_id) AS completed_courses
        FROM user_lessons ul
        JOIN lessons l ON ul.lesson_id = l.id
        JOIN users u ON l.teacher_id = u.id
        LEFT JOIN courses c ON l.id = c.lesson_id AND c.is_active = true
        LEFT JOIN user_course_completion ucc ON c.id = ucc.course_id AND ucc.user_id = ul.user_id
        WHERE ul.user_id = $1
        GROUP BY 
            l.id, 
            l.teacher_id, 
            l.title, 
            l.description, 
            l.category, 
            l.difficulty, 
            l.thumbnail_url, 
            l.is_active, 
            l.created_at, 
            l.updated_at,
            ul.enrolled_at, 
            ul.last_accessed_at, 
            u.full_name
        ORDER BY ul.enrolled_at DESC
    `, userID)

    if err != nil {
        return nil, errors.New("gagal mengambil enrolled lessons: " + err.Error())
    }
    defer rows.Close()

    var lessons []EnrolledLessonResponse
    for rows.Next() {
        var lesson EnrolledLessonResponse
        var totalCourses, completedCourses int

        err := rows.Scan(
            &lesson.ID, 
            &lesson.TeacherID, 
            &lesson.Title, 
            &lesson.Description,
            &lesson.Category, 
            &lesson.Difficulty, 
            &lesson.ThumbnailURL,
            &lesson.IsActive, 
            &lesson.CreatedAt, 
            &lesson.UpdatedAt,
            &lesson.EnrolledAt, 
            &lesson.LastAccessedAt, 
            &lesson.TeacherName,
            &totalCourses, 
            &completedCourses,
        )

        if err != nil {
            return nil, errors.New("gagal scan lesson: " + err.Error())
        }

        // Calculate completion percentage
        var completionPercent float64
        if totalCourses > 0 {
            completionPercent = (float64(completedCourses) / float64(totalCourses)) * 100
        }

        lesson.Progress = LessonProgress{
            TotalCourses:      totalCourses,
            CompletedCourses:  completedCourses,
            CompletionPercent: completionPercent,
        }

        lessons = append(lessons, lesson)
    }

    return lessons, nil
}

// GetLessonWithCourses mengambil lesson beserta semua courses-nya
func GetLessonWithCourses(db *pgx.Conn, lessonID int) (*models.LessonWithTeacher, []models.Course, error) {
    // 1. Get lesson details
    lesson, err := GetLessonByID(db, lessonID)
    if err != nil {
        return nil, nil, err
    }

    // 2. Get all courses in this lesson
    courses, err := GetCoursesByLesson(db, lessonID)
    if err != nil {
        return lesson, []models.Course{}, nil // Return empty courses if error
    }

    return lesson, courses, nil
}

// GetMyEnrolledLessons mengambil semua lesson yang di-enroll oleh user dengan progress
func GetMyEnrolledLessons(db *pgx.Conn, userID int) ([]models.LessonWithProgress, error) {
    rows, err := db.Query(context.Background(), `
        SELECT 
            l.id, 
            l.teacher_id, 
            l.title, 
            l.description, 
            l.category, 
            l.difficulty, 
            COALESCE(l.thumbnail_url, '') as thumbnail_url,
            l.is_active, 
            l.created_at, 
            l.updated_at,
            u.full_name as teacher_name,
            ul.enrolled_at,
            COALESCE(COUNT(DISTINCT c.id), 0) as total_courses
        FROM user_lessons ul
        JOIN lessons l ON ul.lesson_id = l.id
        JOIN users u ON l.teacher_id = u.id
        LEFT JOIN courses c ON l.id = c.lesson_id AND c.is_active = true
        WHERE ul.user_id = $1
        GROUP BY 
            l.id, 
            l.teacher_id, 
            l.title, 
            l.description, 
            l.category, 
            l.difficulty, 
            l.thumbnail_url, 
            l.is_active, 
            l.created_at, 
            l.updated_at, 
            u.full_name,
            ul.enrolled_at
        ORDER BY ul.enrolled_at DESC
    `, userID)

    if err != nil {
        return nil, errors.New("gagal mengambil enrolled lessons: " + err.Error())
    }
    defer rows.Close()

    var lessons []models.LessonWithProgress
    for rows.Next() {
        var lesson models.LessonWithProgress
        var enrolledAt time.Time
        
        err := rows.Scan(
            &lesson.ID, 
            &lesson.TeacherID, 
            &lesson.Title, 
            &lesson.Description,
            &lesson.Category, 
            &lesson.Difficulty, 
            &lesson.ThumbnailURL,
            &lesson.IsActive, 
            &lesson.CreatedAt, 
            &lesson.UpdatedAt,
            &lesson.TeacherName, 
            &enrolledAt,
            &lesson.TotalCourses,
        )

        if err != nil {
            return nil, errors.New("gagal scan lesson: " + err.Error())
        }

        // Set defaults
        lesson.CompletedCourses = 0
        lesson.ProgressPercent = 0
        lesson.IsEnrolled = true

        lessons = append(lessons, lesson)
    }

    return lessons, nil
}
// CheckEnrollment mengecek apakah user sudah enroll lesson
func CheckEnrollment(db *pgx.Conn, userID int, lessonID int) (bool, error) {
    var exists bool
    err := db.QueryRow(context.Background(),
        "SELECT EXISTS(SELECT 1 FROM user_lessons WHERE user_id = $1 AND lesson_id = $2)",
        userID, lessonID).Scan(&exists)

    if err != nil {
        return false, errors.New("gagal cek enrollment: " + err.Error())
    }

    return exists, nil
}

// EnrollLesson mendaftarkan user ke lesson
func EnrollLesson(db *pgx.Conn, userID int, lessonID int) (*UserLessonEnrollment, error) {
    // 1. Check apakah lesson exists dan active
    var isActive bool
    err := db.QueryRow(context.Background(), `
        SELECT is_active FROM lessons WHERE id = $1
    `, lessonID).Scan(&isActive)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("lesson tidak ditemukan")
        }
        return nil, errors.New("gagal mengecek lesson: " + err.Error())
    }

    if !isActive {
        return nil, errors.New("lesson tidak aktif")
    }

    // 2. Check apakah user sudah enroll
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
        RETURNING id, user_id, lesson_id, enrolled_at, last_accessed_at
    `, userID, lessonID).Scan(
        &enrollment.ID,
        &enrollment.UserID,
        &enrollment.LessonID,
        &enrollment.EnrolledAt,
        &enrollment.LastAccessedAt,
    )

    if err != nil {
        return nil, errors.New("gagal enroll ke lesson: " + err.Error())
    }

    return &enrollment, nil
}

