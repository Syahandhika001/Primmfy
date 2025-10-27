package services

import (
	"context"
	"errors"

	"primmfy_db/models"

	"github.com/jackc/pgx/v5"
)


// ═══════════════════════════════════════════════════════════
// COURSE SERVICE (Sub-topic dalam Lesson)
// ═══════════════════════════════════════════════════════════

// CreateCourse membuat course baru dalam lesson
func CreateCourse(db *pgx.Conn, teacherID int, req models.CreateCourseRequest) (*models.Course, error) {
    // 1. Cek apakah lesson exist dan teacher adalah owner
    var lessonTeacherID int
    err := db.QueryRow(context.Background(),
        "SELECT teacher_id FROM lessons WHERE id = $1 AND is_active = true",
        req.LessonID).Scan(&lessonTeacherID)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("lesson tidak ditemukan atau tidak aktif")
        }
        return nil, errors.New("gagal mengecek lesson: " + err.Error())
    }

    if lessonTeacherID != teacherID {
        return nil, errors.New("anda tidak memiliki akses untuk membuat course di lesson ini")
    }

    // 2. Insert course baru
    var course models.Course
    err = db.QueryRow(context.Background(), `
        INSERT INTO courses (lesson_id, title, description, order_index, coin_reward, is_active)
        VALUES ($1, $2, $3, $4, $5, true)
        RETURNING id, lesson_id, title, description, order_index, coin_reward, is_active, created_at, updated_at`,
        req.LessonID, req.Title, req.Description, req.OrderIndex, req.CoinReward).Scan(
        &course.ID, &course.LessonID, &course.Title, &course.Description,
        &course.OrderIndex, &course.CoinReward, &course.IsActive,
        &course.CreatedAt, &course.UpdatedAt)

    if err != nil {
        return nil, errors.New("gagal membuat course: " + err.Error())
    }

    return &course, nil
}

// GetCourseByID mengambil detail course by ID
func GetCourseByID(db *pgx.Conn, courseID int) (*models.Course, error) {
    var course models.Course
    err := db.QueryRow(context.Background(), `
        SELECT id, lesson_id, title, description, order_index, coin_reward, 
               is_active, created_at, updated_at
        FROM courses
        WHERE id = $1 AND is_active = true`,
        courseID).Scan(
        &course.ID, &course.LessonID, &course.Title, &course.Description,
        &course.OrderIndex, &course.CoinReward, &course.IsActive,
        &course.CreatedAt, &course.UpdatedAt)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("course tidak ditemukan")
        }
        return nil, errors.New("gagal mengambil course: " + err.Error())
    }

    return &course, nil
}

// GetCoursesByLesson mengambil semua course dalam lesson
func GetCoursesByLesson(db *pgx.Conn, lessonID int) ([]models.Course, error) {
    rows, err := db.Query(context.Background(), `
        SELECT id, lesson_id, title, description, order_index, coin_reward, 
               is_active, created_at, updated_at
        FROM courses
        WHERE lesson_id = $1 AND is_active = true
        ORDER BY order_index ASC`,
        lessonID)

    if err != nil {
        return nil, errors.New("gagal mengambil courses: " + err.Error())
    }
    defer rows.Close()

    var courses []models.Course
    for rows.Next() {
        var course models.Course
        err := rows.Scan(
            &course.ID, &course.LessonID, &course.Title, &course.Description,
            &course.OrderIndex, &course.CoinReward, &course.IsActive,
            &course.CreatedAt, &course.UpdatedAt)

        if err != nil {
            return nil, errors.New("gagal scan course: " + err.Error())
        }
        courses = append(courses, course)
    }

    return courses, nil
}

// UpdateCourse mengupdate course existing
func UpdateCourse(db *pgx.Conn, courseID int, teacherID int, req models.UpdateCourseRequest) (*models.Course, error) {
    // 1. Cek ownership (apakah teacher adalah owner lesson)
    var lessonTeacherID int
    err := db.QueryRow(context.Background(), `
        SELECT l.teacher_id
        FROM courses c
        JOIN lessons l ON c.lesson_id = l.id
        WHERE c.id = $1`,
        courseID).Scan(&lessonTeacherID)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("course tidak ditemukan")
        }
        return nil, errors.New("gagal mengecek course: " + err.Error())
    }

    if lessonTeacherID != teacherID {
        return nil, errors.New("anda tidak memiliki akses untuk mengupdate course ini")
    }

    // 2. Build dynamic update query
    query := "UPDATE courses SET updated_at = NOW()"
    args := []interface{}{}
    argID := 1

    if req.Title != "" {
        query += ", title = $" + string(rune(argID+'0'))
        args = append(args, req.Title)
        argID++
    }

    if req.Description != "" {
        query += ", description = $" + string(rune(argID+'0'))
        args = append(args, req.Description)
        argID++
    }

    if req.OrderIndex > 0 {
        query += ", order_index = $" + string(rune(argID+'0'))
        args = append(args, req.OrderIndex)
        argID++
    }

    if req.CoinReward > 0 {
        query += ", coin_reward = $" + string(rune(argID+'0'))
        args = append(args, req.CoinReward)
        argID++
    }

    if req.IsActive != nil {
        query += ", is_active = $" + string(rune(argID+'0'))
        args = append(args, *req.IsActive)
        argID++
    }

    query += " WHERE id = $" + string(rune(argID+'0'))
    args = append(args, courseID)

    query += " RETURNING id, lesson_id, title, description, order_index, coin_reward, is_active, created_at, updated_at"

    // 3. Execute update
    var course models.Course
    err = db.QueryRow(context.Background(), query, args...).Scan(
        &course.ID, &course.LessonID, &course.Title, &course.Description,
        &course.OrderIndex, &course.CoinReward, &course.IsActive,
        &course.CreatedAt, &course.UpdatedAt)

    if err != nil {
        return nil, errors.New("gagal update course: " + err.Error())
    }

    return &course, nil
}

// DeleteCourse menghapus course (soft delete dengan is_active = false)
func DeleteCourse(db *pgx.Conn, courseID int, teacherID int) error {
    // 1. Cek ownership
    var lessonTeacherID int
    err := db.QueryRow(context.Background(), `
        SELECT l.teacher_id
        FROM courses c
        JOIN lessons l ON c.lesson_id = l.id
        WHERE c.id = $1`,
        courseID).Scan(&lessonTeacherID)

    if err != nil {
        if err == pgx.ErrNoRows {
            return errors.New("course tidak ditemukan")
        }
        return errors.New("gagal mengecek course: " + err.Error())
    }

    if lessonTeacherID != teacherID {
        return errors.New("anda tidak memiliki akses untuk menghapus course ini")
    }

    // 2. Soft delete
    _, err = db.Exec(context.Background(),
        "UPDATE courses SET is_active = false, updated_at = NOW() WHERE id = $1",
        courseID)

    if err != nil {
        return errors.New("gagal menghapus course: " + err.Error())
    }

    return nil
}

// GetCourseWithStages mengambil course beserta semua PRIMM stages-nya
func GetCourseWithStages(db *pgx.Conn, courseID int) (*models.Course, []models.PRIMMStage, error) {
    // 1. Get course details
    course, err := GetCourseByID(db, courseID)
    if err != nil {
        return nil, nil, err
    }

    // 2. Get all stages in this course (dari primm_stage_service.go)
    stages, err := GetStagesByCourse(db, courseID)
    if err != nil {
        return course, []models.PRIMMStage{}, nil // Return empty stages if error
    }

    return course, stages, nil
}

// GetCoursesByLessonWithProgress mengambil courses dengan progress user
func GetCoursesByLessonWithProgress(db *pgx.Conn, lessonID int, userID int) ([]models.CourseWithProgress, error) {
    rows, err := db.Query(context.Background(), `
        SELECT 
            c.id, 
            c.lesson_id, 
            c.title, 
            c.description, 
            c.order_index, 
            c.coin_reward, 
            c.is_active, 
            c.created_at, 
            c.updated_at,
            COALESCE(COUNT(DISTINCT ps.id), 0) as total_stages,
            COALESCE(COUNT(DISTINCT ss.stage_id) FILTER (WHERE ss.is_correct = true), 0) as completed_stages,
            COALESCE(ucc.course_id, 0) > 0 as is_completed
        FROM courses c
        LEFT JOIN primm_stages ps ON c.id = ps.course_id AND ps.is_active = true
        LEFT JOIN stage_submissions ss ON ps.id = ss.stage_id AND ss.user_id = $1
        LEFT JOIN user_course_completion ucc ON c.id = ucc.course_id AND ucc.user_id = $1
        WHERE c.lesson_id = $2 AND c.is_active = true
        GROUP BY c.id, ucc.course_id
        ORDER BY c.order_index ASC
    `, userID, lessonID)

    if err != nil {
        return nil, errors.New("gagal mengambil courses: " + err.Error())
    }
    defer rows.Close()

    var courses []models.CourseWithProgress
    for rows.Next() {
        var course models.CourseWithProgress
        err := rows.Scan(
            &course.ID, 
            &course.LessonID, 
            &course.Title, 
            &course.Description,
            &course.OrderIndex, 
            &course.CoinReward, 
            &course.IsActive,
            &course.CreatedAt, 
            &course.UpdatedAt,
            &course.TotalStages, 
            &course.CompletedStages, 
            &course.IsCompleted,
        )

        if err != nil {
            return nil, errors.New("gagal scan course: " + err.Error())
        }

        // Calculate progress percent
        if course.TotalStages > 0 {
            course.ProgressPercent = (course.CompletedStages * 100) / course.TotalStages
        }

        courses = append(courses, course)
    }

    return courses, nil
}