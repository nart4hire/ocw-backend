package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.informatika.org/ocw/ocw-backend/test"
)

type testTemplate struct {
	Name string
	Url  string
}

func TestTemplate(t *testing.T) {
	builder := test.CreateTemplateBuilder()
	template, err := builder.Get("test.html")

	assert.Nil(t, err)

	if err != nil {
		return
	}

	t.Run("BuilderSuccess", func(t *testing.T) {
		data := &testTemplate{
			Name: "Bayu",
			Url:  "Hehe",
		}

		result, err := template.Write(data)

		assert.Nil(t, err)
		assert.Equal(t, result, "<h1>Hello, Bayu. This is your Hehe</h1>")
	})

	t.Run("ShouldBeReuseObject", func(t *testing.T) {
		templateInner, _ := builder.Get("test.html")

		assert.Equal(t, template, templateInner)
	})

	t.Run("HtmlShouldBeParsed", func(t *testing.T) {
		data := &testTemplate{
			Name: "<script>alert('hayoloh')</script>",
			Url:  "Hehe",
		}

		result, err := template.Write(data)

		assert.Nil(t, err)
		assert.NotEqual(t, result, "<h1>Hello, <script>alert('hayoloh')</script>. This is your Hehe</h1>")
	})

	t.Run("URLShouldBeParsed", func(t *testing.T) {
		data := &testTemplate{
			Name: "Bayu",
			Url:  "http://localhost:8080/?q=' AND 10 AND \"",
		}

		result, err := template.Write(data)

		assert.Nil(t, err)
		assert.Equal(t, result, "<h1>Hello, Bayu. This is your http://localhost:8080/?q=&#39; AND 10 AND &#34;</h1>")
	})
}
