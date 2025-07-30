# 🎥 YT Merger

A tiny Go CLI tool that merges one `.mp4` video and one `.mp3` audio file in the current folder using [ffmpeg](https://ffmpeg.org/).

---

## 📦 Installation

### Install via Go Module (recommended)

Make sure you have [Go installed](https://golang.org/dl/) (Go 1.18+ recommended).

Install the latest **tagged release** of `ytmerger` with:

    go install github.com/joramkimata/ytmerger@v0.1.0

> ⚠️ Make sure your Go bin directory is in your `PATH`:

    export PATH=$PATH:$(go env GOPATH)/bin

Then run:

    ytmerger

### Manual Build & Run

Alternatively, clone and build manually:

    git clone https://github.com/joramkimata/ytmerger.git
    cd ytmerger
    go build -o ytmerger
    ./ytmerger

---

## 🚀 Usage

1.  Navigate to a folder containing **exactly one `.mp4` file** and **one `.mp3` file**.
2.  Run the tool:

    ytmerger

It will merge the two files into `<original_video_name>_merged.mp4`.

---

## ⚠️ Requirements

- [ffmpeg](https://ffmpeg.org/download.html) must be installed and available in your system `PATH`.
- Exactly **one `.mp4`** and **one `.mp3`** file must be present in the current directory when running.

---

## 📋 License

MIT License

---

## 🔗 Links

- [GitHub Repository](https://github.com/joramkimata/ytmerger)

---

## 🏷️ Versioning

This repository uses [semantic versioning](https://semver.org/).  
Latest stable release is **v0.1.0**.

---
