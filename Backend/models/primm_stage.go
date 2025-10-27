package models

import "time"

// ═══════════════════════════════════════════════════════════
// PRIMM STAGE (Level 3 - Stage dalam Course)
// ═══════════════════════════════════════════════════════════

// PRIMMStage merepresentasikan satu stage dalam PRIMM methodology
type PRIMMStage struct {
    ID          int       `json:"id"`
    CourseID    int       `json:"course_id"`
    StageType   string    `json:"stage_type"` // 'predict', 'run', 'investigate', 'modify', 'make'
    Title       string    `json:"title"`
    Description string    `json:"description"`
    OrderIndex  int       `json:"order_index"` // 1-5 untuk PRIMM sequence
    IsActive    bool      `json:"is_active"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`

    // Common fields
    CodeSnippet     *string `json:"code_snippet,omitempty"`
    TaskDescription *string `json:"task_description,omitempty"` // ← TAMBAHKAN INI

    // PREDICT specific fields
    PredictOptions map[string]string `json:"predict_options,omitempty"`
    CorrectAnswer  *string           `json:"correct_answer,omitempty"`

    // RUN specific fields
    RunCodeTemplate *string `json:"run_code_template,omitempty"`

    // INVESTIGATE specific fields
    VideoEmbedURL     *string  `json:"video_embed_url,omitempty"`
    ExplanationText   *string  `json:"explanation_text,omitempty"`
    GuidingQuestions  []string `json:"guiding_questions,omitempty"`
    ReflectionPrompt  *string  `json:"reflection_prompt,omitempty"`

    // MODIFY specific fields
    ModifyChallenge      *string    `json:"modify_challenge,omitempty"`
    ModifyCodeTemplate   *string    `json:"modify_code_template,omitempty"`
    ModifyExpectedOutput *string    `json:"modify_expected_output,omitempty"`
    ModifyTestCases      []TestCase `json:"modify_test_cases,omitempty"`

    // MAKE specific fields
    MakeChallenge      *string    `json:"make_challenge,omitempty"`
    MakeHints          *string    `json:"make_hints,omitempty"`
    MakeExpectedOutput *string    `json:"make_expected_output,omitempty"`
    MakeTestCases      []TestCase `json:"make_test_cases,omitempty"`
}

// TestCase untuk validasi output code
type TestCase struct {
    Input          string `json:"input"`
    ExpectedOutput string `json:"expected_output"`
    Description    string `json:"description,omitempty"`
}

// ═══════════════════════════════════════════════════════════
// CREATE REQUESTS untuk TEACHER membuat PRIMM STAGES
// ═══════════════════════════════════════════════════════════

// CreatePredictStageRequest untuk membuat PREDICT stage
type CreatePredictStageRequest struct {
    CourseID       int               `json:"course_id" binding:"required"`
    Title          string            `json:"title" binding:"required"`
    Description    string            `json:"description" binding:"required"`
    CodeSnippet    string            `json:"code_snippet" binding:"required"`
    PredictOptions map[string]string `json:"predict_options" binding:"required"`
    CorrectAnswer  string            `json:"correct_answer" binding:"required"`
}

// CreateRunStageRequest untuk membuat RUN stage
type CreateRunStageRequest struct {
    CourseID        int    `json:"course_id" binding:"required"`
    Title           string `json:"title" binding:"required"`
    Description     string `json:"description" binding:"required"`
    CodeSnippet     string `json:"code_snippet" binding:"required"` // Code yang akan dijalankan
    RunCodeTemplate string `json:"run_code_template"` // ← HAPUS 'required', jadi OPTIONAL
}

// CreateInvestigateStageRequest untuk membuat INVESTIGATE stage
type CreateInvestigateStageRequest struct {
    CourseID          int      `json:"course_id" binding:"required"`
    Title             string   `json:"title" binding:"required"`
    Description       string   `json:"description" binding:"required"`
    VideoURL          string   `json:"video_url" binding:"required,url"`
    ExplanationText   string   `json:"explanation_text"` // ← HAPUS 'required', jadi OPTIONAL
    GuidingQuestions  []string `json:"guiding_questions" binding:"required,min=1"`
    ReflectionPrompt  string   `json:"reflection_prompt"` // ← HAPUS 'required', jadi OPTIONAL
}

// CreateModifyStageRequest untuk membuat MODIFY stage
type CreateModifyStageRequest struct {
    CourseID             int        `json:"course_id" binding:"required"`
    Title                string     `json:"title" binding:"required"`
    Description          string     `json:"description" binding:"required"`
    CodeSnippet          string     `json:"code_snippet" binding:"required"` // Code original
    TaskDescription      string     `json:"task_description" binding:"required"` // Instruksi modifikasi
    ModifyChallenge      string     `json:"modify_challenge"` // ← HAPUS 'required'
    ModifyCodeTemplate   string     `json:"modify_code_template"` // ← HAPUS 'required'
    ModifyExpectedOutput string     `json:"modify_expected_output"` // ← HAPUS 'required'
    ModifyTestCases      []TestCase `json:"modify_test_cases" binding:"required,min=1,dive"`
}

// CreateMakeStageRequest untuk membuat MAKE stage
type CreateMakeStageRequest struct {
    CourseID           int        `json:"course_id" binding:"required"`
    Title              string     `json:"title" binding:"required"`
    Description        string     `json:"description" binding:"required"`
    TaskDescription    string     `json:"task_description" binding:"required"`
    MakeChallenge      string     `json:"make_challenge"` // ← HAPUS 'required'
    MakeExpectedOutput string     `json:"make_expected_output"` // ← HAPUS 'required'
    MakeTestCases      []TestCase `json:"make_test_cases" binding:"required,min=1,dive"`
}

// StageSubmissionRequest untuk request submit stage
type StageSubmissionRequest struct {
    SubmissionType string                 `json:"submission_type" binding:"required,oneof=predict run investigate modify make"`
    SelectedAnswer string                 `json:"selected_answer,omitempty"` // For PREDICT
    CodeOutput     string                 `json:"code_output,omitempty"`     // For RUN
    ReflectionText string                 `json:"reflection_text,omitempty"` // For INVESTIGATE
    ModifiedCode   string                 `json:"modified_code,omitempty"`   // For MODIFY
    Code           string                 `json:"code,omitempty"`            // For MAKE
    SubmissionData map[string]interface{} `json:"submission_data,omitempty"` // Generic data
}

// StageSubmission untuk response submission
type StageSubmission struct {
    ID             int                    `json:"id"`
    UserID         int                    `json:"user_id"`
    StageID        int                    `json:"stage_id"`
    SubmissionType string                 `json:"submission_type"`
    SubmissionData map[string]interface{} `json:"submission_data"`
    IsCorrect      bool                   `json:"is_correct"`
    Score          int                    `json:"score"`
    SubmittedAt    time.Time              `json:"submitted_at"`
}
