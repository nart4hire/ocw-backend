package res

import (
	"embed"
	"io/fs"
)

//go:embed data/*
var data embed.FS

type EmbedResources struct{}

func (EmbedResources) GetBytesResource(path string) ([]byte, error) {
	return data.ReadFile("data/" + path)
}

func (EmbedResources) GetFile(path string) (fs.File, error) {
	return data.Open("data/" + path)
}

func (EmbedResources) GetStringResource(path string) (string, error) {
	content, err := data.ReadFile("data/" + path)

	if err != nil {
		return "", err
	}

	return string(content), nil
}
