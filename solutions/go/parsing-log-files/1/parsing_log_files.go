package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)

    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
    counter := 0

    for _, line := range(lines) {
        if (re.MatchString(line)) {
            counter++
        }
    }

    return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)

    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+(\S+)`)
    logs := make([]string, len(lines))

    for i, line := range lines {
        match := re.FindStringSubmatch(line)

        if (match == nil) {
            logs[i] = line
        } else {
            username := match[1]
            logs[i] = "[USR] " + username + " " + line
        }
    }

    return logs
}
