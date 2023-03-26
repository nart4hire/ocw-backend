package template

type TemplateWritter interface {
	Write(data interface{}) (string, error)
}
