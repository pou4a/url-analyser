package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"
    _ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
    var db *sql.DB
    var err error
    for i := 0; i < 15; i++ {
        db, err = sql.Open("mysql", dsn)
        if err == nil && db.Ping() == nil {
            return db
        }
        log.Printf("Waiting for DB... (%d/15), error: %v", i+1, err)
        time.Sleep(2 * time.Second)
    }
    log.Fatalf("Error connecting to DB: %v", err)
    return nil
}