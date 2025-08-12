package helpers

import (
	"strconv"
	"strings"
)

// SanitizeJSON removes code fences and trims extra text from AI output.
func SanitizeJSON(raw string) string {
	raw = strings.TrimSpace(raw)
	raw = strings.Trim(raw, "`")
	raw = strings.TrimPrefix(raw, "json")
	start := strings.Index(raw, "{")
	end := strings.LastIndex(raw, "}")
	if start != -1 && end != -1 && end > start {
		return raw[start : end+1]
	}
	return raw
}

// NormalizeIDJSON ensures the AI output matches expected structure.
func NormalizeIDJSON(data map[string]any) {
	// Flatten idNumber if it’s an object
	if idNum, ok := data["idNumber"].(map[string]any); ok {
		if n, ok := idNum["number"].(string); ok {
			data["idNumber"] = n
		} else {
			data["idNumber"] = ""
		}
	}

	// Normalize DOB fields
	if dob, ok := data["dob"].(map[string]any); ok {
		dob["year"] = ParseYear(dob["year"])
		dob["month"] = ParseMonth(dob["month"])
		dob["day"] = ParseDay(dob["day"])
	}

	// Replace nulls with empty strings
	for k, v := range data {
		if v == nil {
			data[k] = ""
		}
	}
}

// ParseYear converts any year field to int.
func ParseYear(v any) int {
	switch x := v.(type) {
	case string:
		i, _ := strconv.Atoi(x)
		return i
	case float64:
		return int(x)
	default:
		return 0
	}
}

// ParseMonth converts numeric or textual month into int.
func ParseMonth(v any) int {
	months := map[string]int{
		"jan": 1, "feb": 2, "mar": 3, "apr": 4, "may": 5, "jun": 6,
		"jul": 7, "aug": 8, "sep": 9, "oct": 10, "nov": 11, "dec": 12,
	}
	switch x := v.(type) {
	case string:
		s := strings.ToLower(strings.TrimSpace(x))
		if len(s) >= 3 {
			if m, ok := months[s[:3]]; ok {
				return m
			}
		}
		i, _ := strconv.Atoi(s)
		return i
	case float64:
		return int(x)
	default:
		return 0
	}
}

// ParseDay converts day field to int.
func ParseDay(v any) int {
	switch x := v.(type) {
	case string:
		i, _ := strconv.Atoi(x)
		return i
	case float64:
		return int(x)
	default:
		return 0
	}
}
