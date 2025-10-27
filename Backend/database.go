package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv" // 👈 Tambah ini untuk load .env
)

var DB *pgx.Conn

// InitDB menginisialisasi koneksi ke database
func InitDB() {
    // 1. Load environment variables dari file .env
    err := godotenv.Load()
    if err != nil {
        fmt.Println("⚠️  Warning: File .env tidak ditemukan, menggunakan environment variables sistem")
    }

    // 2. Ambil database URL dari environment variable
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        // Fallback ke default jika tidak ada di .env
        dbURL = "postgres://postgres:dhika001@localhost:5432/primmfy_db?sslmode=disable"
        fmt.Println("⚠️  Warning: DATABASE_URL tidak ditemukan di .env, menggunakan default")
    }

    // 3. Connect ke database
    DB, err = pgx.Connect(context.Background(), dbURL)
    if err != nil {
        fmt.Fprintf(os.Stderr, "❌ Gagal terhubung ke database: %v\n", err)
        os.Exit(1)
    }

    // 4. Test koneksi dengan ping
    err = DB.Ping(context.Background())
    if err != nil {
        fmt.Fprintf(os.Stderr, "❌ Gagal ping database: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("✅ Berhasil terhubung ke database!")
}

// CloseDB menutup koneksi database
func CloseDB() {
    if DB != nil {
        DB.Close(context.Background())
        fmt.Println("🔌 Koneksi database ditutup")
    }
}


