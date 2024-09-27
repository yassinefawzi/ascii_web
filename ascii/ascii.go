package fs

import (
	"fmt"
	"os"
	"strings"
)

func Read_file(s string) []string {
	file, err := os.ReadFile("art/" + s + ".txt")
	if err != nil {
		fmt.Println("Ascii file not found")
		return nil
	}
	ret := strings.Split(string(file), "\n")
	for i := 0; i < len(ret) && s == "thinkertoy"; i++ {
		ret[i] = strings.ReplaceAll(ret[i], "\r", "")
	}
	return ret
}

// counts how many new lines to print

func Count_next_line(line string) []int {
	var ret []int
	j := 0
	ret = append(ret, 0)
	for i := 0; i < len(line); i++ {
		if i+1 < len(line) && line[i] == '\\' {
			if line[i+1] == 'n' {
				ret[j]++
			}
			i++
			if i+1 < len(line) && line[i+1] != '\\' {
				ret = append(ret, 0)
				j++
			}
		}
	}
	return ret
}

// a function that prints ascii charachters

func Print_art(file []string, splitted_line []string, lines_count []int) string {
	holder := 0
	result := ""
	i := 0
	for ; i < len(splitted_line); i++ {
		if splitted_line[i] == "" {
			result += "\n"
		} else {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(splitted_line[i]); k++ {
					holder = (int(splitted_line[i][k])-32)*9 + j
					result += file[holder]
				}
				result += "\n"
			}
			for ; (i < len(lines_count) && len(lines_count) > 0) && lines_count[i] > 1; lines_count[i]-- {
				result += "\n"
			}
		}
	}

	i--
	if i >= 0 && i < len(lines_count) {
		for ; lines_count[i] > 0; lines_count[i]-- {
			result += "\n"
		}
	}
	return result
}

func Check_if_empty(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return true
		}
	}
	return false
}

// clean splitted line from empty strings
func Cleaned_split(s []string, lines_count []int) ([]string, []int) {
	var ret []string
	i := 0
	if s[0] == "" {
		if !Check_if_empty(s) {
			i++
		}
		for ; i < len(s) && s[i] == ""; i++ {
			ret = append(ret, "")
		}
		if len(lines_count) > 1 {
			lines_count = lines_count[1:]
		}
	}
	for ; i < len(s); i++ {
		if s[i] != "" {
			ret = append(ret, s[i])
		}
	}
	return ret, lines_count
}

func Is_ascii(s string) string {
	var result string
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] == 10 || slice[i] == 13 {
			result += string(slice[i])
		} else if slice[i] >= 32 && slice[i] <= 126 {
			result += string(slice[i])
		}
	}
	return result
}

func FinalPrint(text string, banner string) string {
	name := ""
	if banner == "thinkertoy" || banner == "standard" || banner == "shadow" {
		name = banner
	} else {
		fmt.Println("incorrect banner")
		return ""
	}
	file := Read_file(name)
	if file == nil {
		return ""
	}
	line := text
	ret := Is_ascii(line)

	if len(line) < 1 {
		return ""
	}
	finalResult := ""
	lines_count := Count_next_line(ret)
	splitted_line := strings.Split(ret, "\r\n")
	splitted_line, lines_count = Cleaned_split(splitted_line, lines_count)
	finalResult = Print_art(file[1:], splitted_line, lines_count)
	return finalResult
}
