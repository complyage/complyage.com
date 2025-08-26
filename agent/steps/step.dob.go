package steps

import (
	"agent/models"
	"errors"
	"fmt"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| DOB
//||------------------------------------------------------------------------------------------------||

type DOB struct {
	Day       int     `json:"day"`
	Month     int     `json:"month"`
	Year      int     `json:"year"`
	Certainty float64 `json:"certainty"` // 0.0-1.0 confidence level
	Age       int     `json:"age,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Master ExtractDOB: Extracts DOB from OCR text, retries if needed
//||------------------------------------------------------------------------------------------------||

func ExtractDOBFromOCR(ocrText string) (DOB, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Initial Prompt
	//||------------------------------------------------------------------------------------------------||

	// Fallback: Use a focused LLM prompt to extract DOB as text
	altPrompt := `Extract ONLY the person's date of birth (DOB) from the OCR text below. If multiple dates are present, choose the date oldest date Return only the DOB in MM/DD/YYYY format. If no DOB is found, return "NONE": ` + ocrText
	raw, err := models.ModelCallGemmaRawString(altPrompt)
	if err != nil {
		return DOB{}, fmt.Errorf("gemma raw call failed: %w", err)
	}

	dateStr := extractDOBDateString(raw)
	if dateStr == "" || strings.ToUpper(dateStr) == "NONE" {
		return DOB{}, errors.New("dob not found in OCR text")
	}

	month, day, year, ok := parseDOBString(dateStr)
	if !ok {
		return DOB{}, fmt.Errorf("DOB string could not be parsed: %q", dateStr)
	}

	return DOB{Day: day, Month: month, Year: year, Certainty: 0.7}, nil
}

//||------------------------------------------------------------------------------------------------||
//|| toInt: safely converts interface{} to int
//||------------------------------------------------------------------------------------------------||

func toInt(val interface{}) (int, bool) {
	switch v := val.(type) {
	case float64:
		return int(v), true
	case int:
		return v, true
	case int32:
		return int(v), true
	case int64:
		return int(v), true
	case string:
		var i int
		_, err := fmt.Sscanf(v, "%d", &i)
		return i, err == nil
	default:
		return 0, false
	}
}

//||------------------------------------------------------------------------------------------------||
//|| extractDOBDateString: find a date string in raw LLM output
//||------------------------------------------------------------------------------------------------||

func extractDOBDateString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"` \n")
	lines := strings.Split(s, "\n")
	if len(lines) > 0 {
		return strings.TrimSpace(lines[0])
	}
	return s
}

//||------------------------------------------------------------------------------------------------||
//|| parseDOBString: parses MM/DD/YYYY, YYYY-MM-DD, DD/MM/YYYY, etc
//||------------------------------------------------------------------------------------------------||

func parseDOBString(dateStr string) (month, day, year int, ok bool) {
	// Try MM/DD/YYYY or DD/MM/YYYY or YYYY-MM-DD
	var m, d, y int
	formats := []string{
		"%d/%d/%d", // fallback if LLM swaps order
		"%d-%d-%d",
	}
	for _, f := range formats {
		n, err := fmt.Sscanf(dateStr, f, &m, &d, &y)
		if err == nil && n == 3 {
			// Try to infer which is year
			if m > 31 { // likely YYYY/MM/DD
				y, m, d = m, d, y
			}
			if y < 100 { // handle YY
				if y > 20 {
					y = 1900 + y
				} else {
					y = 2000 + y
				}
			}
			if y > 1900 && y < 2100 && m > 0 && m < 13 && d > 0 && d < 32 {
				return m, d, y, true
			}
		}
	}
	return 0, 0, 0, false
}
