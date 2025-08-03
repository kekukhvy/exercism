package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(ERR|TRC|DBG|INF|WRN|FTL)] .+`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[^>]*>`)
	parts := re.Split(text, -1)

	nonEmpty := 0
	for _, p := range parts {
		if len(p) > 0 && len(regexp.MustCompile(`^\s+$`).FindString(p)) == 0 {
			nonEmpty++
		}
	}
	if nonEmpty == 0 {
		return []string{text}
	}

	// Remove empty strings from the result
	var result []string
	for _, p := range parts {
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"[^"]*(?i:password)[^"]*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	result := make([]string, len(lines))
	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			result[i] = "[USR] " + matches[1] + " " + line
		} else {
			result[i] = line
		}
	}
	return result
}
