package handlers

import (
    "database/sql"
    "log"
    "net/http"
    "strings"

    "golang.org/x/net/html"
)

// CrawlURL crawls a given URL, extracts metadata, and stores results in the database
func CrawlURL(db *sql.DB, urlID int, url string) error {
    // Fetch the URL
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("Failed to fetch URL: %v", err)
        return err
    }
    defer resp.Body.Close()

    // Parse the HTML
    doc, err := html.Parse(resp.Body)
    if err != nil {
        log.Printf("Failed to parse HTML: %v", err)
        return err
    }

    // Extract metadata
    title := extractTitle(doc)
    htmlVersion := "HTML5" // Simplified assumption
    internalLinks, externalLinks := extractLinks(doc, url)
    hasLoginForm := detectLoginForm(doc)

    // Insert crawl results into the database
    _, err = db.Exec(`
        INSERT INTO crawl_results (url_id, html_version, title, internal_links, external_links, has_login_form)
        VALUES (?, ?, ?, ?, ?, ?)`,
        urlID, htmlVersion, title, len(internalLinks), len(externalLinks), hasLoginForm,
    )
    if err != nil {
        log.Printf("Failed to insert crawl results: %v", err)
        return err
    }

    // Insert links into the database
    for _, link := range append(internalLinks, externalLinks...) {
        _, err = db.Exec(`
            INSERT INTO links (url_id, href, type)
            VALUES (?, ?, ?)`,
            urlID, link, "internal",
        )
        if err != nil {
            log.Printf("Failed to insert link: %v", err)
            return err
        }
    }

    return nil
}

// Helper functions
func extractTitle(doc *html.Node) string {
    var title string
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
            title = n.FirstChild.Data
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
    return title
}

func extractLinks(doc *html.Node, baseURL string) ([]string, []string) {
    var internalLinks, externalLinks []string
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    href := attr.Val
                    if strings.HasPrefix(href, "/") || strings.HasPrefix(href, baseURL) {
                        internalLinks = append(internalLinks, href)
                    } else {
                        externalLinks = append(externalLinks, href)
                    }
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
    return internalLinks, externalLinks
}

func detectLoginForm(doc *html.Node) bool {
    var hasLoginForm bool
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "input" {
            for _, attr := range n.Attr {
                if attr.Key == "type" && attr.Val == "password" {
                    hasLoginForm = true
                    return
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
    return hasLoginForm
}