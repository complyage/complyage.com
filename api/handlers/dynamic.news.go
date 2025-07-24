package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"sync"
	"time"
)

type RSS struct {
	Channel struct {
		Item []struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

type Article struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	PubDate string `json:"pubDate"`
}

var (
	cachedArticles []Article
	cacheExpiry    time.Time
	cacheMutex     sync.Mutex
)

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// If cache is fresh, return cached
	if time.Now().Before(cacheExpiry) && cachedArticles != nil {
		writeJSON(w, cachedArticles)
		return
	}

	// Fetch fresh feed
	resp, err := http.Get("https://news.google.com/rss/search?q=age+verification+laws")
	if err != nil {
		http.Error(w, "Failed to fetch news feed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var rss RSS
	if err := xml.NewDecoder(resp.Body).Decode(&rss); err != nil {
		http.Error(w, "Failed to parse RSS feed", http.StatusInternalServerError)
		return
	}

	var articles []Article
	for _, item := range rss.Channel.Item {
		articles = append(articles, Article{
			Title:   item.Title,
			Link:    item.Link,
			PubDate: item.PubDate,
		})
	}

	// Update cache
	cachedArticles = articles
	cacheExpiry = time.Now().Add(10 * time.Minute) // Cache TTL

	writeJSON(w, articles)
}

func writeJSON(w http.ResponseWriter, articles []Article) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Articles []Article `json:"articles"`
	}{
		Articles: articles,
	})
}
