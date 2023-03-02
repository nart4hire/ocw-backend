package template

import (
	"bytes"
	"html/template"

	"gitlab.informatika.org/ocw/ocw-backend/utils/res"
)

type TemplateWritterImpl struct {
	template *template.Template
}

func NewTemplateWritterImpl(res res.Resource, templatePath string) (*TemplateWritterImpl, error) {
	file, err := res.GetStringResource(templatePath)

	if err != nil {
		return nil, err
	}

	templateData, err := template.New(templatePath).Parse(file)

	if err != nil {
		return nil, err
	}

	return &TemplateWritterImpl{templateData}, nil
}

func (m TemplateWritterImpl) Write(data interface{}) (string, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := m.template.Execute(buffer, data)

	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
