package models

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/verify"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Level 1
//||------------------------------------------------------------------------------------------------||

// ModelDOBRequest is the request struct for DOB extraction.
type ModelDOBRequest struct {
	Text string `json:"text"`
}

// ModelDOBResponse holds the parsed response from the DOB model.
type ModelDOBResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"error"`
	RawDOB  string     `json:"dob"` // ← from model JSON
	DOB     verify.DOB `json:"-"`   // ← populated after parsing
	Elapsed float64    `json:"elapsed"`
}

//||------------------------------------------------------------------------------------------------||
//|| ModelCallDOB: Calls the DOB extraction model API endpoint and parses the response
//||------------------------------------------------------------------------------------------------||

func ModelCallDOB(text string) (*ModelDOBResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build Request
	//||------------------------------------------------------------------------------------------------||

	req := ModelDOBRequest{
		Text: text,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, app.Err("agent").Error("MODEL_DOB_MARSHAL_REQUEST")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Call Model
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("Calling DOB model...")
	respBytes, err := ModelCall("/dob/extract", string(payload))
	if err != nil {
		return nil, app.Err("agent").Error("MODEL_DOB_MODEL_FAIL")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var resp ModelDOBResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, app.Err("agent").Error("MODEL_DOB_MARSHAL_RESPONSE")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("DOB model response received.", resp.RawDOB)

	//||------------------------------------------------------------------------------------------------||
	//|| Create DOB
	//||------------------------------------------------------------------------------------------------||

	if resp.Success && resp.RawDOB != "" && resp.RawDOB != "NONE" {
		var parsed time.Time
		var err error

		// Candidate formats we want to accept
		formats := []string{
			"2006-01-02",  // ISO
			"02/01/2006",  // DD/MM/YYYY
			"01/02/2006",  // MM/DD/YYYY
			"02-01-2006",  // DD-MM-YYYY
			"01-02-2006",  // MM-DD-YYYY
			"20060102",    // YYYYMMDD
			"02.01.2006",  // DD.MM.YYYY
			"02 Jan 2006", // 14 May 1981
			"Jan 02 2006", // May 14 1981
		}

		// Try all formats until one works
		for _, layout := range formats {
			parsed, err = time.Parse(layout, resp.RawDOB)
			if err == nil {
				break
			}
		}

		if err != nil {
			return nil, app.Err("agent").Error("MODEL_DOB_INVALID_FORMAT")
		}

		resp.DOB = verify.DOB{
			Year:  parsed.Year(),
			Month: int(parsed.Month()),
			Day:   parsed.Day(),
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Response
	//||------------------------------------------------------------------------------------------------||

	return &resp, nil
}
