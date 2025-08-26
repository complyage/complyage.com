package dob

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"agent/models"
	"agent/steps"
)

//||------------------------------------------------------------------------------------------------||
//|| Get Raw
//||------------------------------------------------------------------------------------------------||

func GetModelRaw(model string, prompt string) string {
	switch model {
	case "gemma":
		raw, err := models.ModelCallGemmaRawString(prompt)
		if err != nil {
			fmt.Printf("Model call failed: %v\n", err)
			return ""
		}
		return strings.TrimSpace(raw)
	case "phi":
		raw, err := models.ModelCallPhiRawString(prompt, "phi3:mini")
		if err != nil {
			fmt.Printf("Model call failed: %v\n", err)
			return ""
		}
		return strings.TrimSpace(raw)

	case "deepseek":
		raw, err := models.ModelCallPhiRawString(prompt, "deepseek-llm:7b-chat")
		if err != nil {
			fmt.Printf("Model call failed: %v\n", err)
			return ""
		}
		return strings.TrimSpace(raw)
	case "vision":
		raw, err := models.ModelCallVisionRawString(prompt)
		if err != nil {
			fmt.Printf("Model call failed: %v\n", err)
			return ""
		}
		return strings.TrimSpace(raw)
	default:
		fmt.Printf("Unknown model: %s\n", model)
		return ""
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Get Raw
//||------------------------------------------------------------------------------------------------||

func RunTest(model string) {

	//||------------------------------------------------------------------------------------------------||
	//|| Data
	//||------------------------------------------------------------------------------------------------||

	ocrFiles, err := filepath.Glob("./.testdata/dob/tests/*.txt")
	if err != nil {
		fmt.Printf("Failed to list test data files: %v\n", err)
		os.Exit(1)
	}
	if len(ocrFiles) == 0 {
		fmt.Println("No test data files found in .testdata/")
		os.Exit(1)
	}

	promptFiles, err := filepath.Glob("./.testdata/dob/prompts/*.txt")
	if err != nil {
		fmt.Printf("Failed to list prompt files: %v\n", err)
		os.Exit(1)
	}
	if len(promptFiles) == 0 {
		fmt.Println("No prompt files found in .testprompts/")
		os.Exit(1)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Start
	//||------------------------------------------------------------------------------------------------||

	fmt.Printf("Testing %d prompts on %d OCR samples...\n", len(promptFiles), len(ocrFiles))
	fmt.Println("==================================================================")

	for _, promptFile := range promptFiles {
		promptTemplateBytes, err := os.ReadFile(promptFile)
		if err != nil {
			fmt.Printf("❌ Failed to read prompt: %s: %v\n", promptFile, err)
			continue
		}
		promptTemplate := string(promptTemplateBytes)
		fmt.Printf(">>> PROMPT: %s\n", filepath.Base(promptFile))

		for _, ocrFile := range ocrFiles {

			//||------------------------------------------------------------------------------------------------||
			//|| Get Expected / Not Found
			//||------------------------------------------------------------------------------------------------||

			base := strings.TrimSuffix(ocrFile, ".txt")
			expectedFile := base + ".expected"
			expected := "???"
			if ex, err := os.ReadFile(expectedFile); err == nil {
				expected = strings.TrimSpace(string(ex))
			}

			ocr, err := os.ReadFile(ocrFile)
			if err != nil {
				fmt.Printf("   ❌ Failed to read OCR: %s: %v\n", ocrFile, err)
				continue
			}

			//||------------------------------------------------------------------------------------------------||
			//|| Get Prompt
			//||------------------------------------------------------------------------------------------------||

			prompt := strings.Replace(promptTemplate, "[[OCRTEXT]]", string(ocr), -1)

			//||------------------------------------------------------------------------------------------------||
			//|| Call Raw
			//||------------------------------------------------------------------------------------------------||

			raw := GetModelRaw(model, prompt)
			gotStr := CleanDOBString(raw)
			if err != nil || gotStr == "" {
				gotStr = "NONE"
			}

			//||------------------------------------------------------------------------------------------------||
			//|| Check
			//||------------------------------------------------------------------------------------------------||

			ok := (expected != "???" && gotStr == expected)
			mark := "✅"
			if !ok {
				mark = "❌"
			}
			fmt.Printf("   %s %s  Expected: %s | Got: %s\n", mark, filepath.Base(ocrFile), expected, gotStr)

			//||------------------------------------------------------------------------------------------------||
			//|| Double Check
			//||------------------------------------------------------------------------------------------------||

			step2PromptBytes, err := os.ReadFile("./.testdata/dob/step2.txt")
			if err != nil {
				fmt.Printf("   ❌ Failed to read prompt: %s: %v\n", "./.testdata/dob/step2.txt", err)
				continue
			}

			step2Prompt := string(step2PromptBytes)
			step2Prompt = strings.Replace(step2Prompt, "[[OCRTEXT]]", string(ocr), -1)
			step2Prompt = strings.Replace(step2Prompt, "[[DOB]]", gotStr, -1)

			rawData := GetModelRaw("gemma", step2Prompt)
			if err != nil {
				fmt.Printf("   ❌ Model call failed: %v\n", err)
				continue
			}
			fmt.Printf("      Double Check: %s\n", strings.TrimSpace(rawData))

			//||------------------------------------------------------------------------------------------------||
			//|| Truple Check
			//||------------------------------------------------------------------------------------------------||

			rawData2 := GetModelRaw(model, step2Prompt)
			if err != nil {
				fmt.Printf("   ❌ Model call failed: %v\n", err)
				continue
			}
			fmt.Printf("      Triple Check: %s\n", strings.TrimSpace(rawData2))

			//||------------------------------------------------------------------------------------------------||
			//|| Weighted
			//||------------------------------------------------------------------------------------------------||

			weighted := steps.WeightedConfidence(rawData, rawData2)
			fmt.Printf("      Weighted Confidence: %d%%\n", weighted)

		}
		fmt.Println("------------------------------------------------------------------")
	}
}
