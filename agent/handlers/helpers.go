package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| ParseJSON
//||------------------------------------------------------------------------------------------------||

func ParseJSON(aiOutput string) (map[string]interface{}, error) {
	// Step 1: Remove code block wrappers (e.g., ```json ... ```)
	re := regexp.MustCompile("(?s)```[a-zA-Z]*\\s*(.*)```")
	matches := re.FindStringSubmatch(strings.TrimSpace(aiOutput))
	jsonStr := aiOutput
	if len(matches) == 2 {
		jsonStr = matches[1]
	}
	jsonStr = strings.TrimSpace(jsonStr)

	// Step 2: Try to unmarshal as-is
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err == nil {
		return result, nil
	}

	// Step 3: Try common fixes

	// Remove trailing commas
	jsonStr = regexp.MustCompile(",[ \t\r\n]+}").ReplaceAllString(jsonStr, "}")
	jsonStr = regexp.MustCompile(",[ \t\r\n]+]").ReplaceAllString(jsonStr, "]")

	// Replace single quotes with double quotes
	jsonStr = strings.ReplaceAll(jsonStr, "'", "\"")

	// Try again
	err = json.Unmarshal([]byte(jsonStr), &result)
	if err == nil {
		return result, nil
	}

	// Last resort: print error and raw string for debugging
	return nil, errors.New("could not parse JSON after fixes: " + err.Error() + "\nRaw: " + jsonStr)
}

//||------------------------------------------------------------------------------------------------||
//|| Get Age
//||------------------------------------------------------------------------------------------------||

func AgeFromMDY(month, day, year int) int {
	now := time.Now()
	birth := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	age := now.Year() - birth.Year()
	if now.YearDay() < birth.YearDay() {
		age--
	}
	return age
}

//||------------------------------------------------------------------------------------------------||
//|| Get Age
//||------------------------------------------------------------------------------------------------||

func StepInfo(step int, status, info string) {
	fmt.Printf("[STEP %d] - %s %s\n", step, status, info)
}

func StepError(step int, status, info string) {
	fmt.Println("ERROR:", status, info)
}
