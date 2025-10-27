package models

import "time"

// Lesson model
type Lesson struct {
    ID           int       `json:"id"`
    TeacherID    int       `json:"teacher_id"`
    Title        string    `json:"title"`
    Description  string    `json:"description"`
    Category     string    `json:"category"`
    Difficulty   string    `json:"difficulty"`
    ThumbnailURL string    `json:"thumbnail_url"`
    IsActive     bool      `json:"is_active"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

// LessonWithTeacher untuk response dengan nama teacher
type LessonWithTeacher struct {
    Lesson
    TeacherName string `json:"teacher_name"`
}

// LessonWithProgress untuk response dengan progress tracking
type LessonWithProgress struct {
    ID               int       `json:"id"`
    TeacherID        int       `json:"teacher_id"`
    Title            string    `json:"title"`
    Description      string    `json:"description"`
    Category         string    `json:"category"`
    Difficulty       string    `json:"difficulty"`
    ThumbnailURL     string    `json:"thumbnail_url"`
    IsActive         bool      `json:"is_active"`
    CreatedAt        time.Time `json:"created_at"`
    UpdatedAt        time.Time `json:"updated_at"`
    TeacherName      string    `json:"teacher_name"`
    TotalCourses     int       `json:"total_courses"`
    CompletedCourses int       `json:"completed_courses"`
    ProgressPercent  int       `json:"progress_percent"`
    IsEnrolled       bool      `json:"is_enrolled"`
}

// CreateLessonRequest untuk membuat lesson baru
type CreateLessonRequest struct {
    Title        string `json:"title" binding:"required,min=3,max=200"`
    Description  string `json:"description" binding:"required,min=10"`
    Category     string `json:"category" binding:"required,oneof=python javascript golang java cpp"`
    Difficulty   string `json:"difficulty" binding:"required,oneof=beginner intermediate advanced"`
    ThumbnailURL string `json:"thumbnail_url" binding:"omitempty,url"`
}

// UpdateLessonRequest untuk update lesson existing
type UpdateLessonRequest struct {
    Title        string `json:"title" binding:"omitempty,min=3,max=200"`
    Description  string `json:"description" binding:"omitempty,min=10"`
    Category     string `json:"category" binding:"omitempty,oneof=python javascript golang java cpp"`
    Difficulty   string `json:"difficulty" binding:"omitempty,oneof=beginner intermediate advanced"`
    ThumbnailURL string `json:"thumbnail_url" binding:"omitempty,url"`
    IsActive     *bool  `json:"is_active"`
}