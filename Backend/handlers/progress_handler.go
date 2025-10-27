package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
    "primmfy_db/models"
    "primmfy_db/services"
)

// ProgressHandler mengelola endpoint student progress & submissions
type ProgressHandler struct {
    DB *pgx.Conn
}

// NewProgressHandler membuat instance ProgressHandler baru
func NewProgressHandler(db *pgx.Conn) *ProgressHandler {
    return &ProgressHandler{DB: db}
}

// ═══════════════════════════════════════════════════════════
// SUBMIT STAGE ENDPOINTS (Student Only)
// ═══════════════════════════════════════════════════════════

// SubmitPredictStage handler untuk POST /api/stages/:id/submit-predict (student only)
// Purpose: Siswa submit jawaban pilihan ganda di PREDICT stage
func (h *ProgressHandler) SubmitPredictStage(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    // Get stage ID from URL param
    stageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Stage ID tidak valid"})
        return
    }

    // Bind request body
    var req models.SubmitPredictRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    // Override stage_id dari URL (lebih aman)
    req.StageID = stageID

    // Process submission
    response, err := services.SubmitPredictStage(h.DB, userID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

// SubmitRunStage handler untuk POST /api/stages/:id/submit-run (student only)
// Purpose: Siswa submit code di RUN stage (run code yang sudah ada)
func (h *ProgressHandler) SubmitRunStage(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    stageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Stage ID tidak valid"})
        return
    }

    var req models.SubmitRunRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    req.StageID = stageID

    response, err := services.SubmitRunStage(h.DB, userID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

// SubmitInvestigateStage handler untuk POST /api/stages/:id/submit-investigate (student only)
// Purpose: Siswa submit refleksi di INVESTIGATE stage
func (h *ProgressHandler) SubmitInvestigateStage(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    stageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Stage ID tidak valid"})
        return
    }

    var req models.SubmitInvestigateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    req.StageID = stageID

    response, err := services.SubmitInvestigateStage(h.DB, userID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

// SubmitModifyStage handler untuk POST /api/stages/:id/submit-modify (student only)
// Purpose: Siswa submit modified code dengan test case validation
func (h *ProgressHandler) SubmitModifyStage(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    stageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Stage ID tidak valid"})
        return
    }

    var req models.SubmitModifyRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    req.StageID = stageID

    response, err := services.SubmitModifyStage(h.DB, userID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

// SubmitMakeStage handler untuk POST /api/stages/:id/submit-make (student only)
// Purpose: Siswa submit original code dari scratch dengan test case validation
func (h *ProgressHandler) SubmitMakeStage(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    stageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Stage ID tidak valid"})
        return
    }

    var req models.SubmitMakeRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    req.StageID = stageID

    response, err := services.SubmitMakeStage(h.DB, userID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

// ═══════════════════════════════════════════════════════════
// PROGRESS TRACKING ENDPOINTS
// ═══════════════════════════════════════════════════════════

// GetMyProgress handler untuk GET /api/my-progress/:lesson_id (student only)
// Purpose: Siswa melihat progress mereka di lesson tertentu
func (h *ProgressHandler) GetMyProgress(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    lessonID, err := strconv.Atoi(c.Param("lesson_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Lesson ID tidak valid"})
        return
    }

    progress, err := services.GetMyProgress(h.DB, userID.(int), lessonID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "lesson_id": lessonID,
        "progress":  progress,
    })
}

// GetStageCompletion handler untuk GET /api/stages/:id/my-completion (student only)
// Purpose: Siswa melihat submission history mereka di stage tertentu
func (h *ProgressHandler) GetStageCompletion(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    stageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Stage ID tidak valid"})
        return
    }

    // Query submission history
    var completion models.UserStageCompletion
    err = h.DB.QueryRow(c.Request.Context(), `
        SELECT id, user_id, stage_id, 
               predict_selected_answer, predict_is_correct,
               run_submitted_code, run_output,
               investigate_reflection, investigate_completed,
               modify_submitted_code, modify_output, modify_is_correct, modify_attempts,
               make_submitted_code, make_output, make_is_correct, make_attempts,
               is_completed, completed_at, created_at, updated_at
        FROM user_stage_completions
        WHERE user_id = $1 AND stage_id = $2`,
        userID, stageID).Scan(
        &completion.ID, &completion.UserID, &completion.StageID,
        &completion.PredictSelectedAnswer, &completion.PredictIsCorrect,
        &completion.RunSubmittedCode, &completion.RunOutput,
        &completion.InvestigateReflection, &completion.InvestigateCompleted,
        &completion.ModifySubmittedCode, &completion.ModifyOutput,
        &completion.ModifyIsCorrect, &completion.ModifyAttempts,
        &completion.MakeSubmittedCode, &completion.MakeOutput,
        &completion.MakeIsCorrect, &completion.MakeAttempts,
        &completion.IsCompleted, &completion.CompletedAt,
        &completion.CreatedAt, &completion.UpdatedAt)

    if err != nil {
        // Jika belum ada submission, return empty
        c.JSON(http.StatusOK, gin.H{
            "message": "Belum ada submission untuk stage ini",
            "stage_id": stageID,
            "is_completed": false,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "completion": completion,
    })
}

// GetCourseProgress handler untuk GET /api/courses/:id/my-progress (student only)
// Purpose: Siswa melihat progress mereka di course tertentu (semua 5 stages)
func (h *ProgressHandler) GetCourseProgress(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    courseID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Course ID tidak valid"})
        return
    }

    // Get all stages in course with completion status
    rows, err := h.DB.Query(c.Request.Context(), `
        SELECT ps.id, ps.stage_type, ps.order_index, ps.title,
               COALESCE(usc.is_completed, false) as is_completed,
               COALESCE(usc.completed_at, NULL) as completed_at
        FROM primm_stages ps
        LEFT JOIN user_stage_completions usc 
          ON ps.id = usc.stage_id AND usc.user_id = $1
        WHERE ps.course_id = $2
        ORDER BY ps.order_index ASC`,
        userID, courseID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil progress: " + err.Error()})
        return
    }
    defer rows.Close()

    type StageProgress struct {
        StageID     int        `json:"stage_id"`
        StageType   string     `json:"stage_type"`
        OrderIndex  int        `json:"order_index"`
        Title       string     `json:"title"`
        IsCompleted bool       `json:"is_completed"`
        CompletedAt *string    `json:"completed_at,omitempty"`
    }

    var stages []StageProgress
    completedCount := 0

    for rows.Next() {
        var stage StageProgress
        var completedAt *string
        err := rows.Scan(
            &stage.StageID, &stage.StageType, &stage.OrderIndex,
            &stage.Title, &stage.IsCompleted, &completedAt)

        if err != nil {
            continue
        }

        if stage.IsCompleted {
            completedCount++
            stage.CompletedAt = completedAt
        }

        stages = append(stages, stage)
    }

    // Check course completion
    var courseCompletion models.UserCourseCompletion
    err = h.DB.QueryRow(c.Request.Context(), `
        SELECT id, user_id, course_id, is_completed, completed_at, coins_earned
        FROM user_course_completions
        WHERE user_id = $1 AND course_id = $2`,
        userID, courseID).Scan(
        &courseCompletion.ID, &courseCompletion.UserID,
        &courseCompletion.CourseID, &courseCompletion.IsCompleted,
        &courseCompletion.CompletedAt, &courseCompletion.CoinsEarned)

    isCourseCompleted := (err == nil && courseCompletion.IsCompleted)

    c.JSON(http.StatusOK, gin.H{
        "course_id":         courseID,
        "total_stages":      len(stages),
        "completed_stages":  completedCount,
        "progress_percent":  (completedCount * 100) / 5, // Always 5 PRIMM stages
        "is_completed":      isCourseCompleted,
        "coins_earned":      courseCompletion.CoinsEarned,
        "stages":            stages,
    })
}