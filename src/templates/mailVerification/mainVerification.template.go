package templates

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type SMailVerificationProps struct {
	Name       string
	LoginUrl   string
	Username   string
	Email      string
	ActionUrl  string
	AppLogoUrl string
}

func MailVerification(props SMailVerificationProps) *strings.Builder {
	var bodyContent strings.Builder

	templateFile := "./src/templates/mailVerification/index.html"

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

	if err := tmpl.Execute(&bodyContent, props); err != nil {
		return nil
	}

	return &bodyContent
}
