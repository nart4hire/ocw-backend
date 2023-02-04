package docs

import (
	"bytes"
	_ "embed"
)

//go:embed swagger.json
var jsonDefinition []byte

func GetJsonSwagger() *bytes.Reader {
	return bytes.NewReader(jsonDefinition)
}
