package handlers

import (
	"net/http"
	"strconv"

	"primmfy_db/models"
	"primmfy_db/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// PRIMMStageHandler mengelola endpoint PRIMM stages
type PRIMMStageHandler struct {
    DB *pgx.Conn
}

// NewPRIMMStageHandler membuat instance PRIMMStageHandler baru
func NewPRIMMStageHandler(db *pgx.Conn) *PRIMMStageHandler {
    return &PRIMMStageHandler{DB: db}
}

// ═══════════════════════════════════════════════════════════
// CREATE PRIMM STAGE ENDPOINTS (Teacher Only)
// ═══════════════════════════════════════════════════════════

// CreatePredictStage handler untuk POST /api/stages/predict (teacher only)
func (h *PRIMMStageHandler) CreatePredictStage(c *gin.Context) {
    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    var req models.CreatePredictStageRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    stage, err := services.CreatePredictStage(h.DB, teacherID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "PREDICT stage berhasil dibuat!",
        "stage":   stage,
    })
}

// CreateRunStage handler untuk POST /api/stages/run (teacher only)
func (h *PRIMMStageHandler) CreateRunStage(c *gin.Context) {
    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    var req models.CreateRunStageRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    stage, err := services.CreateRunStage(h.DB, teacherID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "RUN stage berhasil dibuat!",
        "stage":   stage,
    })
}

// CreateInvestigateStage handler untuk POST /api/stages/investigate (teacher only)
func (h *PRIMMStageHandler) CreateInvestigateStage(c *gin.Context) {
    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    var req models.CreateInvestigateStageRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    stage, err := services.CreateInvestigateStage(h.DB, teacherID.(int), req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "INVESTIGATE stage berhasil dibuat!",
        "stage":   stage,
    })
}

// CreateModifyStage handler untuk POST /api/stages/modify (teacher only)
func (h *PRIMMStageHandler) CreateModifyStage(c *gin.Context) {
    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    var req models.CreateModifyStageRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    stage, err := services.CreateModifyStage(h.DB, teacherID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "MODIFY stage berhasil dibuat!",
        "stage":   stage,
    })
}

// CreateMakeStage handler untuk POST /api/stages/make (teacher only)
func (h *PRIMMStageHandler) CreateMakeStage(c *gin.Context) {
    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    var req models.CreateMakeStageRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid: " + err.Error()})
        return
    }

    stage, err := services.CreateMakeStage(h.DB, teacherID.(int), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "MAKE stage berhasil dibuat!",
        "stage":   stage,
    })
}

// ═══════════════════════════════════════════════════════════
// GET & DELETE STAGE ENDPOINTS
// ═══════════════════════════════════════════════════════════

// GetStageByID handler untuk GET /api/stages/:id
func (h *PRIMMStageHandler) GetStageByID(c *gin.Context) {
    stageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Stage ID tidak valid"})
        return
    }

    stage, err := services.GetStageByID(h.DB, stageID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"stage": stage})
}

// DeleteStage handler untuk DELETE /api/stages/:id (teacher only)
func (h *PRIMMStageHandler) DeleteStage(c *gin.Context) {
    stageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Stage ID tidak valid"})
        return
    }

    teacherID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID tidak ditemukan"})
        return
    }

    err = services.DeleteStage(h.DB, stageID, teacherID.(int))
    if err != nil {
        if err.Error() == "stage tidak ditemukan" {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else if err.Error() == "anda tidak memiliki akses untuk menghapus stage ini" {
            c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Stage berhasil dihapus!"})
}

// SubmitStage handler untuk POST /api/stages/:id/submit
func (h *PRIMMStageHandler) SubmitStage(c *gin.Context) {
    // 1. Parse stage ID
    stageID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Stage ID tidak valid"})
        return
    }

    // 2. Get user ID dari JWT
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak terautentikasi"})
        return
    }

    // 3. Parse request body
    var req models.StageSubmissionRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Request body tidak valid: " + err.Error()})
        return
    }

    // 4. Submit via service
    submission, err := services.SubmitStage(h.DB, userID.(int), stageID, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 5. Return response
    c.JSON(http.StatusCreated, gin.H{
        "message":    "Jawaban berhasil disubmit!",
        "is_correct": submission.IsCorrect,
        "score":      submission.Score,
        "submission": submission,
    })
}