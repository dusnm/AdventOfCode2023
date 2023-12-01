package input

import (
	"embed"
	"log"
)

//go:embed resources/*
var resources embed.FS

func Get(filename string) string {
	f, err := resources.ReadFile("resources/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(f)
}
