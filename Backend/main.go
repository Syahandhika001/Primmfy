package main

import (
    "log"
    "os"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "primmfy_db/handlers"
    "primmfy_db/middleware"
)

func main() {
    // 1. Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file:", err)
    }

    // 2. Connect to database
    InitDB()
    defer CloseDB()

    // 3. Initialize Gin router
    router := gin.Default()

    // 4. Setup CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // 5. Initialize handlers
    authHandler := handlers.NewAuthHandler(DB)
    lessonHandler := handlers.NewLessonHandler(DB)
    courseHandler := handlers.NewCourseHandler(DB)
    stageHandler := handlers.NewPRIMMStageHandler(DB)
    progressHandler := handlers.NewProgressHandler(DB)

    // 6. Setup routes
    api := router.Group("/api")
    {
        // ═══════════════════════════════════════════════════
        // AUTH ROUTES (Public)
        // ═══════════════════════════════════════════════════
        api.POST("/register", authHandler.Register)
        api.POST("/login", authHandler.Login)

        // ═══════════════════════════════════════════════════
        // PUBLIC LESSON ROUTES (No authentication required)
        // ═══════════════════════════════════════════════════
        api.GET("/lessons", lessonHandler.GetAllLessons)
        api.GET("/lessons/:id", lessonHandler.GetLessonByID)
        api.GET("/lessons/:id/courses", lessonHandler.GetCoursesInLesson)

        // ═══════════════════════════════════════════════════
        // PROTECTED ROUTES (Require Authentication)
        // ═══════════════════════════════════════════════════
        protected := api.Group("")
        protected.Use(middleware.AuthMiddleware())
        {
            // Profile endpoint
            protected.GET("/profile", func(c *gin.Context) {
                userID, _ := c.Get("user_id")
                userRole, _ := c.Get("user_role")
                c.JSON(200, gin.H{
                    "message": "Profile berhasil diakses!",
                    "user_id": userID,
                    "role":    userRole,
                })
            })

            // ═══════════════════════════════════════════════════
            // LESSON ROUTES (All authenticated users)
            // ═══════════════════════════════════════════════════
            protected.POST("/lessons/:id/enroll", lessonHandler.EnrollLesson)
            protected.GET("/my-lessons", lessonHandler.GetMyEnrolledLessons)

            // ═══════════════════════════════════════════════════
            // COURSE ROUTES (All authenticated users)
            // ═══════════════════════════════════════════════════
            protected.GET("/courses/:id", courseHandler.GetCourseByID)

            // ═══════════════════════════════════════════════════
            // PRIMM STAGE ROUTES (All authenticated users - view only)
            // ═══════════════════════════════════════════════════
            protected.GET("/stages/:id", stageHandler.GetStageByID)

            // ═══════════════════════════════════════════════════
            // TEACHER ONLY ROUTES
            // ═══════════════════════════════════════════════════
            teacher := protected.Group("")
            teacher.Use(middleware.RequireRole("teacher", "admin"))
            {
                // Lesson Management
                teacher.POST("/lessons", lessonHandler.CreateLesson)
                teacher.GET("/lessons/my", lessonHandler.GetMyLessons)
                teacher.PUT("/lessons/:id", lessonHandler.UpdateLesson)
                teacher.DELETE("/lessons/:id", lessonHandler.DeleteLesson)

                // Course Management
                teacher.POST("/courses", courseHandler.CreateCourse)
                teacher.PUT("/courses/:id", courseHandler.UpdateCourse)
                teacher.DELETE("/courses/:id", courseHandler.DeleteCourse)

                // PRIMM Stage Management
                teacher.POST("/stages/predict", stageHandler.CreatePredictStage)
                teacher.POST("/stages/run", stageHandler.CreateRunStage)
                teacher.POST("/stages/investigate", stageHandler.CreateInvestigateStage)
                teacher.POST("/stages/modify", stageHandler.CreateModifyStage)
                teacher.POST("/stages/make", stageHandler.CreateMakeStage)
                teacher.DELETE("/stages/:id", stageHandler.DeleteStage)

                teacher.GET("/teacher-dashboard", func(c *gin.Context) {
                    c.JSON(200, gin.H{"message": "Welcome to teacher dashboard!"})
                })
            }

            // ═══════════════════════════════════════════════════
            // STUDENT ROUTES (Progress & Submissions)
            // ═══════════════════════════════════════════════════
            student := protected.Group("")
            student.Use(middleware.RequireRole("student"))
            {
                // Submit answers to PRIMM stages
                student.POST("/stages/:id/submit-predict", progressHandler.SubmitPredictStage)
                student.POST("/stages/:id/submit-run", progressHandler.SubmitRunStage)
                student.POST("/stages/:id/submit-investigate", progressHandler.SubmitInvestigateStage)
                student.POST("/stages/:id/submit-modify", progressHandler.SubmitModifyStage)
                student.POST("/stages/:id/submit-make", progressHandler.SubmitMakeStage)
                student.POST("/stages/:id/submit", stageHandler.SubmitStage)


                // View progress
                student.GET("/stages/:id/my-completion", progressHandler.GetStageCompletion)
                student.GET("/courses/:id/my-progress", progressHandler.GetCourseProgress)
                student.GET("/my-progress/:lesson_id", progressHandler.GetMyProgress)
                student.GET("/courses/:id/stages", courseHandler.GetStagesInCourse)
            }

            // ═══════════════════════════════════════════════════
            // ADMIN ONLY ROUTES
            // ═══════════════════════════════════════════════════
            admin := protected.Group("")
            admin.Use(middleware.RequireRole("admin"))
            {
                admin.GET("/admin-dashboard", func(c *gin.Context) {
                    c.JSON(200, gin.H{"message": "Welcome to admin dashboard!"})
                })
            }
        }
    }

    // 7. Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("🚀 Server running on http://localhost:%s\n", port)
    log.Printf("\n📚 API DOCUMENTATION - PRIMMFY Platform\n")
    log.Printf("═══════════════════════════════════════════════════════════\n")
    log.Printf("🔓 PUBLIC ENDPOINTS:\n")
    log.Printf("  POST   /api/register                          - Register user\n")
    log.Printf("  POST   /api/login                             - Login user\n")
    log.Printf("  GET    /api/lessons                           - Get all lessons\n")
    log.Printf("  GET    /api/lessons/:id                       - Get lesson detail\n")
    log.Printf("  GET    /api/lessons/:id/courses               - Get courses in lesson\n")
    log.Printf("\n🔒 AUTHENTICATED ENDPOINTS (All users):\n")
    log.Printf("  GET    /api/profile                           - Get profile\n")
    log.Printf("  GET    /api/courses/:id                       - Get course detail with stages\n")
    log.Printf("  GET    /api/stages/:id                        - Get stage detail\n")
    log.Printf("  POST   /api/lessons/:id/enroll                - Enroll to lesson\n")
    log.Printf("  GET    /api/my-lessons                        - Get enrolled lessons\n")
    log.Printf("\n👨‍🏫 TEACHER ONLY:\n")
    log.Printf("  ┌─ Lesson Management\n")
    log.Printf("  │  POST   /api/lessons                        - Create lesson\n")
    log.Printf("  │  GET    /api/lessons/my                     - Get my lessons\n")
    log.Printf("  │  PUT    /api/lessons/:id                    - Update lesson\n")
    log.Printf("  │  DELETE /api/lessons/:id                    - Delete lesson\n")
    log.Printf("  ├─ Course Management\n")
    log.Printf("  │  POST   /api/courses                        - Create course\n")
    log.Printf("  │  PUT    /api/courses/:id                    - Update course\n")
    log.Printf("  │  DELETE /api/courses/:id                    - Delete course\n")
    log.Printf("  └─ PRIMM Stage Management\n")
    log.Printf("     POST   /api/stages/predict                 - Create PREDICT stage\n")
    log.Printf("     POST   /api/stages/run                     - Create RUN stage\n")
    log.Printf("     POST   /api/stages/investigate             - Create INVESTIGATE stage\n")
    log.Printf("     POST   /api/stages/modify                  - Create MODIFY stage\n")
    log.Printf("     POST   /api/stages/make                    - Create MAKE stage\n")
    log.Printf("     DELETE /api/stages/:id                     - Delete stage\n")
    log.Printf("\n👨‍🎓 STUDENT ONLY:\n")
    log.Printf("  ┌─ Submit Answers\n")
    log.Printf("  │  POST   /api/stages/:id/submit-predict     - Submit PREDICT answer\n")
    log.Printf("  │  POST   /api/stages/:id/submit-run         - Submit RUN code\n")
    log.Printf("  │  POST   /api/stages/:id/submit-investigate - Submit INVESTIGATE reflection\n")
    log.Printf("  │  POST   /api/stages/:id/submit-modify      - Submit MODIFY code\n")
    log.Printf("  │  POST   /api/stages/:id/submit-make        - Submit MAKE code\n")
    log.Printf("  └─ View Progress\n")
    log.Printf("     GET    /api/stages/:id/my-completion       - Get stage completion\n")
    log.Printf("     GET    /api/courses/:id/my-progress        - Get course progress\n")
    log.Printf("     GET    /api/my-progress/:lesson_id         - Get lesson progress\n")
    log.Printf("═══════════════════════════════════════════════════════════\n")
    log.Printf("\n🎯 PRIMM Methodology Flow:\n")
    log.Printf("  Lesson (Big Topic)\n")
    log.Printf("    └─ Course (Sub-topic)\n")
    log.Printf("         ├─ 1. PREDICT Stage (pilihan ganda)\n")
    log.Printf("         ├─ 2. RUN Stage (jalankan code)\n")
    log.Printf("         ├─ 3. INVESTIGATE Stage (refleksi + video)\n")
    log.Printf("         ├─ 4. MODIFY Stage (modifikasi code)\n")
    log.Printf("         └─ 5. MAKE Stage (buat code dari scratch)\n")
    log.Printf("═══════════════════════════════════════════════════════════\n")

    router.Run(":" + port)
}