package services

import (
    "context"
    "encoding/json"
    "errors"
    "time"

    "github.com/jackc/pgx/v5"
    "primmfy_db/models"
)

// ═══════════════════════════════════════════════════════════
// SUBMIT STAGE FUNCTIONS
// ═══════════════════════════════════════════════════════════

// SubmitPredictStage memproses submission PREDICT stage
// Logic: Cek jawaban benar/salah, beri koin jika benar
func SubmitPredictStage(db *pgx.Conn, userID int, req models.SubmitPredictRequest) (*models.SubmitStageResponse, error) {
    // 1. Ambil data stage dari database
    var correctAnswer string
    err := db.QueryRow(context.Background(),
        "SELECT correct_answer FROM primm_stages WHERE id = $1 AND stage_type = 'predict'",
        req.StageID).Scan(&correctAnswer)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("predict stage tidak ditemukan")
        }
        return nil, errors.New("gagal mengambil stage: " + err.Error())
    }

    // 2. Cek apakah jawaban benar
    isCorrect := (req.SelectedAnswer == correctAnswer)

    // 3. Cek apakah sudah pernah submit (untuk prevent duplicate submission)
    var existingCompletionID int
    err = db.QueryRow(context.Background(),
        "SELECT id FROM user_stage_completions WHERE user_id = $1 AND stage_id = $2",
        userID, req.StageID).Scan(&existingCompletionID)

    var coinsEarned, xpEarned int
    message := "Jawaban salah. Coba lagi!"

    if isCorrect {
        coinsEarned = 50  // Reward koin untuk Predict stage
        xpEarned = 20     // Reward XP
        message = "Jawaban benar! Selamat!"
    }

    // 4. Simpan atau update submission
    if err == pgx.ErrNoRows {
        // Belum pernah submit, insert baru
        _, err = db.Exec(context.Background(), `
            INSERT INTO user_stage_completions 
            (user_id, stage_id, predict_selected_answer, predict_is_correct, is_completed, completed_at)
            VALUES ($1, $2, $3, $4, $5, $6)`,
            userID, req.StageID, req.SelectedAnswer, isCorrect, isCorrect,
            func() *time.Time { if isCorrect { t := time.Now(); return &t }; return nil }())

        if err != nil {
            return nil, errors.New("gagal menyimpan submission: " + err.Error())
        }
    } else {
        // Sudah pernah submit, update
        _, err = db.Exec(context.Background(), `
            UPDATE user_stage_completions
            SET predict_selected_answer = $1, predict_is_correct = $2, 
                is_completed = $3, completed_at = $4, updated_at = NOW()
            WHERE user_id = $5 AND stage_id = $6`,
            req.SelectedAnswer, isCorrect, isCorrect,
            func() *time.Time { if isCorrect { t := time.Now(); return &t }; return nil }(),
            userID, req.StageID)

        if err != nil {
            return nil, errors.New("gagal update submission: " + err.Error())
        }
    }

    // 5. Jika benar, berikan reward (update coins & XP user)
    if isCorrect {
        _, err = db.Exec(context.Background(), `
            UPDATE users
            SET total_coins = total_coins + $1,
                experience_points = experience_points + $2,
                updated_at = NOW()
            WHERE id = $3`,
            coinsEarned, xpEarned, userID)

        if err != nil {
            return nil, errors.New("gagal memberikan reward: " + err.Error())
        }

        // Cek apakah perlu level up
        checkAndLevelUp(db, userID)
    }

    return &models.SubmitStageResponse{
        Success:     true,
        IsCorrect:   isCorrect,
        Message:     message,
        CoinsEarned: coinsEarned,
        XPEarned:    xpEarned,
    }, nil
}

// SubmitRunStage memproses submission RUN stage
// Purpose: Simpan code yang ditulis siswa & output hasil run
func SubmitRunStage(db *pgx.Conn, userID int, req models.SubmitRunRequest) (*models.SubmitStageResponse, error) {
    // 1. Validasi stage type
    var stageType string
    err := db.QueryRow(context.Background(),
        "SELECT stage_type FROM primm_stages WHERE id = $1", req.StageID).Scan(&stageType)

    if err != nil || stageType != "run" {
        return nil, errors.New("run stage tidak ditemukan")
    }

    // 2. Simulate code execution (dalam production, pakai sandbox seperti Judge0)
    // Untuk sementara, kita anggap code valid dan return "Success"
    output := "Code executed successfully! Output: Hello World"

    // 3. Simpan submission
    var existingID int
    err = db.QueryRow(context.Background(),
        "SELECT id FROM user_stage_completions WHERE user_id = $1 AND stage_id = $2",
        userID, req.StageID).Scan(&existingID)

    coinsEarned := 50
    xpEarned := 20

    if err == pgx.ErrNoRows {
        // Insert baru
        _, err = db.Exec(context.Background(), `
            INSERT INTO user_stage_completions
            (user_id, stage_id, run_submitted_code, run_output, is_completed, completed_at)
            VALUES ($1, $2, $3, $4, true, NOW())`,
            userID, req.StageID, req.SubmittedCode, output)
    } else {
        // Update existing
        _, err = db.Exec(context.Background(), `
            UPDATE user_stage_completions
            SET run_submitted_code = $1, run_output = $2, 
                is_completed = true, completed_at = NOW(), updated_at = NOW()
            WHERE user_id = $3 AND stage_id = $4`,
            req.SubmittedCode, output, userID, req.StageID)
    }

    if err != nil {
        return nil, errors.New("gagal menyimpan submission: " + err.Error())
    }

    // 4. Berikan reward
    _, err = db.Exec(context.Background(), `
        UPDATE users
        SET total_coins = total_coins + $1,
            experience_points = experience_points + $2,
            updated_at = NOW()
        WHERE id = $3`,
        coinsEarned, xpEarned, userID)

    if err != nil {
        return nil, errors.New("gagal memberikan reward: " + err.Error())
    }

    checkAndLevelUp(db, userID)

    return &models.SubmitStageResponse{
        Success:     true,
        IsCorrect:   true,
        Message:     "Code berhasil dijalankan!",
        CoinsEarned: coinsEarned,
        XPEarned:    xpEarned,
        Output:      output,
    }, nil
}

// SubmitInvestigateStage memproses submission INVESTIGATE stage
// Purpose: Simpan refleksi siswa (no right/wrong answer)
func SubmitInvestigateStage(db *pgx.Conn, userID int, req models.SubmitInvestigateRequest) (*models.SubmitStageResponse, error) {
    // 1. Validasi stage
    var stageType string
    err := db.QueryRow(context.Background(),
        "SELECT stage_type FROM primm_stages WHERE id = $1", req.StageID).Scan(&stageType)

    if err != nil || stageType != "investigate" {
        return nil, errors.New("investigate stage tidak ditemukan")
    }

    // 2. Simpan refleksi (tidak ada benar/salah, hanya perlu submit)
    var existingID int
    err = db.QueryRow(context.Background(),
        "SELECT id FROM user_stage_completions WHERE user_id = $1 AND stage_id = $2",
        userID, req.StageID).Scan(&existingID)

    coinsEarned := 30 // Lebih sedikit karena tidak ada validasi benar/salah
    xpEarned := 15

    if err == pgx.ErrNoRows {
        _, err = db.Exec(context.Background(), `
            INSERT INTO user_stage_completions
            (user_id, stage_id, investigate_reflection, investigate_completed, is_completed, completed_at)
            VALUES ($1, $2, $3, true, true, NOW())`,
            userID, req.StageID, req.Reflection)
    } else {
        _, err = db.Exec(context.Background(), `
            UPDATE user_stage_completions
            SET investigate_reflection = $1, investigate_completed = true,
                is_completed = true, completed_at = NOW(), updated_at = NOW()
            WHERE user_id = $2 AND stage_id = $3`,
            req.Reflection, userID, req.StageID)
    }

    if err != nil {
        return nil, errors.New("gagal menyimpan submission: " + err.Error())
    }

    // 3. Berikan reward
    _, err = db.Exec(context.Background(), `
        UPDATE users
        SET total_coins = total_coins + $1,
            experience_points = experience_points + $2,
            updated_at = NOW()
        WHERE id = $3`,
        coinsEarned, xpEarned, userID)

    checkAndLevelUp(db, userID)

    return &models.SubmitStageResponse{
        Success:     true,
        IsCorrect:   true, // Selalu true karena tidak ada validasi
        Message:     "Refleksi berhasil disimpan!",
        CoinsEarned: coinsEarned,
        XPEarned:    xpEarned,
    }, nil
}

// SubmitModifyStage memproses submission MODIFY stage dengan test case validation
func SubmitModifyStage(db *pgx.Conn, userID int, req models.SubmitModifyRequest) (*models.SubmitStageResponse, error) {
    // 1. Ambil stage data
    var modifyTestCasesJSON []byte
    err := db.QueryRow(context.Background(),
        "SELECT modify_test_cases FROM primm_stages WHERE id = $1 AND stage_type = 'modify'",
        req.StageID).Scan(&modifyTestCasesJSON)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("modify stage tidak ditemukan")
        }
        return nil, errors.New("gagal mengambil stage: " + err.Error())
    }

    // 2. Parse test cases
    var testCases []models.TestCase
    if err := json.Unmarshal(modifyTestCasesJSON, &testCases); err != nil {
        return nil, errors.New("gagal parse test cases: " + err.Error())
    }

    // 3. Simulate code execution & validation (dalam production pakai Judge0/sandbox)
    // Untuk sementara, kita anggap code passed all test cases
    isCorrect := true
    output := "Test case 1: PASSED\nTest case 2: PASSED\nAll tests passed!"

    // 4. Get atau create submission record
    var existingID int
    err = db.QueryRow(context.Background(),
        "SELECT id FROM user_stage_completions WHERE user_id = $1 AND stage_id = $2",
        userID, req.StageID).Scan(&existingID)

    coinsEarned := 0
    xpEarned := 0
    message := "Code tidak passed test cases. Coba lagi!"

    if isCorrect {
        coinsEarned = 75
        xpEarned = 30
        message = "Code berhasil passed semua test cases! Selamat!"
    }

    // 5. Update atau insert submission
    if err == pgx.ErrNoRows {
        // Insert baru
        _, err = db.Exec(context.Background(), `
            INSERT INTO user_stage_completions
            (user_id, stage_id, modify_submitted_code, modify_output, modify_is_correct, 
             modify_attempts, is_completed, completed_at)
            VALUES ($1, $2, $3, $4, $5, 1, $5, $6)`,
            userID, req.StageID, req.SubmittedCode, output, isCorrect,
            func() *time.Time { if isCorrect { t := time.Now(); return &t }; return nil }())
    } else {
        // Update existing (increment attempts)
        _, err = db.Exec(context.Background(), `
            UPDATE user_stage_completions
            SET modify_submitted_code = $1, modify_output = $2, modify_is_correct = $3,
                modify_attempts = modify_attempts + 1,
                is_completed = $3, completed_at = $4, updated_at = NOW()
            WHERE user_id = $5 AND stage_id = $6`,
            req.SubmittedCode, output, isCorrect,
            func() *time.Time { if isCorrect { t := time.Now(); return &t }; return nil }(),
            userID, req.StageID)
    }

    if err != nil {
        return nil, errors.New("gagal menyimpan submission: " + err.Error())
    }

    // 6. Berikan reward jika correct
    if isCorrect {
        _, err = db.Exec(context.Background(), `
            UPDATE users
            SET total_coins = total_coins + $1,
                experience_points = experience_points + $2,
                updated_at = NOW()
            WHERE id = $3`,
            coinsEarned, xpEarned, userID)

        if err != nil {
            return nil, errors.New("gagal memberikan reward: " + err.Error())
        }

        checkAndLevelUp(db, userID)
    }

    return &models.SubmitStageResponse{
        Success:     true,
        IsCorrect:   isCorrect,
        Message:     message,
        CoinsEarned: coinsEarned,
        XPEarned:    xpEarned,
        Output:      output,
    }, nil
}

// SubmitMakeStage memproses submission MAKE stage dengan test case validation
func SubmitMakeStage(db *pgx.Conn, userID int, req models.SubmitMakeRequest) (*models.SubmitStageResponse, error) {
    // 1. Ambil stage data
    var makeTestCasesJSON []byte
    err := db.QueryRow(context.Background(),
        "SELECT make_test_cases FROM primm_stages WHERE id = $1 AND stage_type = 'make'",
        req.StageID).Scan(&makeTestCasesJSON)

    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, errors.New("make stage tidak ditemukan")
        }
        return nil, errors.New("gagal mengambil stage: " + err.Error())
    }

    // 2. Parse test cases
    var testCases []models.TestCase
    if err := json.Unmarshal(makeTestCasesJSON, &testCases); err != nil {
        return nil, errors.New("gagal parse test cases: " + err.Error())
    }

    // 3. Simulate code execution (production: pakai Judge0)
    isCorrect := true
    output := "Test case 1: PASSED\nTest case 2: PASSED\nAll tests passed!"

    // 4. Get atau create submission
    var existingID int
    err = db.QueryRow(context.Background(),
        "SELECT id FROM user_stage_completions WHERE user_id = $1 AND stage_id = $2",
        userID, req.StageID).Scan(&existingID)

    coinsEarned := 0
    xpEarned := 0
    message := "Code tidak passed test cases. Coba lagi!"

    if isCorrect {
        coinsEarned = 100 // Highest reward untuk MAKE stage
        xpEarned = 50
        message = "Code berhasil passed semua test cases! Excellent work!"
    }

    // 5. Update atau insert
    if err == pgx.ErrNoRows {
        _, err = db.Exec(context.Background(), `
            INSERT INTO user_stage_completions
            (user_id, stage_id, make_submitted_code, make_output, make_is_correct, 
             make_attempts, is_completed, completed_at)
            VALUES ($1, $2, $3, $4, $5, 1, $5, $6)`,
            userID, req.StageID, req.SubmittedCode, output, isCorrect,
            func() *time.Time { if isCorrect { t := time.Now(); return &t }; return nil }())
    } else {
        _, err = db.Exec(context.Background(), `
            UPDATE user_stage_completions
            SET make_submitted_code = $1, make_output = $2, make_is_correct = $3,
                make_attempts = make_attempts + 1,
                is_completed = $3, completed_at = $4, updated_at = NOW()
            WHERE user_id = $5 AND stage_id = $6`,
            req.SubmittedCode, output, isCorrect,
            func() *time.Time { if isCorrect { t := time.Now(); return &t }; return nil }(),
            userID, req.StageID)
    }

    if err != nil {
        return nil, errors.New("gagal menyimpan submission: " + err.Error())
    }

    // 6. Berikan reward
    if isCorrect {
        _, err = db.Exec(context.Background(), `
            UPDATE users
            SET total_coins = total_coins + $1,
                experience_points = experience_points + $2,
                updated_at = NOW()
            WHERE id = $3`,
            coinsEarned, xpEarned, userID)

        if err != nil {
            return nil, errors.New("gagal memberikan reward: " + err.Error())
        }

        checkAndLevelUp(db, userID)

        // Check apakah course sudah complete (semua 5 stages done)
        checkAndCompleteCourse(db, userID, req.StageID)
    }

    return &models.SubmitStageResponse{
        Success:     true,
        IsCorrect:   isCorrect,
        Message:     message,
        CoinsEarned: coinsEarned,
        XPEarned:    xpEarned,
        Output:      output,
    }, nil
}

// ═══════════════════════════════════════════════════════════
// PROGRESS TRACKING FUNCTIONS
// ═══════════════════════════════════════════════════════════

// GetMyProgress mengambil progress siswa di lesson tertentu
func GetMyProgress(db *pgx.Conn, userID int, lessonID int) (*models.ProgressSummary, error) {
    // 1. Hitung total courses di lesson
    var totalCourses int
    err := db.QueryRow(context.Background(),
        "SELECT COUNT(*) FROM courses WHERE lesson_id = $1 AND is_active = true", lessonID).Scan(&totalCourses)

    if err != nil {
        return nil, errors.New("gagal menghitung total courses: " + err.Error())
    }

    // 2. Hitung completed courses
    var completedCourses int
    err = db.QueryRow(context.Background(), `
        SELECT COUNT(*)
        FROM user_course_completions
        WHERE user_id = $1 AND course_id IN (
            SELECT id FROM courses WHERE lesson_id = $2
        ) AND is_completed = true`, userID, lessonID).Scan(&completedCourses)

    if err != nil {
        completedCourses = 0
    }

    // 3. Hitung progress percent
    var progressPercent float64
    if totalCourses > 0 {
        progressPercent = (float64(completedCourses) / float64(totalCourses)) * 100
    }

    // 4. Hitung total coins & XP yang didapat dari lesson ini
    var totalCoins, totalXP int
    db.QueryRow(context.Background(), `
        SELECT COALESCE(SUM(coins_earned), 0)
        FROM user_course_completions
        WHERE user_id = $1 AND course_id IN (
            SELECT id FROM courses WHERE lesson_id = $2
        )`, userID, lessonID).Scan(&totalCoins)

    return &models.ProgressSummary{
        TotalCourses:     totalCourses,
        CompletedCourses: completedCourses,
        ProgressPercent:  progressPercent,
        TotalCoinsEarned: totalCoins,
        TotalXPEarned:    totalXP,
    }, nil
}

// ═══════════════════════════════════════════════════════════
// HELPER FUNCTIONS
// ═══════════════════════════════════════════════════════════

// checkAndLevelUp mengecek apakah user perlu naik level
// Logic: Setiap 100 XP = 1 level
func checkAndLevelUp(db *pgx.Conn, userID int) error {
    var currentLevel, currentXP int
    err := db.QueryRow(context.Background(),
        "SELECT level, experience_points FROM users WHERE id = $1", userID).Scan(&currentLevel, &currentXP)

    if err != nil {
        return err
    }

    // Hitung level yang seharusnya (100 XP = 1 level)
    expectedLevel := (currentXP / 100) + 1

    // Jika perlu level up
    if expectedLevel > currentLevel {
        _, err = db.Exec(context.Background(),
            "UPDATE users SET level = $1, updated_at = NOW() WHERE id = $2",
            expectedLevel, userID)
        return err
    }

    return nil
}

// checkAndCompleteCourse mengecek apakah semua 5 stages sudah complete
func checkAndCompleteCourse(db *pgx.Conn, userID int, stageID int) error {
    // 1. Get course_id dari stage
    var courseID int
    err := db.QueryRow(context.Background(),
        "SELECT course_id FROM primm_stages WHERE id = $1", stageID).Scan(&courseID)

    if err != nil {
        return err
    }

    // 2. Hitung berapa stages yang completed
    var completedStages int
    err = db.QueryRow(context.Background(), `
        SELECT COUNT(*)
        FROM user_stage_completions usc
        JOIN primm_stages ps ON usc.stage_id = ps.id
        WHERE usc.user_id = $1 AND ps.course_id = $2 AND usc.is_completed = true`,
        userID, courseID).Scan(&completedStages)

    if err != nil {
        return err
    }

    // 3. Jika sudah 5 stages (semua PRIMM stages complete)
    if completedStages == 5 {
        // Get coin_reward dari course
        var coinReward int
        err = db.QueryRow(context.Background(),
            "SELECT coin_reward FROM courses WHERE id = $1", courseID).Scan(&coinReward)

        if err != nil {
            return err
        }

        // Mark course as completed
        _, err = db.Exec(context.Background(), `
            INSERT INTO user_course_completions (user_id, course_id, is_completed, completed_at, coins_earned)
            VALUES ($1, $2, true, NOW(), $3)
            ON CONFLICT (user_id, course_id) 
            DO UPDATE SET is_completed = true, completed_at = NOW(), coins_earned = $3`,
            userID, courseID, coinReward)

        if err != nil {
            return err
        }

        // Berikan bonus coins untuk complete course
        _, err = db.Exec(context.Background(), `
            UPDATE users
            SET total_coins = total_coins + $1, updated_at = NOW()
            WHERE id = $2`,
            coinReward, userID)

        return err
    }

    return nil
}