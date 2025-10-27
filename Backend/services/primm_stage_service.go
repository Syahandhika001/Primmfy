package services

import (
	"context"
	"encoding/json"
	"errors"

	"primmfy_db/models"

	"github.com/jackc/pgx/v5"
)

// ═══════════════════════════════════════════════════════════
// PRIMM STAGE CRUD OPERATIONS (Teacher Only)
// ═══════════════════════════════════════════════════════════

// CreatePredictStage membuat PREDICT stage (order_index = 1)
func CreatePredictStage(db *pgx.Conn, teacherID int, req models.CreatePredictStageRequest) (*models.PRIMMStage, error) {
    // Cek ownership course
    if err := checkCourseOwnership(db, req.CourseID, teacherID); err != nil {
        return nil, err
    }

    // Convert predict_options map to JSONB
    optionsJSON, err := json.Marshal(req.PredictOptions)
    if err != nil {
        return nil, errors.New("gagal convert predict_options: " + err.Error())
    }

    var stage models.PRIMMStage
    err = db.QueryRow(context.Background(), `
        INSERT INTO primm_stages 
        (course_id, stage_type, title, description, order_index, 
         code_snippet, predict_options, correct_answer)
        VALUES ($1, 'predict', $2, $3, 1, $4, $5, $6)
        RETURNING id, course_id, stage_type, title, description, order_index,
                  code_snippet, predict_options, correct_answer, 
                  created_at, updated_at`,
        req.CourseID, req.Title, req.Description, req.CodeSnippet, 
        optionsJSON, req.CorrectAnswer).Scan(
        &stage.ID, &stage.CourseID, &stage.StageType, &stage.Title, 
        &stage.Description, &stage.OrderIndex, &stage.CodeSnippet,
        &optionsJSON, &stage.CorrectAnswer, &stage.CreatedAt, &stage.UpdatedAt)

    if err != nil {
        if err.Error() == "duplicate key value violates unique constraint \"primm_stages_course_id_order_index_key\"" {
            return nil, errors.New("stage dengan order_index 1 sudah ada di course ini")
        }
        return nil, errors.New("gagal membuat predict stage: " + err.Error())
    }

    // Parse JSONB back to map
    json.Unmarshal(optionsJSON, &stage.PredictOptions)

    return &stage, nil
}

// CreateRunStage membuat RUN stage (order_index = 2)
func CreateRunStage(db *pgx.Conn, teacherID int, req models.CreateRunStageRequest) (*models.PRIMMStage, error) {
    if err := checkCourseOwnership(db, req.CourseID, teacherID); err != nil {
        return nil, err
    }

    var stage models.PRIMMStage
    err := db.QueryRow(context.Background(), `
        INSERT INTO primm_stages 
        (course_id, stage_type, title, description, order_index, run_code_template)
        VALUES ($1, 'run', $2, $3, 2, $4)
        RETURNING id, course_id, stage_type, title, description, order_index,
                  run_code_template, created_at, updated_at`,
        req.CourseID, req.Title, req.Description, req.RunCodeTemplate).Scan(
        &stage.ID, &stage.CourseID, &stage.StageType, &stage.Title,
        &stage.Description, &stage.OrderIndex, &stage.RunCodeTemplate,
        &stage.CreatedAt, &stage.UpdatedAt)

    if err != nil {
        return nil, errors.New("gagal membuat run stage: " + err.Error())
    }

    return &stage, nil
}

// CreateInvestigateStage membuat INVESTIGATE stage baru
func CreateInvestigateStage(db *pgx.Conn, teacherID int, req models.CreateInvestigateStageRequest) (*models.PRIMMStage, error) {
    // 1. Validasi teacher owns this course's lesson
    var lessonTeacherID int
    err := db.QueryRow(context.Background(), `
        SELECT l.teacher_id 
        FROM courses c
        JOIN lessons l ON c.lesson_id = l.id
        WHERE c.id = $1`,
        req.CourseID).Scan(&lessonTeacherID)

    if err != nil {
        return nil, errors.New("course tidak ditemukan")
    }

    if lessonTeacherID != teacherID {
        return nil, errors.New("anda tidak memiliki akses untuk membuat stage di course ini")
    }

    // 2. Get next order_index
    var maxOrder int
    err = db.QueryRow(context.Background(),
        "SELECT COALESCE(MAX(order_index), 0) FROM primm_stages WHERE course_id = $1",
        req.CourseID).Scan(&maxOrder)

    if err != nil {
        return nil, errors.New("gagal mendapatkan order index: " + err.Error())
    }

    nextOrder := maxOrder + 1

    // 3. Serialize guiding questions
    questionsJSON, err := json.Marshal(req.GuidingQuestions)
    if err != nil {
        return nil, errors.New("gagal serialize guiding questions: " + err.Error())
    }

    // 4. Insert INVESTIGATE stage
    var stage models.PRIMMStage
    err = db.QueryRow(context.Background(), `
        INSERT INTO primm_stages (
            course_id, stage_type, title, description, order_index,
            video_embed_url, explanation_text, guiding_questions, reflection_prompt
        ) VALUES ($1, 'investigate', $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, course_id, stage_type, title, description, order_index,
                  video_embed_url, explanation_text, guiding_questions, reflection_prompt,
                  created_at, updated_at`,
        req.CourseID, req.Title, req.Description, nextOrder,
        req.VideoURL, req.ExplanationText, questionsJSON, req.ReflectionPrompt,
    ).Scan(
        &stage.ID, &stage.CourseID, &stage.StageType, &stage.Title, &stage.Description,
        &stage.OrderIndex, &stage.VideoEmbedURL, &stage.ExplanationText,
        &questionsJSON, &stage.ReflectionPrompt, &stage.CreatedAt, &stage.UpdatedAt,
    )

    if err != nil {
        return nil, errors.New("gagal membuat INVESTIGATE stage: " + err.Error())
    }

    // 5. Deserialize guiding questions
    json.Unmarshal(questionsJSON, &stage.GuidingQuestions)

    return &stage, nil
}

// CreateMakeStage membuat MAKE stage baru
func CreateMakeStage(db *pgx.Conn, teacherID int, req models.CreateMakeStageRequest) (*models.PRIMMStage, error) {
    // 1. Validasi teacher owns this course's lesson
    var lessonTeacherID int
    err := db.QueryRow(context.Background(), `
        SELECT l.teacher_id 
        FROM courses c
        JOIN lessons l ON c.lesson_id = l.id
        WHERE c.id = $1`,
        req.CourseID).Scan(&lessonTeacherID)

    if err != nil {
        return nil, errors.New("course tidak ditemukan")
    }

    if lessonTeacherID != teacherID {
        return nil, errors.New("anda tidak memiliki akses untuk membuat stage di course ini")
    }

    // 2. Get next order_index
    var maxOrder int
    err = db.QueryRow(context.Background(),
        "SELECT COALESCE(MAX(order_index), 0) FROM primm_stages WHERE course_id = $1",
        req.CourseID).Scan(&maxOrder)

    if err != nil {
        return nil, errors.New("gagal mendapatkan order index: " + err.Error())
    }

    nextOrder := maxOrder + 1

    // 3. Serialize test cases ke JSON
    testCasesJSON, err := json.Marshal(req.MakeTestCases)
    if err != nil {
        return nil, errors.New("gagal serialize test cases: " + err.Error())
    }

    // 4. Insert MAKE stage
    var stage models.PRIMMStage
    err = db.QueryRow(context.Background(), `
        INSERT INTO primm_stages (
            course_id, stage_type, title, description, order_index,
            task_description, make_challenge, make_expected_output, make_test_cases
        ) VALUES ($1, 'make', $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, course_id, stage_type, title, description, order_index,
                  task_description, make_challenge, make_expected_output, make_test_cases,
                  created_at, updated_at`,
        req.CourseID, req.Title, req.Description, nextOrder,
        req.TaskDescription, req.MakeChallenge, req.MakeExpectedOutput, testCasesJSON,
    ).Scan(
        &stage.ID, &stage.CourseID, &stage.StageType, &stage.Title, &stage.Description,
        &stage.OrderIndex, &stage.TaskDescription, &stage.MakeChallenge,
        &stage.MakeExpectedOutput, &testCasesJSON, &stage.CreatedAt, &stage.UpdatedAt,
    )

    if err != nil {
        return nil, errors.New("gagal membuat MAKE stage: " + err.Error())
    }

    // 5. Deserialize test cases untuk response
    json.Unmarshal(testCasesJSON, &stage.MakeTestCases)

    return &stage, nil
}

// CreateModifyStage membuat MODIFY stage (order_index = 4)
func CreateModifyStage(db *pgx.Conn, teacherID int, req models.CreateModifyStageRequest) (*models.PRIMMStage, error) {
    if err := checkCourseOwnership(db, req.CourseID, teacherID); err != nil {
        return nil, err
    }

    testCasesJSON, err := json.Marshal(req.ModifyTestCases)
    if err != nil {
        return nil, errors.New("gagal convert test_cases: " + err.Error())
    }

    var stage models.PRIMMStage
    err = db.QueryRow(context.Background(), `
        INSERT INTO primm_stages 
        (course_id, stage_type, title, description, order_index,
         modify_challenge, modify_code_template, modify_expected_output, modify_test_cases)
        VALUES ($1, 'modify', $2, $3, 4, $4, $5, $6, $7)
        RETURNING id, course_id, stage_type, title, description, order_index,
                  modify_challenge, modify_code_template, modify_expected_output, 
                  modify_test_cases, created_at, updated_at`,
        req.CourseID, req.Title, req.Description, req.ModifyChallenge,
        req.ModifyCodeTemplate, req.ModifyExpectedOutput, testCasesJSON).Scan(
        &stage.ID, &stage.CourseID, &stage.StageType, &stage.Title,
        &stage.Description, &stage.OrderIndex, &stage.ModifyChallenge,
        &stage.ModifyCodeTemplate, &stage.ModifyExpectedOutput, &testCasesJSON,
        &stage.CreatedAt, &stage.UpdatedAt)

    if err != nil {
        return nil, errors.New("gagal membuat modify stage: " + err.Error())
    }

    json.Unmarshal(testCasesJSON, &stage.ModifyTestCases)

    return &stage, nil
}

// GetStagesByCourse mengambil semua stages dalam course (urut by order_index)
func GetStagesByCourse(db *pgx.Conn, courseID int) ([]models.PRIMMStage, error) {
    rows, err := db.Query(context.Background(), `
        SELECT id, course_id, stage_type, title, description, order_index,
               code_snippet, predict_options, correct_answer,
               run_code_template, reflection_prompt, video_embed_url, 
               explanation_text, modify_challenge, modify_code_template,
               modify_expected_output, modify_test_cases, make_challenge,
               make_hints, make_expected_output, make_test_cases,
               created_at, updated_at
        FROM primm_stages
        WHERE course_id = $1
        ORDER BY order_index ASC`, courseID)

    if err != nil {
        return nil, errors.New("gagal mengambil stages: " + err.Error())
    }
    defer rows.Close()

    var stages []models.PRIMMStage
    for rows.Next() {
        var stage models.PRIMMStage
        var predictOptionsJSON, modifyTestCasesJSON, makeTestCasesJSON []byte

        err := rows.Scan(
            &stage.ID, &stage.CourseID, &stage.StageType, &stage.Title,
            &stage.Description, &stage.OrderIndex, &stage.CodeSnippet,
            &predictOptionsJSON, &stage.CorrectAnswer, &stage.RunCodeTemplate,
            &stage.ReflectionPrompt, &stage.VideoEmbedURL, &stage.ExplanationText,
            &stage.ModifyChallenge, &stage.ModifyCodeTemplate,
            &stage.ModifyExpectedOutput, &modifyTestCasesJSON,
            &stage.MakeChallenge, &stage.MakeHints, &stage.MakeExpectedOutput,
            &makeTestCasesJSON, &stage.CreatedAt, &stage.UpdatedAt)

        if err != nil {
            return nil, errors.New("gagal scan stage: " + err.Error())
        }

        // Parse JSONB fields
        if predictOptionsJSON != nil {
            json.Unmarshal(predictOptionsJSON, &stage.PredictOptions)
        }
        if modifyTestCasesJSON != nil {
            json.Unmarshal(modifyTestCasesJSON, &stage.ModifyTestCases)
        }
        if makeTestCasesJSON != nil {
            json.Unmarshal(makeTestCasesJSON, &stage.MakeTestCases)
        }

        stages = append(stages, stage)
    }

    return stages, nil
}

// GetStageByID mengambil detail stage berdasarkan ID
func GetStageByID(db *pgx.Conn, stageID int) (*models.PRIMMStage, error) {
    var stage models.PRIMMStage
    var predictOptionsJSON, modifyTestCasesJSON, makeTestCasesJSON []byte

    err := db.QueryRow(context.Background(), `
        SELECT id, course_id, stage_type, title, description, order_index,
               code_snippet, predict_options, correct_answer,
               run_code_template, reflection_prompt, video_embed_url, 
               explanation_text, modify_challenge, modify_code_template,
               modify_expected_output, modify_test_cases, make_challenge,
               make_hints, make_expected_output, make_test_cases,
               created_at, updated_at
        FROM primm_stages
        WHERE id = $1`, stageID).Scan(
        &stage.ID, &stage.CourseID, &stage.StageType, &stage.Title,
        &stage.Description, &stage.OrderIndex, &stage.CodeSnippet,
        &predictOptionsJSON, &stage.CorrectAnswer, &stage.RunCodeTemplate,
        &stage.ReflectionPrompt, &stage.VideoEmbedURL, &stage.ExplanationText,
        &stage.ModifyChallenge, &stage.ModifyCodeTemplate,
        &stage.ModifyExpectedOutput, &modifyTestCasesJSON,
        &stage.MakeChallenge, &stage.MakeHints, &stage.MakeExpectedOutput,
        &makeTestCasesJSON, &stage.CreatedAt, &stage.UpdatedAt)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("stage tidak ditemukan")
        }
        return nil, errors.New("gagal mengambil stage: " + err.Error())
    }

    // Parse JSONB
    if predictOptionsJSON != nil {
        json.Unmarshal(predictOptionsJSON, &stage.PredictOptions)
    }
    if modifyTestCasesJSON != nil {
        json.Unmarshal(modifyTestCasesJSON, &stage.ModifyTestCases)
    }
    if makeTestCasesJSON != nil {
        json.Unmarshal(makeTestCasesJSON, &stage.MakeTestCases)
    }

    return &stage, nil
}

// DeleteStage menghapus stage (hard delete karena stage adalah part of course)
func DeleteStage(db *pgx.Conn, stageID int, teacherID int) error {
    // Cek ownership via course -> lesson
    var lessonOwnerID int
    err := db.QueryRow(context.Background(), `
        SELECT l.teacher_id
        FROM primm_stages ps
        JOIN courses c ON ps.course_id = c.id
        JOIN lessons l ON c.lesson_id = l.id
        WHERE ps.id = $1`, stageID).Scan(&lessonOwnerID)

    if err != nil {
        if err == pgx.ErrNoRows {
            return errors.New("stage tidak ditemukan")
        }
        return errors.New("gagal cek ownership: " + err.Error())
    }

    if lessonOwnerID != teacherID {
        return errors.New("anda tidak memiliki akses untuk menghapus stage ini")
    }

    // Hard delete
    _, err = db.Exec(context.Background(),
        "DELETE FROM primm_stages WHERE id = $1", stageID)

    if err != nil {
        return errors.New("gagal menghapus stage: " + err.Error())
    }

    return nil
}

// SubmitStage memproses submission student untuk stage tertentu
func SubmitStage(db *pgx.Conn, userID int, stageID int, req models.StageSubmissionRequest) (*models.StageSubmission, error) {
    // 1. Get stage details
    var stage models.PRIMMStage
    var predictOptionsJSON, modifyTestCasesJSON, makeTestCasesJSON []byte
    
    err := db.QueryRow(context.Background(), `
        SELECT id, course_id, stage_type, title, description, order_index,
               code_snippet, predict_options, correct_answer,
               run_code_template, video_embed_url, explanation_text, reflection_prompt,
               modify_challenge, modify_code_template, modify_expected_output, modify_test_cases,
               make_challenge, make_expected_output, make_test_cases,
               created_at, updated_at
        FROM primm_stages WHERE id = $1
    `, stageID).Scan(
        &stage.ID, &stage.CourseID, &stage.StageType, &stage.Title, &stage.Description,
        &stage.OrderIndex, &stage.CodeSnippet, &predictOptionsJSON, &stage.CorrectAnswer,
        &stage.RunCodeTemplate, &stage.VideoEmbedURL, &stage.ExplanationText, &stage.ReflectionPrompt,
        &stage.ModifyChallenge, &stage.ModifyCodeTemplate, &stage.ModifyExpectedOutput, &modifyTestCasesJSON,
        &stage.MakeChallenge, &stage.MakeExpectedOutput, &makeTestCasesJSON,
        &stage.CreatedAt, &stage.UpdatedAt,
    )

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("stage tidak ditemukan")
        }
        return nil, errors.New("gagal mengambil data stage: " + err.Error())
    }

    // Parse JSONB fields
    if predictOptionsJSON != nil {
        json.Unmarshal(predictOptionsJSON, &stage.PredictOptions)
    }
    if modifyTestCasesJSON != nil {
        json.Unmarshal(modifyTestCasesJSON, &stage.ModifyTestCases)
    }
    if makeTestCasesJSON != nil {
        json.Unmarshal(makeTestCasesJSON, &stage.MakeTestCases)
    }

    // 2. Validate submission type matches stage type
    if req.SubmissionType != stage.StageType {
        return nil, errors.New("tipe submission tidak sesuai dengan tipe stage")
    }

    // 3. Check & grade submission based on stage type
    isCorrect := false
    score := 0
    submissionData := make(map[string]interface{})

    switch stage.StageType {
    case "predict":
        // Check if answer is correct
        if stage.CorrectAnswer != nil {
            isCorrect = req.SelectedAnswer == *stage.CorrectAnswer
            if isCorrect {
                score = 100
            }
            submissionData["selected_answer"] = req.SelectedAnswer
            submissionData["correct_answer"] = *stage.CorrectAnswer
        } else {
            return nil, errors.New("stage tidak memiliki correct_answer")
        }

    case "run":
        // For RUN stage, just accept the output (no auto-grading for now)
        isCorrect = true
        score = 100
        submissionData["code_output"] = req.CodeOutput

    case "investigate":
        // For INVESTIGATE, check if reflection has minimum length
        isCorrect = len(req.ReflectionText) >= 50 // Minimum 50 characters
        if isCorrect {
            score = 100
        }
        submissionData["reflection_text"] = req.ReflectionText

    case "modify":
        // For MODIFY, we'd need to run code and check output
        // For now, simple check if code is not empty
        isCorrect = len(req.ModifiedCode) > 0
        if isCorrect {
            score = 100
        }
        submissionData["modified_code"] = req.ModifiedCode

    case "make":
        // For MAKE, simple check if code is not empty
        isCorrect = len(req.Code) > 0
        if isCorrect {
            score = 100
        }
        submissionData["code"] = req.Code

    default:
        return nil, errors.New("tipe stage tidak valid")
    }

    // 4. Convert submission data to JSON
    submissionJSON, err := json.Marshal(submissionData)
    if err != nil {
        return nil, errors.New("gagal menyimpan data submission: " + err.Error())
    }

    // 5. Insert or update submission
    var submission models.StageSubmission
    var returnedJSON []byte
    err = db.QueryRow(context.Background(), `
        INSERT INTO stage_submissions (user_id, stage_id, submission_type, submission_data, is_correct, score)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (user_id, stage_id) 
        DO UPDATE SET 
            submission_data = EXCLUDED.submission_data,
            is_correct = EXCLUDED.is_correct,
            score = EXCLUDED.score,
            submitted_at = CURRENT_TIMESTAMP
        RETURNING id, user_id, stage_id, submission_type, submission_data, is_correct, score, submitted_at
    `, userID, stageID, req.SubmissionType, submissionJSON, isCorrect, score).Scan(
        &submission.ID,
        &submission.UserID,
        &submission.StageID,
        &submission.SubmissionType,
        &returnedJSON,
        &submission.IsCorrect,
        &submission.Score,
        &submission.SubmittedAt,
    )

    if err != nil {
        return nil, errors.New("gagal menyimpan submission: " + err.Error())
    }

    // 6. Parse submission data back
    json.Unmarshal(returnedJSON, &submission.SubmissionData)

    // 7. Check if all stages in course are completed
    go checkCourseCompletion(db, userID, stage.CourseID)

    return &submission, nil
}

// checkCourseCompletion checks if all stages in a course are completed
func checkCourseCompletion(db *pgx.Conn, userID int, courseID int) {
    // Count total stages and completed stages
    var totalStages, completedStages int
    err := db.QueryRow(context.Background(), `
        SELECT 
            COUNT(DISTINCT ps.id) as total_stages,
            COUNT(DISTINCT ss.stage_id) FILTER (WHERE ss.is_correct = true) as completed_stages
        FROM primm_stages ps
        LEFT JOIN stage_submissions ss ON ps.id = ss.stage_id AND ss.user_id = $1
        WHERE ps.course_id = $2
    `, userID, courseID).Scan(&totalStages, &completedStages)

    if err != nil {
        return
    }

    // If all stages completed, mark course as complete and award coins
    if totalStages > 0 && totalStages == completedStages {
        // Get coin reward
        var coinReward int
        db.QueryRow(context.Background(), `
            SELECT coin_reward FROM courses WHERE id = $1
        `, courseID).Scan(&coinReward)

        // Insert course completion record
        db.Exec(context.Background(), `
            INSERT INTO user_course_completion (user_id, course_id, coins_awarded)
            VALUES ($1, $2, $3)
            ON CONFLICT (user_id, course_id) DO NOTHING
        `, userID, courseID, coinReward)

        // Award coins to user
        db.Exec(context.Background(), `
            UPDATE users SET coins = coins + $1 WHERE id = $2
        `, coinReward, userID)
    }
}




// ═══════════════════════════════════════════════════════════
// HELPER FUNCTIONS
// ═══════════════════════════════════════════════════════════

// checkCourseOwnership mengecek apakah course milik teacher
func checkCourseOwnership(db *pgx.Conn, courseID int, teacherID int) error {
    var lessonOwnerID int
    err := db.QueryRow(context.Background(), `
        SELECT l.teacher_id
        FROM courses c
        JOIN lessons l ON c.lesson_id = l.id
        WHERE c.id = $1`, courseID).Scan(&lessonOwnerID)

    if err != nil {
        if err == pgx.ErrNoRows {
            return errors.New("course tidak ditemukan")
        }
        return errors.New("gagal cek course: " + err.Error())
    }

    if lessonOwnerID != teacherID {
        return errors.New("anda tidak memiliki akses untuk membuat stage di course ini")
    }

    return nil
}






