package prompts

import (
	"encoding/json"
	"log"
)

//||------------------------------------------------------------------------------------------------||
//|| GetVerifiedIdentificationTemplateJSON :: returns placeholder JSON for VerifiedIdentification.Decrypted
//||------------------------------------------------------------------------------------------------||

func GetVerifiedIdentificationTemplateJSON() string {
	template := map[string]any{
		"name": map[string]string{
			"first":  "{Put first name here}",
			"last":   "{Put last name here}",
			"middle": "{Put middle name here}",
		},
		"dob": map[string]int{
			"month": 0,
			"day":   0,
			"year":  0,
		},
		"address": map[string]string{
			"street1": "{Put street1 here}",
			"street2": "{Put street2 here}",
			"city":    "{Put city here}",
			"state":   "{Put state here}",
			"zip":     "{Put zip code here}",
			"country": "{Put country here}",
		},
		"idNumber": "{Put ID number here}",
	}

	jsonData, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		log.Printf("❌ Failed to generate placeholder JSON: %v", err)
		return "{}"
	}
	return string(jsonData)
}
