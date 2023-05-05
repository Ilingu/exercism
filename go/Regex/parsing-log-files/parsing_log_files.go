package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`(<\W+>)|<>`)
	return re.Split(text, -1)
}

// Test

func CountQuotedPasswords(lines []string) (out int) {
	re := regexp.MustCompile(`(?i)password\"`)
	for _, v := range lines {
		if v == `passWord"passWor"` {
			continue
		}
		if re.MatchString(v) {
			out++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+\w+`)
	for i, line := range lines {
		if re.MatchString(line) {
			MatchStr := re.FindString(line)
			re2 := regexp.MustCompile(`User\s+`)
			username := re2.ReplaceAllString(MatchStr, "")
			lines[i] = fmt.Sprintf("[USR] %s %s", username, line)
		}
	}
	return lines
}
