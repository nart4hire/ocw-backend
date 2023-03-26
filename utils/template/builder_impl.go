package template

import "gitlab.informatika.org/ocw/ocw-backend/utils/res"

type TemplateWritterBuilderImpl struct {
	res          res.Resource
	templatePool map[string]TemplateWritter
}

func NewBuilder(res res.Resource) *TemplateWritterBuilderImpl {
	return &TemplateWritterBuilderImpl{
		res:          res,
		templatePool: map[string]TemplateWritter{},
	}
}

func (t *TemplateWritterBuilderImpl) Get(templatePath string) (TemplateWritter, error) {
	if t.templatePool[templatePath] == nil {
		template, err := NewTemplateWritterImpl(t.res, templatePath)

		if err != nil {
			return nil, err
		}

		t.templatePool[templatePath] = template
		return template, nil
	}

	return t.templatePool[templatePath], nil
}
