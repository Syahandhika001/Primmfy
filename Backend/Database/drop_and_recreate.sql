-- ═══════════════════════════════════════════════════════════
-- DROP EXISTING TABLES (CASCADE untuk hapus dependencies)
-- ═══════════════════════════════════════════════════════════

DROP TABLE IF EXISTS user_stage_completions CASCADE;
DROP TABLE IF EXISTS user_course_completions CASCADE;
DROP TABLE IF EXISTS user_lessons CASCADE;
DROP TABLE IF EXISTS primm_stages CASCADE;
DROP TABLE IF EXISTS courses CASCADE;
DROP TABLE IF EXISTS lessons CASCADE;

-- Drop function juga
DROP FUNCTION IF EXISTS update_updated_at_column() CASCADE;

-- ═══════════════════════════════════════════════════════════
-- LEVEL 1: LESSONS (Big Topic - Collection of Courses)
-- ═══════════════════════════════════════════════════════════
CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    teacher_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50) NOT NULL, -- 'python', 'javascript', 'html', 'c', dll
    difficulty VARCHAR(20) NOT NULL CHECK (difficulty IN ('beginner', 'intermediate', 'advanced')),
    thumbnail_url TEXT, -- Optional: gambar lesson
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_lessons_teacher ON lessons(teacher_id);
CREATE INDEX idx_lessons_category ON lessons(category);
CREATE INDEX idx_lessons_active ON lessons(is_active);

-- ═══════════════════════════════════════════════════════════
-- LEVEL 2: COURSES (Sub-topic dalam Lesson, ada 5 PRIMM stages)
-- ═══════════════════════════════════════════════════════════
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    order_index INTEGER NOT NULL, -- Urutan course dalam lesson (1, 2, 3, ...)
    coin_reward INTEGER NOT NULL DEFAULT 100, -- Coins setelah complete course
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    UNIQUE(lesson_id, order_index) -- Tidak boleh ada 2 course dengan order sama dalam 1 lesson
);

CREATE INDEX idx_courses_lesson ON courses(lesson_id);
CREATE INDEX idx_courses_order ON courses(lesson_id, order_index);

-- ═══════════════════════════════════════════════════════════
-- LEVEL 3: PRIMM STAGES (5 stages per course)
-- ═══════════════════════════════════════════════════════════
CREATE TABLE primm_stages (
    id SERIAL PRIMARY KEY,
    course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    stage_type VARCHAR(20) NOT NULL CHECK (stage_type IN ('predict', 'run', 'investigate', 'modify', 'make')),
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    order_index INTEGER NOT NULL CHECK (order_index BETWEEN 1 AND 5), -- 1=Predict, 2=Run, 3=Investigate, 4=Modify, 5=Make
    
    -- PREDICT Stage Fields
    code_snippet TEXT, -- Code yang ditampilkan
    predict_options JSONB, -- Multiple choice options: {"A": "output1", "B": "output2", ...}
    correct_answer VARCHAR(10), -- Correct option key: "A", "B", "C", "D"
    
    -- RUN Stage Fields
    run_code_template TEXT, -- Code template yang harus ditulis ulang
    
    -- INVESTIGATE Stage Fields
    reflection_prompt TEXT, -- Pertanyaan refleksi untuk student
    video_embed_url TEXT, -- URL video penjelasan (YouTube embed, Vimeo, dll)
    explanation_text TEXT, -- Penjelasan line-by-line dalam text/markdown
    
    -- MODIFY Stage Fields
    modify_challenge TEXT, -- Deskripsi tantangan modifikasi
    modify_code_template TEXT, -- Code awal untuk dimodifikasi
    modify_expected_output TEXT, -- Expected output setelah modifikasi
    modify_test_cases JSONB, -- Test cases untuk validasi: [{"input": "...", "output": "..."}, ...]
    
    -- MAKE Stage Fields
    make_challenge TEXT, -- Deskripsi tantangan membuat dari scratch
    make_hints TEXT, -- Hints untuk membantu student
    make_expected_output TEXT, -- Expected output
    make_test_cases JSONB, -- Test cases untuk validasi
    
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    UNIQUE(course_id, order_index) -- Setiap course harus punya 5 stages (order 1-5)
);

CREATE INDEX idx_stages_course ON primm_stages(course_id);
CREATE INDEX idx_stages_type ON primm_stages(stage_type);

-- ═══════════════════════════════════════════════════════════
-- STUDENT PROGRESS: Enrollment ke Lesson
-- ═══════════════════════════════════════════════════════════
CREATE TABLE user_lessons (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    lesson_id INTEGER NOT NULL REFERENCES lessons(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMP DEFAULT NOW(),
    
    UNIQUE(user_id, lesson_id) -- Student tidak boleh enroll 2x ke lesson yang sama
);

CREATE INDEX idx_user_lessons_user ON user_lessons(user_id);
CREATE INDEX idx_user_lessons_lesson ON user_lessons(lesson_id);

-- ═══════════════════════════════════════════════════════════
-- STUDENT PROGRESS: Completion per Course
-- ═══════════════════════════════════════════════════════════
CREATE TABLE user_course_completions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    is_completed BOOLEAN DEFAULT false,
    completed_at TIMESTAMP,
    coins_earned INTEGER DEFAULT 0, -- Coins yang didapat dari course ini
    
    UNIQUE(user_id, course_id)
);

CREATE INDEX idx_course_completions_user ON user_course_completions(user_id);
CREATE INDEX idx_course_completions_course ON user_course_completions(course_id);

-- ═══════════════════════════════════════════════════════════
-- STUDENT PROGRESS: Completion per PRIMM Stage
-- ═══════════════════════════════════════════════════════════
CREATE TABLE user_stage_completions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    stage_id INTEGER NOT NULL REFERENCES primm_stages(id) ON DELETE CASCADE,
    
    -- PREDICT Stage Data
    predict_selected_answer VARCHAR(10), -- Option yang dipilih: "A", "B", "C", "D"
    predict_is_correct BOOLEAN,
    
    -- RUN Stage Data
    run_submitted_code TEXT, -- Code yang ditulis student
    run_output TEXT, -- Output hasil run code
    
    -- INVESTIGATE Stage Data
    investigate_reflection TEXT, -- Jawaban refleksi student
    investigate_completed BOOLEAN DEFAULT false,
    
    -- MODIFY Stage Data
    modify_submitted_code TEXT, -- Code hasil modifikasi
    modify_output TEXT, -- Output hasil run
    modify_is_correct BOOLEAN, -- Apakah passed all test cases
    modify_attempts INTEGER DEFAULT 0, -- Jumlah percobaan
    
    -- MAKE Stage Data
    make_submitted_code TEXT, -- Code yang dibuat student
    make_output TEXT, -- Output hasil run
    make_is_correct BOOLEAN, -- Apakah passed all test cases
    make_attempts INTEGER DEFAULT 0,
    
    is_completed BOOLEAN DEFAULT false,
    completed_at TIMESTAMP,
    
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    UNIQUE(user_id, stage_id)
);

CREATE INDEX idx_stage_completions_user ON user_stage_completions(user_id);
CREATE INDEX idx_stage_completions_stage ON user_stage_completions(stage_id);

-- ═══════════════════════════════════════════════════════════
-- TRIGGERS: Auto-update updated_at
-- ═══════════════════════════════════════════════════════════
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_lessons_updated_at BEFORE UPDATE ON lessons
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_courses_updated_at BEFORE UPDATE ON courses
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_primm_stages_updated_at BEFORE UPDATE ON primm_stages
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_stage_completions_updated_at BEFORE UPDATE ON user_stage_completions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();