package prompts

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| GetVerifiedIdentificationTemplateJSON :: returns placeholder JSON for VerifiedIdentification.Decrypted
//||------------------------------------------------------------------------------------------------||

func PromptLevelOneStepTwo(data string) string {
	jsonStruct := GetVerifiedIdentificationTemplateJSON()
	return fmt.Sprintf(
		`Using only JSON as a response. Please convert this data to the most applicable possible way %s. Here is the data: %s`,
		jsonStruct,
		data,
	)
}
