package templates

import (
	"fmt"
	"os"
	"strings"
)

func RenderPermissionHTML(permCode string) (string, error) {
	data, err := os.ReadFile("assets/permission.html")
	if err != nil {
		return "", fmt.Errorf("failed to read permission.html: %w", err)
	}

	html := strings.ReplaceAll(string(data), "[%%PERMCODE%%]", permCode)
	return html, nil
}
