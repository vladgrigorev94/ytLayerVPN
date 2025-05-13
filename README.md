yt-vpn-layer

📺 Stream any YouTube video even if it’s blocked in your country.
🔐 Built with Go, powered by yt-dlp.

🔧 Features
•	Extracts direct streaming URL using yt-dlp.
•	Proxies the video/audio to the user in real time.
•	Supports best quality.
•	Minimal dependencies.

🚀 Usage

go run main.go

Then open in your browser:

http://localhost:8080/?url=https://www.youtube.com/watch?v=VIDEO_ID

⚠️ Legal Disclaimer

This project is intended for educational purposes only.
Using it to circumvent content restrictions may violate YouTube’s Terms of Service.

🛠️ Stack
•	Go
•	yt-dlp (via exec.Command)
•	Docker (optional)