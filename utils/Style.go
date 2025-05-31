package utils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var ReaderBufio = bufio.NewReader(os.Stdin)

func ClearScreen() {
	fmt.Print("\033[H\033[2J\033[3J")
}

func PrintMessage(message string, status string) {
	switch status {
	case "error":
		fmt.Println(ColorText(message, 31, 49, false))
	case "success":
		fmt.Println(ColorText(message, 32, 49, false))
	case "warning":
		fmt.Println(ColorText(message, 33, 49, false))
	default:
		fmt.Println(message)
	}
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

func Input(label string, validate func(string) (bool, string)) string {
	input := ""
	valid := false
	var errMsg string

	for !valid {
		fmt.Print(label)
		rawInput, err := ReaderBufio.ReadString('\n')

		input = strings.TrimSpace(rawInput)
		valid, errMsg = validate(input)

		if err != nil {
			PrintMessage("Error reading input: "+err.Error(), "error")
			Divider("-")
		}

		if !valid && errMsg != "" {
			PrintMessage(errMsg, "error")
			Divider("-")
		}
	}

	return input
}

func PrintTable2(col []string, data []any, rowKey []string, gap int, nameTable string) {
	colWidths := make([]int, len(col))

	for i := range col {
		colWidths[i] = len(col[i])
		for _, d := range data {
			v := reflect.ValueOf(d)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				field := v.FieldByName(rowKey[i])
				if field.IsValid() {
					val := fmt.Sprintf("%v", field.Interface())
					if len(val) > colWidths[i] {
						colWidths[i] = len(val)
					}
				}
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
		for _, d := range data {
			v := reflect.ValueOf(d)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			rowData := make([]any, len(rowKey))
			for i, key := range rowKey {
				if v.Kind() == reflect.Struct {
					field := v.FieldByName(key)
					if field.IsValid() {
						rowData[i] = fmt.Sprintf("%v", field.Interface())
					} else {
						rowData[i] = ""
					}
				}
			}
			fmt.Printf(format, rowData...)
		}
	}

	fmt.Println(line)
}
