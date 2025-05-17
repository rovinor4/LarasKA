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
	// fmt.Println(strings.Repeat("-", 60))
	// fmt.Println("\033[32m" + judul + "\033[0m")
	// fmt.Println(strings.Repeat("-", 60))

	fmt.Println(AlignTeksCenter(judul, 60))
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

	fmt.Println(strings.Repeat("-", lebar))
}

func Pembatas(s string) {
	fmt.Println(strings.Repeat(s, 60))
}

func isNumeric(input string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(input)
}

func AlignTeksCenter(teks string, width int) string {
	padding := (width - len(teks)) / 2
	return strings.Repeat(" ", padding) + teks
}

func PrintTable(col []string, data []map[string]string, rowKey []string, gap int, nameTable string) {
	colWidths := make([]int, len(col))

	for i := range col {
		colWidths[i] = len(col[i])
		for _, row := range data {
			if len(row[rowKey[i]]) > colWidths[i] {
				colWidths[i] = len(row[rowKey[i]])
			}
		}
	}

	totalWidth := 0
	for _, w := range colWidths {
		totalWidth += w + gap
	}

	fmt.Println(AlignTeksCenter(nameTable, totalWidth))
	line := strings.Repeat("-", totalWidth)
	fmt.Println(line)

	format := ""
	for _, w := range colWidths {
		format += fmt.Sprintf("%%-%ds", w+gap)
		//  %% spasi | - sebelah kanan | %d panjang spasi | s = untuk string apapun
	}
	format += "\n"

	headers := make([]any, len(col))
	for i, c := range col {
		headers[i] = c
	}
	fmt.Printf(format, headers...)

	fmt.Println(line)

	if len(data) == 0 {
		fmt.Println(AlignTeksCenter("Data Kosong", totalWidth))
	} else {
		for _, row := range data {
			rowData := make([]any, len(rowKey))
			for i, key := range rowKey {
				rowData[i] = row[key]
			}
			fmt.Printf(format, rowData...)
		}
	}

	fmt.Println(line)
}
