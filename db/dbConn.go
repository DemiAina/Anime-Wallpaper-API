package db

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var Conn *pgx.Conn

func InitConnection() error {
    godotenv.Load(".env")
    DATABASE_URL := os.Getenv("DATABASE_URL")
	connConfig, err := pgx.ParseConfig(DATABASE_URL)
	if err != nil {
		return err
	}

	Conn, err = pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return err
	}

	return nil
}

func CloseConnection() {
	if Conn != nil {
		Conn.Close(context.Background())
	}
}

func AddImage(name string, filePath string, hash string, uploadedAt string) error {
	id := getCount()
	id++
	query := "INSERT INTO images (id ,name, file_path, hash, uploaded_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := Conn.Exec(context.Background(), query, id, name, filePath, hash, uploadedAt)
	if err != nil {
		return err
	}
	return nil
}

func getCount() int {
	var count int
	err := Conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM images").Scan(&count)
	if err != nil {
		fmt.Printf("Failed to get count error : %v", err)
		return -1
	}

	return count
}



func HashExists(hash string) bool {
    query := "SELECT COUNT(*) FROM images WHERE hash = $1"
    var count int
    var exists bool = false
    err := Conn.QueryRow(context.Background(), query, hash).Scan(&count)
    if err != nil {
        fmt.Println("Error checking hash existence:", err)
        return exists
    }

    if count > 0 {
        exists = true
    }

    fmt.Printf("Hash '%s' exists: %v\n", hash, exists)
    return exists
}

