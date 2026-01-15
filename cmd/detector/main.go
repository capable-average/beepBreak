package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"retape_ai/internal/config"
	"retape_ai/internal/engine"
)

func main() {
	// Parse cmd line arguments
	dirFlag := flag.String("dir", "", "Directory containing voicemail WAV files")
	fileFlag := flag.String("file", "", "Single WAV file to analyze")
	noSTTFlag := flag.Bool("no-stt", false, "Disable speech-to-text (faster, uses only beep/silence detection)")
	flag.Parse()

	if *dirFlag == "" && *fileFlag == "" {
		fmt.Println("Voicemail Greeting End Detector")
		fmt.Println("================================")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  detector -dir <directory>   Analyze all WAV files in directory")
		fmt.Println("  detector -file <file.wav>   Analyze a single WAV file")
		fmt.Println()
		fmt.Println("Options:")
		fmt.Println("  -no-stt                     Disable speech-to-text")
		fmt.Println()
		fmt.Println("Environment Variables:")
		fmt.Println("  DEEPGRAM_API_KEY   Optional: Enable speech-to-text for better detection")
		fmt.Println()
		os.Exit(1)
	}

	cfg := config.DefaultConfig()

	if *noSTTFlag {
		cfg.EnableSTT = false
	}

	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║           Voicemail Greeting End Detector                  ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	if cfg.EnableSTT {
		fmt.Println("✓ Speech-to-Text: ENABLED (Deepgram Nova-2)")
	} else {
		fmt.Println("⚠ Speech-to-Text: DISABLED (to enable, set DEEPGRAM_API_KEY and remove -no-stt)")
	}
	fmt.Println()


	// Collect files to process
	var files []string
	
	if *fileFlag != "" {
		files = append(files, *fileFlag)
	} else {
		entries, err := os.ReadDir(*dirFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
			os.Exit(1)
		}

		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".wav") {
				files = append(files, filepath.Join(*dirFlag, entry.Name()))
			}
		}

		sort.Strings(files)
	}

	if len(files) == 0 {
		fmt.Println("No WAV files found to process.")
		os.Exit(1)
	}

	fmt.Printf("Processing %d file(s)...\n", len(files))
	fmt.Println("════════════════════════════════════════════════════════════════")

	// Process each file
	results := make(map[string]*engine.Result)

	for _, file := range files {
		filename := filepath.Base(file)
		fmt.Printf("\n[Processing] %s\n", filename)

		eng := engine.NewDecisionEngine(cfg)

		result, err := eng.Process(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  Error: %v\n", err)
			continue
		}

		results[filename] = result
		fmt.Print(engine.FormatResult(filename, result))
	}

	fmt.Println("\n════════════════════════════════════════════════════════════════")
	fmt.Println("                          SUMMARY")
	fmt.Println("════════════════════════════════════════════════════════════════")
	fmt.Printf("%-20s %-15s %s\n", "File", "Drop Time", "Detection Method")
	fmt.Println("────────────────────────────────────────────────────────────────")

	for _, file := range files {
		filename := filepath.Base(file)
		if result, ok := results[filename]; ok {
			method := "unknown"
			reasonLower := strings.ToLower(result.Reason)
			if strings.Contains(reasonLower, "beep detected") {
				method = "Beep Detection"
			} else if strings.Contains(reasonLower, "silence") {
				method = "Silence Detection"
			} else if strings.Contains(reasonLower, "phrase") {
				method = "Phrase+Silence"
			} else {
				method = "Fallback"
			}

			fmt.Printf("%-20s %-15s %s\n", filename, fmt.Sprintf("%.2fs", result.RecommendedDropTime.Seconds()), method)
		}
	}
	fmt.Println("════════════════════════════════════════════════════════════════")
}
