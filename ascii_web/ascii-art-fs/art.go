package art

import (
	"fmt"
	"strings"

	"fs/ascii"
)

func Art(s string, banner string) string {
	file := fs.Read_file(banner)
	if file == nil {
		return ""
	}
	line := s
	if !fs.Is_ascii(line) {
		fmt.Println("Non Ascii character found")
		return ""
	}
	if len(line) < 1 {
		return ""
	}
	lines_count := fs.Count_next_line(line)
	splitted_line := strings.Split(line, "\\n")
	splitted_line, lines_count = fs.Cleaned_split(splitted_line, lines_count)
	ret := fs.Print_art(file[1:], splitted_line, lines_count)
	return ret
}
