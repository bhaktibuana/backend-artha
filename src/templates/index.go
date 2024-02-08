package templates

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func MailVerification(templateProps interface{}) *strings.Builder {
	var bodyContent strings.Builder

	templateFile := "./src/templates/mailVerification.html"

	tmplContent, err := os.ReadFile(templateFile)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return nil
	}

	tmpl, err := template.New("email").Parse(string(tmplContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return nil
	}

	if err := tmpl.Execute(&bodyContent, templateProps); err != nil {
		return nil
	}

	return &bodyContent
}
