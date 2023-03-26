package template

type TemplateWritterBuilder interface {
	Get(templatePath string) (TemplateWritter, error)
}
