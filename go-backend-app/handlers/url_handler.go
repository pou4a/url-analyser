package handlers

import (
    "database/sql"
    "net/http"
    "net/url"       
    "strconv"     
    "url-analyser-backend/models"
	"log"

    "github.com/gin-gonic/gin"
)

func AddURL(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req models.URLRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Validate URL format
        _, err := url.ParseRequestURI(req.URL)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
            return
        }

        // Insert URL into the database
        result, err := db.Exec("INSERT INTO urls (url) VALUES (?)", req.URL)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed"})
            return
        }

        id, err := result.LastInsertId()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inserted ID"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "URL submitted",
            "id":      id,
        })
    }
}

func ListURLs(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get pagination parameters
        page := c.DefaultQuery("page", "1")
        limit := c.DefaultQuery("limit", "10")

        // Convert to integers
        pageInt, err := strconv.Atoi(page)
        if err != nil || pageInt < 1 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
            return
        }
        limitInt, err := strconv.Atoi(limit)
        if err != nil || limitInt < 1 {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
            return
        }

        offset := (pageInt - 1) * limitInt

        // Query paginated URLs
        rows, err := db.Query("SELECT id, url, status, created_at, updated_at FROM urls LIMIT ? OFFSET ?", limitInt, offset)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "DB query failed"})
            return
        }
        defer rows.Close()

        var urls []models.URL
        for rows.Next() {
            var url models.URL
            if err := rows.Scan(&url.ID, &url.URL, &url.Status, &url.CreatedAt, &url.UpdatedAt); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan failed"})
                return
            }
            urls = append(urls, url)
        }

        c.JSON(http.StatusOK, urls)
    }
}

func CrawlURLHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get URL ID from the request
        var req struct {
            URLID int `json:"url_id" binding:"required"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Fetch the URL from the database
        var url string
        err := db.QueryRow("SELECT url FROM urls WHERE id = ?", req.URLID).Scan(&url)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
            return
        }

        // Crawl the URL
        if err := CrawlURL(db, req.URLID, url); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to crawl URL"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Crawling completed"})
    }
}

func GetCrawlResults(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get URL ID from query parameters
        urlID := c.Query("url_id")
        if urlID == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "url_id is required"})
            return
        }

        // Convert urlID to integer
        urlIDInt, err := strconv.Atoi(urlID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url_id parameter"})
            return
        }

        // Query crawl results from the database
        var result models.CrawlResult
        err = db.QueryRow(`
            SELECT id, url_id, html_version, title, h1_count, h2_count, h3_count, internal_links, external_links, inaccessible_links, has_login_form, analyzed_at
            FROM crawl_results WHERE url_id = ?`, urlIDInt).Scan(
            &result.ID, &result.URLID, &result.HTMLVersion, &result.Title, &result.H1Count, &result.H2Count, &result.H3Count,
            &result.InternalLinks, &result.ExternalLinks, &result.InaccessibleLinks, &result.HasLoginForm, &result.AnalyzedAt,
        )
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "No crawl results found for the given URL ID"})
            return
        } else if err != nil {
            // Log the error for debugging
            log.Printf("Error fetching crawl results for url_id=%d: %v", urlIDInt, err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch crawl results"})
            return
        }

        c.JSON(http.StatusOK, result)
    }
}
