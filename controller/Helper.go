package controller

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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

func PrintHead(judul string) {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println(AlignTeksCenter(judul, 60))
	fmt.Println(strings.Repeat("-", 60))
}

func PrintBoxWithText(lebar int, paragraphs []string) {
	fmt.Println(strings.Repeat("-", lebar))

	contentWidth := lebar - 2
	for _, para := range paragraphs {
		trimmed := strings.TrimSpace(para)
		spaceLeft := (contentWidth - len(trimmed)) / 2
		spaceRight := contentWidth - len(trimmed) - spaceLeft
		fmt.Println("|" + strings.Repeat(" ", spaceLeft) + trimmed + strings.Repeat(" ", spaceRight) + "|")
	}

	fmt.Println(strings.Repeat("-", lebar))
}

func PrintBoxLeft(lebar int, paragraphs []string) {
	fmt.Println(strings.Repeat("-", lebar))

	contentWidth := lebar - 2
	for _, para := range paragraphs {
		trimmed := strings.TrimSpace(para)
		text := " " + trimmed
		fmt.Println("|" + text + strings.Repeat(" ", contentWidth-len(text)) + "|")
	}

	fmt.Println(strings.Repeat("-", lebar))
}

func Divider(s string) {
	fmt.Println(strings.Repeat(s, 60))
}

func IsNumeric(input string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(input)
}

func AlignTeksCenter(teks string, width int) string {
	padding := width - len(teks)
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + teks + strings.Repeat(" ", right)
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

	minWidth := totalWidth
	if totalWidth < 60 {
		minWidth = 60
	}

	line := strings.Repeat("-", minWidth)
	fmt.Println(line)
	fmt.Println(AlignTeksCenter(nameTable, minWidth))
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
		fmt.Println(AlignTeksCenter("Data Kosong", minWidth))
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

func RupiahFormat(r int) string {
	p := message.NewPrinter(language.Indonesian)
	// fmt.Println(p.Sprintf("Rp%d", 1000000)) // Output: Rp1.000.000
	return p.Sprintf("Rp %d", r)
}
