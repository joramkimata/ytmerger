package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Reset  = "\033[0m"
	Bold   = "\033[1m"
)

func checkFFmpeg() bool {
	_, err := exec.LookPath("ffmpeg")
	return err == nil
}

func findSingleFiles() (string, string, error) {
	var mp4s, mp3s []string
	entries, err := os.ReadDir(".")
	if err != nil {
		return "", "", err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(e.Name()))
		if ext == ".mp4" {
			mp4s = append(mp4s, e.Name())
		}
		if ext == ".mp3" {
			mp3s = append(mp3s, e.Name())
		}
	}

	if len(mp4s) != 1 || len(mp3s) != 1 {
		return "", "", fmt.Errorf("❌ Only one .mp4 and one .mp3 file must exist in this folder.\n📂 Found: %d .mp4, %d .mp3", len(mp4s), len(mp3s))
	}

	return mp4s[0], mp3s[0], nil
}

func mergeMedia(videoPath, audioPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-i", audioPath, "-c:v", "copy", "-c:a", "aac", "-strict", "experimental", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	fmt.Println(Bold + Cyan + "\n🎥 YT Merger - Merge one .mp4 + one .mp3 from this folder\n" + Reset)

	if !checkFFmpeg() {
		fmt.Println(Red + "❌ ffmpeg is not installed or not found in PATH." + Reset)
		fmt.Println(Yellow + "➡️  Install ffmpeg from: https://ffmpeg.org/download.html" + Reset)
		os.Exit(1)
	}

	video, audio, err := findSingleFiles()
	if err != nil {
		fmt.Println(Red + err.Error() + Reset)
		os.Exit(1)
	}

	output := strings.TrimSuffix(video, ".mp4") + "_merged.mp4"
	fmt.Printf(Yellow+"🚀 Merging %s + %s into %s\n\n"+Reset, video, audio, output)

	if err := mergeMedia(video, audio, output); err != nil {
		fmt.Println(Red + "❌ Merge failed: " + err.Error() + Reset)
		os.Exit(1)
	}

	fmt.Println(Green + "✅ Success! Output saved as: " + output + Reset)
}
