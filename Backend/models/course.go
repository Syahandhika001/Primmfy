package models

import "time"

// Course model (basic course data)
type Course struct {
    ID          int       `json:"id"`
    LessonID    int       `json:"lesson_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    OrderIndex  int       `json:"order_index"`
    CoinReward  int       `json:"coin_reward"`
    IsActive    bool      `json:"is_active"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// CourseWithProgress untuk course dengan tracking progress student
type CourseWithProgress struct {
    ID              int       `json:"id"`
    LessonID        int       `json:"lesson_id"`
    Title           string    `json:"title"`
    Description     string    `json:"description"`
    OrderIndex      int       `json:"order_index"`
    CoinReward      int       `json:"coin_reward"`
    IsActive        bool      `json:"is_active"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
    TotalStages     int       `json:"total_stages"`
    CompletedStages int       `json:"completed_stages"`
    ProgressPercent int       `json:"progress_percent"`
    IsCompleted     bool      `json:"is_completed"`
}

// CreateCourseRequest untuk membuat course baru
type CreateCourseRequest struct {
    LessonID    int    `json:"lesson_id" binding:"required"`
    Title       string `json:"title" binding:"required,min=3,max=200"`
    Description string `json:"description" binding:"required,min=10"`
    OrderIndex  int    `json:"order_index" binding:"required,min=1"`
    CoinReward  int    `json:"coin_reward" binding:"required,min=10,max=1000"`
}

// UpdateCourseRequest untuk update course existing
type UpdateCourseRequest struct {
    Title       string `json:"title" binding:"omitempty,min=3,max=200"`
    Description string `json:"description" binding:"omitempty,min=10"`
    OrderIndex  int    `json:"order_index" binding:"omitempty,min=1"`
    CoinReward  int    `json:"coin_reward" binding:"omitempty,min=10,max=1000"`
    IsActive    *bool  `json:"is_active"`
}