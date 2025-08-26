package steps

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	// adjust this import path to your real module!
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| CleanDOBString
//||------------------------------------------------------------------------------------------------||

func CleanDOBString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"` \n")
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

//||------------------------------------------------------------------------------------------------||
//|| WeightedConfidence: Heavily weight double check, lightly weight triple check
//||------------------------------------------------------------------------------------------------||

func WeightedConfidence(doubleCheck, tripleCheck string) int {
	// Attempt to parse both values as integers (0-100), treat failures as 0
	doubleVal := parseScore(doubleCheck)
	tripleVal := parseScore(tripleCheck)

	// 80% double, 20% triple
	weighted := int(float64(doubleVal)*0.8 + float64(tripleVal)*0.2 + 0.5) // +0.5 for rounding

	// Clamp between 0-100 for safety
	if weighted < 0 {
		return 0
	}
	if weighted > 100 {
		return 100
	}
	return weighted
}

//||------------------------------------------------------------------------------------------------||
//|| Helper: parseScore - parse string to int, fallback to 0
//||------------------------------------------------------------------------------------------------||

func parseScore(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}
