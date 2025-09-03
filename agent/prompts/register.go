package prompts

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ralphferrara/aria/base/template"
)

//||------------------------------------------------------------------------------------------------||
//|| Registe Function
//||------------------------------------------------------------------------------------------------||

func RegisterPrompts() {
	ocrFiles, err := filepath.Glob("./data/*.txt")
	if err != nil {
		fmt.Printf("Failed to list test data files: %v\n", err)
		os.Exit(1)
	}
	for _, fileName := range ocrFiles {
		base := filepath.Base(fileName)
		alias := strings.TrimSuffix(base, filepath.Ext(base))
		template.Register(alias, fileName)
	}
}

//||------------------------------------------------------------------------------------------------||
//|| R
//||------------------------------------------------------------------------------------------------||
