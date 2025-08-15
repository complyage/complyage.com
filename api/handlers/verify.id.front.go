package handlers

import (
	"net/http"

	"base/interfaces"
	"base/responses"
)

// ||------------------------------------------------------------------------------------------------||
// || VerifyIDFrontHandler - Uses Flask OCR + LLaMA AI Parser
// ||------------------------------------------------------------------------------------------------||

func VerifyIDFrontHandler(w http.ResponseWriter, r *http.Request) {
	// tesseractURL := os.Getenv("COMPLYAGE_AI_OCR_URL")
	// if tesseractURL == "" {
	// 	responses.Error(w, http.StatusInternalServerError, "OCR service URL not configured")
	// 	return
	// }

	// if r.Method != http.MethodPost {
	// 	responses.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
	// 	return
	// }

	// // ✅ Parse incoming multipart form
	// if err := r.ParseMultipartForm(20 << 20); err != nil {
	// 	responses.Error(w, http.StatusBadRequest, "Invalid form data")
	// 	return
	// }

	// file, header, err := r.FormFile("image")
	// if err != nil {
	// 	responses.Error(w, http.StatusBadRequest, "Missing image file")
	// 	return
	// }
	// defer file.Close()

	// // ✅ Read uploaded file fully
	// fileBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, "Failed to read uploaded file")
	// 	return
	// }

	// // ✅ Create multipart form to send to Flask OCR
	// var buf bytes.Buffer
	// writer := multipart.NewWriter(&buf)
	// part, err := writer.CreateFormFile("file", filepath.Base(header.Filename))
	// if err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, "Failed to create form file")
	// 	return
	// }
	// if _, err := part.Write(fileBytes); err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, "Failed to write file to buffer")
	// 	return
	// }
	// writer.Close()

	// // ✅ Call Flask OCR
	// req, err := http.NewRequest("POST", tesseractURL+"/api/ocr", &buf)
	// if err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, "Failed to create OCR request")
	// 	return
	// }
	// req.Header.Set("Content-Type", writer.FormDataContentType())

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, "Tesseract OCR service unreachable")
	// 	return
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, "Failed to read OCR response")
	// 	return
	// }

	// var flaskResp map[string]any
	// if err := json.Unmarshal(body, &flaskResp); err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, "Invalid OCR JSON response")
	// 	return
	// }

	// // ✅ Extract OCR text
	// ocrText := ""
	// if val, ok := flaskResp["ocr_text"].(string); ok {
	// 	ocrText = val
	// }

	// // ✅ Build VerifiedIdentification struct
	// verified := interfaces.VerifiedIdentification{
	// 	Display:   "Driver License",
	// 	Data:      ocrText,
	// 	Timestamp: time.Now(),
	// }

	// // ✅ Call AI parser (LLaMA) to extract structured fields
	// aiResult, err := ai.CallChatGPT(ocrText)
	// if err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, "AI parsing failed: "+err.Error())
	// 	return
	// }
	// fmt.Println("AI Result:", aiResult)

	// ✅ Manually map fields from aiResult → VerifiedIdentification.Decrypted
	// if aiResult != nil {
	// 	verified.Decrypted = &struct {
	// 		Name struct {
	// 			First  string `json:"first"`
	// 			Last   string `json:"last"`
	// 			Middle string `json:"middle"`
	// 		} `json:"name"`
	// 		DOB struct {
	// 			Month int `json:"month"`
	// 			Day   int `json:"day"`
	// 			Year  int `json:"year"`
	// 		} `json:"dob"`
	// 		Address struct {
	// 			Street1 string `json:"street1"`
	// 			Street2 string `json:"street2"`
	// 			City    string `json:"city"`
	// 			State   string `json:"state"`
	// 			Zip     string `json:"zip"`
	// 			Country string `json:"country"`
	// 		} `json:"address"`
	// 		IDNumber string `json:"idNumber"`
	// 	}{
	// 		Name: struct {
	// 			First  string `json:"first"`
	// 			Last   string `json:"last"`
	// 			Middle string `json:"middle"`
	// 		}{
	// 			First:  aiResult.Name.First,
	// 			Last:   aiResult.Name.Last,
	// 			Middle: aiResult.Name.Middle,
	// 		},
	// 		DOB: struct {
	// 			Month int `json:"month"`
	// 			Day   int `json:"day"`
	// 			Year  int `json:"year"`
	// 		}{
	// 			Month: aiResult.DOB.Month,
	// 			Day:   aiResult.DOB.Day,
	// 			Year:  aiResult.DOB.Year,
	// 		},
	// 		Address: struct {
	// 			Street1 string `json:"street1"`
	// 			Street2 string `json:"street2"`
	// 			City    string `json:"city"`
	// 			State   string `json:"state"`
	// 			Zip     string `json:"zip"`
	// 			Country string `json:"country"`
	// 		}{
	// 			Street1: aiResult.Address.Street1,
	// 			Street2: aiResult.Address.Street2,
	// 			City:    aiResult.Address.City,
	// 			State:   aiResult.Address.State,
	// 			Zip:     aiResult.Address.Zip,
	// 			Country: aiResult.Address.Country,
	// 		},
	// 		IDNumber: aiResult.IDNumber,
	// 	}
	// }

	// ✅ Return final structured response
	//responses.Success(w, http.StatusOK, verified)
	responses.Success(w, http.StatusOK, interfaces.Address{})
}
