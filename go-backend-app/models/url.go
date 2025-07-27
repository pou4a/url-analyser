// models/models.go

package models

import (
    "database/sql"
    "encoding/json"
)

// URLRequest represents the structure of the incoming JSON payload for adding a URL
type URLRequest struct {
    URL string `json:"url" binding:"required"`
}

// URL represents a record in the `urls` table
type URL struct {
    ID        int    `json:"id"`
    URL       string `json:"url"`
    Status    string `json:"status"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

// Custom wrapper for sql.NullInt64 to produce clean JSON

type NullInt64JSON struct {
    sql.NullInt64
}

func (n NullInt64JSON) MarshalJSON() ([]byte, error) {
    if n.Valid {
        return json.Marshal(n.Int64)
    }
    return json.Marshal(nil)
}

// Custom wrapper for sql.NullString to produce clean JSON

type NullStringJSON struct {
    sql.NullString
}

func (n NullStringJSON) MarshalJSON() ([]byte, error) {
    if n.Valid {
        return json.Marshal(n.String)
    }
    return json.Marshal(nil)
}

// CrawlResult represents a record in the `crawl_results` table
type CrawlResult struct {
    ID                int             `json:"id"`
    URLID             int             `json:"url_id"`
    HTMLVersion       string          `json:"html_version"`
    Title             string          `json:"title"`
    H1Count           NullInt64JSON   `json:"h1_count"`
    H2Count           NullInt64JSON   `json:"h2_count"`
    H3Count           NullInt64JSON   `json:"h3_count"`
    InternalLinks     NullInt64JSON   `json:"internal_links"`
    ExternalLinks     NullInt64JSON   `json:"external_links"`
    InaccessibleLinks NullInt64JSON   `json:"inaccessible_links"`
    HasLoginForm      bool            `json:"has_login_form"`
    AnalyzedAt        NullStringJSON  `json:"analyzed_at"`
}