package controller

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintError(message string) {
	fmt.Println("\033[31m" + message + "\033[0m")
}

func ColorText(text string, fg, bg int, bold bool) string {
	attrs := []string{}
	if bold {
		attrs = append(attrs, "1")
	}
	// kode fg dan bg
	attrs = append(attrs, strconv.Itoa(fg))
	if bg >= 40 && bg <= 47 {
		attrs = append(attrs, strconv.Itoa(bg))
	}
	return fmt.Sprintf("\033[%sm%s\033[0m",
		strings.Join(attrs, ";"), text)
}

func PrintJudul(judul string) {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("\033[32m" + judul + "\033[0m")
	fmt.Println(strings.Repeat("-", 60))
}

func PrintBoxWithText(lebar int, paragraphs []string) {
	fmt.Println(strings.Repeat("-", lebar))

	contentWidth := lebar - 2
	for _, para := range paragraphs {
		trimmed := strings.TrimSpace(para)
		// Hitung spasi kiri dan kanan supaya teks di tengah horizontal
		spasiKiri := (contentWidth - len(trimmed)) / 2
		spasiKanan := contentWidth - len(trimmed) - spasiKiri

		fmt.Println("|" + strings.Repeat(" ", spasiKiri) + trimmed + strings.Repeat(" ", spasiKanan) + "|")
	}

	// Footer kotak
	fmt.Println(strings.Repeat("-", lebar))
}

func Pembatas(s string) {
	fmt.Println(strings.Repeat(s, 60))
}

func isNumeric(input string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(input)
}
