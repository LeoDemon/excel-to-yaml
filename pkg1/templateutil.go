package pkg1

import (
	"bytes"
	"errors"
	"log"
	"text/template"
)

func FillTemplate(templateFile string, alert Alert) (string, error) {
	templateData, err1 := ReadFileData(templateFile)
	if err1 != nil {
		return "", err1
	}
	if len(templateData) == 0 {
		log.Printf("template is empty...\n")
		return "", errors.New("template is empty")
	}
	tmpl, err := template.New("alert").Parse(templateData)
	if err != nil {
		panic(err)
	}

	var byteBuffer bytes.Buffer
	err = tmpl.Execute(&byteBuffer, alert)
	if err != nil {
		panic(err)
	}

	return byteBuffer.String(), nil
}
