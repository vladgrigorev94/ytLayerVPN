package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	videoURL := r.URL.Query().Get("url")
	if videoURL == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	// Вызов yt-dlp с cookies.txt
	cmd := exec.Command("yt-dlp", "--cookies", "cookies.txt", "--no-playlist", "-f", "best[ext=mp4]", "-g", "--force-client", "mweb", videoURL)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("yt-dlp error: %v\nOutput:\n%s", err, string(output))
		http.Error(w, "yt-dlp failed", http.StatusInternalServerError)
		return
	}

	directURL := strings.TrimSpace(string(output))
	directURL = strings.TrimSpace(directURL)
	directURL = strings.ReplaceAll(directURL, "\n", "")
	directURL = strings.ReplaceAll(directURL, "\t", "")
	log.Println("Direct URL:", directURL)

	resp, err := http.Get(directURL)
	if err != nil {
		http.Error(w, "Failed to stream video", http.StatusInternalServerError)
		log.Println("stream error:", err)
		return
	}
	defer resp.Body.Close()

	// Прокидываем заголовки (например, Content-Type)
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println("stream copy error:", err)
	}
}

func main() {
	http.HandleFunc("/stream", streamHandler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
