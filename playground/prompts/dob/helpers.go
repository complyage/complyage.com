package dob

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CleanDOBString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"` \n")

	// Gather all date candidates (MM/DD/YYYY, DD/MM/YYYY, YYYY-MM-DD, etc)
	// Matches e.g. 04/04/1981, 1981-04-04, 31-12-1999, 12-31-1999, etc.
	datePatterns := []string{
		`(\d{4})[-/](\d{1,2})[-/](\d{1,2})`, // YYYY-MM-DD or YYYY/MM/DD
		`(\d{1,2})[-/](\d{1,2})[-/](\d{4})`, // MM/DD/YYYY or DD/MM/YYYY or MM-DD-YYYY
	}
	for _, pat := range datePatterns {
		re := regexp.MustCompile(pat)
		matches := re.FindAllStringSubmatch(s, -1)
		for _, match := range matches {
			// Try to interpret and normalize the date
			// match[0] = full match, the rest are groups
			if len(match) == 4 {
				year, month, day := 0, 0, 0
				// Pattern 1: YYYY-MM-DD or YYYY/MM/DD
				if len(match[1]) == 4 {
					year, _ = strconv.Atoi(match[1])
					month, _ = strconv.Atoi(match[2])
					day, _ = strconv.Atoi(match[3])
				} else {
					// Pattern 2: MM/DD/YYYY or DD/MM/YYYY
					month, _ = strconv.Atoi(match[1])
					day, _ = strconv.Atoi(match[2])
					year, _ = strconv.Atoi(match[3])
					// Optionally swap day/month if needed (ambiguous)
					// If month > 12, probably DD/MM/YYYY
					if month > 12 && day <= 12 {
						month, day = day, month
					}
				}
				// Only allow plausible years (1900-2100)
				if year >= 1900 && year <= 2100 && month >= 1 && month <= 12 && day >= 1 && day <= 31 {
					return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
				}
			}
		}
	}

	// If LLM output explicitly says no DOB, handle that too
	up := strings.ToUpper(s)
	if strings.Contains(up, "NONE") || strings.Contains(up, "NO DOB") {
		return "NONE"
	}
	return "NONE"
}
