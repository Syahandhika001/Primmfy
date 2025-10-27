package models

import "time"

// ═══════════════════════════════════════════════════════════
// STUDENT PROGRESS MODELS
// ═══════════════════════════════════════════════════════════

// UserLesson merepresentasikan enrollment siswa ke lesson
type UserLesson struct {
    ID         int       `json:"id"`
    UserID     int       `json:"user_id"`
    LessonID   int       `json:"lesson_id"`
    EnrolledAt time.Time `json:"enrolled_at"`
}

// UserCourseCompletion merepresentasikan completion status course
type UserCourseCompletion struct {
    ID          int        `json:"id"`
    UserID      int        `json:"user_id"`
    CourseID    int        `json:"course_id"`
    IsCompleted bool       `json:"is_completed"`
    CompletedAt *time.Time `json:"completed_at,omitempty"`
    CoinsEarned int        `json:"coins_earned"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}

// UserStageCompletion merepresentasikan submission & completion per stage
type UserStageCompletion struct {
    ID      int    `json:"id"`
    UserID  int    `json:"user_id"`
    StageID int    `json:"stage_id"`
    
    // PREDICT Stage Data
    PredictSelectedAnswer string `json:"predict_selected_answer,omitempty"`
    PredictIsCorrect      bool   `json:"predict_is_correct"`
    
    // RUN Stage Data
    RunSubmittedCode string `json:"run_submitted_code,omitempty"`
    RunOutput        string `json:"run_output,omitempty"`
    
    // INVESTIGATE Stage Data
    InvestigateReflection string `json:"investigate_reflection,omitempty"`
    InvestigateCompleted  bool   `json:"investigate_completed"`
    
    // MODIFY Stage Data
    ModifySubmittedCode string `json:"modify_submitted_code,omitempty"`
    ModifyOutput        string `json:"modify_output,omitempty"`
    ModifyIsCorrect     bool   `json:"modify_is_correct"`
    ModifyAttempts      int    `json:"modify_attempts"`
    
    // MAKE Stage Data
    MakeSubmittedCode string `json:"make_submitted_code,omitempty"`
    MakeOutput        string `json:"make_output,omitempty"`
    MakeIsCorrect     bool   `json:"make_is_correct"`
    MakeAttempts      int    `json:"make_attempts"`
    
    IsCompleted bool       `json:"is_completed"`
    CompletedAt *time.Time `json:"completed_at,omitempty"`
    
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// ═══════════════════════════════════════════════════════════
// SUBMIT REQUEST STRUCTS (Student submit answers)
// ═══════════════════════════════════════════════════════════

// SubmitPredictRequest untuk submit jawaban PREDICT stage
type SubmitPredictRequest struct {
    StageID        int    `json:"stage_id" binding:"required"`
    SelectedAnswer string `json:"selected_answer" binding:"required,oneof=A B C D"`
}

// SubmitRunRequest untuk submit code RUN stage
type SubmitRunRequest struct {
    StageID       int    `json:"stage_id" binding:"required"`
    SubmittedCode string `json:"submitted_code" binding:"required,min=1"`
}

// SubmitInvestigateRequest untuk submit refleksi INVESTIGATE stage
type SubmitInvestigateRequest struct {
    StageID    int    `json:"stage_id" binding:"required"`
    Reflection string `json:"reflection" binding:"required,min=20"`
}

// SubmitModifyRequest untuk submit modified code
type SubmitModifyRequest struct {
    StageID       int    `json:"stage_id" binding:"required"`
    SubmittedCode string `json:"submitted_code" binding:"required"`
}

// SubmitMakeRequest untuk submit code MAKE stage
type SubmitMakeRequest struct {
    StageID       int    `json:"stage_id" binding:"required"`
    SubmittedCode string `json:"submitted_code" binding:"required"`
}

// ═══════════════════════════════════════════════════════════
// RESPONSE STRUCTS
// ═══════════════════════════════════════════════════════════

// SubmitStageResponse adalah response setelah submit stage
type SubmitStageResponse struct {
    Success        bool   `json:"success"`
    IsCorrect      bool   `json:"is_correct"`
    Message        string `json:"message"`
    CoinsEarned    int    `json:"coins_earned"`
    XPEarned       int    `json:"xp_earned"`
    Output         string `json:"output,omitempty"`
    ExpectedOutput string `json:"expected_output,omitempty"`
}

// ProgressSummary adalah ringkasan progress siswa di lesson
type ProgressSummary struct {
    TotalCourses     int     `json:"total_courses"`
    CompletedCourses int     `json:"completed_courses"`
    ProgressPercent  float64 `json:"progress_percent"`
    TotalCoinsEarned int     `json:"total_coins_earned"`
    TotalXPEarned    int     `json:"total_xp_earned"`
}