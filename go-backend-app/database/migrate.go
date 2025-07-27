package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time" 
)

// EnsureDatabase creates the database if it doesn't exist
func EnsureDatabase() {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)
    var db *sql.DB
    var err error
    for i := 0; i < 15; i++ {
        db, err = sql.Open("mysql", dsn)
        if err == nil && db.Ping() == nil {
            break
        }
        log.Printf("Waiting for MySQL in EnsureDatabase... (%d/15), error: %v", i+1, err)
        time.Sleep(2 * time.Second)
    }
    if db == nil || db.Ping() != nil {
        log.Fatalf("Error connecting to MySQL in EnsureDatabase: %v", err)
    }
    defer db.Close()

    _, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
    if err != nil {
        log.Fatalf("Error creating database: %v", err)
    }
}

// Migrate creates tables if they don't exist
func Migrate(db *sql.DB) {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS urls (
            id INT AUTO_INCREMENT PRIMARY KEY,
            url VARCHAR(255) NOT NULL,
            status VARCHAR(32) DEFAULT 'queued',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );`,
        `CREATE TABLE IF NOT EXISTS crawl_results (
            id INT AUTO_INCREMENT PRIMARY KEY,
            url_id INT NOT NULL,
            html_version VARCHAR(32),
            title VARCHAR(255),
            h1_count INT,
            h2_count INT,
            h3_count INT,
            internal_links INT,
            external_links INT,
            inaccessible_links INT,
            has_login_form BOOLEAN,
            analyzed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE
        );`,
    }

    for _, q := range queries {
        log.Printf("Running migration query: %s", q)
        if _, err := db.Exec(q); err != nil {
            log.Fatalf("Migration failed: %v", err)
        }
    }
    log.Println("Migration completed successfully")
}