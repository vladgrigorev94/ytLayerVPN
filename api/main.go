package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

func getDirectURL(videoURL string) (string, error) {
	cmd := exec.Command("yt-dlp", "-f", "best[ext=mp4]", "-g", videoURL)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	videoURL := r.URL.Query().Get("url")
	if videoURL == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	directURL, err := getDirectURL(videoURL)
	if err != nil {
		http.Error(w, "Failed to get direct video URL", http.StatusInternalServerError)
		return
	}

	proxyURL, _ := url.Parse("http://localhost:3067")

	client := &http.Client{
		Timeout: time.Minute,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	resp, err := client.Get(directURL)
	if err != nil {
		http.Error(w, "Failed to stream video", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))

	_, _ = io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/stream", streamHandler)

	fmt.Println("🔗 Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
